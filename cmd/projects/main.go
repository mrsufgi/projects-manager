package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	ehd "github.com/mrsufgi/projects-manager/internal/events/delivery/http"
	thd "github.com/mrsufgi/projects-manager/internal/projects/delivery/http"
	"github.com/rs/cors"

	er "github.com/mrsufgi/projects-manager/internal/events/repository/pg"
	es "github.com/mrsufgi/projects-manager/internal/events/service"
	tr "github.com/mrsufgi/projects-manager/internal/projects/repository/pg"
	ts "github.com/mrsufgi/projects-manager/internal/projects/service"
	"github.com/mrsufgi/projects-manager/pkg/helpers"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	// logger setup
	setupLogger(true)

	conf, err := sqlx.Connect("postgres", helpers.GetConnectionString())

	if err != nil {
		log.Fatalln(err)
	}

	pusherClient := helpers.GetPusherClient()

	trepo := tr.NewPgProjectsRepository(conf)
	erepo := er.NewPgEventsRepository(conf)
	mservice := ts.NewProjectService(trepo, pusherClient)
	eservice := es.NewEventService(erepo, mservice, pusherClient)

	port := ":3000"

	if err != nil {
		log.Fatalf("socket server unexpected error %v", err)
	}
	router := httprouter.New()
	router.GET("/metrics", Metrics(promhttp.Handler()))
	router.GET("/health", Health)

	// add CORS
	handler := cors.Default().Handler(router)

	s := &http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,

		// MaxHeaderBytes: 1 << 20,
	}

	// handlers setup their own router
	thd.NewProjectsHandler(router, mservice)
	ehd.NewEventsHandler(router, eservice)

	go func() {
		log.Infof("start http server on port %s", port)
		if err := s.ListenAndServe(); err != nil {
			log.Println("HTTP server shutting down")
			if err != http.ErrServerClosed {
				log.Fatalf("closed unexpected error %v", err)
			}

			s.Close()
		}
	}()

	gracefulShutdown(s)
}

func Metrics(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}

func Health(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func gracefulShutdown(s *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	ctxTimeout := time.Second * 10

	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Error while shutting down %v", err)
	}

	cancel()
	os.Exit(0)
}

func setupLogger(debug bool) {
	log.SetFormatter(&log.JSONFormatter{})

	if debug {
		log.SetLevel(log.DebugLevel)
	}
}
