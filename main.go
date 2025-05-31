package main

import (
	"fmt"
	"net/http"
	"os"

	accountDelivery "github.com/ariefsn/intrans/apps/account/delivery"
	accountRepository "github.com/ariefsn/intrans/apps/account/repository"
	accountService "github.com/ariefsn/intrans/apps/account/service"
	trxDelivery "github.com/ariefsn/intrans/apps/transaction/delivery"
	trxRepository "github.com/ariefsn/intrans/apps/transaction/repository"
	trxService "github.com/ariefsn/intrans/apps/transaction/service"
	"github.com/ariefsn/intrans/db"
	_ "github.com/ariefsn/intrans/docs"
	"github.com/ariefsn/intrans/entities"
	"github.com/ariefsn/intrans/logger"
	"github.com/ariefsn/intrans/validator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	logger.InitLogger()
	validator.InitValidator()
	godotenv.Load()
}

// @title Swagger Transaction API
// @version 1.0
// @description API Transaction server.
// @contact.name API Support
// @contact.url https://ariefsn.dev
// @contact.email hello@ariefsn.dev
// @host localhost:3000
// @BasePath /
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		logger.Fatal(fmt.Errorf("unable to start application. error: missing DB_DSN"))
	}

	dbConn := db.InitDB(dsn)
	err := dbConn.Ping()
	if err != nil {
		logger.Fatal(fmt.Errorf("unable to connect to database. error: %s", err.Error()))
	}

	logger.Info("database connected, checking migrations...")

	_, _, err = db.MigrateDB(dbConn)
	if err != nil {
		logger.Fatal(fmt.Errorf("migrate database failed. error: %v", err))
	}

	logger.Info("setup database completed ✅")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")

			next.ServeHTTP(w, r)
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res := entities.Response{
			Status:  true,
			Message: "running",
		}
		res.Send(w)
	})

	// Swagger
	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/index.html", http.StatusMovedPermanently)
	})
	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/docs/doc.json"),
	))

	accountRepo := accountRepository.New(dbConn)
	accountSvc := accountService.New(accountRepo)
	accountHandler := accountDelivery.New(accountSvc)

	trxRepo := trxRepository.New(dbConn)
	trxSvc := trxService.New(trxRepo)
	trxHandler := trxDelivery.New(trxSvc)

	r.Mount("/accounts", accountHandler)
	r.Mount("/transactions", trxHandler)

	address := ":" + port

	logger.Info("app started at " + address)

	http.ListenAndServe(address, r)
}
