GitHub Link: https://github.com/zingenzombie/BORE
Video Link: https://youtu.be/HQmp3ZwiFuk

Work Completed in Sprint 2:
Frontend:
-          Created a form to enter the username and a button to submit.
-          Add username to a list of usernames which is displayed.
-          Created unit tests for the components.
-          Created a Cypress test to test if the About button displays a description when clicked.
Backend:
Function to createRoom
Changed request handling (JSON)
getName function 
Unit Testing

Frontend Unit Tests:
Unit tests (using Jasmine)
-          DescriptionBoxComponent
o   The About button toggles whether the description is shown or not.
-          UserService
o   The service should set and retrieve the username successfully through the setUsername and getUsername functions.
o   The service should post the username when the postUsername function is called.
-          UsernameComponent
o   The addUser function adds a new user to the list.
o   The submit button sets the submitted flag to true.
o   The component is able to call the setUsername function of the UserService.
o   The component is able to call the postUsername function of the UserService.
Cypress tests
-          DescriptionBoxComponent
o   The button, when clicked, should toggle whether the description is shown or not.


Backend Unit Tests:
TestNewUser tests the newUser function by confirming that the connectedDevice with the IP is added to connectedDevs
TestSetName tests the setName function by checking that the name of the connectedDevice is updated in connectedDevs
TestJoinRoom tests the joinRoom function by confirming that the connectedDevice stores the joined room and that the Room stores the connectedDevice in connectedDevs
TestCheckIn tests the CheckIn function by ensuring the connectedDevice’s active bool is set to true
TestCreateRoom tests the CreateRoom function by confirming that a new room is added to rooms with the correct name and password

Backend Documentation:

DOCUMENTATION:

BORE consists of a series of functions that can be called over http. The functions and their usage are as follows:

joinRoom - Allows the user to join a specific room. They must provide a room name and password (if needed) in the request body as follows:

{
    "name":"NAMEHERE",
    "password":"PASSWORDHERE"
}


getRooms - Returns all rooms and their members to the user. Request body will be ignored.

setName - Sets the user's name. They must provide a name in the request body as follows:

{
    "name":"NAMEHERE"
}

checkIn - This should be called once a second and resets the user's activity counter so they are not removed for inactivity. In addition, it returns data about BORE (how many rooms there are, who is in them, and what file requests may be available).

getRoomMembers - Returns all other users in the user's room. Request body will be ignored.

createRoom - Allows the user to create a specific room. They must provide a room name and password (if needed) in the request body as follows:

{
    "name":"NAMEHERE",
    "password":"PASSWORDHERE"
}