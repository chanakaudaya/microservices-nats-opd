module example.com/nats-microservices-opd/common-data-service

go 1.16

replace example.com/nats-microservices-opd/shared => ../shared

require (
	example.com/nats-microservices-opd/shared v0.0.0-00010101000000-000000000000
	github.com/nats-io/nats.go v1.11.1-0.20210623165838-4b75fc59ae30
	github.com/nats-io/nuid v1.0.1
)
