package fakeDataGenerator

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type FakeElement struct {
	Name          string    `fake:"{name}"`              // Any available function all lowercase
	Sentence      string    `fake:"{sentence:100,3000}"` // Can call with parameters
	RandStr       string    `fake:"{randomstring}"`
	PhoneNumber   string    `fake:"{phone}"` // Comma separated for multiple values
	Email         string    `fake:"{email}"` // Generate string from regex
	Longitude     float64   `fake:"{longitude}"`
	Latitude      float64   `fake:"{latitude}"`
	Zip           string    `fake:"{zip}"`
	CarModel      string    `fake:"{carmodel}"`
	Created       time.Time `fake:"{date}"`
	Updated       time.Time
	MessageId     string `fake:"{uuid}"`
	FavoriteFood  string `fake:"{dinner}"`
	FavoriteColor string `fake:"{color}"`
}

func GenerateFakeJson() []byte {

	var fakeElement FakeElement
	//populates empty struct
	gofakeit.Struct(&fakeElement)
	fakeElement.Updated = time.Now()
	b, err := json.Marshal(fakeElement)
	if err != nil {
		fmt.Println(err)
	}

	return b
}

func FakeDataGenerator(channel chan []byte, generatorFrequency *int) {
	// fills channel until it reaches capacity

	if *generatorFrequency != 0 {
		for {
			time.Sleep(time.Duration(1 / *generatorFrequency) * time.Second)
			channel <- GenerateFakeJson()
		}

	} else {
		for {
			channel <- GenerateFakeJson()
		}
	}

}
