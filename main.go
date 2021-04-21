package main

import (
	"bytes"
	"compress/gzip"
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
var isGzip *bool

func dataStreaming(w http.ResponseWriter, req *http.Request) {
	var message bytes.Buffer
	//infinite data stream
	if req.Method == "GET" {
		for {
			//stream in the response body

			if *isGzip {
				gz := gzip.NewWriter(&message)
				gz.Write(<-c)
				gz.Close()
			} else {
				message.Write(<-c)
			}

			fmt.Fprintf(w, message.String()+"\n")
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

}

func main() {
	generatorThrottling := flag.Int("generator-throttling-rate", 0, "json generator throttling (element/second)")
	handlerThrottling = flag.Int("handler-throttling-rate", 0, "handler throttling (element/second)")
	channelSize := flag.Int("channel-size", 100000, "channel size (number of messages)")
	isGzip = flag.Bool("is-gzip", false, "is gzip")

	flag.Parse()

	c = make(chan []byte, *channelSize)

	//initiate fakeDataGeneration in a separate goroutine
	go fakeDataGenerator.FakeDataGenerator(c, generatorThrottling)
	//Handles get requests
	http.HandleFunc("/dataStreamingEndpoint", dataStreaming)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
