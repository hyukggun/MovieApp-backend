package apis

import (
	"fmt"
	"testing"
)

func TestNowPlaying(t *testing.T) {
	data := NowPlayingMovies()
	fmt.Printf("size of the data is %d", len(data))
	if len(data) <= 0 {
		t.Error("Response data is empty")
	}
}
