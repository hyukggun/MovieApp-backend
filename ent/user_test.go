package ent

import (
	"context"
	"movie-app/ent/migrate"
	"testing"

	_ "github.com/lib/pq"
)

func TestClient(t *testing.T) {
	client, err := Open("postgres", "host=127.0.0.1 port=5432 user=papayetoo dbname=movieapp password=rhkdgus0322 sslmode=disable")
	if err != nil {
		t.Error("Ent Client error")
	}
	client.Close()
}

func TestMigration(t *testing.T) {
	client, err := Open("postgres", "host=127.0.0.1 port=5432 user=papayetoo dbname=movieapp password=rhkdgus0322 sslmode=disable")
	defer client.Close()
	if err != nil {
		t.Error("Ent client error")
	}
	ctx := context.TODO()
	if err := client.Schema.Create(ctx); err != nil {
		t.Errorf("failed creating schem resources: %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	client, err := Open("postgres", "host=127.0.0.1 port=5432 user=papayetoo dbname=movieapp password=rhkdgus0322 sslmode=disable")
	if err != nil {
		t.Error("Ent client error")
	}
	context := context.TODO()
	if err := client.Schema.Create(
		context,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		t.Error(err)
	}

	newUser := client.User.Create().SetName("tester").SetEmail("tester@naver.com").SaveX(context)
	if newUser == nil {
		t.Error("Creating user failed")
	}
	t.Logf("new User %v\n", newUser)
	client.Close()
}

func TestUpdateUser(t *testing.T) {
	client, err := Open("postgres", "host=127.0.0.1 port=5432 user=papayetoo dbname=movieapp password=rhkdgus0322 sslmode=disable")
	if err != nil {
		t.Error("Ent client error")
	}
	context := context.TODO()
	updatedUser, err := client.User.UpdateOneID(1).SetEmail("rhkdgus0826@gmail.com").Save(context)
	if err != nil {
		t.Error("Ent update error")
	}
	t.Log(updatedUser)
}
