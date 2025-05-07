package notificationrepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

func (repo *notificationRepo) FindPushTokenByAccountId(ctx context.Context, accId uuid.UUID) (*notificationdomain.PushToken, error) {
	var accdto PushTokenDTO
	where := "account_id=?"
	query := common.GenerateSQLQueries(common.FIND_WITH_CREATED_AT, TABLE_PUSHTOKEN, GET_PUSHTOKEN, &where)
	if err := repo.db.Get(&accdto, query, accId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return accdto.ToPushTokenEntity()
}
