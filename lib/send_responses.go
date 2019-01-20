package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rhymen/go-whatsapp"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ResponseToServer struct {
	Phone         string
	ID            string
	AutoId        int
	Message       string
	DateTime      string
	Timestamp     int64
	LicenseId     int
	AppointmentId int
	FromMe        bool
	Status        whatsapp.MessageStatus
}

type DataResponses struct {
	Responses []ResponseToServer
}

func SendResponsesToServer(data []ResponseToServer) {
	url := getUrl(strconv.Itoa(105), "RESPONSES_URL")
	fmt.Println("URL:>", url)

	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)

	jsonStr, err := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	if true{
		fmt.Println("response Body:", string(body))
	}
}
