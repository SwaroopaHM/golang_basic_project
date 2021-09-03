package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "King"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//converting a deck of cards to slice of cards, and retuning string
//of cards using 'strings' package "Join" function
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//writing a string of card names to a file using ioutil package's
//WriteFile function, see its syntax in the official doc

//this function is returning an error because 'WriteFile' function
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//option #1: Log the error and call myDeck()
		//option #2: Log the error and exit the propram

		//choosing option #2
		fmt.Println("Error:", err)

		//'os' is the platform independant package
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

//the toughest part of this project is generating random number

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
