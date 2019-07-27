package models

import "time"

// Student - struc that models the student database table
type Student struct {
	ID               uint      `gorm:"primary_key"`
	Firstname        string    `gorm:"size:100" json:"firstname"`
	Lastname         string    `gorm:"size:100" json:"lastname"`
	Gender           string    `gorm:"size:10" json:"gender"`
	DateOfBirth      time.Time `gorm:"size:10" json:"dateofbirth"`
	ParentOrGuardian string    `gorm:"size:200" json:"parentorguardian"`
	Address          string    `gorm:"type:varchar(500)" json:"address"`
	ContactNumber    string    `gorm:"size:16" json:"contactnumber"`
}
