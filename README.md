
### Starting the Registration Service
`cd registration-service`

`go run cmd/main.go -dbName opd_data -dbUser root -dbPassword Root@1985`

2021/07/11 13:56:12 Starting NATS Microservices OPD Sample - Registration Service version 0.1.0
2021/07/11 13:56:12 Listening for HTTP requests on 0.0.0.0:9090


### Starting the Inspection Service
`cd inspection-service`

`go run cmd/main.go -dbName opd_data -dbUser root -dbPassword Root@1985`

2021/07/11 13:56:18 Starting NATS Microservices OPD Sample - Inspection Service version 0.1.0
2021/07/11 13:56:18 Listening for HTTP requests on 0.0.0.0:9091


### Starting the Treatment Service
`cd treatment-service`

`go run cmd/main.go -dbName opd_data -dbUser root -dbPassword Root@1985`

2021/07/11 13:56:26 Starting NATS Microservices OPD Sample - Treatment Service version 0.1.0
2021/07/11 13:56:26 Listening for HTTP requests on 0.0.0.0:9092


### Starting the Release Service
`cd release-service`

`go run cmd/main.go -dbName opd_data -dbUser root -dbPassword Root@1985`

2021/07/11 13:54:25 Starting NATS Microservices OPD Sample - Release Service version 0.1.0
2021/07/11 13:54:25 Listening for HTTP requests on 0.0.0.0:9093

### Register a Patient
- Request
`curl "http://localhost:9090/opd/patient/register" -X POST -d '{"full_name":"fernando","address":"44, liyanagemulla, seeduwa","id":100, "sex":"male", "phone":222222222}'`
- Response
{"id":100,"token":1}

### View the Patient
- Request
`curl "http://localhost:9090/opd/patient/view/100"`
- Response
{"full_name":"fernando","address":"44, liyanagemulla, seeduwa","id":100,"sex":"male","phone":222222222}

### Update a Patient
- Request
`curl "http://localhost:9090/opd/patient/update" -X PUT -d '{"full_name":"fernando","address":"667/280/6, liyanagemulla, seeduwa","id":100, "sex":"male", "phone":222222222}'`
- Response
"Record for Patient updated sucessfully"

### Generate a Token
- Request
`curl "http://localhost:9090/opd/patient/token/100"`
- Response
{"id":100,"token":1}



### Port Numbers for microservices

- Registration Service - 9090
- Inspection Service - 9091
- Treatment Service - 9092
- Release Service - 9093
- Common Data Service -9094
