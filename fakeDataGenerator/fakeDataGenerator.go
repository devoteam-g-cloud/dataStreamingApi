package fakeDataGenerator

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type FakeStruct struct {
	Name          string    `fake:"{name}"`             // Any available function all lowercase
	Sentence      string    `fake:"{sentence:100,300}"` // Can call with parameters
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

	var fakeStruct FakeStruct
	//populates empty struct
	gofakeit.Struct(&fakeStruct)
	fakeStruct.Updated = time.Now()
	b, err := json.Marshal(fakeStruct)
	if err != nil {
		fmt.Println(err)
	}

	return b
}

func FakeDataGenerator(channel chan []byte) {
	// fills channel until it reaches capacity
	for {
		//time.Sleep(time.Second)
		channel <- GenerateFakeJson()
	}

}
