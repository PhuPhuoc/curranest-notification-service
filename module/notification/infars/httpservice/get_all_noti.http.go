package notihttpservice

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationqueries "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/quries"
)

// @Summary		get notification
// @Description	get notification
// @Tags			notifications
// @Accept			json
// @Produce		json
// @Param			account-id	query		string					true	"service ID (UUID)"
// @Param			had-read	query		bool					false	"apply pagination not not"
// @Param			page		query		int						false	"current page index"
// @Param			page-size	query		int						false	"number of items per page"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/notifications [get]
// @Security		ApiKeyAuth
func (s *notiHttpService) handleGetAllNoti() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filter := notificationqueries.FilterGetNoti{}

		if accId := ctx.Query("account-id"); accId != "" {
			accUUID, err := uuid.Parse(accId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("service-id invalid (not a UUID)"))
				return
			}
			filter.AccountID = accUUID
		}

		if hadRead := ctx.Query("had-read"); hadRead != "" {
			hadReadBool, err := strconv.ParseBool(hadRead)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("had-read must be a boolean value"))
				return
			}
			filter.HadRead = &hadReadBool
		}

		paging := common.Paging{}
		if page := ctx.Query("page"); page != "" {
			pageInt, err := strconv.Atoi(page)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("page must be a integer"))
				return
			}
			paging.Page = pageInt
		}

		if pageSize := ctx.Query("page-size"); pageSize != "" {
			pageSizeInt, err := strconv.Atoi(pageSize)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("page-size must be a integer"))
				return
			}
			paging.Size = pageSizeInt
		}
		filter.Paging = &paging

		dtos, err := s.query.GetAllNoti.Handle(ctx.Request.Context(), &filter)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseGetWithPagination(ctx, dtos, filter.Paging, filter)
	}
}
