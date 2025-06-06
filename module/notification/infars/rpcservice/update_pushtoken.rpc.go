package notificationrpcservice

import (
	"github.com/gin-gonic/gin"
	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
	notificationcommands "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/commands"
)

// @Summary		create/update push-token
// @Description	create/update push-token
// @Tags			rpc: notifications
// @Accept			json
// @Produce		json
// @Param			update	form		body					notificationcommands.RequestPushTokenDTO	true	"push-token data to create/update"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/external/rpc/notifications/push-token [post]
// @Security		ApiKeyAuth
func (s *notificationRPCService) handleUpdatePushToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto notificationcommands.RequestPushTokenDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		pushTokenDTO, err := s.query.GetPushTokenByAccId.Handle(ctx, dto.AccountId)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		var pushTokenEntity *notificationdomain.PushToken
		if pushTokenDTO != nil {
			pushTokenEntity, _ = pushTokenDTO.ToPushTokenEntity()
		}

		err = s.cmd.UpdateAccountPushToken.Handle(ctx.Request.Context(), &dto, pushTokenEntity)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}
