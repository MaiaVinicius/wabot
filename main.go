package main

import (
	"encoding/json"
	"github.com/MaiaVinicius/wabot/input"
	"github.com/MaiaVinicius/wabot/model"
	"github.com/MaiaVinicius/wabot/service"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func StartProjects(w http.ResponseWriter, r *http.Request) {
	service.StartProjects()

	json.NewEncoder(w).Encode(true)
	return
}

func main() {
	//service.StartProjects()
	startCron()
}

func startWebServer() {
	router := mux.NewRouter()

	router.HandleFunc("/start", StartProjects).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func startCron() {
	c := cron.New()
	model.LogMessage(1, "Cron inicializado.", 0)
	RunJob()

	c.AddFunc("@every 15m", RunJob)
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func RunJob() {
	input.Feed()
	println("===============================")
	println("|     Iniciando Cron job.     |")
	println("===============================")
	service.StartProjects()
}
