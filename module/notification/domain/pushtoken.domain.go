package notificationdomain

import (
	"time"

	"github.com/google/uuid"
)

type PushToken struct {
	accountId         uuid.UUID
	exponentPushToken string
	createdAt         *time.Time
}

func (p *PushToken) GetAccountID() uuid.UUID {
	return p.accountId
}

func (p *PushToken) GetExponentPushToken() string {
	return p.exponentPushToken
}

func (p *PushToken) GetCreatedAt() *time.Time {
	return p.createdAt
}

func NewPushToken(
	accountId uuid.UUID,
	exponentPushToken string,
	createdAt *time.Time,
) (*PushToken, error) {
	return &PushToken{
		accountId:         accountId,
		exponentPushToken: exponentPushToken,
		createdAt:         createdAt,
	}, nil
}
