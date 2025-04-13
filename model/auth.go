package model

import "time"

type LoginResponse struct {
	UserData              User       `json:"user_data"`
	AccessToken           string     `json:"access_token"`
	AccessTokenExpiresAt  *time.Time `json:"access_token_expires_at"`
	RefreshToken          string     `json:"refresh_token"`
	RefreshTokenExpiresAt *time.Time `json:"refresh_token_expires_at"`
}
