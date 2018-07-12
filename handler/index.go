package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/helenant007/naproject/model"
	"github.com/helenant007/naproject/utils/nsq"
	utilRedis "github.com/helenant007/naproject/utils/redis"
	"github.com/helenant007/naproject/utils/render"
)

type IndexModel struct {
	Users        []*model.User
	VisitorCount int
	QueryParam   string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n\nRENDERING INDEX PAGE\n\n")

	filterName := r.FormValue("filterName")
	if filterName == "" {
		nsq.PublishData("naproject-hn-visitor-topic", "todo_count")
	}

	users, err := model.GetUsers(filterName)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	// get visitor count
	visitorCount, err := redis.Int(utilRedis.GET("naproject:helen:visitorcount"))
	if err != nil && err.Error() != "redigo: nil returned" {
		log.Fatal(err)
	}

	err = render.RenderTemplate(w, "index", IndexModel{
		Users:        users,
		VisitorCount: visitorCount,
		QueryParam:   filterName,
	})

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
}
