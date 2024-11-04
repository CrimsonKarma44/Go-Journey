package utils

import (
	"UrlShortner/models"
	"encoding/json"
	"log"
	"os"
)

func JsonUrlReader() []models.Url {
	var urls []models.Url
	file, err := os.ReadFile("storage/storage.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &urls)
	if err != nil {
		log.Fatal(err)
	}
	return urls
}

func JsonUrlSaver(new models.Url) error {
	content := JsonUrlReader()
	content = append(content, new)

	jsonType, err := json.Marshal(content)
	if err != nil {
		return err
	}
	// creates if not exists
	file, err := os.Create("storage/storage.json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonType)
	if err != nil {
		return err
	}
	return nil
}

func DeleteJson(url models.Url) error {
	var newUrls []models.Url
	urls := JsonUrlReader()
	for _, u := range urls {
		if url != u {
			newUrls = append(newUrls, u)
		}
	}
	jsonType, err := json.Marshal(newUrls)
	if err != nil {
		return err
	}
	file, err := os.Create("storage/storage.json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonType)
	if err != nil {
		return err
	}

	return nil
}
