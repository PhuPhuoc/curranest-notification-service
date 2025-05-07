package notihttpservice

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"githuh.com/PhuPhuoc/curanest-notification-service/common"
)

//	@Summary		update notification read at (user seen message)
//	@Description	update notification read at (user seen message)
//	@Tags			notifications
//	@Accept			json
//	@Produce		json
//	@Param			notification-id	path		string					true	"notification ID (UUID)"
//	@Success		200				{object}	map[string]interface{}	"data"
//	@Failure		400				{object}	error					"Bad request error"
//	@Router			/api/v1/notifications/{notification-id} [patch]
//	@Security		ApiKeyAuth
func (s *notiHttpService) handleUpdateReadAt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var notiUUID uuid.UUID
		var err error
		if notiId := ctx.Param("notification-id"); notiId != "" {
			notiUUID, err = uuid.Parse(notiId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("service-id invalid (not a UUID)"))
				return
			}
		}

		notiDTO, err := s.query.FindNotiById.Handle(ctx.Request.Context(), notiUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		notiEntity, _ := notiDTO.ToNotificationEntity()

		if err := s.cmd.UpdateReadAt.Handle(ctx.Request.Context(), notiEntity); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}
