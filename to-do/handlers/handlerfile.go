package handlers

import (
	"net/http"
	"to-do/database"
	"to-do/utils"
)

func DisplayHandler(w http.ResponseWriter, r *http.Request) {
	db, _ := database.GormInit()
	lists := utils.RetrievingAllList(db)

	utils.RenderTemplate(w, "templates/home", lists)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {}
