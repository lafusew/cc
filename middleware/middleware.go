package middleware

import (
	"net/http"
	"strings"

	"github.com/lafusew/cc/utils"
)

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		var tokenStr string;

		if len(strings.Split(bearerToken, " ")) != 2 {
			utils.JsonResponse(w, http.StatusBadRequest, map[string]interface{}{
				"error": "authorization token was not provided",
			})
			return
		}

		tokenStr = strings.Split(bearerToken, " ")[1]

		_, err := utils.CheckJWT(tokenStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		next(w, r)
	}
}