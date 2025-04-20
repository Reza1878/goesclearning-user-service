package usecases

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Reza1878/goesclearning/user-service/helper/fault"
	"github.com/Reza1878/goesclearning/user-service/proto/product"
)

type productUseCase struct {
	serverRPC product.ProductServiceClient
}

func NewProductUsecase(serverRPC product.ProductServiceClient) *productUseCase {
	return &productUseCase{
		serverRPC: serverRPC,
	}
}

var _ ProductUsecases = &productUseCase{}

type ProductUsecases interface {
	InsertProduct(ctx context.Context, req *product.ProductInsertRequest) (*product.ProductInsertResponse, error)
	ListProduct(ctx context.Context, req *product.ListProductRequest) (*product.ListProductResponse, error)
}

func (u *productUseCase) InsertProduct(ctx context.Context, req *product.ProductInsertRequest) (*product.ProductInsertResponse, error) {
	insertOk, err := u.serverRPC.InsertProduct(ctx, req)
	if err != nil {
		return nil, fault.Custom(
			http.StatusUnprocessableEntity,
			fault.ErrUnprocessable,
			fmt.Sprintf("failed insert product: %v", err),
		)
	}

	return insertOk, nil
}

func (u *productUseCase) ListProduct(ctx context.Context, req *product.ListProductRequest) (*product.ListProductResponse, error) {
	product, err := u.serverRPC.ListProduct(ctx, req)
	if err != nil {
		return nil, fault.Custom(
			http.StatusUnprocessableEntity,
			fault.ErrUnprocessable,
			fmt.Sprintf("failed retrieve list product: %v", err),
		)
	}

	return product, nil
}
