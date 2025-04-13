package main

import (
	"database/sql"

	"github.com/Reza1878/goesclearning/user-service/config"
	handlers "github.com/Reza1878/goesclearning/user-service/handler/user"
	repository "github.com/Reza1878/goesclearning/user-service/repository/user"
	"github.com/Reza1878/goesclearning/user-service/routes"
	usecases "github.com/Reza1878/goesclearning/user-service/usecases/user"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		return
	}

	db, err := config.InitPostgreSQL(cfg.Postgres)
	if err != nil {
		return
	}
	defer db.Close()

	redis, err := config.InitRedis(cfg.Redis)
	if err != nil {
		return
	}

	rpc, err := config.RPCDial(cfg.Grpc)
	if err != nil {
		return
	}

	routes := initDepedencies(db, rpc, redis)
	routes.SetupRoutes()
	routes.Run(cfg.Port)
}

func initDepedencies(db *sql.DB, rpc *grpc.ClientConn, redis *redis.Client) *routes.Routes {
	userRepo := repository.NewStore(db)
	userUC := usecases.NewUserUsecase(userRepo, redis)
	userHandler := handlers.NewHandler(userUC)

	return &routes.Routes{
		User: userHandler,
	}
}
