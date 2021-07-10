package inspection

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"example.com/nats-microservices-opd/shared"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/nats-io/nuid"
	"github.com/nats-io/nats.go"
)

const (
	Version = "0.1.0"
)

// Server is a component.
type Server struct {
	*shared.Component
}


func dbConn()(db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Root@1985"
	dbName := "opd_data"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (s *Server) ListenRegisterEvents() error {
	nc := s.NATS()
	nc.Subscribe("patient.register", func(msg *nats.Msg) {
		var req *shared.RegistrationEvent
		err := json.Unmarshal(msg.Data, &req)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		
		log.Printf("New Patient Registration Event received for PatientID %d with Token  %d\n",
			req.ID, req.Token)

			// Insert data to the database
		db := dbConn()

		insForm, err := db.Prepare("INSERT INTO patient_registrations(id, token) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(req.ID, req.Token)
		//log.Println("INSERT: Name: " + name + " | City: " + city)
		
		defer db.Close()

	})

	return nil

}


// HandleRegister processes patient registration requests.
func (s *Server) HandleRecord(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var inspection *shared.InspectionRequest
	err = json.Unmarshal(body, &inspection)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Insert data to the database
	db := dbConn()

	insForm, err := db.Prepare("INSERT INTO inspection_details(id, time, observations, medication, tests, notes) VALUES(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(inspection.ID, inspection.Time, inspection.Observations, inspection.Medication, inspection.Tests, inspection.Notes)
	//log.Println("INSERT: Name: " + name + " | City: " + city)
    
    defer db.Close()

	// Tag the request with an ID for tracing in the logs.
	inspection.RequestID = nuid.Next()
	fmt.Println(inspection)

	// Publish event to the NATS server
	nc := s.NATS()

	//var registration_event shared.RegistrationEvent
	inspection_event := shared.InspectionEvent{inspection.ID, inspection.Medication, inspection.Tests, inspection.Notes}
	reg_event, err := json.Marshal(inspection_event)

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("requestID:%s - Publishing inspection event with patientID %d\n", inspection.RequestID, inspection.ID)
	// Publishing the message to NATS Server
	nc.Publish("patient.treatment", reg_event)

	json.NewEncoder(w).Encode(inspection_event)
}

// HandleView processes requests to view patient data.
func (s *Server) HandleHistory(w http.ResponseWriter, r *http.Request) {
	patientID := mux.Vars(r)["id"]
	// Insert data to the database
	db := dbConn()

	selDB, err := db.Query("SELECT * FROM patient_details WHERE ID=?", patientID)
    if err != nil {
        panic(err.Error())
    }

    registration := shared.RegistrationRequest{}
    for selDB.Next() {
        var id, phone int
        var full_name, address, sex, remarks string
        err = selDB.Scan(&id, &full_name, &address, &sex, &phone, &remarks)
        if err != nil {
            panic(err.Error())
        }
        registration.ID = id
        registration.FullName = full_name
        registration.Address = address
		registration.Sex = sex
		registration.Phone = phone
		registration.Remarks = remarks
    }

	fmt.Println(registration)
	json.NewEncoder(w).Encode(registration)
    defer db.Close()
}

func (s *Server) HandleHomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("Inspection Service v%s\n", Version))
}

// ListenAndServe takes the network address and port that
// the HTTP server should bind to and starts it.
func (s *Server) ListenAndServe(addr string) error {

	// Start listening to patient registration events
	s.ListenRegisterEvents()

	r := mux.NewRouter()
	router := r.PathPrefix("/opd/inspection/").Subrouter()

	// Handle base path requests
	// GET /opd/inspection
	router.HandleFunc("/", s.HandleHomeLink)
	
	// Handle inspection record requests
	// POST /opd/inspection/record/{id}
	router.HandleFunc("/record", s.HandleRecord).Methods("POST")

	// Handle history view requests
	// GET /opd/inspection/history/{id}
	router.HandleFunc("/history/{id}", s.HandleHistory).Methods("GET")

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