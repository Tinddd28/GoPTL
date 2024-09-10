package service

import (
	"errors"
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/Tinddd28/GoPTL/internal/repository"
	"github.com/Tinddd28/GoPTL/pkg/hash"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	tokenTTL          = time.Hour * 12
	signingKey        = "kjdbgbsj#@j141fodjsvbsdv"
	issupuerUserEmpty = false
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId      int  `json:"user_id"`
	Issuperuser bool `json:"is_superuser"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = hash.GeneratePassHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	usr, err := s.repo.GetUser(email, hash.GeneratePassHash(password))
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
		usr.Issuperuser,
	})
	//log_.Info("token is: ", slog.Any("token", token))
	return token.SignedString([]byte(signingKey))
}

//func (s *AuthService) generatePassHash(password string) string {
//	hash := sha1.New()
//	hash.Write([]byte(password))
//
//	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
//}

func (s *AuthService) ParseToken(accessToken string) (int, bool, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, issupuerUserEmpty, err
	}

	if claims, ok := token.Claims.(*tokenClaims); !ok {
		return 0, issupuerUserEmpty, errors.New("token claims are not of type *tokenClaims")
	} else {
		return claims.UserId, claims.Issuperuser, nil
	}
}
