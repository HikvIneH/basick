
# REST API with GOLANG, Gorilla Mux and Mysql
_____

## GO
Go Version:
- `GO1.12`
> This version of go is already supporting 
> the go module, so we can work outside 
> of the GOPATH src directory.

## SQL
The SQL Schema is on `schema.sql` as well as the trigger to hold back rating `trigger.sql`

## Run
```sh
$ go run main.go
```

## API Endpoints
- http://localhost:3000/user-reviews
	- `GET`: get all user_review records
- http://localhost:8000/user-review/{id}
	- `GET`: get user_review record by id
- http://localhost:3000/user-review
	- `POST`: create new user_review record
- http://localhost:3000/user-review?id={id}
	- `PUT`: update a user_review record by id
	- `DELETE`: delete user_review record by id