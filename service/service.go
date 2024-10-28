package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/shiibs/fitness-app/models"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthService {
	return AuthService{db}
}

func (s *AuthService) HandleLogin(googleID, name, email, imageURL string) (*models.User, string, error) {
	var user models.User
	err := s.db.Where("google_id = ?", googleID).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// User does not exist, create a new user
			user = models.User{
				GoogleID: googleID,
				Name:     name,
				Email:    email,
				ImageURL: imageURL,
			}
			if err := s.db.Create(&user).Error; err != nil {
				return nil, "", err
			}
		} else {
			return nil, "", err
		}
	}

	tokenString, err := createJWTToken(user)
	if err != nil {
		return nil, "", err
	}

	return &user, tokenString, nil
}

func createJWTToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 720).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
