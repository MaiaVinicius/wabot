package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getUrl(licenseId string, key string) string {
	//[LICENSE_ID]

	err2 := godotenv.Load()
	if err2 != nil {
		log.Fatal("Error loading .env file")
	}

	var replacer = strings.NewReplacer("[LICENSE_ID]", licenseId)

	url := os.Getenv(key)

	url = replacer.Replace(url)

	return url
}

type Sent struct {
	LicenseId     int
	AppointmentId int
	EventId       int
}

type DataRemove struct {
	DataList []Sent
}

func RemoveQueue(data []Sent) {
	println("Removing from server sms queue")

	url := getUrl(strconv.Itoa(105), "REMOVE_QUEUE_URL")
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

	fmt.Println("Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, 	_ := ioutil.ReadAll(resp.Body)
	if false {
		fmt.Println("response Body:", string(body))
	}
}
