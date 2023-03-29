package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strconv"
	"testing"
)

// checks that new connectedDevice with ip is added to connectedDevs
func TestNewUser(t *testing.T) {
	w := httptest.NewRecorder()

	roomAndNames := &RoomAndNames{
		connectedDevs: make(map[string]*connectedDevice),
	}

	reqBody := bytes.NewBufferString(`{""}`)
	req := httptest.NewRequest("POST", "/", reqBody)
	req.RemoteAddr = "1"

	newUser(w, roomAndNames, req)
	expectedIP := "1"
	actualIP := roomAndNames.connectedDevs["1"].ip

	if actualIP != expectedIP {
		t.Errorf("Expected ip to be %s, but got %s", expectedIP, actualIP)
	}
}

// checks that name is updated in connectedDevs
func TestSetName(t *testing.T) {
	roomAndNames := &RoomAndNames{
		connectedDevs: make(map[string]*connectedDevice),
	}

	roomAndNames.connectedDevs["1"] = &connectedDevice{name: "name"}

	reqBody := bytes.NewBufferString(`{"name": "New Device Name"}`)
	req := httptest.NewRequest("POST", "/setName", reqBody)
	req.RemoteAddr = "1"

	setName(roomAndNames, req)

	expectedName := "New Device Name"
	actualName := roomAndNames.connectedDevs["1"].name
	if actualName != expectedName {
		t.Errorf("Expected name to be %s, but got %s", expectedName, actualName)
	}
}

// checks that connectedDevice stores the room and that the room stores the connectedDevice
func TestJoinRoom(t *testing.T) {
	w := httptest.NewRecorder()

	roomAndNames := &RoomAndNames{
		connectedDevs: make(map[string]*connectedDevice),
		rooms:         make(map[string]*Room),
	}

	roomAndNames.rooms["Room"] = &Room{password: "123", connectedDevs: make(map[string]*connectedDevice)}
	roomAndNames.connectedDevs["1"] = &connectedDevice{ip: "1", room: &Room{}}

	reqJoinRoom := requestJoinRoom{"Room", "123"}

	payload, _ := json.Marshal(reqJoinRoom)

	reqBody := bytes.NewBuffer(payload)
	req := httptest.NewRequest("POST", "/joinRoom", reqBody)
	req.RemoteAddr = "1"

	joinRoom(w, roomAndNames, req)

	expectedIP := "1"
	actualIP := roomAndNames.rooms["Room"].connectedDevs["1"].ip

	expectedRoomPassword := "123"
	actualRoomPassword := roomAndNames.connectedDevs["1"].room.password

	if actualIP != expectedIP {
		t.Errorf("Expected ip to be %s, but got %s", expectedIP, actualIP)
	}

	if actualRoomPassword != expectedRoomPassword {
		t.Errorf("Expected the room password to be %s, but got %s", expectedRoomPassword, actualRoomPassword)
	}

}

// checks that active bool of connectedDevice is set to true
func TestCheckIn(t *testing.T) {
	roomAndNames := &RoomAndNames{
		connectedDevs: make(map[string]*connectedDevice),
	}

	roomAndNames.connectedDevs["1"] = &connectedDevice{active: false}

	reqBody := bytes.NewBufferString(`{""}`)
	req := httptest.NewRequest("POST", "/checkIn", reqBody)
	req.RemoteAddr = "1"

	checkIn(roomAndNames, req)

	expectedBool := true
	actualBool := roomAndNames.connectedDevs["1"].active

	if actualBool != expectedBool {
		t.Errorf("Expected active to be %s, but got %s", strconv.FormatBool(expectedBool), strconv.FormatBool(actualBool))
	}

}

// checks that room is added to rooms
func TestCreateRoom(t *testing.T) {
	roomAndNames := &RoomAndNames{
		rooms:         make(map[string]*Room),
		connectedDevs: make(map[string]*connectedDevice),
	}

	roomAndNames.connectedDevs["1"] = &connectedDevice{}

	w := httptest.NewRecorder()

	reqJoinRoom := requestJoinRoom{"Room", "123"}

	payload, _ := json.Marshal(reqJoinRoom)

	reqBody := bytes.NewBuffer(payload)
	req := httptest.NewRequest("POST", "/createRoom", reqBody)
	req.RemoteAddr = "1"

	createRoom(roomAndNames, req, w)

	expectedPassword := "123"
	actualPassword := roomAndNames.rooms["Room"].password

	if actualPassword != expectedPassword {
		t.Errorf("Expected the room password to be %s, but got %s", expectedPassword, actualPassword)
	}

}

