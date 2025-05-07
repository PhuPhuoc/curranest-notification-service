package notihttpservice

import (
	"github.com/gin-gonic/gin"
	notificationcommands "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/commands"
	notificationqueries "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/quries"
)

type notiHttpService struct {
	cmd   notificationcommands.Commands
	query notificationqueries.Queries
}

func NewNotiHTTPService(cmd notificationcommands.Commands, query notificationqueries.Queries) *notiHttpService {
	return &notiHttpService{
		cmd:   cmd,
		query: query,
	}
}

func (s *notiHttpService) Routes(g *gin.RouterGroup) {
	noti_route := g.Group("/notifications")
	{
		noti_route.GET(
			"",
			// middleware.RequireAuth(s.auth),
			s.handleGetAllNoti(),
		)
		noti_route.PATCH(
			"/:notification-id",
			s.handleUpdateReadAt(),
		)
	}
}
