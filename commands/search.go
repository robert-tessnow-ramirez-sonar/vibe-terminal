package commands

import (
	"database/sql"
	"fmt"
	"net/http"
)

// Search handles HTTP search requests and queries the DB directly with user input.
func Search(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	// SQL injection: user input concatenated directly into query
	rows, err := db.Query("SELECT * FROM items WHERE name = '" + query + "'")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	fmt.Fprintf(w, "results: %v", rows)
}
