package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/Tinddd28/GoPTL/internal/repository"
	"github.com/Tinddd28/GoPTL/internal/user"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "nf289^ho3h2t2hfh32fhei&^E"
	tokenTTL   = time.Hour * 12
	signingKey = "kjdbgbsj#@j141fodjsvbsdv"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user user.User) (int, error) {
	user.Password = s.generatePassHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	usr, err := s.repo.GetUser(email, s.generatePassHash(password))
	//log_ := logger.SetupPrettyLogger()
	if err != nil {
		//log_.Info("GenerateToken", slog.Any("error", err))
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		usr.Id,
		usr.Email,
	})
	//log_.Info("token is: ", slog.Any("token", token))
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) generatePassHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(*tokenClaims); !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	} else {
		return claims.UserId, claims.Email, nil
	}
}
