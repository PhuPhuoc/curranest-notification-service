package notificationcommands

import "github.com/google/uuid"

type RequestPushTokenDTO struct {
	AccountId uuid.UUID `json:"account-id"`
	PushToken string    `json:"push-token"`
}
