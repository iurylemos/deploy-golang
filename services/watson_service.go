package services

import (
	assistant "github.com/watson-developer-cloud/go-sdk/v2/assistantv1"
)

func GetContext(service *assistant.AssistantV1, contexto *assistant.Context, options *assistant.MessageOptions) (ctx *assistant.Context, err error) {
	if contexto != nil {
		return contexto, nil
	}

	result, _, err := service.Message(options)

	// Check successful call
	if err != nil {
		return nil, err
	}

	context := result.Context

	return context, nil
}
