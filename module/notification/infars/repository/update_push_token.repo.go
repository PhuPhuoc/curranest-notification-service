package notificationrepository

import (
	"context"

	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

func (repo *notificationRepo) UpdatePushToken(ctx context.Context, entity *notificationdomain.PushToken) error {
	dto := ToPushTokenDTO(entity)
	where := "account_id=:account_id"
	query := common.GenerateSQLQueries(common.UPDATE, TABLE_PUSHTOKEN, UPDATE_PUSHTOKEN, &where)
	_, err := repo.db.NamedExec(query, dto)
	return err
}
