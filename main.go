package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/helenant007/naproject/model"
	"github.com/helenant007/naproject/utils/database"
)

func main() {
	// init database
	connStr := "postgres://tkpdtraining:1tZCjrIcYeR1uQblQz0gBlIFU@devel-postgre.tkpd/tokopedia-user?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	database.Init(db)

	http.HandleFunc("/", test)

	fmt.Println("SERVING...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Tokopedia Tower"))

	users, err := model.GetUsers()
	if err != nil {
		fmt.Printf("Error %s", err.Error())
	} else {
		res2B, _ := json.Marshal(users)
		fmt.Printf("%+v\n", string(res2B))
	}
}
