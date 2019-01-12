package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

type SMS struct {
	Id            int64  `json:"id"`
	LicenseId     int64  `json:"LicencaID"`
	AppointmentId int64  `json:"AgendamentoID"`
	EventId       int64  `json:"EventoID"`
	Message       string `json:"Mensagem"`
	DateTime      string `json:"DataHora"`
	Phone         string `json:"Celular"`
	CreatedAt     string `json:"sysDate"`
}

type Queue struct {
	SMSList []SMS `json:"sms"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {

	err2 := godotenv.Load()
	if err2 != nil {
		log.Fatal("Error loading .env file")
	}
	url := os.Getenv("QUEUE_URL")

	println("Envio finalizado. \n")
	queue := new(Queue)
	getJson(url, queue)

	for _, element := range queue.SMSList {

		println(element.Phone)
	}

}
