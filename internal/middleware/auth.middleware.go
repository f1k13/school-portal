package middleware

import (
	"context"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

const UserIDKey string = "userID"

type AuthMiddleWare struct {
	JwtSecret string
}

func NewAuthMiddleware() *AuthMiddleWare {
	return &AuthMiddleWare{JwtSecret: os.Getenv("JWT_SECRET_KEY")}
}
func (am *AuthMiddleWare) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenBearer := strings.Split(authHeader, " ")
		if len(tokenBearer) != 2 || tokenBearer[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
		}
		tokenString := tokenBearer[1]
		claims := jwt.MapClaims{}
		t, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil || !t.Valid {
			logger.Log.Error("Token invalid", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		userID, ok := claims["sub"].(string)
		if !ok {
			http.Error(w, "Invalid token user", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
