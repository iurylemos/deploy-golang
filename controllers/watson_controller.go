package controllers

import (
	"deploy-golang/auth"
	"deploy-golang/models"
	"deploy-golang/services"
	"encoding/json"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/gofiber/fiber/v2"
	assistant "github.com/watson-developer-cloud/go-sdk/v2/assistantv1"
)

func SendMessageController(c *fiber.Ctx) error {
	var data models.RequestWatson
	// var context *assistant.Context

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	service, workspaceID, err := auth.GetCredentials()

	if err != nil {
		return err
	}

	messageOptions := service.NewMessageOptions(*workspaceID)

	contx, err := services.GetContext(service, data.Context, messageOptions)

	if err != nil {
		return err
	}

	input := &assistant.MessageInput{
		Text: core.StringPtr(data.Input),
	}

	messageOptions.SetContext(contx).SetInput(input)

	_, response, responseErr := service.Message(messageOptions)

	// Check successful call
	if responseErr != nil {
		return responseErr
	}

	byteData, _ := json.Marshal(response.Result)
	var responseWatson models.ResponseWatson

	if err := json.Unmarshal(byteData, &responseWatson); err != nil {
		return err
	}

	return c.JSON(responseWatson)
}
