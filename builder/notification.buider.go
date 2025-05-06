package builder

import (
	"github.com/jmoiron/sqlx"
	"githuh.com/PhuPhuoc/curanest-notification-service/module/notification/infars/externalapi/exponoti"
	notificationrepository "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/infars/repository"
	notificationcommands "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/commands"
	notificationqueries "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/quries"
)

type builderOfNotification struct {
	db              *sqlx.DB
	urlPathExpoNoti string
}

func (s builderOfNotification) AddUrlExpoNotiUrl(url string) builderOfNotification {
	s.urlPathExpoNoti = url
	return s
}

func NewNotificationBuilder(db *sqlx.DB) builderOfNotification {
	return builderOfNotification{db: db}
}

func (s builderOfNotification) BuildNotificationCmdRepo() notificationcommands.NoticationCommandRepo {
	return notificationrepository.NewNotificationRepo(s.db)
}

func (s builderOfNotification) BuildNotificationQueryRepo() notificationqueries.NotificationQueryRepo {
	return notificationrepository.NewNotificationRepo(s.db)
}

func (s builderOfNotification) BuildExpoFetcher() notificationcommands.ExpoNotiFetcher {
	return exponoti.NewExternalExpoNotiRPC(s.urlPathExpoNoti)
}
