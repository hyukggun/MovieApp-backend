package db

import (
	"fmt"
	"movie-app/model"
	"testing"
)

func TestConnector(t *testing.T) {
	db, err := ConnectDatabase()
	defer db.Close()
	if err != nil {
		t.Error("Failed to connecting DB")
	}
}

func TestFindUsers(t *testing.T) {
	db, err := ConnectDatabase()
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
	if err := FindUsers(db); err != nil {
		t.Error(err)
	}
}

func TestInsertUser(t *testing.T) {
	user := model.CreateUser("tester2", "1234", "tester2@naver.com")
	db, err := ConnectDatabase()
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
	if _, err := InsertUser(db, user); err != nil {
		t.Error(err)
	}
}

func TestDeleteUser(t *testing.T) {
	db, err := ConnectDatabase()
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
	// user := model.CreateUser("tester2", "1234", "tester2@naver.com")
	// if _, err := InsertUser(db, user); err != nil {
	// 	t.Error(err)
	// }
	if _, err := DeleteUser(db, "babayetoo"); err != nil {
		t.Error(err)
	}
	t.Cleanup(func() {
		fmt.Println("Clean up funciton called")
	})
}
