package notificationqueries

import (
	"context"

	"github.com/google/uuid"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

type Queries struct {
	GetPushTokenByAccId *getPushTokenByAccountIdHandler

	GetAllNoti   *getNotiByFilterHandler
	FindNotiById *findNotiByIdHandler
}

type Builder interface {
	BuildNotificationQueryRepo() NotificationQueryRepo
}

func NewNotificationQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetPushTokenByAccId: NewGetPushTokenByAccountIdHandler(
			b.BuildNotificationQueryRepo(),
		),
		GetAllNoti: NewGetNotiByFilterHandler(
			b.BuildNotificationQueryRepo(),
		),
		FindNotiById: NewFindNotiByIdHandler(
			b.BuildNotificationQueryRepo(),
		),
	}
}

type NotificationQueryRepo interface {
	FindPushTokenByAccountId(ctx context.Context, accId uuid.UUID) (*notificationdomain.PushToken, error)
	GetAllNotification(ctx context.Context, filter *FilterGetNoti) ([]notificationdomain.Notification, error)
	FindById(ctx context.Context, appointmentId uuid.UUID) (*notificationdomain.Notification, error)
}
