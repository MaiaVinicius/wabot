package service

import (
	"fmt"
	"github.com/MaiaVinicius/wabot/lib"
	"github.com/MaiaVinicius/wabot/model"
	w "github.com/rhymen/go-whatsapp"
	"os"
	"strings"
	"time"
)

func StartProjects() {
	projects := model.GetProjects()

	for _, element := range projects {

		prepareQueue(element)
		//prepareReception(element)
	}
	println("Envio finalizado. \n")
}

func prepareQueue(project model.Project) {
	print(fmt.Sprintf("Enviando do Celular: %s id: %d \n", project.Phone, int(project.SenderID)))

	queue := model.GetQueue(project.SenderID)

	if len(queue) > 0 {
		msg := fmt.Sprintf("Envio iniciado. %d", len(queue))

		model.LogMessage(1, msg, project.ID)

		sendQueue(project.Phone, queue)
	}
}

func prepareReception(project model.Project) {
	print(fmt.Sprintf("Recebendo respostas do Celular: %s \n", project.Phone))

	var datetimeLastSent string

	datetimeLastSent = model.GetLastSent(project.ID)

	t, err := time.Parse("2006-01-02 15:04:05", datetimeLastSent)

	if err != nil {
		fmt.Println(err)
	}

	var responses []lib.Response

	responses = lib.Receive(project.Phone)

	for _, element := range responses {

		if element.Timestamp-10800 > t.Unix() || datetimeLastSent == "" {
			model.InsertResponse(project.ID, element.Phone, element.ID, element.Message, element.Datetime)
		}
	}
}

func sendQueue(senderPhone string, queue []model.Queue) {
	//connect

	wac, err := lib.Connect(senderPhone)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending message: %v", err)
	}

	for _, element := range queue {
		element.Message = strings.Replace(element.Message, "[METADATA]", element.Metadata2, -1)

		sendMessage(wac, element)
	}
}

func sendMessage(wac *w.Conn, message model.Queue) int {
	print(fmt.Sprintf("-		Enviando para: %s \n", message.Phone))

	status := lib.Send(wac, message.Phone, message.Message)

	if status == 200 {
		removeFromQueue(message.ID)
	}

	return status
}

func removeFromQueue(messageId int) {
	model.RemoveFromQueue(messageId)
	return
}
