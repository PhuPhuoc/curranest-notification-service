package notificationqueries

import (
	"time"

	"github.com/google/uuid"

	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

type PushTokenDTO struct {
	AccountId         uuid.UUID  `json:"account-id"`
	ExponentPushToken string     `json:"exponent-push-token"`
	CreatedAt         *time.Time `json:"created-at"`
}

func toPushTokenDTO(data *notificationdomain.PushToken) *PushTokenDTO {
	dto := &PushTokenDTO{
		AccountId:         data.GetAccountID(),
		ExponentPushToken: data.GetExponentPushToken(),
		CreatedAt:         data.GetCreatedAt(),
	}
	return dto
}

func (p *PushTokenDTO) ToPushTokenEntity() (*notificationdomain.PushToken, error) {
	return notificationdomain.NewPushToken(
		p.AccountId,
		p.ExponentPushToken,
		p.CreatedAt,
	)
}

type NotificationDTO struct {
	Id        uuid.UUID  `json:"id"`
	AccountId uuid.UUID  `json:"account-id"`
	Content   string     `json:"content"`
	SubId     uuid.UUID  `json:"sub-id"`
	Route     string     `json:"route"`
	CreatedAt *time.Time `json:"created-at"`
	ReadAt    *time.Time `json:"read-at"`
}

func ToNotificationDTO(data *notificationdomain.Notification) *NotificationDTO {
	dto := &NotificationDTO{
		Id:        data.GetID(),
		AccountId: data.GetAccountID(),
		Content:   data.GetContent(),
		Route:     data.GetRoute(),
		CreatedAt: data.GetCreatedAt(),
		ReadAt:    data.GetReadAt(),
	}
	return dto
}

func (p *NotificationDTO) ToNotificationEntity() (*notificationdomain.Notification, error) {
	return notificationdomain.NewNotification(
		p.Id,
		p.AccountId,
		p.SubId,
		p.Content,
		p.Route,
		p.CreatedAt,
		p.ReadAt,
	)
}

type FilterGetNoti struct {
	AccountID uuid.UUID      `json:"account-id"`
	HadRead   *bool          `json:"had-read,omitempty"`
	Paging    *common.Paging `json:"-"`
}
