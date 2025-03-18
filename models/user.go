package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user document in MongoDB
type User struct {
	ID                 primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	Email              string                 `bson:"email" json:"email"`
	EmployeeID         string                 `bson:"employeeId,omitempty" json:"employeeId"`
	Password           string                 `bson:"password,omitempty" json:"-"`
	Name               string                 `bson:"name" json:"name"`
	Designation        string                 `bson:"designation" json:"designation"`
	IsReset            bool                   `bson:"is_reset" json:"is_reset"`
	IsDefault          bool                   `bson:"is_default" json:"is_default"`
	RoleID             primitive.ObjectID     `bson:"roleId" json:"roleId"`
	AddressLine1       string                 `bson:"addressLine1,omitempty" json:"addressLine1,omitempty"`
	AddressLine2       string                 `bson:"addressLine2,omitempty" json:"addressLine2,omitempty"`
	City               string                 `bson:"city,omitempty" json:"city,omitempty"`
	State              string                 `bson:"state,omitempty" json:"state,omitempty"`
	Zipcode            string                 `bson:"zipcode,omitempty" json:"zipcode,omitempty"`
	Country            string                 `bson:"country,omitempty" json:"country,omitempty"`
	JoiningDate        *time.Time             `bson:"joining_date,omitempty" json:"joining_date,omitempty"`
	Bio                string                 `bson:"bio,omitempty" json:"bio,omitempty"`
	LinkedIn           string                 `bson:"linkedin,omitempty" json:"linkedin,omitempty"`
	GitHub             string                 `bson:"github,omitempty" json:"github,omitempty"`
	Adhaar             string                 `bson:"adhaar,omitempty" json:"adhaar,omitempty"`
	PAN                string                 `bson:"pan,omitempty" json:"pan,omitempty"`
	DOB                *time.Time             `bson:"dob,omitempty" json:"dob,omitempty"`
	PhoneNumber        string                 `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Image              string                 `bson:"image,omitempty" json:"image,omitempty"`
	ProjectIDs         []primitive.ObjectID   `bson:"projectId" json:"projectId"`
	EnabledAccess      map[string]interface{} `bson:"enabled_access" json:"enabled_access"`
	IsClient           bool                   `bson:"is_client" json:"is_client"`
	LocationID         primitive.ObjectID     `bson:"location" json:"location"`
	ClientID           primitive.ObjectID     `bson:"clientId" json:"clientId"`
	IsFirstLogin       bool                   `bson:"isFirstLogin" json:"isFirstLogin"`
	EmpStatus          string                 `bson:"empStatus" json:"empStatus"`
	ProjectStartDate   *time.Time             `bson:"pjctStartDate,omitempty" json:"pjctStartDate,omitempty"`
	ProjectEndDate     *time.Time             `bson:"pjctEndDate,omitempty" json:"pjctEndDate,omitempty"`
	RelieveDate        *time.Time             `bson:"reliveDate,omitempty" json:"reliveDate,omitempty"`
	Reason             string                 `bson:"reason,omitempty" json:"reason,omitempty"`
	ResumeLink         string                 `bson:"resumeLink,omitempty" json:"resumeLink,omitempty"`
	RemainingReminders map[string]int         `bson:"remainingReminders" json:"remainingReminders"`
	CreatedAt          time.Time              `bson:"createdAt,omitempty"` // Timestamp when created
	UpdatedAt          time.Time              `bson:"updatedAt,omitempty"` // Timestamp when updated
}

// HashPassword hashes the user's password before storing it in the database
func (u *User) HashPassword() error {
	if u.Password == "" {
		return errors.New("password cannot be empty")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// ComparePassword compares a given password with the hashed password stored in the database
func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

// ValidateAadhaar validates Aadhaar number format
func (u *User) ValidateAadhaar() error {
	if u.Adhaar != "" && len(u.Adhaar) != 12 {
		return errors.New("invalid Aadhaar number: must be 12 digits")
	}
	return nil
}

// ValidatePAN validates PAN number format
func (u *User) ValidatePAN() error {
	if u.PAN != "" {
		if len(u.PAN) != 10 {
			return errors.New("invalid PAN number: must be 10 characters")
		}
	}
	return nil
}
