package utils

import (
	"IPS/img/models"
	"encoding/json"

	"gorm.io/datatypes"
)

func TransformJSONParser(t models.TransformRequest) datatypes.JSON {
	// Placeholder for JSON transformation logic

	jsonData, _ := json.Marshal(t)

	return datatypes.JSON(jsonData)
}
