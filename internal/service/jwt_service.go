package service

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/reinanhs/golang-web-api-structure/internal/core/config"
	"strconv"
	"time"
)

type JWTService interface {
	GenerateToken(id uint) (string, error)
	ValidateToken(token string) bool
	GetIDFromToken(t string) (int64, error)
}

type jwtService struct {
	ctx       context.Context
	secretKey string
	issuer    string
	expiresIn int
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

//NewJWTService is creates a new instance of JWTService
func NewJWTService(ctx context.Context) JWTService {
	return &jwtService{
		ctx:       ctx,
		secretKey: ctx.Value("config").(*config.AppConfig).JwtSecret,
		expiresIn: ctx.Value("config").(*config.AppConfig).JwtExpiresIn,
		issuer:    "auth-login",
	}
}

func (s *jwtService) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(s.expiresIn)).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}

func (s *jwtService) GetIDFromToken(t string) (int64, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid Token: %v", t)
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := fmt.Sprintf("%v", claims["sum"])
		val, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return 0, err
		}

		return val, nil
	}

	return 0, err
}
