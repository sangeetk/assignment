package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"assignment/matrix"
)

func Sum(w http.ResponseWriter, r *http.Request) {

	// Parse CSV from the request body.
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	m, err := matrix.Parse(records)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, m.Sum())
}
