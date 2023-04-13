package main

import (
	"log"

	protov2 "service-fasilitas/proto/v2"
	serverv2 "service-fasilitas/server/v2"
	"service-fasilitas/util/run"
)

func main() {
	address := "localhost"
	port := "8080"

	var server protov2.FasilitasServiceServer
	{
		// todo create repo (redis, micro, db, etc)

		// todo create brockers

		// include middleware as server
		server = serverv2.NewFasilitasServiceServer()
	}

	grpcServer := run.CreateGRPCServer()
	protov2.RegisterFasilitasServiceServer(grpcServer, server)

	go run.ServeHTTP(address, port, protov2.RegisterFasilitasServiceHandlerFromEndpoint, run.DefaultHTTPOption(), run.DefaultHTTPHandler)

	log.Println("server started", address, port)
	run.OnShutdown(func() {
		grpcServer.GracefulStop()
	})
}
