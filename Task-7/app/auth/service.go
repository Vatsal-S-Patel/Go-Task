package auth

import (
	"encoding/json"
	"jwt-go/app"
	"jwt-go/model"
	"jwt-go/utils"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// service struct contains Database pointer
type service struct {
	DB *gorm.DB
}

// Service interface with all CRUD functions signature
type Service interface {
	GenerateJWTToken(http.HandlerFunc) http.HandlerFunc
	VerifyJWTToken(http.HandlerFunc) http.HandlerFunc
}

func NewService(app *app.App) Service {
	return &service{
		DB: app.DB,
	}
}

// Credentials for checking the user while login
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claims for adding the claims in JWT token
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Generate the JWT Token
func (s *service) GenerateJWTToken(endpointHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Initialize Credentials and decode it from input JSON data
		var credential Credentials
		err := json.NewDecoder(r.Body).Decode(&credential)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Match the provided JSON data to the database
		// If data match then proceed for JWT token generation
		var profile model.Profile
		s.DB.Find(&profile, "email = ?", credential.Email)
		if profile.Email != credential.Email || profile.Password != credential.Password {
			log.Println("EMail or Password not matched")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Initializing Claims
		claims := &Claims{
			Email: credential.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    profile.Email,
				ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Minute * 5)},
			},
		}

		// generating token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

		// Getting Secret Key from env file
		secretKey, err := utils.GetFromEnvFile("SECRET_KEY")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Generate signed token string
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set token in Cookie
		http.SetCookie(w, &http.Cookie{
			Name:    "Token",
			Value:   tokenString,
			Expires: claims.ExpiresAt.Time,
		})
	}
}

func (s *service) VerifyJWTToken(endpointHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve Token from Cookie
		storedToken, err := r.Cookie("Token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Parse the token to verify that it is Valid or not with claims
		token, err := jwt.ParseWithClaims(storedToken.Value, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			secretKey, err := utils.GetFromEnvFile("SECRET_KEY")
			if err != nil {
				log.Println(err.Error())
				return nil, err
			}
			return []byte(secretKey), nil
		})

		// Handling if the Token is invalid
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Execute passed handler
		endpointHandler(w, r)
	}
}
