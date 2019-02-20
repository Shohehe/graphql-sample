package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Shohehe/graphql-sample/gateway/mygraphql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(r.Body)
	query := bufBody.String()

	result := mygraphql.ExecuteQuery(query, mygraphql.Schema)
	json.NewEncoder(w).Encode(result)
}
