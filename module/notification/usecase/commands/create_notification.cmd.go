package notificationcommands

import (
	"context"
	"log"

	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
)

type createNotificationHandler struct {
	cmdRepo         NoticationCommandRepo
	expoNotiFetcher ExpoNotiFetcher
}

func NewCreateNotificationHandler(cmdRepo NoticationCommandRepo, expoNotiFetcher ExpoNotiFetcher) *createNotificationHandler {
	return &createNotificationHandler{
		cmdRepo:         cmdRepo,
		expoNotiFetcher: expoNotiFetcher,
	}
}

func (h *createNotificationHandler) Handle(ctx context.Context, ptoken *notificationdomain.PushToken, req *RequestCreateNotification) error {
	notiId := common.GenUUID()
	newNoti, _ := notificationdomain.NewNotification(
		notiId,
		req.AccountId,
		req.SubId,
		req.Content,
		req.Route,
		nil,
		nil,
	)

	if err := h.cmdRepo.CreateNotification(ctx, newNoti); err != nil {
		return common.NewInternalServerError().
			WithReason("error when create notification").
			WithInner(err.Error())
	}

	if ptoken != nil {
		notiReq := PushNotification{
			To:    ptoken.GetExponentPushToken(),
			Sound: "default",
			Title: "CuraNest",
			Body:  newNoti.GetContent(),
			Data: NotificationData{
				SubId:  newNoti.GetSubID(),
				Screen: newNoti.GetRoute(),
				Title:  "CuraNest notification",
			},
			Android: AndroidConfig{
				ChannelID: "curanest_channel",
			},
			IOS: IOSConfig{
				Badge: 1,
			},
		}

		err := h.expoNotiFetcher.PushNoti(ctx, &notiReq)
		if err != nil {
			log.Println("error when push notification by expo: ", err)
		}
	}

	return nil
}
