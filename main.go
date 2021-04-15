package main

import (
	"fmt"
	"net/http"

	"main/fakeDataGenerator"
)

func dataStreaming(w http.ResponseWriter, req *http.Request) {
	c := make(chan []byte, 10000)
	stopChannel := make(chan bool)
	defer close(c)
	for name, values := range req.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
	i := 0

	go fakeDataGenerator.FakeDataGenerator(c, stopChannel)

	for {
		fmt.Fprintf(w, string(<-c)+"\n")

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		i++
		if i%100000 == 0 {
			fmt.Println(i)
		}
		select {
		case <-req.Context().Done():
			stopChannel <- true
			req.Body.Close()

			return

		default:
		}
	}
}

func main() {
	http.HandleFunc("/dataStreamingEndpoint", dataStreaming)

	http.ListenAndServe(":8080", nil)
}
