package exponoti

import (
	"context"

	notificationcommands "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/commands"
)

func (api *externalExpoNotiService) PushNoti(ctx context.Context, req *notificationcommands.PushNotification) error {
	return nil
}
