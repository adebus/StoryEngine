package card_test

import (
	"testing"

	"github.com/adebus/StoryEngine/card"
)

// TestNew calls card.New to init both a 2 sided and 4 sided card
// checking for the approrpiate struct type
func TestNew (t *testing.T) {
	side1 := "side1"
	side2 := "side2"
	side3 := "side3"
	side4 := "side4"

	var s = []string{side1, side2}

	c2, err := card.New(s)
	if err != nil {
		t.Fatalf("Received error initializing c2:"+err.Error())
	}
	s = append(s, side3)
	_, err2 := card.New(s)
	if err2 == nil {
		t.Fatalf("Did not receive an error initializing c3")
	} 
	s = append(s, side4)
	c4, err3 := card.New(s)
	if err != nil {
		t.Fatalf("Received error initializing c4"+err3.Error())
	}

	if _, ok := c2.(card.Card2); !ok {
		t.Fatalf("c2 is not a Card2 type")
	}

	if _, ok := c4.(card.Card4); !ok {
		t.Fatalf("c4 is not a Card4 type")
	}
}