// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type Disease struct {
	DiseaseID   int    `json:"diseaseID"`
	DiseaseName string `json:"diseaseName"`
}

type DiseaseInput struct {
	DiseaseName string `json:"diseaseName"`
}

type Mutation struct {
}

type Patient struct {
	PatientID             int     `json:"patientID"`
	Notes                 *string `json:"notes,omitempty"`
	BloodType             *string `json:"bloodType,omitempty"`
	EmergencyContactName  *string `json:"emergencyContactName,omitempty"`
	EmergencyContactPhone *string `json:"emergencyContactPhone,omitempty"`
	InsuranceProvider     *string `json:"insuranceProvider,omitempty"`
	InsurancePolicyNumber *string `json:"insurancePolicyNumber,omitempty"`
	CreatedAt             int     `json:"createdAt"`
	UpdatedAt             int     `json:"updatedAt"`
	FirstName             string  `json:"firstName"`
	LastName              string  `json:"lastName"`
	Email                 *string `json:"email,omitempty"`
	PhoneNumber           string  `json:"phoneNumber"`
	Address               *string `json:"address,omitempty"`
	DateOfBirth           int     `json:"dateOfBirth"`
	Gender                string  `json:"gender"`
	Status                string  `json:"status"`
	ProfilePicture        *string `json:"profilePicture,omitempty"`
}

type PatientDisease struct {
	PatientID int `json:"patientID"`
	DiseaseID int `json:"diseaseID"`
}

type PatientDiseaseInput struct {
	PatientID int `json:"patientID"`
	DiseaseID int `json:"diseaseID"`
}

type PatientInput struct {
	Notes                 *string `json:"notes,omitempty"`
	BloodType             string  `json:"bloodType"`
	EmergencyContactName  *string `json:"emergencyContactName,omitempty"`
	EmergencyContactPhone *string `json:"emergencyContactPhone,omitempty"`
	InsuranceProvider     *string `json:"insuranceProvider,omitempty"`
	InsurancePolicyNumber *string `json:"insurancePolicyNumber,omitempty"`
	FirstName             string  `json:"firstName"`
	LastName              string  `json:"lastName"`
	Email                 *string `json:"email,omitempty"`
	PhoneNumber           string  `json:"phoneNumber"`
	Address               *string `json:"address,omitempty"`
	DateOfBirth           int     `json:"dateOfBirth"`
	Gender                string  `json:"gender"`
	ProfilePicture        *string `json:"profilePicture,omitempty"`
}

type Prescription struct {
	PrescriptionID int     `json:"prescriptionID"`
	Medications    *string `json:"medications,omitempty"`
	VisitID        int     `json:"visitID"`
	CreatedAt      int     `json:"createdAt"`
	Instructions   *string `json:"instructions,omitempty"`
}

type PrescriptionInput struct {
	Medications  *string `json:"medications,omitempty"`
	VisitID      int     `json:"visitID"`
	Instructions *string `json:"instructions,omitempty"`
}

type PrescriptionItem struct {
	ID             int    `json:"id"`
	MedicationName string `json:"medicationName"`
	Dosage         string `json:"dosage"`
	Frequency      string `json:"frequency"`
	Duration       *int   `json:"duration,omitempty"`
	Refills        *int   `json:"refills,omitempty"`
}

type PrescriptionItemInput struct {
	MedicationName string `json:"medicationName"`
	Dosage         string `json:"dosage"`
	Frequency      string `json:"frequency"`
	Duration       *int   `json:"duration,omitempty"`
	Refills        *int   `json:"refills,omitempty"`
}

type Query struct {
}

type User struct {
	UserID         int     `json:"userID"`
	Username       string  `json:"username"`
	Password       string  `json:"password"`
	FirstName      string  `json:"firstName"`
	LastName       string  `json:"lastName"`
	Email          string  `json:"email"`
	PhoneNumber    *string `json:"phoneNumber,omitempty"`
	Role           string  `json:"role"`
	Address        *string `json:"address,omitempty"`
	DateOfBirth    *int    `json:"dateOfBirth,omitempty"`
	Gender         *string `json:"gender,omitempty"`
	DateJoined     int     `json:"dateJoined"`
	LastLogin      *int    `json:"lastLogin,omitempty"`
	Status         string  `json:"status"`
	ProfilePicture *string `json:"profilePicture,omitempty"`
	CreatedAt      int     `json:"createdAt"`
	UpdatedAt      int     `json:"updatedAt"`
}

type UserInput struct {
	Username       string  `json:"username"`
	Password       string  `json:"password"`
	FirstName      string  `json:"firstName"`
	LastName       string  `json:"lastName"`
	Email          string  `json:"email"`
	PhoneNumber    *string `json:"phoneNumber,omitempty"`
	Role           string  `json:"role"`
	Address        *string `json:"address,omitempty"`
	DateOfBirth    *int    `json:"dateOfBirth,omitempty"`
	Gender         *string `json:"gender,omitempty"`
	DateJoined     int     `json:"dateJoined"`
	LastLogin      *int    `json:"lastLogin,omitempty"`
	Status         string  `json:"status"`
	ProfilePicture *string `json:"profilePicture,omitempty"`
}

type Visit struct {
	VisitID        int     `json:"visitID"`
	PatientID      int     `json:"patientID"`
	PrescriptionID *int    `json:"prescriptionID,omitempty"`
	Details        *string `json:"details,omitempty"`
	Symptoms       *string `json:"symptoms,omitempty"`
	Diagnosis      *string `json:"diagnosis,omitempty"`
	Prescription   *string `json:"prescription,omitempty"`
	DoctorInquiry  *string `json:"doctorInquiry,omitempty"`
	Notes          *string `json:"notes,omitempty"`
	VisitDate      int     `json:"visitDate"`
}

type VisitInput struct {
	PatientID      int     `json:"patientID"`
	PrescriptionID *int    `json:"prescriptionID,omitempty"`
	Details        *string `json:"details,omitempty"`
	Symptoms       *string `json:"symptoms,omitempty"`
	Diagnosis      *string `json:"diagnosis,omitempty"`
	Prescription   *string `json:"prescription,omitempty"`
	DoctorInquiry  *string `json:"doctorInquiry,omitempty"`
	Notes          *string `json:"notes,omitempty"`
	VisitDate      int     `json:"visitDate"`
}
