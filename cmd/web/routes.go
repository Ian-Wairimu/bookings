package main

import (
	"github.com/Ian-Wairimu/bookings/pkg/config"
	"github.com/Ian-Wairimu/bookings/pkg/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//multiplexer is a http.handler
	mux := chi.NewRouter()

	//using a middleware in chi
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RealIP)
	mux.Use(WriteToConsole)
	mux.Use(SessionLoad)
	mux.Use(NoSurf)

	mux.Get("/home", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	return mux
}
