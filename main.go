package main

import (
	"encoding/json"
	"fmt"
	"github.com/MaiaVinicius/wabot/input"
	"github.com/MaiaVinicius/wabot/model"
	"github.com/MaiaVinicius/wabot/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	//c := cron.New()
	RunJob()

	//c.AddFunc("@every 4m", RunJob)
	//go c.Start()
	//sig := make(chan os.Signal)
	//signal.Notify(sig, os.Interrupt, os.Kill)
	//<-sig
}

func RunJob() {

	defer recoverName()

	model.LogMessage(1, "Processo iniciado.", 0)

	input.Feed()
	println("===============================")
	println("|     Iniciando Cron job.     |")
	println("===============================")
	service.StartProjects()
}

func recoverName() {
	if r := recover(); r != nil {
		msg := fmt.Sprintf("Panic: ", r)
		println(msg)
		model.LogMessage(3, msg, 0)
	}
}
