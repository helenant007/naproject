package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"

	"github.com/helenant007/naproject/handler"
	"github.com/helenant007/naproject/utils/database"
	utilRedis "github.com/helenant007/naproject/utils/redis"
	"github.com/helenant007/naproject/utils/render"
)

func main() {
	visitorCount, err := redis.Int(utilRedis.GET("naproject:helen:visitorcount"))
	if err != nil && err.Error() != "redigo: nil returned" {
		log.Fatal(err)
	}
	fmt.Printf("FROM REDIS %d\n", visitorCount)

	_, err = utilRedis.INCR("naproject:helen:visitorcount")
	if err != nil {
		log.Fatal(err)
	}

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

	http.HandleFunc("/", handler.IndexHandler)

	fmt.Println("SERVING...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
