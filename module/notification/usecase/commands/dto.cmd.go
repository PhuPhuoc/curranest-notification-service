package notificationcommands

import "github.com/google/uuid"

type RequestPushTokenDTO struct {
	AccountId uuid.UUID `json:"account-id"`
	PushToken string    `json:"push-token"`
}

type RequestCreateNotification struct {
	AccountId uuid.UUID `json:"account-id"`
	SubId     uuid.UUID `json:"sub-id"`
	Content   string    `json:"content"`
	Route     string    `json:"route"`
}

type PushNotification struct {
	To      string           `json:"to"`
	Sound   string           `json:"sound"`
	Title   string           `json:"title"`
	Body    string           `json:"body"`
	Data    NotificationData `json:"data"`
	Android AndroidConfig    `json:"android"`
	IOS     IOSConfig        `json:"ios"`
}

type NotificationData struct {
	SubId  uuid.UUID `json:"sub-id"`
	Screen string    `json:"screen"`
	Title  string    `json:"title"`
}

type AndroidConfig struct {
	ChannelID string `json:"channelId"`
}

type IOSConfig struct {
	Badge int `json:"badge"`
}

type Data struct {
	Status string `json:"status"`
	ID     string `json:"id"`
}

type ExpoResponse struct {
	Data Data `json:"data"`
}
