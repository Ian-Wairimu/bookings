package main

import (
	"fmt"
	"github.com/Ian-Wairimu/bookings/pkg/config"
	"github.com/Ian-Wairimu/bookings/pkg/handler"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const (
	PORT = "localhost:8081"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//using the template cache
	app.InProduction = false
	//scs package for session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	//tc, err := render.CreateTemplateCache()
	//if err != nil {
	//	log.Fatal("cannot create template cache")
	//}
	//app.TemplateCache = tc
	//render.NewTemplates(&app)
	app.UseCache = false
	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	fmt.Println(fmt.Sprintf("-> Starting the application on port: %s", PORT))
	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
