package model

import "testing"

func TestCreateMovie(t *testing.T) {
	movie := CreateMovie("test1", "barbie", "Barbie is a doll", "hello_world.jpg")

	if movie.Id != "test1" {
		t.Errorf("Movie id not matching error : expected id is %s\n", movie.Id)
	}
	if movie.Title != "barbie" {
		t.Errorf("Movie title not matching error : expected title is %s\n", movie.Title)
	}
	if movie.Overview != "Barbie is a doll" {
		t.Errorf("Movie overview not matching error : expected overview is %s\n", movie.Overview)
	}
	if movie.PosterPath != "hello_world.jpg" {
		t.Errorf("Movie postter path not matching error : expected posterPath is %s\n", movie.PosterPath)
	}
}

func TestNowPlaying(t *testing.T) {
	movies := NowPlayingMovies()
	if len(movies) != 8 {
		t.Errorf("Now Playing movies length is not %d\n", len(movies))
	}
}
