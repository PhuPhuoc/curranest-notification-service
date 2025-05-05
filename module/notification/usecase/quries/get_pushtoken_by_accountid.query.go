package notificationqueries

import (
	"context"

	"github.com/google/uuid"
	"githuh.com/PhuPhuoc/curanest-notification-service/common"
)

type getPushTokenByAccountIdHandler struct {
	queryRepo NotificationQueryRepo
}

func NewGetPushTokenByAccountIdHandler(queryRepo NotificationQueryRepo) *getPushTokenByAccountIdHandler {
	return &getPushTokenByAccountIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *getPushTokenByAccountIdHandler) Handle(ctx context.Context, accId uuid.UUID) (*PushTokenDTO, error) {
	entity, err := h.queryRepo.FindPushTokenByAccountId(ctx, accId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get token from db").
			WithInner(err.Error())
	}

	if entity == nil {
		return nil, nil
	}

	return toPushTokenDTO(entity), nil
}
