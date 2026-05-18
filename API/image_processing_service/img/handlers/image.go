package handlers

import (
	"IPS/img/models"
	"IPS/img/services"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ImageHandler struct {
	service services.ImageService

	OriginalPath string
}

func (i ImageHandler) New(path string, service services.ImageService) ImageHandler {
	return ImageHandler{service: service, OriginalPath: path + "/original/"}
}

func (i ImageHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		panic(err)
	}
	content, handler, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	dst, err := os.Create(pwd + i.OriginalPath + handler.Filename)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	err, _ = i.service.SaveImage(handler.Filename, i.OriginalPath)
	if err != nil {
		fmt.Errorf("error saving image: %v", err)
	}

	if _, err := io.Copy(dst, content); err != nil {
		panic(err)
	}

	//file, _ := os.Open(i.OriginalPath + handler.Filename)
}

func (i ImageHandler) TransformHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		imageID, _ := strconv.Atoi(r.PathValue("id"))

		retImageData, err := i.service.GetImageByID(imageID)
		if err != nil {
			fmt.Errorf("error getting image: %v", err)
		}
		imgData := retImageData.ImageData()

		value, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		r.Body.Close()

		request := models.TransformRequest{}
		err = json.Unmarshal(value, &request)
		if err != nil {
			log.Printf("JSON unmarshal error: %v", err)
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = imgData.RunAll(&request)

		if err != nil {
			fmt.Errorf("error running all: %v", err)
		}
		w.Write(imgData.Data)
	}
}

func (i ImageHandler) RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		imageID, _ := strconv.Atoi(r.PathValue("id"))
		format := r.PathValue("format")

		retImageData, err := i.service.GetImageByID(imageID)
		if err != nil {
			fmt.Errorf("error getting image: %v", err)
		}
		imgData := retImageData.ImageData()

		err = imgData.FormatChange(models.Format{FileExt: format})
		if err != nil {
			fmt.Errorf("error formatting image: %v", err)
		}

		w.Write(imgData.HttpResponse())
	}
}

type construct struct {
	PageNos int      `json:"page_nos"`
	Content []string `json:"content"`
}

func (i ImageHandler) RetrievePageHandler(w http.ResponseWriter, r *http.Request) {
	var cst []construct
	if r.Method == http.MethodGet {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		for pg := 1; pg <= page; pg++ {
			var list []string
			offset := (page - 1) * limit
			imageValue, err := i.service.GetImages(limit, offset)
			for _, item := range imageValue {
				list = append(list, item.Path+item.Filename)
			}
			cst = append(cst, construct{pg, list})
			if err != nil {
				fmt.Errorf("error getting images: %v", err)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(cst)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
