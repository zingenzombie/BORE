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

	th := &CounterHandler{counter: 0}
	http.Handle("/count", th)
	http.ListenAndServe(":3621", nil)

	for 1 == 1 {
		break
	}

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
type CounterHandler struct {
	counter       int
	connectedDevs map[string]*connectedDevice
}

// Responds to all HTTP requests.
func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(ct.counter)
	ct.counter++

	//THIS NEEDS TO BE CHANGED AS IT WILL SIMPLY CHECK FOR THE FRONT END'S ADDRESS AND STORE THAT AS A USER
	fmt.Fprintln(w, "Hello ", r.RemoteAddr)

	fmt.Fprintln(w, "Counter:", ct.counter)
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
