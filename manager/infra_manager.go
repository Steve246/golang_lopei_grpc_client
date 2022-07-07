package manager

import (
	"fmt"
	"golang_lopei_grc_client/config"
	"golang_lopei_grc_client/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InfraManager interface {
	LopeiClientConn() service.LopeiPaymentClient
}



type infraManager struct {
	lopeiClient service.LopeiPaymentClient
	cfg config.Config
}

func (i *infraManager) LopeiClientConn() service.LopeiPaymentClient {
	return i.lopeiClient
}

func (i *infraManager) initGrpcClient() {
	dial, err := grpc.Dial(i.cfg.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("did not connect...." , err)
	}

	client := service.NewLopeiPaymentClient(dial)
	i.lopeiClient = client
	fmt.Println("GRPC client connected....")
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config, 
	}
	infra.initGrpcClient()
	return &infra
}