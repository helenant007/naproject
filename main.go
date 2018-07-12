package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/helenant007/naproject/handler"
	"github.com/helenant007/naproject/utils/database"
	"github.com/helenant007/naproject/utils/nsq"
	"github.com/helenant007/naproject/utils/render"
)

func main() {

	nsq.InitConsumer()

	// init database
	connStr := "postgres://tkpdtraining:1tZCjrIcYeR1uQblQz0gBlIFU@devel-postgre.tkpd/tokopedia-user?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	database.Init(db)

	// initialize template
	tpl, err := template.ParseGlob("view/*.html")
	if err != nil {
		log.Fatal(err)
	}
	render.Init(tpl)

	http.HandleFunc("/index", handler.IndexHandler)

	fmt.Println("SERVING...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
