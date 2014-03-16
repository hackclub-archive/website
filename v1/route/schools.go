package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/coopernurse/gorp"
	"github.com/hackedu/backend/v1/model"
)

func GetSchools(db gorp.SqlExecutor, log *log.Logger) (int, string) {
	var schools []model.School
	_, err := db.Select(&schools, "SELECT * FROM School ORDER BY Id")
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, ""
	}

	json, err := json.Marshal(schools)
	if err != nil {
		log.Println(err, "Error marshaling schools to JSON")
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, string(json)
}
