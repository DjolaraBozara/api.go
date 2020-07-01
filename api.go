package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
}

type deck struct {
	ID        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card
}

type allDecks []deck

var decks = allDecks{
	{
		ID:        "1",
		Shuffled:  false,
		Remaining: 2,
		Cards: []Card{
			Card{
				Value: "2",
				Suit:  "DIAMOND",
			},
			Card{
				Value: "KING",
				Suit:  "SPADE",
			},
		},
	},
}

func getAllDecks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(decks)
}

type Deck []Card

func newDeck(w http.ResponseWriter, r *http.Request) {
	var deck Deck

	values := []string{"2", "3", "4", "5", "6", "7",
		"8", "9", "10", "JACK", "QUEEN", "KING", "ACE"}

	suits := []string{"HEART", "DIAMOND", "CLUB", "SPADE"}

	for i := 0; i < len(values); i++ {
		for n := 0; n < len(suits); n++ {

			card := Card{
				Value: values[i],
				Suit:  suits[n],
			}
			deck = append(deck, card)

			json.NewEncoder(w).Encode(card)

		}
	}
	return
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func Deal(d Deck, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(d[i])
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/newDeck", newDeck).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
