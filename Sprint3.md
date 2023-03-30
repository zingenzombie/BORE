Detail work you've completed in Sprint 3
    -Create different views for the home page, and room pages (triangle room, square room, circle room)
    -Changes in the User Interface to allow for tabs to different pages
    -Room Features
        -Refractored the entire main page’s front end UI and redesigned all of the CSS to be CSS flexbox compliant so that the web application can work on all resolutions 
        -Created components for each file exchange room with styled CSS so that routing can display those specific components through a boolean toggle
        -Added a styled table nested within each room to display data pulled from backend to each respective file room 
        -Implemented a functionality which allows for the user to upload a file which changes the icon of a file exchange room
        -Nested a room component within another UI routing function Sade created
        -Modified CSS to become more linear and readable and added comments all over codebase

    Backend
    -Having GET requests return JSON objects
    -File uploading
    -Sanitizing IPs

List frontend unit tests
    Cypress Tests
    DescriptionBox
        The component gets created successfully
        Button for About displays correct information
        The button should toggle the description
    Angular Testing Service (Jasmine)
        UsernameService
            Should set and get the username
            Should be created successfully
            Post request should go through
    UsernameComponent
        Should add new users to the user list
        Should be created successfully
        Should set submitted to true
        Should call postUsername method of UserService
        For getAllMembers()
            Should be able to split the string gotten from get method into a list
            Be able to handle empty string
            Be able to handle single item string
    AppComponent
        Should contain a link to the Home, Circle Room, Triangle Room, and Square Room page
        Html elements should be displayed correctly
        Navigation elements should be displayed 
        Should create the app successfully
        Should not display username when user is not logged in 

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

Show updated documentation for your backend API

joinRoom - Allows the user to join a specific room. They must provide a room name and password (if needed) in the request body as follows:

{ "name":"NAMEHERE", "password":"PASSWORDHERE" }

getRooms - Returns all rooms to the user. 
{“rooms”:“ROOM1,ROOM2,...”}

setName - Sets the user's name. They must provide a name in the request body as follows:

{ "name":"NAMEHERE" }

getName - Returns the user’s name.

{“name”:“NAMEHERE”}

checkIn - This should be called once a second and resets the user's activity counter so they are not removed for inactivity. In addition, it returns data about BORE (how many rooms there are, who is in them, and what file requests may be available).

getRoomMembers - Returns all usernames in the user’s room. 
{“roomUsers”:“USER1,USER2,...”}

createRoom - Allows the user to create a specific room. They must provide a room name and password (if needed) in the request body as follows:

{ "name":"NAMEHERE", "password":"PASSWORDHERE" }

getAllMembers - Returns all usernames.

{“users”:“USER1,USER2,...”}

upload - Uploads a file to the server.


