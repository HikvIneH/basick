package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// userReview represents database entity.
type userReview struct {
	ID           string `json:"id"`
	order_id     string `json:"orderid"`
	product_id   string `json:"productid"`
	user_id      string `json:"userid"`
	rating       string `json:"rating"`
	users_review string `json:"review"`
	created_at   string `json:"createdat"`
	updated_at   string `json:"updatedat"`
}

const port = ":8000"

var err error

// Check for errors
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var db *sql.DB

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "WELCOME")
}

// GetUserReview Returns a list of all database User_Review to the response.
func GetUserReviews(w http.ResponseWriter, r *http.Request) {
	rows, e := db.Query("SELECT * FROM user_review")
	checkErr(e)
	defer rows.Close()
	userReviews := make([]*userReview, 0)

	for rows.Next() {
		c := new(userReview)
		e := rows.Scan(&c.ID, &c.order_id, &c.product_id, &c.user_id,
			&c.rating, &c.users_review, &c.created_at, &c.updated_at)
		checkErr(e)
		userReviews = append(userReviews, c)
		if e = rows.Err(); e != nil {
			log.Fatal(e)
		}
	}
	json.NewEncoder(w).Encode(userReviews)
}

// GetUserReview Returns a single database matching given ID parameter.
func GetUserReview(w http.ResponseWriter, r *http.Request) {
	var userRev userReview
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, http.StatusText(500), 500)
	}
	row := db.QueryRow("SELECT * FROM user_review WHERE ID = ?", id)
	e := row.Scan(&userRev.ID, &userRev.order_id, &userRev.product_id, &userRev.user_id,
		&userRev.rating, &userRev.users_review, &userRev.created_at, &userRev.updated_at)
	checkErr(e)
	json.NewEncoder(w).Encode(userRev)
}

func main() {
	fmt.Println("first line works")
	// Establish a connection to MySQL.
	db, err = sql.Open("mysql", "test:test@tcp(db:3306)/user_review")
	checkErr(err)

	defer db.Close()

	err = db.Ping()
	checkErr(err)

	// Create routes
	r := mux.NewRouter()

	r.HandleFunc("/", Index)
	r.HandleFunc("/user-review", GetUserReviews).Methods("GET")
	r.HandleFunc("/user-review/{id}", GetUserReview).Methods("GET")

	log.Println("Server is up on " + port + " port")
	log.Fatal(http.ListenAndServe(port, r))
}
