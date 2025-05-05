package notificationrepository

import (
	"context"

	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

func (repo *notificationRepo) CreatePushToken(ctx context.Context, entity *notificationdomain.PushToken) error {
	dto := ToPushTokenDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE_PUSHTOKEN, CREAT_PUSHTOKEN, nil)
	if _, err := repo.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}
