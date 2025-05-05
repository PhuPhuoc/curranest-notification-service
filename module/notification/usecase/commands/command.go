package notificationcommands

import (
	"context"

	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

type Commands struct {
	UpdateAccountPushToken *updateAccountPushTokenHandler
}

type Builder interface {
	BuildNotificationCmdRepo() NoticationCommandRepo
}

func NewNotificationCmdWithBuilder(b Builder) Commands {
	return Commands{
		UpdateAccountPushToken: NewUpdateAccountPushTokenHandler(
			b.BuildNotificationCmdRepo(),
		),
	}
}

type NoticationCommandRepo interface {
	CreatePushToken(ctx context.Context, entity *notificationdomain.PushToken) error
	UpdatePushToken(ctx context.Context, entity *notificationdomain.PushToken) error
	CreateNotification(ctx context.Context, entity *notificationdomain.Notification) error
}
