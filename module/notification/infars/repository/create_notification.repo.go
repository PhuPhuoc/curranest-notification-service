package notificationrepository

import (
	"context"

	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

func (repo *notificationRepo) CreateNotification(ctx context.Context, entity *notificationdomain.Notification) error {
	dto := ToNotificationDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE_NOTIFICATION, CREAT_NOTIFICATION, nil)
	if _, err := repo.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}
