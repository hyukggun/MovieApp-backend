package db

import (
	"fmt"
	"movie-app/model"
	"testing"
)

func TestConnector(t *testing.T) {
	db, err := ConnectDatabase()
	if err != nil {
		t.Error("Failed to connecting DB")
	}
	db.Close()
}

func TestFindUsers(t *testing.T) {
	if _, err := FindUsers(); err != nil {
		t.Error(err)
	}
}

func TestInsertUser(t *testing.T) {
	user := model.CreateUser("tester2", "1234", "tester2@naver.com")
	db, err := ConnectDatabase()
	if err != nil {
		t.Error(err)
	}
	if _, err := InsertUser(db, user); err != nil {
		t.Error(err)
	}
	db.Close()
}

func TestDeleteUser(t *testing.T) {
	db, err := ConnectDatabase()
	if err != nil {
		t.Error(err)
	}
	if _, err := DeleteUser(db, "tester2"); err != nil {
		t.Error(err)
	}
	t.Cleanup(func() {
		fmt.Println("Clean up funciton called")
	})
	db.Close()
}
