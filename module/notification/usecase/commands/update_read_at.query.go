package notificationcommands

import (
	"context"
	"time"

	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

type updateNotiReadAtHandler struct {
	cmdRepo NoticationCommandRepo
}

func NewUpdateNotiReadAtHandler(cmdRepo NoticationCommandRepo) *updateNotiReadAtHandler {
	return &updateNotiReadAtHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *updateNotiReadAtHandler) Handle(ctx context.Context, entity *notificationdomain.Notification) error {
	readAtNow := time.Now()
	newEntity, _ := notificationdomain.NewNotification(
		entity.GetID(),
		entity.GetAccountID(),
		entity.GetSubID(),
		entity.GetContent(),
		entity.GetRoute(),
		entity.GetCreatedAt(),
		&readAtNow,
	)

	if err := h.cmdRepo.UpdateNoti(ctx, newEntity); err != nil {
		return common.NewInternalServerError().
			WithReason("error when update notification").
			WithInner(err.Error())
	}
	return nil
}
