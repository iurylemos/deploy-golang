package controllers

import (
	"deploy-golang/auth"
	"deploy-golang/models"
	"log"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/gofiber/fiber/v2"
	assistant "github.com/watson-developer-cloud/go-sdk/v2/assistantv1"
)

func SendMessageController(c *fiber.Ctx) error {
	var data models.RequestWatson
	var context *assistant.Context

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	service, workspaceID, err := auth.GetCredentials()

	if err != nil {
		return err
	}

	messageOptions := service.NewMessageOptions(*workspaceID)

	if data.Context == nil {
		// Call the Message method with no specified context
		messageResult, _, responseErr := service.Message(messageOptions)

		// Check successful call
		if responseErr != nil {
			return err
		}

		// log.Println(response)

		context = messageResult.Context
	} else {
		context = data.Context
	}

	input := &assistant.MessageInput{
		Text: core.StringPtr(data.Input),
	}

	messageOptions.SetContext(context).SetInput(input)

	_, response, responseErr := service.Message(messageOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	log.Println(response)
	return c.JSON([]string{})
}
