package handlers

import (
	"Unit-Converter/helper"
	"Unit-Converter/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var err error

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("...Home <==")
	form := models.Form{}
	if r.Method == http.MethodPost {
		fmt.Println("...Posting ==>")
		form.To = r.FormValue("to")
		form.From = r.FormValue("from")
		form.Type = r.FormValue("type")
		form.Value, _ = strconv.ParseFloat(r.FormValue("val"), 64)

		switch form.Type {
		case "length":
			form.Response, err = helper.ConvLength(form.Value, form.From, form.To)
			if err != nil {
				log.Panic(err)
			}
		case "temp":
			form.Response, err = helper.ConvTemp(form.Value, form.From, form.To)
			if err != nil {
				log.Panic(err)
			}
		case "weight":
			form.Response, err = helper.ConvWeight(form.Value, form.From, form.To)
			if err != nil {
				log.Panic(err)
			}
		}
		helper.RenderTemplate(w, "result", &form)
		return
	}
	helper.RenderTemplate(w, "home", &form)
}
