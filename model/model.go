package model

import (
	"database/sql"
	"fmt"
	"github.com/MaiaVinicius/go-whatsapp"
	"github.com/MaiaVinicius/wabot/lib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type Project struct {
	ID        int
	Label     string
	LicenseId int
	SenderID  int
	Phone     string
}
type Queue struct {
	ID            int
	SenderID      int
	LicenseId     int
	AppointmentId int
	EventId       int
	Message       string
	Phone         string
	//Metadata2     string
}

type Sent struct {
	ID            int
	LicenseId     int
	AppointmentId int
	EventId       int
}

type Config struct {
	SendMinimumTimeout int
	SendTimeRandom     int
	LimitPerExecution  int
	CronTimeout        int
}

var db = dbConn()

func unicode2utf8(source string) string {
	var res = []string{""}
	sUnicode := strings.Split(source, "\\u")
	var context = ""
	for _, v := range sUnicode {
		var additional = ""
		if len(v) < 1 {
			continue
		}
		if len(v) > 4 {
			rs := []rune(v)
			v = string(rs[:4])
			additional = string(rs[4:])
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			context += v
		}
		context += fmt.Sprintf("%c", temp)
		context += additional
	}
	res = append(res, context)
	return strings.Join(res, "")
}

func dbConn() (db *sql.DB) {

	err2 := godotenv.Load()
	if err2 != nil {
		log.Fatal("Error loading .env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_SCHEMA")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetProjects() []Project {
	rows, err := db.Query("SELECT p.id, p.label, s.id senderId, s.phone, p.license_id LicenseId FROM wabot_project p INNER JOIN wabot_sender s ON s.project_id = p.id where s.status_id=1 and s.active=1 and p.active=1")

	if err != nil {
		panic(err.Error())
	}

	var projects []Project
	for rows.Next() {
		var project Project

		err = rows.Scan(&project.ID, &project.Label, &project.SenderID, &project.Phone, &project.LicenseId)

		projects = append(projects, project)
	}

	return projects
}

func GetQueue(senderId int) []Queue {
	config := GetConfig()

	limit := config.LimitPerExecution

	stmt, err := db.Prepare("SELECT id, message, phone, license_id as LicenseId, appointment_id as AppointmentId FROM wabot_queue where active=1 and sender_id=? and send_date=CURDATE() order by send_date, send_time asc LIMIT ?")

	rows, err := stmt.Query(senderId, limit)

	if err != nil {
		panic(err.Error())
	}

	var queues []Queue
	for rows.Next() {
		var queue Queue

		err = rows.Scan(&queue.ID, &queue.Message, &queue.Phone, &queue.LicenseId, &queue.AppointmentId)

		queues = append(queues, queue)
	}

	return queues
}

func RemoveFromQueue(queueId int) {
	//println(queueId)

	stmt, err := db.Prepare("INSERT INTO wabot_sent (sender_id, phone, message, price, appointment_id, license_id, event_id)  (SELECT sender_id, phone, message, price, appointment_id, license_id, event_id FROM wabot_queue where id=?)")

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(queueId)

	stmtDel, errDel := db.Prepare("DELETE FROM wabot_queue WHERE id=?")

	if errDel != nil {
		panic(errDel.Error())
	}

	stmtDel.Exec(queueId)
}

func InsertResponse(projectId int, phone string, id string, message string, timestamp string, statusId whatsapp.MessageStatus, fromMe bool) bool {

	stmt, err := db.Prepare("INSERT INTO wabot_response (id, phone, message, project_id, datetime, status_id,from_me) VALUES (?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	fromMeNumeric := 0
	if fromMe {
		fromMeNumeric = 1
	}

	res, err := stmt.Exec(id, phone, message, projectId, timestamp, statusId, fromMeNumeric)

	if err != nil {
		//fmt.Println(err.Error())
	} else {
		count, err2 := res.RowsAffected()

		if err2 == nil {
			//linha foi adicionada

			if count > 0 {
				//println("Affected row")
				licenseId, appointmentId, eventId := findResponseClient(phone, timestamp)

				if licenseId > 0 {
					//println("Updating response")
					updateResponse(id, licenseId, appointmentId, eventId)

				}
				return true
			}
		}
	}

	return false
}

func updateResponse(responseId string, licenseId int, appointmentId int, eventId int) {
	stmt, err := db.Prepare("UPDATE wabot_response SET event_id=?,license_id=?, appointment_id=? WHERE id=? AND license_id is null and appointment_id is null")

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(eventId, licenseId, appointmentId, responseId)
}

func findResponseClient(phone string, datetime string) (int, int, int) {
	stmt, err := db.Prepare("SELECT ifnull(ID,0)ID, ifnull(license_id,0 ) LicenseId, ifnull(appointment_id,0) AppointmentId, ifnull(event_id, 0)EventId FROM wabot_sent WHERE phone=? ORDER BY datetime DESC LIMIT 1")

	rows := stmt.QueryRow(phone)

	if err != nil {
		panic(err.Error())
	}

	var response Sent

	err2 := rows.Scan(&response.ID, &response.AppointmentId, &response.LicenseId, &response.EventId)

	if err2 != nil {
		//panic(err2.Error()) // proper error handling instead of panic in your app
	}

	appointmentId := response.AppointmentId
	licenseId := response.LicenseId
	eventId := response.EventId

	return appointmentId, licenseId, eventId
}

func LogMessage(logType int, message string, projectId int) {
	stmt, err := db.Prepare("INSERT INTO wabot_log (project_id, log_type_id, message) value (?, ?, ?)")

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(projectId, logType, message)
}

func GetLastSent(projectId int) string {
	stmt, err := db.Prepare("SELECT datetime FROM wabot_response WHERE project_id=? ORDER BY datetime DESC LIMIT 1")

	rows := stmt.QueryRow(projectId)

	if err != nil {
		panic(err.Error())
	}

	var response lib.Response

	rows.Scan(&response.Datetime)

	return response.Datetime
}

func QueueAlreadyAdded(licenseId int64, appointmentId int64, phone string, message string) string {
	stmt, err := db.Prepare("SELECT id FROM wabot_queue WHERE (license_id=? AND appointment_id=?) or (phone=? and message=? and send_date=curdate())")

	rows := stmt.QueryRow(licenseId, appointmentId, phone, message)

	if err != nil {
		panic(err.Error())
	}

	var response lib.Response

	rows.Scan(&response.ID)

	return response.ID
}

func MessageAlreadySent(licenseId int64, appointmentId int64) string {
	stmt, err := db.Prepare("SELECT id FROM wabot_sent WHERE license_id=? AND appointment_id=?")

	rows := stmt.QueryRow(licenseId, appointmentId)

	if err != nil {
		panic(err.Error())
	}

	var response lib.Response

	rows.Scan(&response.ID)

	return response.ID
}

func AddToQueue(phone string, message string, datetime string, senderId int, licenseId int64, appointmentId int64, eventId int64) {
	stmt, err := db.Prepare("INSERT INTO wabot.wabot_queue (sender_id, message, phone, send_date, send_time, license_id, appointment_id, event_id) value (?, ?, ?, DATE(?), TIME(?), ?, ?, ?)")

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(senderId, message, phone, datetime, datetime, licenseId, appointmentId, eventId)
}

func GetResponsesToSync() []lib.ResponseToServer {
	rows, err := db.Query("SELECT id, message, appointment_id AppointmentId, event_id EventId, license_id LicenseId, datetime, phone, auto_id AutoId FROM wabot.wabot_response WHERE sync=0 AND from_me=0 AND license_id is not null")

	if err != nil {
		panic(err.Error())
	}

	var responses []lib.ResponseToServer
	for rows.Next() {
		var response lib.ResponseToServer

		err = rows.Scan(&response.ID, &response.Message, &response.AppointmentId, &response.EventId, &response.LicenseId, &response.DateTime, &response.Phone, &response.AutoId)

		responses = append(responses, response)
	}

	return responses
}

func UpdateSyncedResponses(lastId int) {
	stmt, err := db.Prepare("UPDATE wabot_response SET sync=1, sync_at=NOW() WHERE from_me=0 AND sync=0 AND auto_id<=?")

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(lastId)
}

func RegularizeResponseLicenseId() {
	stmt, err := db.Prepare("UPDATE wabot_response SET license_id=(SELECT license_id) WHERE license_id is null")

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec()
}

func GetConfig() Config {
	stmt, err := db.Prepare("SELECT send_minimum_timeout SendMinimumTimeout, send_time_random SendTimeRandom, limit_per_execution LimitPerExecution, cron_timeout CronTimeout  FROM wabot_config WHERE id=1")

	if err != nil {
		panic(err.Error())
	}

	rows := stmt.QueryRow()

	var response Config

	err = rows.Scan(&response.SendMinimumTimeout, &response.SendTimeRandom, &response.LimitPerExecution, &response.CronTimeout)

	return response
}
