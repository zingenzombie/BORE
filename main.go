package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	fmt.Println("Hello, welcome to BORE, your file-sharing drill!")

	fmt.Println("Starting server...")

	fmt.Print(GetOutboundIP())
	fmt.Println(":3621")

	th := &RoomAndNames{counter: 0}

	th.connectedDevs = make(map[string]*connectedDevice)
	th.rooms = make(map[string]*Room)

	makeRooms(th)
	go checkUsers(th)

	http.Handle("/joinRoom", th)
	http.Handle("/connect", th)
	http.Handle("/debug", th)
	http.Handle("/checkIn", th)

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
	ip          string
	name        string
	active      bool
	lastCheckIn time.Time
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

	if r.RequestURI == "/debug" {
		debug(ct, r)
	}
	if r.RequestURI == "/checkIn" {
		checkIn(ct, r)
	}

}

func newUser(w http.ResponseWriter, roomAndNames *RoomAndNames, r *http.Request) {

	if val, ok := roomAndNames.connectedDevs[r.RemoteAddr]; ok {
		fmt.Println(val.ip)
	} else {
		tmp := connectedDevice{"", "", true, time.Now().UTC()}
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

func debug(roomAndNames *RoomAndNames, r *http.Request) {
	fmt.Println("DEBUG TIME!!!")
}

/*// Saving display name from form entry (not used)
func displayName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	_ = name
}*/

// setting time of last checkin
func checkIn(roomAndNames *RoomAndNames, r *http.Request) {
	roomAndNames.connectedDevs[r.RemoteAddr].lastCheckIn = time.Now().UTC()
}

// looping through connected devices
func checkActive(roomAndNames *RoomAndNames) {
	for _, element := range roomAndNames.connectedDevs {
		// device inactive if no checkin in 10 seconds
		if time.Now().UTC().Sub(element.lastCheckIn) > (10 * time.Second) {
			element.active = false
		}
	}
}

// calling checkActive every second
func checkUsers(roomAndNames *RoomAndNames) {
	for {
		<-time.After(1 * time.Second)
		go checkActive(roomAndNames)
	}
}

func printUsers(w http.ResponseWriter, roomAndNames *RoomAndNames, r *http.Request) {
	for key, element := range roomAndNames.connectedDevs {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
}
