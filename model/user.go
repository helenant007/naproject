package model

import (
	"log"
	"strings"
	"time"

	"github.com/helenant007/naproject/utils/database"
)

type User struct {
	ID                string
	Name              string
	MSISDN            string
	Email             string
	BirthDate         time.Time
	BirthDateString   string
	CreatedTime       time.Time
	CreatedTimeString string
	UpdateTime        time.Time
	UpdateTimeString  string
	UserAge           int
}

// GetUsers : to get User data from DB
func GetUsers(nameWhere string) ([]*User, error) {
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

		if nameWhere != "" && !strings.Contains(user.Name, nameWhere) {
			continue
		}

		user.UserAge = int(time.Since(user.BirthDate).Hours() / (float64(24 * 365))) // count age manually
		dateFormat := "2006/01/02 03:04:05"
		user.BirthDateString = user.BirthDate.Format(dateFormat)
		user.CreatedTimeString = user.CreatedTime.Format(dateFormat)
		user.UpdateTimeString = user.UpdateTime.Format(dateFormat)

		result = append(result, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}
