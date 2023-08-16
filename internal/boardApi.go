package internal

import (
	"context"
	"movie-app/ent"
	"movie-app/ent/user"
	"movie-app/ent/userpost"

	_ "github.com/lib/pq"
)

func Client() (*ent.Client, error) {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=papayetoo dbname=movieapp password=rhkdgus0322 sslmode=disable")
	return client, err
}

func CreatePost(ctx context.Context, images map[string]string, text string) (*ent.UserPost, error) {
	user := ctx.Value("user").(*ent.User)
	client, err := Client()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	return client.UserPost.Create().SetPostImages(images).SetPostText(text).SetUserID(user).Save(ctx)
}

func GetPosts(ctx context.Context) ([](*ent.UserPost), error) {
	client, err := Client()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	return client.UserPost.Query().All(ctx)
}

func GetPostsByUser(ctx context.Context) ([](*ent.UserPost), error) {
	usr := ctx.Value("user").(*ent.User)
	client, err := Client()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	return client.UserPost.Query().Where(userpost.HasUserIDWith(user.IDEQ(usr.ID))).All(ctx)
}
