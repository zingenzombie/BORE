package main

import (
	"fmt"
	"io/ioutil"
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

	mux := http.NewServeMux()

	th := &RoomAndNames{counter: 0}

	th.connectedDevs = make(map[string]*connectedDevice)
	th.rooms = make(map[string]*Room)

	makeRooms(th)
	go checkUsers(th)

	mux.HandleFunc("/", th.ServeHTTP)

	http.ListenAndServe(":3621", corsHandler(mux))

}

func corsHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// If the request method is OPTIONS, send a 200 status code and return
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the original handler function
		h.ServeHTTP(w, r)
	})
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
	room        *Room
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
	checkIn(ct, r)

	switch r.URL.Path {
	case "/debug":
		debug(ct, r)
	case "/joinRoom":
		joinRoom(ct, r)
	case "/getRooms":
		printRooms(w, ct)
	case "/setName":
		setName(ct, r)
	case "/checkIn":
		checkIn(ct, r)
	case "/getRoomMembers":
		printRoomUsers(w, ct.connectedDevs[r.RemoteAddr].room)
	}

	if ct.connectedDevs[r.RemoteAddr].name == "" {
		fmt.Fprintln(w, "Hello ", r.RemoteAddr)
	} else {
		fmt.Fprintln(w, "Hello ", ct.connectedDevs[r.RemoteAddr].name)
	}
}

func newUser(w http.ResponseWriter, roomAndNames *RoomAndNames, r *http.Request) {

	if val, ok := roomAndNames.connectedDevs[r.RemoteAddr]; ok {
		fmt.Println(val.ip)
	} else {
		tmp := connectedDevice{"", "", true, nil, time.Now().UTC()}
		tmp.ip = r.RemoteAddr

		roomAndNames.connectedDevs[r.RemoteAddr] = &tmp
		printUsers(w, roomAndNames)
	}
}

func setName(roomAndNames *RoomAndNames, r *http.Request) {
	body, error := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if error != nil {
		fmt.Println(error)
	}

	roomAndNames.connectedDevs[r.RemoteAddr].name = string(body)
}

/*
// creating a new room
func createRoom(roomAndNames *RoomAndNames) {
	roomAndNames.rooms = append(roomAndNames.rooms, Room{})
}*/

// connect to a room
func joinRoom(roomAndNames *RoomAndNames, r *http.Request) {
	body, error := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if error != nil {
		fmt.Println(error)
	}

	if roomAndNames.connectedDevs[r.RemoteAddr].room != nil {
		leaveRoom(roomAndNames.connectedDevs[r.RemoteAddr])
	}

	roomAndNames.rooms[string(body)].connectedDevs[r.RemoteAddr] = roomAndNames.connectedDevs[r.RemoteAddr]
	roomAndNames.connectedDevs[r.RemoteAddr].room = roomAndNames.rooms[string(body)]

}

// leave room
func leaveRoom(cd *connectedDevice) {

	if cd.room == nil {
		fmt.Println("User " + cd.ip + " does not exist within a room!")
		return
	}

	for key := range cd.room.connectedDevs {
		if key == string(cd.ip) {
			delete(cd.room.connectedDevs, key)
			cd.room = nil
			return
		}
	}
}

func debug(roomAndNames *RoomAndNames, r *http.Request) {
	fmt.Println("DEBUG TIME!!!")
}

// setting time of last checkin
func checkIn(roomAndNames *RoomAndNames, r *http.Request) {
	fmt.Println(r.RemoteAddr, " checked in.")
	roomAndNames.connectedDevs[r.RemoteAddr].lastCheckIn = time.Now().UTC()
	roomAndNames.connectedDevs[r.RemoteAddr].active = true
}

// looping through connected devices
func checkActive(roomAndNames *RoomAndNames) {
	for key, element := range roomAndNames.connectedDevs {
		// device inactive if no checkin in 10 seconds
		if time.Now().UTC().Sub(element.lastCheckIn) > (10 * time.Second) {
			fmt.Println(element.ip, " checking out.")
			leaveRoom(element)
			delete(roomAndNames.connectedDevs, key)
		}
	}
}

// calling checkActive every second
func checkUsers(roomAndNames *RoomAndNames) {
	for {
		<-time.After(1 * time.Second)
		checkActive(roomAndNames)
	}
}

func printUsers(w http.ResponseWriter, roomAndNames *RoomAndNames) {
	for key, element := range roomAndNames.connectedDevs {
		if element.name != "" {
			fmt.Fprintln(w, "Key:", element.name, "=>", "Element:", element)
		} else {
			fmt.Fprintln(w, "Key:", key, "=>", "Element:", element)
		}
	}
}

func printRooms(w http.ResponseWriter, roomAndNames *RoomAndNames) {
	for key := range roomAndNames.rooms {
		fmt.Fprintln(w, key, " room:")
		printRoomUsers(w, roomAndNames.rooms[key])
	}
}

func printRoomUsers(w http.ResponseWriter, room *Room) {
	if room == nil {
		fmt.Fprintln(w, "This user is not in a room or this room does not exist.")
	}

	for key, element := range room.connectedDevs {
		if element.name != "" {
			fmt.Fprintln(w, "Key:", element.name, "=>", "Element:", element)
		} else {
			fmt.Fprintln(w, "Key:", key, "=>", "Element:", element)
		}
	}
}
