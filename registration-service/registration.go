package registration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
	"sync/atomic"
	"github.com/nats-io/nuid"
	"github.com/gorilla/mux"
	"example.com/nats-microservices-opd/shared"
)

const (
	Version = "0.1.0"
)

// Server is a component.
type Server struct {
	*shared.Component
}

var ops uint64

// Generate token number for patient
func generateTokenNumber(start uint64) uint64 {
  if start > 0 {
	ops = start
  }
  atomic.AddUint64(&ops, 1)
  return ops
}

// HandleRides processes requests to find available drivers in an area.
func (s *Server) HandleToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var request *shared.DriverAgentRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Tag the request with an ID for tracing in the logs.
	request.RequestID = nuid.Next()
	req, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	nc := s.NATS()

	// Find a driver available to help with the request.
	log.Printf("requestID:%s - Finding available driver for request: %s\n", request.RequestID, string(body))
	msg, err := nc.Request("drivers.find", req, 5*time.Second)
	if err != nil {
		log.Printf("requestID:%s - Gave up finding available driver for request\n", request.RequestID)
		http.Error(w, "Request timeout", http.StatusRequestTimeout)
		return
	}
	log.Printf("requestID:%s - Response: %s\n", request.RequestID, string(msg.Data))

	var resp *shared.DriverAgentResponse
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if resp.Error != "" {
		http.Error(w, resp.Error, http.StatusServiceUnavailable)
		return
	}

	log.Printf("requestID:%s - Driver with ID %s is available to handle the request", request.RequestID, resp.ID)
	fmt.Fprintf(w, string(msg.Data))
}

// HandleRegister processes patient registration requests.
func (s *Server) HandleRegister(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var registration *shared.RegistrationRequest
	err = json.Unmarshal(body, &registration)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Tag the request with an ID for tracing in the logs.
	registration.RequestID = nuid.Next()
	fmt.Println(registration)
	// req, err := json.Marshal(registration)
	// if err != nil {
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	nc := s.NATS()

	//var registration_event shared.RegistrationEvent
	tokenNo := generateTokenNumber(0)
	registration_event := shared.RegistrationEvent{registration.ID, tokenNo}
	reg_event, err := json.Marshal(registration_event)

	if err != nil {
		log.Fatal(err)
		return
	}
	// Publishing the message to NATS Server
	nc.Publish("patient.register", reg_event)



	// // Find a driver available to help with the request.
	// log.Printf("requestID:%s - Finding available driver for request: %s\n", request.RequestID, string(body))
	// msg, err := nc.Request("drivers.find", req, 5*time.Second)
	// if err != nil {
	// 	log.Printf("requestID:%s - Gave up finding available driver for request\n", request.RequestID)
	// 	http.Error(w, "Request timeout", http.StatusRequestTimeout)
	// 	return
	// }
	// log.Printf("requestID:%s - Response: %s\n", request.RequestID, string(msg.Data))

	// var resp *shared.DriverAgentResponse
	// err = json.Unmarshal(msg.Data, &resp)
	// if err != nil {
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	// if resp.Error != "" {
	// 	http.Error(w, resp.Error, http.StatusServiceUnavailable)
	// 	return
	// }

	// log.Printf("requestID:%s - Driver with ID %s is available to handle the request", request.RequestID, resp.ID)
	// fmt.Fprintf(w, string(msg.Data))
}

// HandleRides processes requests to find available drivers in an area.
func (s *Server) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var request *shared.DriverAgentRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Tag the request with an ID for tracing in the logs.
	request.RequestID = nuid.Next()
	req, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	nc := s.NATS()

	// Find a driver available to help with the request.
	log.Printf("requestID:%s - Finding available driver for request: %s\n", request.RequestID, string(body))
	msg, err := nc.Request("drivers.find", req, 5*time.Second)
	if err != nil {
		log.Printf("requestID:%s - Gave up finding available driver for request\n", request.RequestID)
		http.Error(w, "Request timeout", http.StatusRequestTimeout)
		return
	}
	log.Printf("requestID:%s - Response: %s\n", request.RequestID, string(msg.Data))

	var resp *shared.DriverAgentResponse
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if resp.Error != "" {
		http.Error(w, resp.Error, http.StatusServiceUnavailable)
		return
	}

	log.Printf("requestID:%s - Driver with ID %s is available to handle the request", request.RequestID, resp.ID)
	fmt.Fprintf(w, string(msg.Data))
}

// HandleRides processes requests to find available drivers in an area.
func (s *Server) HandleView(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var request *shared.DriverAgentRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Tag the request with an ID for tracing in the logs.
	request.RequestID = nuid.Next()
	req, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	nc := s.NATS()

	// Find a driver available to help with the request.
	log.Printf("requestID:%s - Finding available driver for request: %s\n", request.RequestID, string(body))
	msg, err := nc.Request("drivers.find", req, 5*time.Second)
	if err != nil {
		log.Printf("requestID:%s - Gave up finding available driver for request\n", request.RequestID)
		http.Error(w, "Request timeout", http.StatusRequestTimeout)
		return
	}
	log.Printf("requestID:%s - Response: %s\n", request.RequestID, string(msg.Data))

	var resp *shared.DriverAgentResponse
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if resp.Error != "" {
		http.Error(w, resp.Error, http.StatusServiceUnavailable)
		return
	}

	log.Printf("requestID:%s - Driver with ID %s is available to handle the request", request.RequestID, resp.ID)
	fmt.Fprintf(w, string(msg.Data))
}

func (s *Server) HandleHomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("Registration Service v%s\n", Version))
}

// ListenAndServe takes the network address and port that
// the HTTP server should bind to and starts it.
func (s *Server) ListenAndServe(addr string) error {

	r := mux.NewRouter()
	router := r.PathPrefix("/opd/patient/").Subrouter()

	// Handle base path requests
	// GET /opd/patient
	router.HandleFunc("/", s.HandleHomeLink)
	// Handle registration requests
	// POST /opd/patient/register
	router.HandleFunc("/register", s.HandleRegister).Methods("POST")

	// Handle update requests
	// PUT /opd/patient/update/{id}
	router.HandleFunc("/update/{id}", s.HandleUpdate).Methods("PUT")

	// Handle view requests
	// GET /opd/patient/view/{id}
	router.HandleFunc("/view/{id}", s.HandleView).Methods("GET")

	// Handle token requests
	// GET /opd/patient/token
	router.HandleFunc("/token", s.HandleToken).Methods("GET")

	//router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	//log.Fatal(http.ListenAndServe(":8080", router))

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	srv := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go srv.Serve(l)

	return nil
}
