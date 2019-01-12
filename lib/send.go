package lib

import (
	"fmt"
	"github.com/rhymen/go-whatsapp"
	w "github.com/rhymen/go-whatsapp"
	"math/rand"
	"time"
)

func Send(wac *w.Conn, sendTo string, message string) int {

	t := 1 + rand.Intn(5)
	//t := 90 + rand.Intn(20)

	fmt.Printf("		timeout: %d \n", t)

	timeout := time.Duration(t)

	<-time.After(timeout * time.Second)

	var numbers []string

	numbers = append(numbers, sendTo)

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: sendTo + "@s.whatsapp.net",
		},
		Text: message,
	}

	wac.Send(msg)
	//println(msg.Text)

	return 200
}
