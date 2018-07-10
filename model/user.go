package model

import "time"

type User struct {
	ID          string
	Name        string
	MSISDN      string
	Email       string
	BirthDate   time.Time
	CreatedTime time.Time
	UpdateTime  time.Time
	UserAge     int
}

// GetUsers : to get User data from DB
func GetUsers() ([]*User, error) {
	result := []*User{}

	return result, nil
}
