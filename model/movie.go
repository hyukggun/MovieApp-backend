package model

type Movie struct {
	Id         string
	Title      string
	Overview   string
	PosterPath string
}

func CreateMovie(id, title, overview, posterPath string) *Movie {
	movie := &Movie{id, title, overview, posterPath}
	return movie
}

func NowPlayingMovies() []*Movie {
	movies := []*Movie{
		CreateMovie("1", "더 문", "babie", ""),
		CreateMovie("2", "콘크리트 유토피아", "babie", ""),
		CreateMovie("3", "밀수", "babie", ""),
		CreateMovie("4", "인시디어스 붉은 문", "babie", ""),
		CreateMovie("5", "잠", "babie", ""),
		CreateMovie("6", "barbie", "babie", ""),
		CreateMovie("7", "barbie", "babie", ""),
		CreateMovie("8", "barbie", "babie", ""),
	}
	return movies
}
