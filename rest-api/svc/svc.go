package svc

import (
	"cesargdd/rest-api-assigment/rgpb"
	"log"

	"google.golang.org/grpc"
)

func RegSrv() rgpb.RegistrationServiceClient {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("service:50050", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	svc := rgpb.NewRegistrationServiceClient(cc)

	return svc
}
