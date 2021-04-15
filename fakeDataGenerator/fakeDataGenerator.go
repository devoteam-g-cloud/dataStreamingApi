package fakeDataGenerator

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type FakeStruct struct {
	Name          string    `fake:"{name}"`     // Any available function all lowercase
	Sentence      string    `fake:"{sentence}"` // Can call with parameters
	RandStr       string    `fake:"{randomstring}"`
	PhoneNumber   string    `fake:"{phone}"` // Comma separated for multiple values
	Email         string    `fake:"{email}"` // Generate string from regex
	Longitude     float64   `fake:"{longitude}"`
	Latitude      float64   `fake:"{latitude}"`
	Zip           string    `fake:"{zip}"`
	CarModel      string    `fake:"{carmodel}"`
	Created       time.Time // Can take in a fake tag as well as a format tag
	CreatedFormat time.Time `fake:"{year}-{month}-{day}" format:"2006-01-02"`
	FavoriteFood  string    `fake:"{dinner}"`
	FavoriteColor string    `fake:"{color}"`
}

func GenerateFakeJson() []byte {
	var fakeStruct FakeStruct
	gofakeit.Struct(&fakeStruct)
	b, err := json.Marshal(fakeStruct)
	if err != nil {
		fmt.Println(err)
	}

	return b
}

func FakeDataGenerator(channel chan []byte, stopChannel chan bool) {
	for {
		channel <- GenerateFakeJson()
		select {
		case <-stopChannel:
			return
		default:
		}

	}

}
