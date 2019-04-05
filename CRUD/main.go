package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// userReview represents database entity.
type userReview struct {
	ID          string `json:"idxx"`
	orderid     string `json:"orderid"`
	productid   string `json:"productid"`
	userid      string `json:"userid"`
	rating      string `json:"rating"`
	usersreview string `json:"review"`
	createdat   string `json:"createdat"`
	updatedat   string `json:"updatedat"`
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
	userReviews := make([]*userReview, 0)

	defer rows.Close()
	for rows.Next() {
		c := new(userReview)
		e := rows.Scan(&c.ID, &c.orderid, &c.productid, &c.userid,
			&c.rating, &c.usersreview, &c.createdat, &c.updatedat)
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
	e := row.Scan(&userRev.ID, &userRev.orderid, &userRev.productid, &userRev.userid,
		&userRev.rating, &userRev.usersreview, &userRev.createdat, &userRev.updatedat)
	checkErr(e)
	json.NewEncoder(w).Encode(userRev)
}

// CreateUserReview Post the database.
func CreateUserReview(w http.ResponseWriter, r *http.Request) {
	var userRev userReview

	if isRatingError(userRev.rating) {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&userRev)
	_, e := db.Exec("INSERT INTO user_review (orderid, productid, userid, rating, usersreview) VALUES (?,?,?,?,?)", userRev.orderid, userRev.productid, userRev.userid, userRev.rating, userRev.usersreview)
	checkErr(e)
	json.NewEncoder(w).Encode(userRev)
}

// UpdateUserReview Update user review (identified by parameter) from the database
func UpdateUserReview(w http.ResponseWriter, r *http.Request) {
	var userRev userReview
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, http.StatusText(500), 500)
	}
	row := db.QueryRow("SELECT * FROM user_review WHERE ID = ?", id)
	e := row.Scan(&userRev.ID, &userRev.orderid, &userRev.productid, &userRev.userid,
		&userRev.rating, &userRev.usersreview, &userRev.createdat, &userRev.updatedat)
	checkErr(e)
	_ = json.NewDecoder(r.Body).Decode(&userRev)
	_, err := db.Exec("UPDATE orderid=?, productid=?, userid=?, rating=?, usersreview=? WHERE ID = ?", userRev.orderid, userRev.productid, userRev.userid, userRev.rating, userRev.usersreview, id)
	checkErr(err)
	json.NewEncoder(w).Encode(userRev)
}

// DeleteUserReview Removes user review (identified by parameter) from the database.
func DeleteUserReview(w http.ResponseWriter, r *http.Request) {
	var userRev userReview
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, http.StatusText(500), 500)
	}
	_, e := db.Exec("DELETE FROM user_review WHERE ID = ?", id)
	checkErr(e)
	json.NewEncoder(w).Encode(userRev)
}

func main() {
	fmt.Println("first line works")
	// Establish a connection to MySQL.
	db, err = sql.Open("mysql", "test:test@/test_golang?charset=utf8")
	checkErr(err)

	defer db.Close()

	err = db.Ping()
	checkErr(err)

	// Create routes
	r := mux.NewRouter()

	r.HandleFunc("/", Index)
	r.HandleFunc("/user-review", GetUserReviews).Methods("GET")
	r.HandleFunc("/user-review/{id}", GetUserReview).Methods("GET")
	r.HandleFunc("/user-review/{id}", CreateUserReview).Methods("POST")
	r.HandleFunc("/user-review/{id}", UpdateUserReview).Methods("PUT")
	r.HandleFunc("/user-review/{id}", DeleteUserReview).Methods("DELETE")

	log.Println("Server is up on " + port + " port")
	log.Fatal(http.ListenAndServe(port, r))
}

func isRatingError(rating string) bool {
	realRating, e := strconv.ParseFloat(rating, 64)
	if e != nil {
		fmt.Print(e.Error())
	}
	return realRating < 1.0 || realRating > 5.0
}
