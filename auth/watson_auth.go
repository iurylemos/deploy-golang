package auth

import (
	"github.com/IBM/go-sdk-core/v5/core"
	assistant "github.com/watson-developer-cloud/go-sdk/v2/assistantv1"
)

func GetCredentials() (*assistant.AssistantV1, *string, error) {
	authenticator := &core.IamAuthenticator{
		ApiKey: "RusZc7bq_xqRALGxKb3yEbRbVCGTHnVjKg--cnyjH7oK",
	}
	service, serviceErr := assistant.NewAssistantV1(&assistant.AssistantV1Options{
		URL:           "https://api.us-east.assistant.watson.cloud.ibm.com/instances/f172ba5f-7e24-40b0-a784-bf4deb7480fe",
		Version:       core.StringPtr("2018-07-10"),
		Authenticator: authenticator,
	})

	// Check successful instantiation
	if serviceErr != nil {
		return nil, nil, serviceErr
	}

	workspace, _, err := service.GetWorkspace(service.NewGetWorkspaceOptions("36986997-5833-49db-9436-ffbd0258ad0c"))

	if err != nil {
		return nil, nil, err
	}

	workspaceID := workspace.WorkspaceID

	return service, workspaceID, nil
}
