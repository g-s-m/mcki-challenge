package main

import (
	"back/data"
	"strings"
	"io/ioutil"
	"encoding/json"
	"github.com/vrischmann/envconfig"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
)

type Config struct {
	HttpPort  string `envconfig:"BACK_HTTP_PORT"`
}

func getConfig() *Config {
	var conf Config
	if err := envconfig.Init(&conf); err != nil {
		log.Fatalln(err)
	}
	return &conf
}

type Server struct {
	reportApi data.ReportApi
}

func (p* Server) reloadReports(w http.ResponseWriter, r *http.Request) {
	p.reportApi = data.CreateReportStub()
	w.WriteHeader(http.StatusOK)
}

func (p* Server) getReports(w http.ResponseWriter, r *http.Request) {
	rl := p.reportApi.GetReports()
	response, err := json.Marshal(rl)
	if err != nil {
		log.Printf("Internal error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (p* Server) markReport(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(strings.TrimLeft(r.URL.Path, "/"), "reports/")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Incorect request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	type ChangeReport struct {
		State  string `json:"ticketState"`
	}

	var changeRequest ChangeReport
	json.Unmarshal(body, &changeRequest)
	log.Printf("Receive %s, extracted %s for id=%s", r.URL.Path, changeRequest.State, id)
	if changeRequest.State == "RESOLVED" {
		p.reportApi.ResolveReport(id)
	}
	if changeRequest.State == "BLOCKED" {
		p.reportApi.BlockReport(id)
	}
	w.WriteHeader(http.StatusOK)
}

func setContentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	conf := getConfig()
	reportApi := data.CreateReportStub()
	backSrv := Server {
		reportApi: reportApi,
	}
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(setContentTypeJson)
	router.Route("/reports", func (r chi.Router) {
		r.Get("/", func (w http.ResponseWriter, r *http.Request) { backSrv.getReports(w, r)})
		r.Post("/reload", func (w http.ResponseWriter, r *http.Request) { backSrv.reloadReports(w, r)})
		r.Put("/{reportId}", func (w http.ResponseWriter, r *http.Request) { backSrv.markReport(w, r)})
	})
	srv := http.Server {
		Addr:    ":" + conf.HttpPort,
		Handler: router,
	}
	srv.ListenAndServe()
}

