package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Dinosaur struct {
	Name           string
	Pronunciation  string
	LengthInMeters int
	Diet           string
}

func main() {
	dinosaurs := []Dinosaur{
		{
			Name:           "Coelophysis",
			Pronunciation:  "seel-OH-fie-sis",
			LengthInMeters: 2,
			Diet:           "carnivorous",
		},
		{
			Name:           "Triceratops",
			Pronunciation:  "tri-SERRA-tops",
			LengthInMeters: 9,
			Diet:           "herbivorous",
		},
		{
			Name:           "Tyrannosaurus",
			Pronunciation:  "tie-RAN-oh-sore-us",
			LengthInMeters: 12,
			Diet:           "carnivorous",
		},
		{
			Name:           "Diplodocus",
			Pronunciation:  "DIP-low DOCK-us",
			LengthInMeters: 26,
			Diet:           "herbivorous",
		},
		{
			Name:           "Panoplosaurus",
			Pronunciation:  "pan-op-loh-sore-us",
			LengthInMeters: 7,
			Diet:           "herbivorous",
		},
	}

	http.HandleFunc("/dinosaurs", func(writer http.ResponseWriter, request *http.Request) {
		jsonResponse, err := json.Marshal(dinosaurs)
		if err != nil {
			log.Println("cannot serialize dinosaurs")
		}

		_, err = fmt.Fprintln(writer, string(jsonResponse))

		if err != nil {
			log.Println("Could not write dinosaurs to the response")
		}
	})

	log.Println("Dinosaur service starting on port 8084")
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		log.Fatal("dinosaur-service could not start")
		return
	}
}
