package notificationrepository

import (
	"context"

	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

func (repo *notificationRepo) UpdateNoti(ctx context.Context, entity *notificationdomain.Notification) error {
	dto := ToNotificationDTO(entity)
	where := "id=:id"
	query := common.GenerateSQLQueries(common.UPDATE, TABLE_NOTIFICATION, UPDATE_NOTIFICATION, &where)

	// If no transaction, use db directly
	_, err := repo.db.NamedExec(query, dto)
	return err
}
