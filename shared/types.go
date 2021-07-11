package shared

// Location is represents the latitude and longitude pair.
type Location struct {
	// Latitude is the latitude of the user making the request.
	Latitude float64 `json:"lat,omitempty"`

	// Longitude is the longitude of the user making the request.
	Longitude float64 `json:"lng,omitempty"`
}

// Address of a person.
type Address struct {
	// Type is the type of agent that is requested.
	House string `json:"house,omitempty"`

	// Type is the type of agent that is requested.
	Street string `json:"street,omitempty"`

	// Type is the type of agent that is requested.
	City string `json:"city,omitempty"`

	// Type is the type of agent that is requested.
	State string `json:"state,omitempty"`

}

// DriverAgentRequest is the request sent to the driver.
type DriverAgentRequest struct {
	// Type is the type of agent that is requested.
	Type string `json:"type,omitempty"`

	// Location is the location of the user that is being
	// served the request.
	Location *Location `json:"location,omitempty"`

	// RequestID is the ID from the request.
	RequestID string `json:"request_id,omitempty"`
}

// RegistrationRequest is the request to register a patient.
type RegistrationRequest struct {
	// Full Name of the patient.
	FullName string `json:"full_name,omitempty"`

	// Address of the patient.
	Address string `json:"address,omitempty"`

	// National Identification Number of the patient.
	ID int `json:"id"`

	// Sexual orientation 
	Sex string `json:"sex,omitempty"`

	// Sexual orientation 
	Email string `json:"email,omitempty"`	

	// Sexual orientation 
	Phone int `json:"phone,omitempty"`	

	// Other details
	Remarks string `json:"remarks,omitempty"`

	// RequestID is the ID from the request.
	RequestID string `json:"request_id,omitempty"`
}

type RegistrationEvent struct {
	// ID of the patient
	ID int `json:"id"`

	// Token of the patient
	Token uint64 `json:"token"`
}

// RegistrationRequest is the request to register a patient.
type InspectionRequest struct {
	// National Identification Number of the patient.
	ID int `json:"id"`

	// Time the inspection was done.
	Time string `json:"time,omitempty"`

	// Observations from the inspection.
	Observations string `json:"observations,omitempty"`

	// Medication schedule 
	Medication string `json:"medication,omitempty"`

	// Tests to be carried out 
	Tests string `json:"tests,omitempty"`	

	// Special notes 
	Notes string `json:"notes,omitempty"`	

	// RequestID is the ID from the request.
	RequestID string `json:"request_id,omitempty"`
}

// TestRequest is the request to do a test on a patient.
type TestRequest struct {
	// National Identification Number of the patient.
	ID int `json:"id"`

	// Time the inspection was done.
	Time string `json:"time,omitempty"`

	// Observations from the inspection.
	TestName string `json:"test_name,omitempty"`

	// Test results 
	Results string `json:"results,omitempty"`

	// Status of the test 
	Status string `json:"status,omitempty"`	

	// Special notes 
	Notes string `json:"notes,omitempty"`	

	// RequestID is the ID from the request.
	RequestID string `json:"request_id,omitempty"`
}

// MedicationRequest is the request to report a medication instance on a patient.
type MedicationRequest struct {
	// National Identification Number of the patient.
	ID int `json:"id"`

	// Time the inspection was done.
	Time string `json:"time,omitempty"`

	// Details of the dose of medication.
	Dose string `json:"test_name,omitempty"`	

	// Special notes 
	Notes string `json:"notes,omitempty"`	

	// RequestID is the ID from the request.
	RequestID string `json:"request_id,omitempty"`
}

// MedicationRequest is the request to report a medication instance on a patient.
type DischargeRequest struct {
	// National Identification Number of the patient.
	ID int `json:"id"`

	// Time the inspection was done.
	Time string `json:"time,omitempty"`

	// Details of the dose of medication.
	State string `json:"test_name,omitempty"`	

	// Details of the dose of medication.
	PostMedication string `json:"post_medication,omitempty"`

	// Special notes 
	Notes string `json:"notes,omitempty"`	

	// Next visit(if any) date 
	NextVisit string `json:"next_visit,omitempty"`	

	// RequestID is the ID from the request.
	RequestID string `json:"request_id,omitempty"`
}

// AdmissionEvent is the request to admit a patient.
type AdmissionEvent struct {
	// National Identification Number of the patient.
	ID int `json:"id"`

	// Medication schedule 
	Time string `json:"time,omitempty"`	

	// Special notes 
	Notes string `json:"notes,omitempty"`	
}

// RegistrationRequest is the request to register a patient.
type ReleaseEvent struct {
	// National Identification Number of the patient.
	ID int `json:"id"`

	// Medication schedule 
	NextState string `json:"next_state,omitempty"`

	// Tests to be carried out 
	PostMedication string `json:"post_medication,omitempty"`	

	// Special notes 
	Notes string `json:"notes,omitempty"`	

	// RequestID is the ID from the request.
	RequestID string `json:"request_id,omitempty"`
}

// RegistrationRequest is the request to register a patient.
type InspectionEvent struct {
	// National Identification Number of the patient.
	ID int `json:"id"`

	// Medication schedule 
	Medication string `json:"medication,omitempty"`

	// Tests to be carried out 
	Tests string `json:"tests,omitempty"`	

	// Special notes 
	Notes string `json:"notes,omitempty"`	
}

// DriverAgentResponse is the response from the driver.
type DriverAgentResponse struct {
	// ID is the identifier of the driver that will accept
	// the request.
	ID string `json:"driver_id,omitempty"`

	// Error is included in case there was an error handling the
	// request.
	Error string `json:"error,omitempty"`
}
