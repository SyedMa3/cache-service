# cache-service

A full-fledged RPC service to get and set data to Redis cache using `Golang`

### Setup:
Run the following commands in terminal.
*       bash init/initialize_redis.sh
    This is for setting up redis with conf
*       go run server/server.go
    Server
*       go run client/connect.go client/raw.go client/user.go
    Client

### Note:
* Operating System used: `Pop!_OS 22.04 LTS`
*       GOPATH=$HOME/go