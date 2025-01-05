package bot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/caiosilvestre/gopher-economy/integrations/awesomeapi/economia"
	"github.com/caiosilvestre/gopher-economy/logger"
)

// Normalization currency abbreviations for request in API
func NomalizationCurrencyAbbreviationToRequest(ctxLogger *logger.AppLogger, currentCurrencyAbbreviation, intoCurrencyAbbreviations string) (currencyAbbreviationsToRequest string) {
	currencyAbbreviationsToRequest = fmt.Sprintf("%s-%s", strings.ToUpper(currentCurrencyAbbreviation), strings.ToUpper(intoCurrencyAbbreviations))
	ctxLogger.Info("Successfully normalized currency abbreviation in the request.")
	return
}

// getCurrencySymbol returns the currency symbol for a given abbreviation.
func GetCurrencySymbol(ctxLogger *logger.AppLogger, currencyCode string) string {
	switch strings.ToLower(currencyCode) {
	case "brl":
		return "R$"
	case "eur":
		return "€"
	case "usd":
		return "$"
	default:
		ctxLogger.Info(
			"Successfully returned the converted currency amount.",
			"parameters", map[string]interface{}{"currencyCode": currencyCode},
		)
		return "$" // Return an empty string if the currency code is $
	}
}

// Return a messagen with convert currency
func ConvertMessage(ctxLogger *logger.AppLogger, message string) string {
	// Convert discord message in Array
	messageArray := strings.Split(message, " ")
	arraySize := len(messageArray)
	if arraySize <= 2 {
		ctxLogger.Error(
			"failed to request quote, arraySize <= 2",
			"parameters", map[string]interface{}{
				"message":   message,
				"arraySize": arraySize,
			})
		return "invalid command"
	}
	// Format currency abbreviation from API
	currencyAbbreviationsToConvert := NomalizationCurrencyAbbreviationToRequest(ctxLogger, messageArray[1], messageArray[2])

	// get currency value
	result, err := economia.GetQuote(ctxLogger, currencyAbbreviationsToConvert)
	if err != nil {
		ctxLogger.Error("failed to get Quote information",
			"parameters", map[string]interface{}{
				"currencyAbbreviationsToConvert": currencyAbbreviationsToConvert,
			})
		return "failed to search quote currency or invalid command"
	}

	// get correct currency symbol to amount convert
	quoteSymbol := GetCurrencySymbol(ctxLogger, messageArray[2])

	amount := 1.0

	// check amount was provided
	if arraySize > 3 {
		// convert amount to float64
		amount, err = strconv.ParseFloat(messageArray[3], 64)
		if err != nil {
			ctxLogger.Error("a string fornecida não pode ser convertida em float64",
				"parameters", map[string]interface{}{
					"message": message,
				})
			return "invalid amount value: the provided value is not a valid number"
		}
	}

	// convert amount currency provided
	convertAmount := result.Bid * amount

	ctxLogger.Info("Successfully returned the converted currency amount.")
	// return convert amount
	return fmt.Sprintf("%s %.2f", quoteSymbol, convertAmount)

}
