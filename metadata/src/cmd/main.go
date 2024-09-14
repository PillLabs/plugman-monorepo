package main

import (
	"context"
	"github.com/PillLabs/plugman-monorepo/metadata/src/cmd/docs"
	_ "github.com/PillLabs/plugman-monorepo/metadata/src/common"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common/log"
	"github.com/PillLabs/plugman-monorepo/metadata/src/conf"
	_ "github.com/PillLabs/plugman-monorepo/metadata/src/conf"
	_ "github.com/PillLabs/plugman-monorepo/metadata/src/controller"
	_ "github.com/PillLabs/plugman-monorepo/metadata/src/database"
	_ "github.com/PillLabs/plugman-monorepo/metadata/src/premise"
	"github.com/PillLabs/plugman-monorepo/metadata/src/route"
	"github.com/PillLabs/plugman-monorepo/metadata/src/service/listener"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title           Plugman service
// @version         0.1.0
// @description     plugman backend service
// @contact.name    Maoma
// @host            plugman-dev.demonitors.xyz
// @BasePath        /api/v1
// @schemes         https
func main() {
	log.Info("[Main] Plugman-backend is initializing...")

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":9111", nil)
		if err != nil {
			log.Fatal(" %v", err)
		}
	}()
	log.Info("[Main] prometheus exporter listening...")

	quitChan := make(chan *struct{})
	l := listener.NewZetaListener()
	go l.Listen(quitChan)

	gin.SetMode(conf.GConfig.GetString("server.runMode"))
	docs.SwaggerInfo.BasePath = "/api/v1"
	endPoint := conf.GConfig.GetString("server.addr")
	handler := route.InitRouter()
	server := &http.Server{
		Addr:    endPoint,
		Handler: handler,
	}
	log.Info("[Main] start http server, listening %s", endPoint)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("[Main] server is not started, error is %s", err)
		}
	}()

	// capture the terminating signal to stop gracefully.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)

	<-quit
	l.GracefulStop()
	<-quitChan

	log.Info("[Main] stopping the http server...")
	httpServerContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(httpServerContext); err != nil {
		log.Fatal("[Main] server forced to shutdown:", err)
	}
	log.Info("[Main] http server stopped")

	log.Info("[Main] server shutting...")

}
