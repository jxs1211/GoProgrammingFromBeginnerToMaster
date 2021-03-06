package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

type result struct {
	value string
}

func first(servers ...*httptest.Server) (result, error) {
	c := make(chan result)
	queryFunc := func(server *httptest.Server) {
		defer fmt.Println("query done")
		url := server.URL
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("http get error: %s\n", err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		c <- result{
			value: string(body),
		}
	}
	for _, serv := range servers {
		go queryFunc(serv)
	}

	select {
	case r := <-c:
		return r, nil
	case <-time.After(2 * time.Second):
		return result{}, errors.New("timeout")
	}
}

func fakeWeatherServer(name string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s receive a http request\n", name)
		time.Sleep(1 * time.Second)
		w.Write([]byte(name + ":ok"))
	}))
}

func main() {
	result, err := first(
		fakeWeatherServer("open-weather-1"),
		fakeWeatherServer("open-weather-2"),
		fakeWeatherServer("open-weather-3"))
	if err != nil {
		log.Println("invoke first error:", err)
		return
	}

	log.Println(result)
}
