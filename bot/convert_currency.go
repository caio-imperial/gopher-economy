package bot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/caiosilvestre/gopher-economy/integrations/awesomeapi/economia"
)

// Normalization currency abbreviations for request in API
func NomalizationCurrencyAbbreviationToRequest(currentCurrencyAbbreviation, intoCurrencyAbbreviations string) (currencyAbbreviationsToRequest string) {
	currencyAbbreviationsToRequest = fmt.Sprintf("%s-%s", strings.ToUpper(currentCurrencyAbbreviation), strings.ToUpper(intoCurrencyAbbreviations))
	return
}

// Fetch information about quote currency values
func GetQuote(currencyAbbreviations string) (result economia.QuoteCurrency, err error) {
	result, err = economia.GetQuote(currencyAbbreviations)
	if err != nil {
		fmt.Printf("failed to request quote, err: %s", err)
	}
	return
}

// getCurrencySymbol returns the currency symbol for a given abbreviation.
func GetCurrencySymbol(currencyCode string) string {
	switch strings.ToLower(currencyCode) {
	case "brl":
		return "R$"
	case "eur":
		return "€"
	case "usd":
		return "$"
	default:
		return "$" // Return an empty string if the currency code is $
	}
}

// Return a messagen with convert currency
func ConvertMessage(message string) string {
	// Convert discord message in Array
	messageArray := strings.Split(message, " ")
	arraySize := len(messageArray)
	if arraySize <= 2 {
		fmt.Printf("invalid command, array size: %d", arraySize)
		return "invalid command"
	}
	// Format currency abbreviation from API
	currencyAbbreviationsToConvert := NomalizationCurrencyAbbreviationToRequest(messageArray[1], messageArray[2])

	// get currency value
	result, err := economia.GetQuote(currencyAbbreviationsToConvert)
	if err != nil {
		fmt.Printf("failed to request quote, err: %s", err)
		return "failed to search quote currency or invalid command"
	}

	// get correct currency symbol to amount convert
	quoteSymbol := GetCurrencySymbol(messageArray[2])

	amount := 1.0

	// check amount was provided
	if arraySize > 3 {
		// convert amount to float64
		amount, err = strconv.ParseFloat(messageArray[3], 64)
		if err != nil {
			fmt.Println("Erro: a string fornecida não pode ser convertida em float64")
			return "invalid amount value: the provided value is not a valid number"
		}
	}

	// convert amount currency provided
	convertAmount := result.Bid * amount

	// return convert amount
	return fmt.Sprintf("%s %.2f", quoteSymbol, convertAmount)

}
