Backend Documentation:

/joinRoom - Allows the user to join a specific room. They must provide a room name and password (if needed) in the request body as follows:

{ "name":"NAMEHERE", "password":"PASSWORDHERE" }

/getRooms - Returns all rooms to the user. 
{“rooms”:“ROOM1,ROOM2,...”}

/getRoom - Returns the name of the room that the user is in.
{ "room":"ROOMNAME" } 

/setName - Sets the user's name. They must provide a name in the request body as follows:

{ "name":"NAMEHERE" }

/getName - Returns the user’s name.

{“name”:“NAMEHERE”}

/checkIn - This should be called once a second and resets the user's activity counter so they are not removed for inactivity. In addition, it returns data about BORE (how many rooms there are, who is in them, and what file requests may be available).

getRoomMembers - Returns all usernames in the user’s room. 
{“roomUsers”:“USER1,USER2,...”}

/createRoom - Allows the user to create a specific room. They must provide a room name and password (if needed) in the request body as follows:

{ "name":"NAMEHERE", "password":"PASSWORDHERE" }

/getAllMembers - Returns all usernames.

{“users”:“USER1,USER2,...”}

/upload - Uploads a file to the server.

/download - Downloads a file to the client.