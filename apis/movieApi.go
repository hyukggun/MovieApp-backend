package apis

import (
	"fmt"
	"io"
	"net/http"
)

const BaseUrl = "https://api.themoviedb.org/3/movie"

func createGetRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJmZmE3NGM2NmE4MTViYTI3NzAzNjdmMTc1ZTgxODhlOSIsInN1YiI6IjY0M2JhYTFjNzE5YWViMDRjYmViZWVkOSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.3XqlRAnhTm9oIAPb_4zT8dWLyySVIfOl-mA8UZqgrzg")
	return req
}

func handleStatusCodeError(statusCode int) error {
	if statusCode != 200 {
		return fmt.Errorf("Status code is not 200")
	}
	return nil
}

func NowPlayingMovies() []byte {
	url := BaseUrl + "/now_playing"
	req := createGetRequest(url)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print("Error occured")
		panic(err)
	}

	if res.StatusCode != 200 {
		err = fmt.Errorf("Response status is not 200 : status code is %d\n", res.StatusCode)
		panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

func PopularMovies() []byte {
	url := BaseUrl + "/popular"
	req := createGetRequest(url)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error occuredd")
		panic(err)
	}

	if err := handleStatusCodeError(res.StatusCode); err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}
