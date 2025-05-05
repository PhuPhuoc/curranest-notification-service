package notificationdomain

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	id        uuid.UUID
	accountId uuid.UUID
	content   string
	route     string
	createdAt *time.Time
	readAt    *time.Time
}

func (n *Notification) GetID() uuid.UUID {
	return n.id
}

func (n *Notification) GetAccountID() uuid.UUID {
	return n.accountId
}

func (n *Notification) GetContent() string {
	return n.content
}

func (n *Notification) GetRoute() string {
	return n.route
}

func (n *Notification) GetCreatedAt() *time.Time {
	return n.createdAt
}

func (n *Notification) GetReadAt() *time.Time {
	return n.readAt
}

func NewNotification(
	id, accountId uuid.UUID,
	content, route string,
	createdAt, readAt *time.Time,
) (*Notification, error) {
	return &Notification{
		id:        id,
		accountId: accountId,
		content:   content,
		route:     route,
		createdAt: createdAt,
		readAt:    readAt,
	}, nil
}
