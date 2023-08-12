package db

import (
	"database/sql"
	"fmt"
	"movie-app/model"

	_ "github.com/lib/pq"
)

const (
	DB_USER             = "papayetoo"
	DB_PASSWORD         = "rhkdgus0322"
	DB_NAME             = "movieapp"
	select_users        = "SELECT name, email FROM users"
	insert_user         = "INSERT INTO USERS(NAME, EMAIL) VALUES ($1, $2)"
	delter_user_by_name = "DELETE FROM USERS WHERE NAME = $1"
)

func ConnectDatabase() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	return sql.Open("postgres", dbinfo)
}

func FindUsers() ([]*model.User, error) {
	db, err := ConnectDatabase()
	if err != nil {
		return []*model.User{}, err
	}
	defer db.Close()

	rows, err := db.Query(select_users)
	if err != nil {
		return []*model.User{}, err
	}
	defer rows.Close()

	users := []*model.User{}

	for rows.Next() {
		u := &model.User{}
		err := rows.Scan(&u.Name, &u.Email)
		if err != nil {
			return []*model.User{}, err
		}
		users = append(users, u)
	}
	return users, nil
}

func FindUserByEmail(email string) (*model.User, error) {
	db, err := ConnectDatabase()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT username, password, email FROM service_user")
	if err != nil {
		return nil, fmt.Errorf("failed to execute the query : %w", err)
	}
	defer rows.Close()

	u := &model.User{}
	found := false
	for rows.Next() {
		if err := rows.Scan(&u.Name, &u.Password, &u.Email); err != nil {
			return nil, err
		}
		found = true
	}
	if !found {
		return nil, fmt.Errorf("user with email '%s' is not found", email)
	}

	return u, nil
}

func InsertUser(db *sql.DB, user *model.User) (sql.Result, error) {
	return db.Exec(insert_user, user.Name, user.Email)
}

func DeleteUser(db *sql.DB, name string) (sql.Result, error) {
	return db.Exec(delter_user_by_name, name)
}
