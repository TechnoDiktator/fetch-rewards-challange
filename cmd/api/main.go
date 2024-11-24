package main

import (
	"context"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/handlers"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/inmemorydb"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/middlewares"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/services"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/utils/constants"
	"github.com/TechnoDiktator/fetch-rewards-challange/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"     // Import swagger files
	ginSwagger "github.com/swaggo/gin-swagger" // Import gin-swagger middleware
	"golang.org/x/net/http2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	// Import the generated docs from the "cmd/api/docs" folder
	_ "github.com/TechnoDiktator/fetch-rewards-challange/cmd/api/docs"
)

// @title           Fetch Rewards API
// @version         1.0
// @description     This is a simple API for processing receipts and calculating rewards.
// @termsOfService  http://swagger.io/terms/

// @contact.name    Tarang Rastogi
// @contact.url     https://technodiktator.github.io/portfolio/
// @contact.email   rastogitarang5@gmail.com

// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT

// @host            localhost:4040
// @BasePath        /
// @schemes         http
func main() {

	// Set up signal handling to gracefully shut down
	logger.InitializeLogger()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Initialize Gin router
	r := gin.Default()

	// Apply the logging middleware globally
	r.Use(middleware.LogRequest)

	// Initialize the in-memory store for receipt data
	store := inmemorydb.NewMemoryStore()

	// Initialize the ReceiptService with the store
	receiptService := services.NewReceiptServiceImpl(store)

	// Initialize the ReceiptHandler with the service
	receiptHandler := handlers.NewReceiptHandler(receiptService)

	// Define routes

	// @Summary Process a receipt
	// @Description Process a receipt to calculate rewards
	// @Accept json
	// @Produce json
	// @Param receipt body Receipt true "Receipt Data"
	// @Success 200 {object} ReceiptResponse "Points calculated"
	// @Failure 400 {object} ErrorResponse "Invalid data"
	// @Router /receipts [post]
	r.POST(constants.ProcessReceipts, receiptHandler.ProcessReceipt)

	// @Summary Get points for a user
	// @Description Get the total points a user has earned
	// @Accept json
	// @Produce json
	// @Param user_id path string true "User ID"
	// @Success 200 {object} PointsResponse "Total points"
	// @Failure 400 {object} ErrorResponse "Invalid user ID"
	// @Router /points/{user_id} [get]
	r.GET(constants.GetPoints, receiptHandler.GetPoints)

	// Start the Gin server on port 8080
	startServer(r)

	// Setup Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Block until a shutdown signal is received
	<-stop
	logrus.Infof("Shutting down...")
}

// to start service
func startServer(router *gin.Engine) {
	/*
		DEFINITION : // Start the HTTP server to listen for incoming API requests on the registered routes.//
	*/
	server := &http.Server{
		Addr:         constants.PORT,
		Handler:      router,
		WriteTimeout: constants.TIMEOUT, // Use this constant constants.TIMEOUT
	}
	err := http2.ConfigureServer(server, nil)
	if err != nil {
		logrus.Errorf("Error while configuring http", err)
		return
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logrus.Error("Server Closed", err)
			return
		}
	}()
	logrus.Infof("Server listening at %s", constants.PORT)
	Gracefullstop(server)

}

// to stop service gracefully
func Gracefullstop(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.Error(err)
	}

}
