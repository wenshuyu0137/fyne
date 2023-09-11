package main

import (
	"testing"
)

func TestGold_GetPrices(t *testing.T) {
	g := Gold{
		Prices: nil,
	}

	_, err := g.get_prices()
	if err != nil {
		t.Error(err)
	}
}
