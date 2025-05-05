package notificationcommands

import (
	"context"

	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

type updateAccountPushTokenHandler struct {
	cmdRepo NoticationCommandRepo
}

func NewUpdateAccountPushTokenHandler(cmdRepo NoticationCommandRepo) *updateAccountPushTokenHandler {
	return &updateAccountPushTokenHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *updateAccountPushTokenHandler) Handle(ctx context.Context, dto *RequestPushTokenDTO, entity *notificationdomain.PushToken) error {
	if dto == nil {
		return common.NewBadRequestError().
			WithReason("cannot get creation information")
	}
	if entity == nil {
		newEntity, _ := notificationdomain.NewPushToken(
			dto.AccountId,
			dto.PushToken,
			nil,
		)

		if err := h.cmdRepo.CreatePushToken(ctx, newEntity); err != nil {
			return common.NewInternalServerError().
				WithReason("error when create push token").
				WithInner(err.Error())
		}

	} else {
		updateEntity, _ := notificationdomain.NewPushToken(
			dto.AccountId,
			dto.PushToken,
			nil,
		)

		if err := h.cmdRepo.UpdatePushToken(ctx, updateEntity); err != nil {
			return common.NewInternalServerError().
				WithReason("error when update push token").
				WithInner(err.Error())
		}
	}

	return nil
}
