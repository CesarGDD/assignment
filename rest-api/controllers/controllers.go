package controllers

import (
	"cesargdd/rest-api-assigment/rgpb"
	"cesargdd/rest-api-assigment/svc"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetRegistrations(ctx *fiber.Ctx) error {
	fmt.Println("Get Registrations rest")
	res, err := svc.RegSrv().ListRegistration(context.Background(), &rgpb.ListRegistrationRequest{})
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "Registration not found",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(res.Registration)
}

func CreateRegistrations(ctx *fiber.Ctx) error {
	fmt.Println("Create Registrations rest")
	body := new(rgpb.Registration)
	err := ctx.BodyParser(body)
	if err != nil {
		return err
	}
	res, err := svc.RegSrv().CreateRegistration(context.Background(), &rgpb.CreateRegistrationRequest{
		Registration: &rgpb.Registration{
			NumberPlate: body.NumberPlate,
			Vehicle: &rgpb.Vehicle{
				Type:   body.Vehicle.Type,
				Make:   body.Vehicle.Make,
				Model:  body.Vehicle.Model,
				Colour: body.Vehicle.Colour,
			},
			Insurer: &rgpb.Insurer{
				Name: body.Insurer.Name,
				Code: body.Insurer.Code,
			},
		},
	})
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "Error creating registration",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(res.Registration)
}
