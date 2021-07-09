
### Starting the API Server
`go run code/api-server/cmd/api-server/main.go -nats "nats://127.0.0.1:4222"`
Starting NATS Rider API Server version 0.1.0
Listening for HTTP requests on 0.0.0.0:9090


### Rides Manager
`go run code/rides-manager/cmd/manager/main.go`
Starting NATS Rider Rides Manager version 0.1.0


### Driver Agent

`go run code/driver-agent/cmd/agent/main.go -type mini`
Starting NATS Rider Driver Agent version 0.1.0


### Request a Ride

`curl "http://127.0.0.1:9090/rides" -X POST -d '{"type": "mini"}'`
 {"driver_id":"mKFJOZzmawAzaxEKkjKRAP"}

### Port Numbers for microservices

- Registration Service - 9090
- Inspection Service - 9091
- Treatment Service - 9092
- Release Service - 9093
- Common Data Service -9094
