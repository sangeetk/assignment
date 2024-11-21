package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"assignment/matrix"
)

// Muitiply - Return the product of the integers in the matrix
//
//	Input:
//	    1,2,3
//	    4,5,6
//	    7,8,9
//
//	Output:
//	    362880
func Multiply(w http.ResponseWriter, r *http.Request) {

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
	fmt.Fprint(w, m.Multiply())
}
