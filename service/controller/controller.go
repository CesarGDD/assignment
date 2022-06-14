package controller

import (
	"cesargdd/assigment/postgres"
	"cesargdd/assigment/rgpb"
	"context"
	"encoding/json"
	"fmt"
)

var conn = postgres.Connect()
var db = postgres.New(conn)

type Server struct {
	rgpb.RegistrationServiceServer
}

func (*Server) CreateRegistration(ctx context.Context, req *rgpb.CreateRegistrationRequest) (*rgpb.CreateRegistrationResponse, error) {
	fmt.Println("Create Registration gRPC")

	vehicle, err := json.Marshal(&rgpb.Vehicle{
		Type:   req.Registration.Vehicle.Type,
		Make:   req.Registration.Vehicle.Make,
		Model:  req.Registration.Vehicle.Model,
		Colour: req.Registration.Vehicle.Colour,
	})
	if err != nil {
		return nil, err
	}

	insurer, _ := json.Marshal(&rgpb.Insurer{
		Name: req.Registration.Insurer.Name,
		Code: req.Registration.Insurer.Code,
	})

	res, err := db.CreateRegistration(ctx, postgres.CreateRegistrationParams{
		NumberPlate: req.Registration.NumberPlate,
		Vehicle:     vehicle,
		Insurer:     insurer,
	})
	if err != nil {
		return nil, err
	}
	var vehicle2 rgpb.Vehicle
	err = json.Unmarshal(res.Vehicle, &vehicle2)
	if err != nil {
		return nil, err
	}
	var insurer2 rgpb.Insurer
	err = json.Unmarshal(res.Insurer, &insurer2)
	if err != nil {
		return nil, err
	}
	return &rgpb.CreateRegistrationResponse{
		Registration: &rgpb.Registration{
			NumberPlate: res.NumberPlate,
			Vehicle:     &vehicle2,
			Insurer:     &insurer2,
		},
	}, nil
}

func (*Server) ListRegistration(ctx context.Context, req *rgpb.ListRegistrationRequest) (*rgpb.ListRegistrationResponse, error) {
	fmt.Println("List Registrations gRPC")
	res, err := db.ListRegistration(ctx)
	if err != nil {
		return nil, err
	}

	var reg []*rgpb.Registration
	for _, data := range res {
		var insurer2 rgpb.Insurer
		var vehicle2 rgpb.Vehicle
		var registration rgpb.Registration
		_ = json.Unmarshal(data.Insurer, &insurer2)
		_ = json.Unmarshal(data.Vehicle, &vehicle2)
		registration.NumberPlate = data.NumberPlate
		registration.Vehicle = &vehicle2
		registration.Insurer = &insurer2

		reg = append(reg, &registration)

	}
	return &rgpb.ListRegistrationResponse{
		Registration: reg,
	}, nil
}
