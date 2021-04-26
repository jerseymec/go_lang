package main

import (
	"encoding/csv"
	"strings"
	"sync"

	"io"
	"log"
	"os"
	"strconv"
)

type City struct {
	Name       string
	Population int
}

func main() {
	f, err := os.Open("cities.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rows := genRows(f)
	filterSmallCity := filterOutPopulation(40000)
	upperRows1 := upperCityName(filterSmallCity(rows))
	upperRows2 := upperCityName(filterSmallCity(rows))
	upperRows3 := upperCityName(filterSmallCity(rows))
	for c := range fanIn(upperRows1, upperRows2, upperRows3) {
		log.Println(c)
	}
}

func filterOutPopulation(min int) func(<-chan City) <-chan City {
	return func(cities <-chan City) <-chan City {
		out := make(chan City)
		go func() {
			for c := range cities {
				if c.Population > min {
					out <- City{Name: strings.ToUpper(c.Name), Population: c.Population}
				}

			}
			close(out)
		}()
		return out
	}
}

func upperCityName(cities <-chan City) <-chan City {
	out := make(chan City)

	go func() {
		for c := range cities {
			out <- City{
				Name: strings.ToUpper(c.Name), Population: c.Population,
			}

		}
		close(out)
	}()
	return out
}

func genRows(r io.Reader) chan City {
	out := make(chan City)

	go func() {
		reader := csv.NewReader(r)
		_, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			populationInt, err := strconv.Atoi(row[9])
			if err != nil {
				continue
			}
			out <- City{
				Name:       row[1],
				Population: populationInt,
			}
		}
		close(out)

	}()
	return out

}

func fanIn(chans ...<-chan City) <-chan City {
	out := make(chan City)
	wg := &sync.WaitGroup{}
	wg.Add(len(chans))
	for _, c := range chans {
		go func(city <-chan City) {
			for r := range city {
				out <- r
			}
			wg.Done()
		}(c)

	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
