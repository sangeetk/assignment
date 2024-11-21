package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"assignment/matrix"
)

// Flatten - Return the matrix as a 1 line string, with values separated by commas.
//
//	Input:
//	    1,2,3
//	    4,5,6
//	    7,8,9
//
//	Output:
//	    1,2,3,4,5,6,7,8,9
func Flatten(w http.ResponseWriter, r *http.Request) {

	// Read file from the request body.
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}
	defer file.Close()

	// Parse CSV records from the file
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	// Parse CSV and create a squre Matrix
	m, err := matrix.Parse(records)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, m.Flatten())
}
