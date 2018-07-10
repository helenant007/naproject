package model

import (
	"log"
	"time"

	"github.com/helenant007/naproject/utils/database"
)

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
	db := database.GetDB()

	rows, err := db.Query(selectUsersQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name, &user.MSISDN, &user.Email, &user.BirthDate, &user.CreatedTime, &user.UpdateTime)
		if err != nil {
			log.Fatal(err)
		}
		user.UserAge = int(time.Since(user.BirthDate).Hours() / (float64(24 * 365))) // count age manually
		result = append(result, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}
