# Handind-5

## Running the program

Download the code in a zip file.

Unzip the project

Navigate to the folder in the terminal and type 
`````go
go run .
``````

This will lead to the following ui to be displayed:
````bash
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select an option: 
  ▸ Start Server
    Start new Bidder
    Exit
`````

### Selecting $Start\:server$:
Will open up a server on one of the given ports.
It will display as following:
````bash
✔ Start Server
2024/12/04 22:00:54 Starting server...
2024/12/04 22:00:54 Server started at port :4000
````
or 
````bash
✔ Start Server
2024/12/04 22:00:55 Starting server...
2024/12/04 22:00:55 Error creating the server %v listen tcp :4000: bind: address already in use
2024/12/04 22:00:55 Server started at port :4001
````
for the second second server, the third server will have tried another port, before starting the server.

### Selecting $Start\:new\:Bidder$:
Shows the following prompt:
````
✔ Enter Name: █
````
This is where you can enter the desired username of the client.

Then the following prompt is shown:
`````
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select an option: 
  ▸ Bid
    Result
    Exit
`````
Selecting "Bid" then allows you to enter the desired bid you want into the followign prompt: 
````bash
✔ Bid
Enter bid: 
````
after entering your desired bid you will recieve responses from the active servers.
This will look as follows:
````bash
-- Response One
2024/12/04 22:21:27 Response from server at port localhost:4000: accepted

-- Response Two
2024/12/04 22:21:27 Response from server at port localhost:4000: rejected

-- Response Three
2024/12/04 22:21:27 Response from server at port localhost:4000: Auction Ended
````
Selecting result will display the result from each active server:
```bash
✔ Result
2024/12/04 22:26:16 Dialing localhost:4000
2024/12/04 22:26:16 Current status: Highest Bid:132 by: hello
2024/12/04 22:26:16 Dialing localhost:4001
2024/12/04 22:26:16 Current status: Highest Bid:132 by: hello
2024/12/04 22:26:16 Dialing localhost:4002
2024/12/04 22:26:16 Failed to join auction: rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing: dial tcp [::1]:4002: connect: connection refused"
2024/12/04 22:26:16 Not all responses are valid
```
as can be seen above, if one of the servers are down or the servers don't have the same values, it will state that not all responses are valid.

Selecting exit, will like the others, exit the program entirely.
```bash
✔ Exit
2024/12/04 22:27:55 Exiting...
```



### Selecting $Exit$:
will exit the program. and display the following:
```bash
✔ Exit
2024/11/27 10:56:07 Exiting...
```


