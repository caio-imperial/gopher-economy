package bot

import (
	"fmt"
	"strings"
)

func NomalizationQuotes(message string) (currencyAbbreviationsToConvert string, QuoteCurrencys []string, err error) {
	messageArray := strings.Split(message, " ")
	if arraySize := len(messageArray); arraySize <= 2 {
		err = fmt.Errorf("invalid command, array size: %d", arraySize)
		return
	}
	currencyAbbreviationsToConvert = fmt.Sprintf("%s-%s", strings.ToUpper(messageArray[1]), strings.ToUpper(messageArray[2]))
	QuoteCurrencys = messageArray[1:3]
	return
}
