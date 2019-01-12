package lib

import (
	"encoding/gob"
	"fmt"
	"github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/rhymen/go-whatsapp"
	w "github.com/rhymen/go-whatsapp"
	"os"
	"time"
)

func NewSession(timeout time.Duration) (*w.Conn, error) {
	wac, err := whatsapp.NewConn(timeout * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
	}

	return wac, err
}

func Connect(phone string) (*w.Conn, error) {
	wac, err := NewSession(3)

	//create new WhatsApp connection
	err = login(wac, phone)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return wac, err
	}
	return wac, err
}

func login(wac *whatsapp.Conn, phone string) error {
	//load saved session
	session, err := readSession(phone)
	if err == nil {
		//restore session
		session, err = wac.RestoreSession(session)
		if err != nil {
			return fmt.Errorf("restoring failed: %v\n", err)
		}
	} else {
		//no saved session -> regular login
		qr := make(chan string)
		go func() {
			terminal := qrcodeTerminal.New()
			terminal.Get(<-qr).Print()
		}()
		session, err = wac.Login(qr)
		if err != nil {
			return fmt.Errorf("error during login: %v\n", err)
		}
	}

	//save session
	err = writeSession(session, phone)
	if err != nil {
		return fmt.Errorf("error saving session: %v\n", err)
	}
	return nil
}

func readSession(phone string) (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Open(getDir(phone))
	if err != nil {
		return session, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

func writeSession(session whatsapp.Session, phone string) error {
	file, err := os.Create(getDir(phone))
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}

func getDir(id string) string {
	var path string
	var sessionPath string

	path = "/Users/MaiaVinicius/go/src/github.com/MaiaVinicius/wabot/storage/session"
	//path = "/root/go/src/github.com/MaiaVinicius/wabot/storage/session"

	sessionPath = fmt.Sprintf("%s/%s-whatsappSession.gob", path, id)

	return sessionPath
}
