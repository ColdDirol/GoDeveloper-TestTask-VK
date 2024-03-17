package middleware

import (
	"GoDeveloperVK-testTask/auth/jwt"
	"GoDeveloperVK-testTask/utils"
	"GoDeveloperVK-testTask/utils/logger"
	"errors"
	"net/http"
)

const (
	user  = "user"
	admin = "admin"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		claims, err := jwt.VerifyToken(token)
		if err != nil {
			utils.LOG.Error("unauthorized user - wrong jwt", logger.Err(err))
			http.Error(w, "unauthorized user - wrong jwt", http.StatusUnauthorized)
			return
		}

		switch claims.Role {
		case user:
			if r.Method != http.MethodGet {
				utils.LOG.Error("not enough rights", logger.Err(errors.New("not enough rights")))
				http.Error(w, "not enough rights", http.StatusForbidden)
			}
		case admin:
		default:
			utils.LOG.Error("not enough rights", logger.Err(err))
			http.Error(w, "not enough rights", http.StatusForbidden)
		}

		next(w, r)
	}
}
