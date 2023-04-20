Detail work you've completed in Sprint 4
Frontend:
    -Button to list all users logged in
    -Three private password-protected file sharing rooms
    -Place to upload files on the home page
    -List of all uploaded files with time-stamp at the time they were uploaded
    -Number of files uploaded displayed 
    -Room display added to home page
Backend:
	-file download
	-returning user’s room


Frontend unit and Cypress tests
Cypress Tests
    DescriptionBox
        -The component gets created successfully
        -Button for About displays correct information
        -The button should toggle the description
Angular Testing Service (Jasmine)
    -UsernameService
        -Should set and get the username
        -Should be created successfully
        -Post request should go through
    -UsernameComponent
        -Should add new users to the user list
        -Should be created successfully
        -Should set submitted to true
        -Should call postUsername method of UserService
        -For getAllMembers()
            -Should be able to split the string gotten from get method into a list
            -Be able to handle empty string
            -Be able to handle single item string
    -AppComponent
        -Should contain a link to the Home, Circle Room, Triangle Room, and Square Room page
        -Html elements should be displayed correctly
        -Navigation elements should be displayed 
        -Should create the app successfully
        -Should not display username when user is not logged in 
    -HomePageComponent (sprint 4 tests)
        -Form should be rendered with upload button to upload files
        -The input text element to display the file name should be displayed
        -The input text element should be of the type “file”
        -The file upload counter should be increased after upload
        -The form element should be set to the correct…
            -Encoding type
            -Method
            -Action Url (for sending a post request to the files after upload)
        -The submit button element should be displayed and contain the correct text
    -SampleRoomPageComponent (sprint 4 tests)
        -Should set the submitted flag to true when submitted
        -If the password is incorrect, should set the corresponding flags


List backend unit tests
    -TestNewUser tests the newUser function by confirming that the connectedDevice with the IP is added to connectedDevs
    -TestSetName tests the setName function by checking that the name of the connectedDevice is updated in connectedDevs
    -TestJoinRoom tests the joinRoom function by confirming that the connectedDevice stores the joined room and that the Room stores the connectedDevice in connectedDevs
    -TestCheckIn tests the CheckIn function by ensuring the connectedDevice’s active bool is set to true
    -TestCreateRoom tests the CreateRoom function by confirming that a new room is added to rooms with the correct name and password
    -TestGetRooms tests the getRooms function by confirming that the names of rooms are returned in a JSON object
    -TestGetName tests the getName function by confirming that the name of the user is returned in a JSON object
    -TestGetRoomMembers tests the getRoomMembers function by confirming that the names of all users in the same room as the requesting user are returned in a JSON object
    -TestGetAllMembers tests the getAllMembers function by confirming that the names of all users are returned in a JSON object
    -TestGetRoom verifies that that GetRoom returns the name of the room a user is in.



Backend Documentation 
/joinRoom - Allows the user to join a specific room. They must provide a room name and password (if needed) in the request body as follows:

{ "name":"NAMEHERE", "password":"PASSWORDHERE" }

/getRooms - Returns all rooms to the user. {“rooms”:“ROOM1,ROOM2,...”}

/getRoom - Returns the name of the room that the user is in. { "room":"ROOMNAME" }

/setName - Sets the user's name. They must provide a name in the request body as follows:

{ "name":"NAMEHERE" }

/getName - Returns the user’s name.

{“name”:“NAMEHERE”}

/checkIn - This should be called once a second and resets the user's activity counter so they are not removed for inactivity. In addition, it returns data about the user (what room they are in, who else is in it, their username, and what files are available.).

{
    "username": "",
    "room": "",
    "usersInRoom": "",
    "files": ""
}

/getRoomMembers - Returns all usernames in the user’s room. {“roomUsers”:“USER1,USER2,...”}

/createRoom - Allows the user to create a specific room. They must provide a room name and password (if needed) in the request body as follows:

{ "name":"NAMEHERE", "password":"PASSWORDHERE" }

/getAllMembers - Returns all usernames.

{“users”:“USER1,USER2,...”}

/upload - Uploads a file to the server.

/download - Downloads a file to the client.
