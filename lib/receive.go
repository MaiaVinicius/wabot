package lib

import (
	"fmt"
	"github.com/MaiaVinicius/go-whatsapp"
	"os"
	"strings"
	"time"
)

type waHandler struct{}

type Response struct {
	Phone     string
	ID        string
	Message   string
	Datetime  string
	Timestamp int64
	FromMe    bool
	Status    whatsapp.MessageStatus
}

var responses []Response

//HandleError needs to be implemented to be a valid WhatsApp handler
func (*waHandler) HandleError(err error) {
	fmt.Fprintf(os.Stderr, "error occoured: %v", err)
}

//Optional to be implemented. Implement HandleXXXMessage for the types you need.
func (*waHandler) HandleTextMessage(message whatsapp.TextMessage) {

	//if message.Info.FromMe != true {
	var response Response

	unixTimeUTC := time.Unix(int64(message.Info.Timestamp), 0)
	unitTimeInRFC3339 := unixTimeUTC.Format("2006-01-02 15:04:05")

	response.ID = message.Info.Id
	response.Message = message.Text
	response.Timestamp = int64(message.Info.Timestamp)
	response.Datetime = unitTimeInRFC3339
	response.Status = message.Info.Status
	response.FromMe = message.Info.FromMe
	response.Phone = strings.Replace(message.Info.RemoteJid, "@s.whatsapp.net", "", -1)

	responses = append(responses, response)
	//}
}

//Example for media handling. Video, Audio, Document are also possible in the same way
func (*waHandler) HandleImageMessage(message whatsapp.ImageMessage) {
	return

	data, err := message.Download()
	if err != nil {
		return
	}
	filename := fmt.Sprintf("%v/%v.%v", os.TempDir(), message.Info.Id, strings.Split(message.Type, "/")[1])
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return
	}
	_, err = file.Write(data)
	if err != nil {
		return
	}
	fmt.Printf("%v %v\n\timage reveived, saved at:%v\n", message.Info.Timestamp, message.Info.RemoteJid, filename)
}

func Receive(phone string) []Response {
	wac, err := NewSession(10)

	//Add handler
	wac.AddHandler(&waHandler{})

	//create new WhatsApp connection
	err = login(wac, phone)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
	}
	<-time.After(3 * time.Second)

	return responses
}
