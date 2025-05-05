package notificationrpcservice

import (
	"github.com/gin-gonic/gin"
	notificationcommands "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/commands"
	notificationqueries "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/quries"
)

type notificationRPCService struct {
	cmd   notificationcommands.Commands
	query notificationqueries.Queries
	// auth  middleware.AuthClient
}

func NewNotificationRPCService(cmd notificationcommands.Commands, query notificationqueries.Queries) *notificationRPCService {
	return &notificationRPCService{
		cmd:   cmd,
		query: query,
	}
}

// func (s *accountRPCService) AddAuth(auth middleware.AuthClient) *accountRPCService {
// 	s.auth = auth
// 	return s
// }

func (s *notificationRPCService) Routes(g *gin.RouterGroup) {
	noti_route := g.Group("/notifications")
	{
		noti_route.POST(
			"/push-token",
			s.handleUpdatePushToken(),
		)
	}
}
