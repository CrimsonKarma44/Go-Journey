package views

import (
	"UrlShortner/models"
	"UrlShortner/utils"
	"log"
	"net/http"
)

func DisplayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		url := r.FormValue("url")
		alias := r.FormValue("alias")
		saveUrl := models.Url{url, alias}
		err := utils.JsonUrlSaver(saveUrl)
		if err != nil {
			panic(err)
		}
	}
	urls := utils.JsonUrlReader()
	utils.RenderTemplate(w, "templates/home", urls)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	alias := r.URL.Path[len("/home/delete/"):]

	urls := utils.JsonUrlReader()
	for _, url := range urls {
		if url.Alias == alias {
			err := utils.DeleteJson(url)
			if err != nil {
				log.Println("error deleting url:", err)
				panic(err)
			}
		}
	}
	http.Redirect(w, r, "/home", http.StatusFound)
}

func CopyHandler(w http.ResponseWriter, r *http.Request) {
	alias := r.URL.Path[len("/home/copy/"):]
	var neededUrl string

	urls := utils.JsonUrlReader()
	for _, url := range urls {
		if url.Alias == alias {
			neededUrl = url.Url
		}
	}

	err := utils.CopyUrl(neededUrl)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/home", http.StatusFound)
}
