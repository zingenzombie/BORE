package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	fmt.Println("Hello, welcome to BORE, your file-sharing drill!")

	fmt.Println("Starting server...")

	fmt.Print(GetOutboundIP())
	fmt.Println(":3621")

	os.Mkdir("temp-files", 0755)

	/*
		fmt.Println("Starting front-end...")

		exec.Command("cd /BORE")
		exec.Command("ng serve")*/

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
	roomAndNames.rooms["star"] = &Room{make(map[string]*connectedDevice), true, "", "star"}
	os.Mkdir("temp-files/star", 0755)

	roomAndNames.rooms["square"] = &Room{make(map[string]*connectedDevice), true, "", "square"}
	os.Mkdir("temp-files/square", 0755)

	roomAndNames.rooms["circle"] = &Room{make(map[string]*connectedDevice), true, "", "circle"}
	os.Mkdir("temp-files/circle", 0755)

	roomAndNames.rooms["triangle"] = &Room{make(map[string]*connectedDevice), true, "", "triangle"}
	os.Mkdir("temp-files/triangle", 0755)
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
	isPersistant  bool
	password      string
	name          string
}

// Structure for the series of rooms that hold users.
type RoomAndNames struct {
	counter       int
	connectedDevs map[string]*connectedDevice
	rooms         map[string]*Room
}

// Responds to HTTP /count request.
func (ct *RoomAndNames) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	r.RemoteAddr = r.RemoteAddr[:strings.IndexByte(r.RemoteAddr, ':')]

	newUser(w, ct, r)
	checkIn(ct, r)

	switch r.URL.Path {
	case "/debug":
		debug(w, ct, r)
	case "/joinRoom":
		joinRoom(w, ct, r)
	case "/getRooms":
		printRooms(w, ct)
	case "/setName":
		setName(ct, r)
	case "/getName":
		getName(ct, r, w)
	case "/checkIn":
		checkIn(ct, r)
	case "/getRoomMembers":
		printRoomUsers(w, ct.connectedDevs[r.RemoteAddr].room)
	case "/getAllMembers":
		printAllUsers(w, ct)
	case "/createRoom":
		createRoom(ct, r, w)
	case "/upload":
		uploadFile(w, ct, r)
	}

	//fmt.Println("OMG HI" + ct.connectedDevs[r.RemoteAddr].name)

	if ct.connectedDevs[r.RemoteAddr].name == "" {
		fmt.Println("OMG HI", r.RemoteAddr)
	} else {
		fmt.Println("OMG HI", ct.connectedDevs[r.RemoteAddr].name)
	}
}

func newUser(w http.ResponseWriter, roomAndNames *RoomAndNames, r *http.Request) {

	if val, ok := roomAndNames.connectedDevs[r.RemoteAddr]; ok {
		fmt.Println(val.ip)
	} else {
		tmp := connectedDevice{"", "", true, nil, time.Now().UTC()}
		tmp.ip = r.RemoteAddr

		roomAndNames.connectedDevs[r.RemoteAddr] = &tmp
		//printUsers(w, roomAndNames)
	}
}

type name struct {
	Name string `json:"name"`
}

func setName(roomAndNames *RoomAndNames, r *http.Request) {

	var requestName name
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestName)
	if err != nil {
		panic(err)
	}

	fmt.Println(requestName.Name)

	roomAndNames.connectedDevs[r.RemoteAddr].name = requestName.Name
}

func getName(roomAndNames *RoomAndNames, r *http.Request, w http.ResponseWriter) {

	if roomAndNames.connectedDevs[r.RemoteAddr] != nil {
		n := name{
			Name: roomAndNames.connectedDevs[r.RemoteAddr].name,
		}

		jsonBytes, err := json.Marshal(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonBytes)
	}

}

type roomRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func uploadFile(w http.ResponseWriter, roomAndNames *RoomAndNames, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	if roomAndNames.connectedDevs[r.RemoteAddr].room == nil {
		fmt.Fprintln(w, "ERROR, user is not in a room!")
		return
	}

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-files/"+roomAndNames.connectedDevs[r.RemoteAddr].room.name, "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

// creating a new room
func createRoom(roomAndNames *RoomAndNames, r *http.Request, w http.ResponseWriter) {
	var requestData roomRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		panic(err)
	}

	if roomAndNames.rooms[requestData.Name] != nil {
		fmt.Fprintln(w, "A room with that name already exists!")
		return
	}

	roomAndNames.rooms[requestData.Name] = &Room{make(map[string]*connectedDevice), false, requestData.Password, requestData.Name}

	os.Mkdir("temp-files/"+requestData.Name, 0755)

	if roomAndNames.connectedDevs[r.RemoteAddr].room != nil {
		leaveRoom(roomAndNames.connectedDevs[r.RemoteAddr])
	}

	roomAndNames.rooms[requestData.Name].connectedDevs[r.RemoteAddr] = roomAndNames.connectedDevs[r.RemoteAddr]
	roomAndNames.connectedDevs[r.RemoteAddr].room = roomAndNames.rooms[requestData.Name]
}

func checkRooms(roomAndNames *RoomAndNames) {
	for key, element := range roomAndNames.rooms {
		if len(element.connectedDevs) == 0 && !element.isPersistant {
			delete(roomAndNames.rooms, key)
		}
	}
}

type requestJoinRoom struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// connect to a room
func joinRoom(w http.ResponseWriter, roomAndNames *RoomAndNames, r *http.Request) {
	var requestData requestJoinRoom
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		panic(err)
	}

	if roomAndNames.rooms[requestData.Name] == nil {
		fmt.Fprintln(w, "That room does not exist")
		return
	}

	if roomAndNames.rooms[requestData.Name].password != requestData.Password {
		fmt.Fprintln(w, "That password is incorrect")
		return
	}

	if roomAndNames.connectedDevs[r.RemoteAddr].room != nil {
		leaveRoom(roomAndNames.connectedDevs[r.RemoteAddr])
	}

	roomAndNames.rooms[requestData.Name].connectedDevs[r.RemoteAddr] = roomAndNames.connectedDevs[r.RemoteAddr]
	roomAndNames.connectedDevs[r.RemoteAddr].room = roomAndNames.rooms[requestData.Name]

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

func debug(w http.ResponseWriter, roomAndNames *RoomAndNames, r *http.Request) {
	fmt.Println("DEBUG TIME!!!")
	n := name{
		Name: "debug",
	}

	jsonBytes, err := json.Marshal(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonBytes)
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
		//checkActive(roomAndNames)
		checkRooms(roomAndNames)
	}
}

/*
func printUsers(w http.ResponseWriter, roomAndNames *RoomAndNames) {
	for key, element := range roomAndNames.connectedDevs {
		if element.name != "" {
			fmt.Fprintln(w, "Key:", element.name, "=>", "Element:", element)
		} else {
			fmt.Fprintln(w, "Key:", key, "=>", "Element:", element)
		}
	}
}
*/

type allRooms struct {
	Rooms string `json:"rooms"`
}

func printRooms(w http.ResponseWriter, roomAndNames *RoomAndNames) {
	delimiter := ","
	r := allRooms{
		Rooms: "",
	}

	count := 0
	for key := range roomAndNames.rooms {
		count++
		r.Rooms += key

		if count != len(roomAndNames.rooms) {
			r.Rooms += delimiter
		}
		//printRoomUsers(w, roomAndNames.rooms[key])
	}

	jsonBytes, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonBytes)
}

type allUsers struct {
	Users string `json:"users"`
}

func printAllUsers(w http.ResponseWriter, roomAndNames *RoomAndNames) {
	delimiter := ","
	a := allUsers{
		Users: "",
	}

	count := 0
	for key, element := range roomAndNames.connectedDevs {
		count++
		if element.name != "" {
			a.Users += element.name
		} else {
			a.Users += key
		}

		if count != len(roomAndNames.connectedDevs) {
			a.Users += delimiter
		}
	}

	jsonBytes, err := json.Marshal(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonBytes)

}

type allRoomUsers struct {
	RoomUsers string `json:"roomUsers"`
}

func printRoomUsers(w http.ResponseWriter, room *Room) {
	delimiter := ","

	if room == nil {
		fmt.Fprintln(w, "This user is not in a room or this room does not exist.")
		return
	}

	r := allRoomUsers{
		RoomUsers: "",
	}

	count := 0
	for key, element := range room.connectedDevs {
		count++
		if element.name != "" {
			r.RoomUsers += element.name
		} else {
			r.RoomUsers += key
		}

		if count != len(room.connectedDevs) {
			r.RoomUsers += delimiter
		}
	}

	jsonBytes, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonBytes)
}
