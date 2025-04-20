package main

import (
	"database/sql"
	"log"

	"github.com/Reza1878/goesclearning/user-service/config"
	productHandlers "github.com/Reza1878/goesclearning/user-service/handler/product"
	handlers "github.com/Reza1878/goesclearning/user-service/handler/user"
	"github.com/Reza1878/goesclearning/user-service/proto/product"
	repository "github.com/Reza1878/goesclearning/user-service/repository/user"
	"github.com/Reza1878/goesclearning/user-service/routes"
	productUC "github.com/Reza1878/goesclearning/user-service/usecases/product"
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
		log.Default().Printf("[ERROR] %v", err)
		return
	}
	defer db.Close()

	redis, err := config.InitRedis(cfg.Redis)
	if err != nil {
		log.Default().Printf("[ERROR] %v", err)
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

	productRPC := product.NewProductServiceClient(rpc)
	productUC := productUC.NewProductUsecase(productRPC)
	productHandler := productHandlers.NewProductUsecase(productUC)

	return &routes.Routes{
		User:    userHandler,
		Product: productHandler,
	}
}
