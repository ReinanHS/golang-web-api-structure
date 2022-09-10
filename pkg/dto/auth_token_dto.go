package dto

import "time"

type AuthTokenDto struct {
	Token     string        `json:"token"`
	ExpiresIn time.Duration `json:"expires_in"`
}
