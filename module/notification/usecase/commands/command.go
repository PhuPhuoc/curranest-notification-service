package notificationcommands

import (
	"context"

	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

type Commands struct {
	UpdateAccountPushToken *updateAccountPushTokenHandler
	CreateNoti             *createNotificationHandler
}

type Builder interface {
	BuildNotificationCmdRepo() NoticationCommandRepo
	BuildExpoFetcher() ExpoNotiFetcher
}

func NewNotificationCmdWithBuilder(b Builder) Commands {
	return Commands{
		UpdateAccountPushToken: NewUpdateAccountPushTokenHandler(
			b.BuildNotificationCmdRepo(),
		),
		CreateNoti: NewCreateNotificationHandler(
			b.BuildNotificationCmdRepo(),
			b.BuildExpoFetcher(),
		),
	}
}

type NoticationCommandRepo interface {
	CreatePushToken(ctx context.Context, entity *notificationdomain.PushToken) error
	UpdatePushToken(ctx context.Context, entity *notificationdomain.PushToken) error
	CreateNotification(ctx context.Context, entity *notificationdomain.Notification) error
}

type ExpoNotiFetcher interface {
	PushNoti(ctx context.Context, req *PushNotification) error
}
