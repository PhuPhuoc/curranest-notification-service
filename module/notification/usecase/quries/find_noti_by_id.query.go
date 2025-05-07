package notificationqueries

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"githuh.com/PhuPhuoc/curanest-notification-service/common"
)

type findNotiByIdHandler struct {
	queryRepo NotificationQueryRepo
}

func NewFindNotiByIdHandler(queryRepo NotificationQueryRepo) *findNotiByIdHandler {
	return &findNotiByIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *findNotiByIdHandler) Handle(ctx context.Context, notiId uuid.UUID) (*NotificationDTO, error) {
	entity, err := h.queryRepo.FindById(ctx, notiId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason(fmt.Sprintf("cannot found notification with id: %v", notiId.String())).
			WithInner(err.Error())
	}

	return ToNotificationDTO(entity), nil
}
