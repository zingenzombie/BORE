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

	th.connectedDevs = make(map[string]*connectedDevice)
	th.rooms = make(map[string]*Room)

	makeRooms()

	http.Handle("/joinRoom", th)
	http.Handle("/connect", th)

	http.ListenAndServe(":3621", nil)

}

func makeRooms(roomAndNames *RoomAndNames) {
	roomAndNames.rooms["star"] = &Room{make(map[string]*connectedDevice)}
	roomAndNames.rooms["square"] = &Room{make(map[string]*connectedDevice)}
	roomAndNames.rooms["circle"] = &Room{make(map[string]*connectedDevice)}
	roomAndNames.rooms["triangle"] = &Room{make(map[string]*connectedDevice)}
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
	ip     string
	name   string
	active bool
}

type Room struct {
	connectedDevs map[string]*connectedDevice
}

// Structure for the series of rooms that hold users.
type RoomAndNames struct {
	counter       int
	connectedDevs map[string]*connectedDevice
	rooms         map[string]*Room
}

// Responds to HTTP /count request.
func (ct *RoomAndNames) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	newUser(w, ct, r)

	fmt.Fprintln(w, "Hello ", r.RemoteAddr)
}

func newUser(w http.ResponseWriter, roomAndNames *RoomAndNames, r *http.Request) {

	if val, ok := roomAndNames.connectedDevs[r.RemoteAddr]; ok {
		fmt.Println(val.ip)
	} else {
		tmp := connectedDevice{"", ""}
		tmp.ip = r.RemoteAddr

		roomAndNames.connectedDevs[r.RemoteAddr] = &tmp
		printUsers(w, roomAndNames, r)
	}
}

/*
// creating a new room
func createRoom(roomAndNames *RoomAndNames) {
	roomAndNames.rooms = append(roomAndNames.rooms, Room{})
}*/

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

func printUsers(w http.ResponseWriter, roomAndNames *RoomAndNames, r *http.Request) {
	for key, element := range roomAndNames.connectedDevs {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
}
