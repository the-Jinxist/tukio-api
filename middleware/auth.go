package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/the-Jinxist/tukio-api/internal/token"
	"github.com/the-Jinxist/tukio-api/pkg"
	"github.com/thedevsaddam/renderer"
)

var rnd = renderer.New()

type HeaderValue string

func VerifyCodeAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		rawToken := r.Header.Get("Authorization")
		tokenStr := strings.TrimPrefix(rawToken, "Bearer ")
		if tokenStr == "" {
			rnd.JSON(w, http.StatusUnauthorized, pkg.GenericResponse{
				Message: "you are unauthorized to access this endpoint",
				Status:  "error",
			})
			return
		}

		claims := jwt.MapClaims{}
		t, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
			return token.VERIFY_TOKEN_SECRET, nil
		})

		if err != nil {
			rnd.JSON(w, http.StatusUnauthorized, pkg.GenericResponse{
				Message: "you are unauthorized to access this endpoint",
				Status:  "error",
			})
			return
		}

		if !t.Valid {
			rnd.JSON(w, http.StatusUnauthorized, pkg.GenericResponse{
				Message: "you are unauthorized to access this endpoint",
				Status:  "error",
			})
			return
		}

		if claims["token_id"] != "verify_otp" {
			rnd.JSON(w, http.StatusUnauthorized, pkg.GenericResponse{
				Message: "you are unauthorized to access this endpoint",
				Status:  "error",
			})
			return
		}

		c := context.WithValue(r.Context(), HeaderValue("uID"), claims["user_id"])
		r = r.WithContext(c)
		next.ServeHTTP(w, r)
	})
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenStr := strings.TrimPrefix("Bearer ", r.Header.Get("Authorization"))
		if tokenStr == "" {
			rnd.JSON(w, http.StatusUnauthorized, pkg.GenericResponse{
				Message: "you are unauthorized to access this endpoint",
				Status:  "error",
			})
			return
		}

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return token.AUTH_TOKEN_SECRET, nil
		})

		if err != nil {
			rnd.JSON(w, http.StatusUnauthorized, pkg.GenericResponse{
				Message: "you are unauthorized to access this endpoint",
				Status:  "error",
			})
			return
		}

		if claims["token_id"] != "verify_otp" {
			rnd.JSON(w, http.StatusUnauthorized, pkg.GenericResponse{
				Message: "you are unauthorized to access this endpoint",
				Status:  "error",
			})
			return
		}

		c := context.WithValue(r.Context(), HeaderValue("uID"), claims["user_id"])
		r = r.WithContext(c)
		next.ServeHTTP(w, r)
	})
}

func GetUserID(ctx context.Context) string {
	return ctx.Value(HeaderValue("uID")).(string)
}
