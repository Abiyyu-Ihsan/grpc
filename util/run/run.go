package run

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RegisterHTTPHandler register endpoint to http server
type RegisterHTTPHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

// HTTPMiddleware is middleware for http handler
type HTTPMiddleware func(handler http.Handler) http.Handler

// HTTPOption are settings for http server
type HTTPOption struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func defaultServerOptions() []grpc.ServerOption {
	handler := func(p interface{}) (err error) {
		return fmt.Errorf("internal server error : %v", err)
	}
	opts := []recovery.Option{
		recovery.WithRecoveryHandler(handler),
	}
	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(opts...),
		)}

	return serverOptions
}

func CreateGRPCServer() *grpc.Server {
	return grpc.NewServer(defaultServerOptions()...)
}

// DefaultHTTPOption return default option read and write timeout
func DefaultHTTPOption() HTTPOption {
	return HTTPOption{ReadTimeout: 60 * time.Second, WriteTimeout: 60 * time.Second}
}

// DefaultHTTPHandler specifies default http handler
func DefaultHTTPHandler(handler http.Handler) http.Handler {
	handler = CORSHandler(handler)
	handler = HealthCheckHandler(handler)
	return handler
}

// HealthCheckHandler returns http 200 for root path
func HealthCheckHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.WriteHeader(http.StatusOK)
		} else {
			handler.ServeHTTP(w, r)
		}
	})
}

// CORSHandler enables cross-origin resource sharing
func CORSHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				headers := []string{"Content-Type", "Accept", "Authorization"}
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
				methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
				w.Header().Set("Access-Control-Max-Age", "86400")
				return
			}
		}
		handler.ServeHTTP(w, r)
	})
}

// ServeHTTP is server Rest API from grpc proto
func ServeHTTP(address, port string, register RegisterHTTPHandler, option HTTPOption, handlers ...HTTPMiddleware) {
	// opts := []grpc.DialOption{grpc.WithInsecure()} // deprecated
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// create server mux
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}))

	// register
	err := register(ctx, mux, address, opts)
	if err != nil {
		panic(err)
	}

	// add handle middleware
	var handler http.Handler
	handler = mux
	for _, hm := range handlers {
		handler = hm(handler)
	}

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), handler)
	if err != nil {
		panic(err)
	}
}

// OnShutdown calls shutdown on signal interrupt
func OnShutdown(shutdown func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	<-c
	if shutdown != nil {
		shutdown()
	}
}
