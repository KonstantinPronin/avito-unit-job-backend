package internal

import (
	http2 "github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/delivery/http"
	repository2 "github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/repository"
	usecase2 "github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/usecase"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user/delivery/http"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user/repository"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user/usecase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
	"go.uber.org/zap"
)

type Server struct {
	e      *echo.Echo
	db     *gorm.DB
	logger *zap.Logger
	port   string
}

func NewServer(e *echo.Echo, db *gorm.DB, logger *zap.Logger, port string) *Server {
	sanitizer := bluemonday.UGCPolicy()

	userRepository := repository.NewUser(db, logger)
	userUsecase := usecase.NewUser(userRepository, logger)
	http.NewUserHandler(userUsecase, e, logger, sanitizer)

	transactionRepository := repository2.NewTransaction(db, logger)
	transactionUsecase := usecase2.NewTransaction(10, transactionRepository, logger)
	http2.NewTransactionHandler(transactionUsecase, e, logger, sanitizer)

	return &Server{
		e:      e,
		db:     db,
		logger: logger,
		port:   port,
	}
}

func (s *Server) Start() error {
	return s.e.Start(s.port)
}
