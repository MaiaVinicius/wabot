package input

import (
	"encoding/json"
	"fmt"
	"github.com/MaiaVinicius/wabot/model"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	_ "log"
	"net/http"
	"os"
	_ "os"
	"strconv"
	"strings"
	_ "strings"
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

type SMSList struct {
	SMSList []SMS `json:"sms"`
}

type Queue struct {
	Data SMSList `json:"data"`
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

func Feed() {

	projects := model.GetProjects()

	for _, proj := range projects {

		println(proj.LicenseId)

		url := getUrl(strconv.Itoa(proj.LicenseId))

		println("Iniciando importação. \n")

		queue := new(Queue)
		println(url)

		getJson(url, queue)

		for _, element := range queue.Data.SMSList {

			phone := element.Phone
			licenseId := element.LicenseId
			appointmentId := element.AppointmentId
			eventId := element.EventId
			message := element.Message
			datetime := element.DateTime
			senderId := proj.SenderID

			alreadyAddedId := model.QueueAlreadyAdded(licenseId, appointmentId, phone, message)
			if alreadyAddedId == "" {
				messageAlreadySent := model.MessageAlreadySent(licenseId, appointmentId)

				if messageAlreadySent == "" {
					model.AddToQueue(phone, message, datetime, senderId, licenseId, appointmentId, eventId)
				}
			}
		}

		println(fmt.Sprintf("---->  	%d importados", len(queue.Data.SMSList)))
	}

}
