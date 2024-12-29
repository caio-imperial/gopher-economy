package economia

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	//baseUrl represent api URL
	baseUrl string
)

// QuoteCurrency represents the details of a currency quote in a financial system.
// @Code is the standard currency code (e.g., "USD").
// @CodeIn is the code of the currency being quoted against (e.g., "BRL" for Brazilian Real).
// @Name is the name of the currency (e.g., "United States Dollar").
// @High is the highest price of the currency in the current trading session.
// @Low is the lowest price of the currency in the current trading session.
// @VarBid is the variation in the bid price of the currency compared to the previous session.
// @PctChange is the percentage change in the currency price.
// @Bid is the current bid price of the currency.
// @Ask is the current ask price of the currency.
// @Timestamp is the Unix timestamp of when the quote was recorded.
// @CreateDate is the date and time when the quote was created, formatted as a string.
type QuoteCurrency struct {
	Code       string  `json:"code"`
	CodeIn     string  `json:"codein"`
	Name       string  `json:"name"`
	High       float64 `json:"high,string"`
	Low        float64 `json:"low,string"`
	VarBid     float64 `json:"varBid,string"`
	PctChange  float64 `json:"pctChange,string"`
	Bid        float64 `json:"bid,string"`
	Ask        float64 `json:"ask,string"`
	Timestamp  int64   `json:"timestamp,string"`
	CreateDate string  `json:"create_date"`
}

// Init initializes all environment variables using in package
func Init() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env", err)
	}
	baseUrl = os.Getenv("BASE_URL")
	if baseUrl == "" {
		fmt.Printf("BASE_URL not set in environment\n")
		err = errors.New("BASE_URL not set in environment")
		return
	}
	return
}

// concatUrlPath concat url path.
// @param path: Path you want to add in url.
// @return: urlFull: Url with path or nil, err: Error or nil.
func concatUrlPath(path string) (urlFull string, err error) {
	urlFull, err = url.JoinPath(baseUrl, path)
	return
}

// GetQuote get the information about quotes.
// @param currencySymbol: Simbol of quote.
// @return: quoteCurrency: quote currency.
func GetQuote(currencySymbol string) (quoteCurrency QuoteCurrency, err error) {
	// Concatenate the base URL with the currency symbol to form the full URL.
	urlFull, err := concatUrlPath(fmt.Sprintf("/json/last/%s", currencySymbol))
	if err != nil {
		return
	}

	// Get information about quote currency
	resp, err := http.Get(urlFull)
	if err != nil {
		fmt.Println("Error request:", err)
		return
	}

	defer resp.Body.Close()

	// Convert body information into []byte
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Convert bodyByte information into map[string]QuoteCurrency
	var result map[string]QuoteCurrency
	err = json.Unmarshal(bodyByte, &result)
	if err != nil {
		fmt.Println("Error convert map[string]QuoteCurrency:", err)
		return
	}

	// Convert result information into QuoteCurrency
	quoteCurrency, isValid := result[strings.Replace(currencySymbol, "-", "", -1)]
	if !isValid {
		fmt.Print("Error reading response currencySymbol is Invalid")
		return
	}
	return
}
