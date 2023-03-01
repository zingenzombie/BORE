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