package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Authorization() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, auth := checkLoggin(r)
			ctx := context.WithValue(r.Context(), "userid", id)
			ctx = context.WithValue(ctx, "auth", auth)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

//Проверка на наличие аутентификации
func checkLoggin(r *http.Request) (id int, auth bool) {
	c, err := r.Cookie("auth")
	if err != nil || c.Value == "" {
		return 0, false
	}
	token, claims, err := parseToken(c.Value)
	if err != nil {
		return 0, false
	}
	var exp float64
	var jti int
	for key, val := range *claims {
		switch key {
		case "exp":
			exp, _ = strconv.ParseFloat(fmt.Sprint(val), 64)
		case "jti":
			jti, _ = strconv.Atoi(fmt.Sprint(val))
		}
	}
	if token.Valid {
		if time.Now().After(time.Unix(int64(exp), 0)) {
			fmt.Print("Token expired")
			return 0, false
		}
		return jti, true
	}
	return 0, false
}
