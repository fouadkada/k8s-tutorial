package service

import (
	"encoding/json"
	"github.com/fouadkada/k8s-tutorial/aggregator/internal/model"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Aggregator struct {
	moviesUrl  string
	weatherUrl string
}

func NewAggregatorService(moviesUrl string, weatherUrl string) Aggregator {
	return Aggregator{
		moviesUrl:  moviesUrl,
		weatherUrl: weatherUrl,
	}
}

func (s Aggregator) Aggregate() model.Result {
	result := model.Result{}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := http.Get(s.moviesUrl)
		if err != nil {
			log.Printf("could not fetching movies [%v]\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Println("did not get status code 200 while fetching movies")
			return
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("could not read response body while fetching movies  [%v]\n", err)
			return
		}

		var movies []model.Movie
		err = json.Unmarshal(b, &movies)
		if err != nil {
			log.Printf("could not unmarshal response while fetching movies  [%v]\n", err)
			return
		}

		result.Movies = movies
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := http.Get(s.weatherUrl)
		if err != nil {
			log.Printf("could not fetching weather  [%v]\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Println("did not get status code 200 while fetching weather")
			return
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("could not read response body while fetching weather  [%v]\n", err)
			return
		}

		var weather []model.Weather
		err = json.Unmarshal(b, &weather)
		if err != nil {
			log.Printf("could not unmarshal response while fetching weather  [%v]\n", err)
			return
		}

		result.Weather = weather
	}()

	wg.Wait()

	return result
}
