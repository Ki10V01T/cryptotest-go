# cryptotest-go
This is a very simple client-server application, that working as chat betweet client (which sending a requests) and server (which accepts the requests and sending some information to client)

Functionality:
- logging into terminal
- all new clients forking a parallel threads
- supports a few commands

Supported commands:
- @exit - for closing connection and exit from client utility
- @conf - for requesting a some config variables from the server

Launching steps:
1. In server folder, execute command: go build && ./server
2. In client folder, execute command: go build && ./client
