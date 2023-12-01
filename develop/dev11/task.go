package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(NewErrorResponse(err))
		return
	}

	err = v.ValidateEventCreate(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewErrorResponse(err))
		return
	}

	err = s.CreateEvent(body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(NewErrorResponse(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(NewEditResponse("event created"))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(NewErrorResponse(err))
		return
	}

	err = v.ValidateEventID(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewErrorResponse(err))
		return
	}

	err = s.UpdateEvent(body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(NewErrorResponse(err))
		return
	}

	w.Write(NewEditResponse("event updated"))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(NewErrorResponse(err))
		return
	}

	err = v.ValidateEventID(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewErrorResponse(err))
		return
	}

	err = s.DeleteEvent(body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(NewErrorResponse(err))
		return
	}

	w.Write(NewEditResponse("event deleted"))
}

func dayHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	err := v.ValidateDate(params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewErrorResponse(err))
		return
	}

	jsonResp, err := s.GetEventsDay(params)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(NewErrorResponse(err))
		return
	}

	w.Write(jsonResp)
}

func weekHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	err := v.ValidateDate(params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewErrorResponse(err))
		return
	}

	jsonResp, err := s.GetEventsWeek(params)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(NewErrorResponse(err))
		return
	}

	w.Write(jsonResp)
}

func monthHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	err := v.ValidateMonth(params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewErrorResponse(err))
		return
	}

	jsonResp, err := s.GetEventsMonth(params)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(NewErrorResponse(err))
		return
	}

	w.Write(jsonResp)
}

var s = NewService()
var v = NewValidator()

const configPath = "config.json"

func main() {

	cfg, err := ParseConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	http.HandleFunc("/create_event", MakeLoggingHandler(createHandler))
	http.HandleFunc("/update_event", MakeLoggingHandler(updateHandler))
	http.HandleFunc("/delete_event", MakeLoggingHandler(deleteHandler))
	http.HandleFunc("/events_for_day", MakeLoggingHandler(dayHandler))
	http.HandleFunc("/events_for_week", MakeLoggingHandler(weekHandler))
	http.HandleFunc("/events_for_month", MakeLoggingHandler(monthHandler))

	server := &http.Server{Addr: cfg.Port}
	go func() {
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	logrus.WithField("port", cfg.Port).Info("server started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	logrus.Info("server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatal(err)
	}
}
