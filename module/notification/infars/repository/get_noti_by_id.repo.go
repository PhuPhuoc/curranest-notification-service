package notificationrepository

import (
	"context"

	"github.com/google/uuid"
	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

func (repo *notificationRepo) FindById(ctx context.Context, appointmentId uuid.UUID) (*notificationdomain.Notification, error) {
	var dto NotificationDTO
	where := "id=?"
	query := common.GenerateSQLQueries(common.FIND_WITH_CREATED_AT, TABLE_NOTIFICATION, GET_NOTIFICATION, &where)
	if err := repo.db.Get(&dto, query, appointmentId); err != nil {
		return nil, err
	}
	return dto.ToNotificationEntity()
}
