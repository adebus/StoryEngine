package card

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Card2 struct {
	Side1 string
	Side2 string
}

type Card4 struct {
	Side1 string
	Side2 string
	Side3 string
	Side4 string
}

type Card interface {
	RandomSide() (string, error)
}

func (c Card2) RandomSide() (string, error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	n := r.Intn(2)

	switch n {
	case 0:
		return c.Side1, nil
	case 1:
		return c.Side2, nil
	}
	return "", fmt.Errorf("received an unexpected side number: %v", n)
}

func (c Card4) RandomSide() (string, error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	n := r.Intn(4)

	switch n {
	case 0:
		return c.Side1, nil
	case 1:
		return c.Side2, nil
	case 2:
		return c.Side3, nil
	case 3:
		return c.Side4, nil
	}
	return "", fmt.Errorf("received an unexpected side number: %v", n)
}

func New(sides []string) (Card, error) {
	var c Card

	if len(sides) == 4 {
		// This is a 4 sided card
		c = Card4{Side1: sides[0], Side2: sides[1], Side3: sides[2], Side4: sides[3]}
	} else if len(sides) == 2 {
		// This is a 2 sided card
		c = Card2{Side1: sides[0], Side2: sides[1]}
	} else {
		return nil, errors.New("sides must be 2 or 4")
	}
	return c, nil
}