package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	fmt.Println("Hello, welcome to BORE, your file-sharing drill!")

	fmt.Println("Starting server...")

	fmt.Print(GetOutboundIP())
	fmt.Println(":3621")

	th := &RoomAndNames{counter: 0}

	http.Handle("/count", th)

	http.ListenAndServe(":3621", nil)

}

// Gets the device's local address, which is returned to the user.
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// Structure for a user's address and string.
type connectedDevice struct {
	ip   string
	name string
}

// Structure for the series of rooms that hold users.
type RoomAndNames struct {
	counter       int
	connectedDevs map[string]*connectedDevice
}

// Responds to HTTP /count request.
func (ct *RoomAndNames) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(ct.counter)
	ct.counter++

	fmt.Fprintln(w, "Hello ", r.RemoteAddr)

	newUser(ct, r)
	//printUsers(w, ct, r)

	fmt.Fprintln(w, "Counter:", ct.counter)
}

func newUser(roomAndNames *RoomAndNames, r *http.Request) {

	if val, ok := roomAndNames.connectedDevs[r.RemoteAddr]; ok {
		fmt.Println(val.ip)
	} else {
		var tmp *connectedDevice
		tmp.ip = r.RemoteAddr

		roomAndNames.connectedDevs[r.RemoteAddr] = tmp
		fmt.Println("F")
	}
}

// Saving display name from form entry
func displayName(w http.ResponseWriter, r *http.Request) {
	connectedDev := connectedDevice{
		string(GetOutboundIP()),
		r.FormValue("name"),
	}
	// need to store information in some way
	_ = connectedDev

}

/*
func printUsers(w http.ResponseWriter, roomAndNames *RoomAndNames, r *http.Request) {
	for key, element := range roomAndNames.connectedDevs {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
}*/
