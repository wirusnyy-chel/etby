package routers

import (
	"context"
	middlewares "etby/services/gateway/internal/middlwares"
	authService "etby/services/login/client"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func Login(path string) chi.Router {
	router := chi.NewRouter()
	middlewares.AddToNoAuth(path)
	l := authService.GetLoginClient()
	router.Post("/", func(rw http.ResponseWriter, r *http.Request) {
		login := r.PostFormValue("login")
		password := r.PostFormValue("password")
		auth, refresh, err := l.Login(context.Background(), login, password)
		if err != nil {
			logrus.Warn(err)
		}
		http.SetCookie(rw, &http.Cookie{
			Name:    "auth",
			Value:   auth,
			Expires: time.Now().Add(time.Hour),
			Path:    "/",
		})
		http.SetCookie(rw, &http.Cookie{
			Name:    "refresh",
			Value:   refresh,
			Expires: time.Now().Add(time.Hour),
			Path:    "/",
		})
		http.Redirect(rw, r, "/", http.StatusFound)
	})
	return router
}
