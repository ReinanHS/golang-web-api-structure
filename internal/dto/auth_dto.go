package dto

import "time"

type AuthDto struct {
	Token     string        `json:"token"`
	ExpiresIn time.Duration `json:"expires_in"`
}
