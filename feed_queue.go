package main

import (
	"encoding/json"
	"github.com/MaiaVinicius/wabot/model"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func getUrl(licenseId string) string {
	//[LICENSE_ID]

	err2 := godotenv.Load()
	if err2 != nil {
		log.Fatal("Error loading .env file")
	}

	var replacer = strings.NewReplacer("[LICENSE_ID]", licenseId)

	url := os.Getenv("QUEUE_URL")

	url = replacer.Replace(url)

	return url
}

func main() {
	url := getUrl(strconv.Itoa(105))

	println("Iniciando importação. \n")

	queue := new(Queue)
	getJson(url, queue)

	for _, element := range queue.SMSList {

		phone := element.Phone
		licenseId := element.LicenseId
		appointmentId := element.AppointmentId
		message := element.Message
		datetime := element.DateTime
		senderId := 1

		alreadyAddedId := model.QueueAlreadyAdded(licenseId, appointmentId)
		if alreadyAddedId == "" {
			messageAlreadySent := model.MessageAlreadySent(licenseId, appointmentId)

			if messageAlreadySent == "" {
				model.AddToQueue(phone, message, datetime, senderId, licenseId, appointmentId)
			}
		}
	}

}