// tests that the names of all rooms are returned in json string
func TestGetRoooms(t *testing.T) {
	roomAndNames := &RoomAndNames{
		rooms:         make(map[string]*Room),
		connectedDevs: make(map[string]*connectedDevice),
	}

	roomAndNames.rooms["star"] = &Room{}
	roomAndNames.rooms["triangle"] = &Room{}

	w := httptest.NewRecorder()

	printRooms(w, roomAndNames)

	var requestData = allRooms{}
	decoder := json.NewDecoder(w.Body)
	decoder.Decode(&requestData)

	expectedResponse1 := allRooms{Rooms: "star,triangle"}
	expectedResponse2 := allRooms{Rooms: "triangle,star"}
	actualResponse := requestData

	if expectedResponse1 != actualResponse && expectedResponse2 != actualResponse {
		t.Errorf("error")
	}

}

// tests that name of device with ip is returned
func TestGetName(t *testing.T) {
	roomAndNames := &RoomAndNames{
		rooms:         make(map[string]*Room),
		connectedDevs: make(map[string]*connectedDevice),
	}

	roomAndNames.connectedDevs["1"] = &connectedDevice{name: "name"}

	reqBody := bytes.NewBufferString(`{""}`)
	req := httptest.NewRequest("GET", "/getName", reqBody)
	req.RemoteAddr = "1"

	w := httptest.NewRecorder()

	getName(roomAndNames, req, w)

	var requestData = name{}
	decoder := json.NewDecoder(w.Body)
	decoder.Decode(&requestData)

	expectedResponse := name{Name: "name"}
	actualResponse := requestData

	if expectedResponse != actualResponse {
		t.Errorf("error")
	}
}

// tests that all devices in room are returned
func TestGetRoomMembers(t *testing.T) {
	room := Room{connectedDevs: make(map[string]*connectedDevice)}

	room.connectedDevs["1"] = &connectedDevice{name: "joe"}
	room.connectedDevs["2"] = &connectedDevice{name: "kate"}

	w := httptest.NewRecorder()

	printRoomUsers(w, &room)

	var requestData = allRoomUsers{}
	decoder := json.NewDecoder(w.Body)
	decoder.Decode(&requestData)

	expectedResponse1 := allRoomUsers{RoomUsers: "joe,kate"}
	expectedResponse2 := allRoomUsers{RoomUsers: "kate,joe"}
	actualResponse := requestData

	if expectedResponse1 != actualResponse && expectedResponse2 != actualResponse {
		t.Errorf("error")
	}
}

// tests that all users are returned
func TestGetAllMembers(t *testing.T) {
	roomAndNames := &RoomAndNames{
		rooms:         make(map[string]*Room),
		connectedDevs: make(map[string]*connectedDevice),
	}

	roomAndNames.connectedDevs["1"] = &connectedDevice{name: "joe"}
	roomAndNames.connectedDevs["2"] = &connectedDevice{name: "kate"}

	w := httptest.NewRecorder()

	printAllUsers(w, roomAndNames)

	var requestData = allUsers{}
	decoder := json.NewDecoder(w.Body)
	decoder.Decode(&requestData)

	expectedResponse1 := allUsers{Users: "joe,kate"}
	expectedResponse2 := allUsers{Users: "kate,joe"}
	actualResponse := requestData

	if expectedResponse1 != actualResponse && expectedResponse2 != actualResponse {
		t.Errorf("error")
	}
}

/*func TestUpload(t *testing.T) {
	roomAndNames := &RoomAndNames{
		rooms:         make(map[string]*Room),
		connectedDevs: make(map[string]*connectedDevice),
	}

	roomAndNames.connectedDevs["1"] = &connectedDevice{room : &Room{name: "star"}}

	pr, pw := io.Pipe()

	writer := multipart.NewWriter(pw)

	go func() {
		defer writer.Close()
		// We create the form data field 'fileupload'
		// which returns another writer to write the actual file
		_, err := writer.CreateFormFile("fileupload", "someimg.png")
		if err != nil {
			t.Error(err)
		}

	}()

	req := httptest.NewRequest("POST", "/upload", pr)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	response := httptest.NewRecorder()

	uploadFile(response, roomAndNames, req)

	t.Log("It should respond with an HTTP status code of 200")
	if response.Code != 200 {
		t.Errorf("Error")
	}
	t.Log("It should create a file named 'someimg.png' in uploads folder")
	if _, err := os.Stat("./uploads/someimg.png"); os.IsNotExist(err) {
		t.Error("Expected file ./uploads/someimg.png' to exist")
	}

}*/
