package main

import "github.com/MaiaVinicius/wabot/lib"

func main() {

	var s []lib.Sent

	item := lib.Sent{}

	item.LicenseId = 0
	item.AppointmentId = 1

	// append works on nil slices.
	s = append(s, item)

	lib.RemoveQueue(s)
}
