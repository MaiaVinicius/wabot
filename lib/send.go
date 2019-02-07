package lib

import (
	"fmt"
	"github.com/MaiaVinicius/go-whatsapp"
	w "github.com/MaiaVinicius/go-whatsapp"
	"math/rand"
	"time"
)

func Send(wac *w.Conn, sendTo string, message string, sendMinimumTimeout int, sendTimeRandom int) int {
	//10 segundos minimo bloqueia o chip

	//30 seg. minimo segura tranquilo
	//random de 20 seg

	t := sendMinimumTimeout + rand.Intn(sendTimeRandom)
	//t := 90 + rand.Intn(20)

	println(fmt.Sprintf("	/	timeout: %d s ...", t))

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
