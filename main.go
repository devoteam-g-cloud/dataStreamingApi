package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"main/fakeDataGenerator"
)

//Make a global channel
var c chan []byte
var handlerThrottling *int

func dataStreaming(w http.ResponseWriter, req *http.Request) {

	//infinite data stream
	for {
		//stream in the response body
		fmt.Fprintf(w, string(<-c)+"\n")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		//Checks if the client closes connection or not (not supported by cloud run yet)
		select {
		case <-req.Context().Done():
			req.Body.Close()
			return

		default:
		}
		if *handlerThrottling != 0 {
			time.Sleep(time.Duration(1 / *handlerThrottling) * time.Second)
		}

	}
}

func main() {
	generatorThrottling := flag.Int("generator-throttling-rate", 0, "json generator throttling (element/second)")
	handlerThrottling = flag.Int("handler-throttling-rate", 0, "handler throttling (element/second)")
	channelSize := flag.Int("channel-size", 100000, "channel size (number of messages)")
	flag.Parse()

	c = make(chan []byte, *channelSize)

	//initiate fakeDataGeneration in a separate goroutine
	go fakeDataGenerator.FakeDataGenerator(c, generatorThrottling)
	//Handles get requests
	http.HandleFunc("/dataStreamingEndpoint", dataStreaming)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
