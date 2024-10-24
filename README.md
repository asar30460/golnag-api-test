# Golang Gorilla WebSocket Use Case

In this scenario, there is one server with multiple clients. When one of the connected clients sends a message, all other clients can receive it. To achieve this, we can utilize Golang's channels and goroutines.

Let's say there are two clients connected to the server via WebSocket (e.g., A and B). The data flow is as follows:

- A send message via WS.
- The WebSocket reads the message from client A and sends it to the server's broadcast channel.
- The server broadcasts this message to all connected clients (including A).
- Each client has its own channel to receive data
- WS write message to client

There are 2 version in this repo

- General version that send string message
- JSON version that send structured data
