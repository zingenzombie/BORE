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

	http.Handle("", th)

	//http.Handle("/newUser", )

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

type Room struct {
	connectedDevs map[string]*connectedDevice
}

// Structure for the series of rooms that hold users.
type RoomAndNames struct {
	counter       int
	connectedDevs map[string]*connectedDevice
	rooms         []Room
}

// Responds to HTTP /count request.
func (ct *RoomAndNames) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(ct.counter)
	ct.counter++

	fmt.Fprintln(w, "Hello ", r.RemoteAddr)

	newUser(ct, r)

	fmt.Fprintln(w, "Counter:", ct.counter)
}

func newUser(roomAndNames *RoomAndNames, r *http.Request) {

	val, ok := roomAndNames.connectedDevs[r.RemoteAddr]

	if ok {
		fmt.Println(val)
		return
	}

	roomAndNames.connectedDevs[r.RemoteAddr] = &connectedDevice{ip: r.RemoteAddr, name: ""}
}

// creating a new room
func createRoom(roomAndNames *RoomAndNames) {
	roomAndNames.rooms = append(roomAndNames.rooms, Room{})
}

// connect to a room
func joinRoom(room *Room, cd *connectedDevice) {
	room.connectedDevs[cd.ip] = cd
	// actual connection here
}

// leave room
func leaveRoom(room *Room, cd *connectedDevice) {
	delete(room.connectedDevs, cd.ip)

}

// Saving display name from form entry (not used)
func displayName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	_ = name
}
