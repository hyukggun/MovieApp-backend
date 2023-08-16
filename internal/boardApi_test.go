package internal

import (
	"context"
	"testing"
)

func TestBoardApiCreatePost(t *testing.T) {
	client, err := Client()
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	ctx := context.Background()
	usr, err := client.User.Get(ctx, 3)
	if err != nil {
		t.Error(err)
	}
	ctx = context.WithValue(ctx, "user", usr)
	post, err := CreatePost(ctx, map[string]string{"Uniform": "https://pds.joongang.co.kr/news/component/htmlphoto_mmdata/202306/29/c837a164-4d93-4956-8393-7fd8953b6a55.jpg"}, "Test2")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Create Post %v %v\n", post.ID, post.PostText)
}

func TestBoardApiGetPosts(t *testing.T) {
	ctx := context.Background()
	posts, err := GetPosts(ctx)
	if err != nil {
		t.Error(err)
	}
	for _, post := range posts {
		t.Logf("Post id %v image %v text %v\n", post.ID, post.PostImages, post.PostText)
	}
}

func TestBoardApiGetPostsByUserId(t *testing.T) {
	ctx := context.Background()
	client, err := Client()
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	usr, err := client.User.Get(ctx, 3)
	if err != nil {
		t.Error(usr)
	}
	ctx = context.WithValue(ctx, "user", usr)
	posts, err := GetPostsByUser(ctx)
	if err != nil {
		t.Error(err)
	}
	if len(posts) == 0 {
		t.Error("Post count is zero")
	}
}
