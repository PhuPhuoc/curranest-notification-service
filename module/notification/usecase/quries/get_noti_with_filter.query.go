package notificationqueries

import (
	"context"

	"githuh.com/PhuPhuoc/curanest-notification-service/common"
)

type getNotiByFilterHandler struct {
	queryRepo NotificationQueryRepo
}

func NewGetNotiByFilterHandler(queryRepo NotificationQueryRepo) *getNotiByFilterHandler {
	return &getNotiByFilterHandler{
		queryRepo: queryRepo,
	}
}

func (h *getNotiByFilterHandler) Handle(ctx context.Context, filter *FilterGetNoti) ([]NotificationDTO, error) {
	entities, err := h.queryRepo.GetAllNotification(ctx, filter)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get token from db").
			WithInner(err.Error())
	}

	if len(entities) == 0 {
		return []NotificationDTO{}, nil
	}

	dtos := make([]NotificationDTO, len(entities))
	for i, entity := range entities {
		dtos[i] = *ToNotificationDTO(&entity)
	}

	return dtos, nil
}
