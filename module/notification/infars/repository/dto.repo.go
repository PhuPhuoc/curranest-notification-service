package notificationrepository

import (
	"time"

	"github.com/google/uuid"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

var (
	TABLE_NOTIFICATION = `notifications`

	CREAT_NOTIFICATION = []string{
		"id",
		"account_id",
		"content",
		"route",
	}

	GET_NOTIFICATION = []string{
		"id",
		"account_id",
		"content",
		"route",
		"created_at",
		"read_at",
	}

	UPDATE_NOTIFICATION = []string{
		"read_at",
	}

	TABLE_PUSHTOKEN = `push_tokens`

	CREAT_PUSHTOKEN = []string{
		"account_id",
		"exponent_push_token",
	}

	GET_PUSHTOKEN = []string{
		"account_id",
		"exponent_push_token",
		"created_at",
	}

	UPDATE_PUSHTOKEN = []string{
		"exponent_push_token",
	}
)

type NotificationDTO struct {
	Id        uuid.UUID  `db:"id"`
	AccountId uuid.UUID  `db:"account_id"`
	Content   string     `db:"content"`
	Route     string     `db:"route"`
	CreatedAt *time.Time `db:"created_at"`
	ReadAt    *time.Time `db:"read_at"`
}

func (dto *NotificationDTO) ToNotificationEntity() (*notificationdomain.Notification, error) {
	return notificationdomain.NewNotification(
		dto.Id,
		dto.AccountId,
		dto.Content,
		dto.Route,
		dto.CreatedAt,
		dto.ReadAt,
	)
}

func ToNotificationDTO(data *notificationdomain.Notification) *NotificationDTO {
	return &NotificationDTO{
		Id:        data.GetID(),
		AccountId: data.GetAccountID(),
		Content:   data.GetContent(),
		Route:     data.GetRoute(),
		CreatedAt: data.GetCreatedAt(),
		ReadAt:    data.GetReadAt(),
	}
}

type PushTokenDTO struct {
	AccountId         uuid.UUID  `db:"account_id"`
	ExponentPushToken string     `db:"exponent_push_token"`
	CreatedAt         *time.Time `db:"created_at"`
}

func (dto *PushTokenDTO) ToPushTokenEntity() (*notificationdomain.PushToken, error) {
	return notificationdomain.NewPushToken(
		dto.AccountId,
		dto.ExponentPushToken,
		dto.CreatedAt,
	)
}

func ToPushTokenDTO(data *notificationdomain.PushToken) *PushTokenDTO {
	return &PushTokenDTO{
		AccountId:         data.GetAccountID(),
		ExponentPushToken: data.GetExponentPushToken(),
		CreatedAt:         data.GetCreatedAt(),
	}
}
