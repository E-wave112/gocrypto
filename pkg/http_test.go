package pkg

import (
	"reflect"
	"testing"
)

func TestSupportedCurrencies(t *testing.T) {
	t.Helper()

	t.Run("validate the list of supported currencies", func(t *testing.T) {
		result := ListSupportedCryptoCurrencies()
		expected := []string{"Bitcoin (BTC)", "Ethereum(ETH)", "Dogecoin(DOGE)", "Solana(SOL)", "Shiba(SHIB)", "Tether(USDT)"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}

	})

}

func TestNumberFormatter(t *testing.T) {
	t.Helper()

	t.Run("validate that the formatter runs for non negative decimal numbers", func(t *testing.T) {
		result := formatNumber("12.34000000")
		expected := "0.081037"

		if result != expected {
			t.Errorf("expected %q, got %q", expected, result)
		}
	})

	t.Run("validate that the formatter runs for negative numbers", func(t *testing.T) {
		result := formatNumber("-0.999")
		expected := "-1.001001"

		if result != expected {
			t.Errorf("expected %q, got %q", expected, result)
		}

	})

	t.Run("validate that the formatter runs for integer numbers", func(t *testing.T) {
		result := formatNumber("5000")
		expected := "0.0002"

		if result != expected {
			t.Errorf("expected %q, got %q", expected, result)
		}

	})
}
