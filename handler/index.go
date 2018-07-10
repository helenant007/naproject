package handler

import (
	"fmt"
	"net/http"

	"github.com/helenant007/naproject/model"
	"github.com/helenant007/naproject/utils/render"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetUsers()
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	err = render.RenderTemplate(w, "index", users)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
}
