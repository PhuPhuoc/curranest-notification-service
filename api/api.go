package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"githuh.com/PhuPhuoc/curanest-notification-service/builder"
	"githuh.com/PhuPhuoc/curanest-notification-service/config"
	"githuh.com/PhuPhuoc/curanest-notification-service/docs"
	"githuh.com/PhuPhuoc/curanest-notification-service/middleware"
	notificationrpcservice "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/infars/rpcservice"
	notificationcommands "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/commands"
	notificationqueries "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/quries"
)

type server struct {
	port string
	db   *sqlx.DB
}

func InitServer(port string, db *sqlx.DB) *server {
	return &server{
		port: port,
		db:   db,
	}
}

const (
	env_local     = "local"
	env_vps       = "vps"
	url_acc_local = "http://localhost:8001"
	url_acc_prod  = "http://auth_service:8080"

	url_nursing_local = "http://localhost:8003"
	url_nursing_prod  = "http://nurse_service:8080"

	url_appointment_local = "http://localhost:8004"
	url_appointment_prod  = "http://appointment_service:8080"

	url_expo_noti = "https://exp.host/--/api/v2/push/send"
)

// @Summary		ping server
// @Description	ping server
// @Tags			ping
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]any	"message success"
// @Failure		400	{object}	error			"Bad request error"
// @Router			/ping [get]
func (sv *server) RunApp() error {
	envDevlopment := config.AppConfig.EnvDev
	if envDevlopment == env_local {
		// gin.SetMode(gin.ReleaseMode)
		docs.SwaggerInfo.BasePath = "/"
	}

	if envDevlopment == env_vps {
		gin.SetMode(gin.ReleaseMode)
		docs.SwaggerInfo.BasePath = "/notification"
	}

	router := gin.New()

	configcors := cors.DefaultConfig()
	configcors.AllowAllOrigins = true
	configcors.AllowMethods = []string{"POST", "GET", "PUT", "DELETE", "PATCH", "OPTIONS"}
	configcors.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	configcors.ExposeHeaders = []string{"Content-Length"}
	configcors.AllowCredentials = true
	configcors.MaxAge = 12 * time.Hour

	router.Use(cors.New(configcors))
	router.Use(middleware.SkipSwaggerLog(), gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "curanest-notification-service - pong"}) })

	// authClient := common.NewJWTx(config.AppConfig.Key)
	noti_cmd_builder := notificationcommands.NewNotificationCmdWithBuilder(
		builder.NewNotificationBuilder(sv.db).AddUrlExpoNotiUrl(url_expo_noti),
	)
	noti_query_builder := notificationqueries.NewNotificationQueryWithBuilder(
		builder.NewNotificationBuilder(sv.db),
	)

	// api := router.Group("/api/v1")
	// {
	// }

	rpc := router.Group("/external/rpc")
	{
		notificationrpcservice.NewNotificationRPCService(
			noti_cmd_builder,
			noti_query_builder,
		).Routes(rpc)
	}

	log.Println("server start listening at port: ", sv.port)
	return router.Run(sv.port)
}
