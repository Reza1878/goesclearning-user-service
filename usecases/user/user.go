package usecases

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Reza1878/goesclearning/user-service/helper/cache"
	"github.com/Reza1878/goesclearning/user-service/helper/fault"
	"github.com/Reza1878/goesclearning/user-service/helper/jwt"
	"github.com/Reza1878/goesclearning/user-service/middlewares"
	"github.com/Reza1878/goesclearning/user-service/model"
	repository "github.com/Reza1878/goesclearning/user-service/repository/user"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type userUsecase struct {
	user  repository.UserRepository
	redis *redis.Client
}

func NewUserUsecase(repository repository.UserRepository, redis *redis.Client) *userUsecase {
	return &userUsecase{
		user:  repository,
		redis: redis,
	}
}

type UserUsecases interface {
	UserRegister(body model.RegisterUser) (*model.LoginResponse, error)
}

func (u *userUsecase) UserRegister(body model.RegisterUser) (*model.LoginResponse, error) {
	ctx := context.TODO()
	cacheKey := fmt.Sprintf("login:%s", body.Email)

	if cacheExists, err := cache.Exist(ctx, u.redis, cacheKey); err != nil {
		return nil, fault.Custom(
			http.StatusInternalServerError,
			fault.ErrInternalServer,
			fmt.Sprintf("failed to check Redis key existence for '%s': %v", cacheKey, err),
		)
	} else if cacheExists {
		tokenValue, err := cache.Get(ctx, u.redis, cacheKey)
		if err != nil {
			return nil, fault.Custom(
				http.StatusInternalServerError,
				fault.ErrInternalServer,
				fmt.Sprintf("failed to retrieve access token from Redis for key '%s': %v", cacheKey, err),
			)
		}

		accessToken, ok := tokenValue.(string)
		if !ok {
			return nil, fault.Custom(
				http.StatusInternalServerError,
				fault.ErrInternalServer,
				fmt.Sprintf("cached access token is not a string for key '%s'", cacheKey),
			)
		}

		user, err := u.user.GetUserDetail(model.GetUserDetailRequest{Email: body.Email})
		if err != nil {
			return nil, err
		}

		return &model.LoginResponse{
			UserData:    *user,
			AccessToken: accessToken,
		}, nil
	}

	exist, err := u.user.UserExistsByName(body.Name)

	if err != nil {
		return nil, err
	}

	var user *model.User
	var userId uuid.UUID

	if !exist {
		body.Password = middlewares.GenerateHashed(body.Password)

		createdId, err := u.user.InsertUser(body)
		if err != nil {
			return nil, err
		}
		userId = *createdId

		user, err = u.user.GetUserDetail(model.GetUserDetailRequest{
			UserId: userId,
		})
		if err != nil {
			return nil, err
		}
	} else {
		user, err = u.user.GetUserDetail(model.GetUserDetailRequest{
			Email: body.Email,
		})
		if err != nil {
			return nil, err
		}
		userId = user.Id
	}

	accessToken, payload, err := jwt.CreateAccessToken(user.Name, user.Email, userId.String())
	if err != nil {
		return nil, err
	}

	refreshToken, refreshPayload, err := jwt.CreateRefreshToken(user.Name, user.Email, userId.String())
	if err != nil {
		return nil, err
	}

	if err := cache.Set(ctx, u.redis, cacheKey, *accessToken, 10*time.Minute); err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		UserData:              *user,
		AccessToken:           *accessToken,
		AccessTokenExpiresAt:  &payload.ExpiresAt.Time,
		RefreshToken:          *refreshToken,
		RefreshTokenExpiresAt: &refreshPayload.ExpiresAt.Time,
	}, nil
}
