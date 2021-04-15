package main

import (
	"fmt"
	"net/http"

	"main/fakeDataGenerator"
)

//Make a global channel
var c = make(chan []byte, 100000)

func dataStreaming(w http.ResponseWriter, req *http.Request) {
	// prints header values
	for name, values := range req.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
	// initiate a counter
	i := 0

	//infinite data stream
	for {
		//stream in the response body
		fmt.Fprintf(w, string(<-c)+"\n")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		i++
		if i%100000 == 0 {
			fmt.Println(i)
		}
		//Checks if the client closes connection or not (not supported by cloud run yet)
		select {
		case <-req.Context().Done():
			req.Body.Close()
			return

		default:
		}
	}
}

func main() {
	//initiate fakeDataGeneration in a separate goroutine
	go fakeDataGenerator.FakeDataGenerator(c)
	//Handles get requests
	http.HandleFunc("/dataStreamingEndpoint", dataStreaming)

	http.ListenAndServe(":8080", nil)
}
