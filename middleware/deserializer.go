package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/herizal95/gomux/helper"
	"github.com/herizal95/gomux/utils"
)

type JWTClaims struct {
	Username string
	jwt.RegisteredClaims
}

func Deserializer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				reponse := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, reponse)
				return
			}
		}

		// mengambil token
		tokenString := c.Value

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return utils.JWT_KEY, nil

		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// token invalid
				reponse := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, reponse)
				return
			case jwt.ValidationErrorExpired:
				// token expired
				reponse := map[string]string{"message": "Unauthorized, token expired"}
				helper.ResponseJson(w, http.StatusUnauthorized, reponse)
				return
			default:
				reponse := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, reponse)
				return
			}
		}

		if !token.Valid {
			reponse := map[string]string{"message": "Unauthorized"}
			helper.ResponseJson(w, http.StatusUnauthorized, reponse)
			return
		}

		next.ServeHTTP(w, r)

	})
}
