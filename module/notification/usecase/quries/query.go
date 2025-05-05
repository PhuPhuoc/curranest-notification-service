package notificationqueries

import (
	"context"

	"github.com/google/uuid"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

type Queries struct {
	GetPushTokenByAccId *getPushTokenByAccountIdHandler
}

type Builder interface {
	BuildNotificationQueryRepo() NotificationQueryRepo
}

func NewNotificationQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetPushTokenByAccId: NewGetPushTokenByAccountIdHandler(
			b.BuildNotificationQueryRepo(),
		),
	}
}

type NotificationQueryRepo interface {
	FindPushTokenByAccountId(ctx context.Context, accId uuid.UUID) (*notificationdomain.PushToken, error)
}
