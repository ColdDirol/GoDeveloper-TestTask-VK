package auth

import (
	"GoDeveloperVK-testTask/auth/jwt"
	"GoDeveloperVK-testTask/utils"
)

func InitAuth(AuthConfig *utils.Auth) {
	jwt.SecretKey = []byte(AuthConfig.SecretKey)
	jwt.Salt = AuthConfig.Salt

	initRegistrationHandlers()
	initLoginHandlers()
}
