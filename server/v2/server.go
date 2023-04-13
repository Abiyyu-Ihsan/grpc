package serverv2

import (
	"context"
	"fmt"

	protov2 "service-fasilitas/proto/v2"
)

type fasilitasService struct {
	protov2.UnimplementedFasilitasServiceServer
}

func NewFasilitasServiceServer() protov2.FasilitasServiceServer {
	return new(fasilitasService)
}

func (fm *fasilitasService) GetMasterRegionByParamReq(ctx context.Context, msg *protov2.GetMasterRegionByParamReq) (*protov2.GetMasterRegionByParamReq, error) {
	const (
		funcName = "GetMasterRegionByParamReq"
		tagID    = 1
	)
	fmt.Println(tagID, " - ", funcName)
	return msg, nil
}
