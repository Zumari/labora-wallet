package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Zumari/labora-wallet/pruebas/models"
	"github.com/joho/godotenv"
)

const (
	urlTruoraAPI = "https://api.checks.truora.com/v1/checks"
	contentType  = "application/x-www-form-urlencoded"
	timeOut      = 5 * time.Second
)

// PostRequestTruora consult the background of a person in the Truora API.
func PostRequestTruora(national_id, country, entityType string, user_authorized bool) (string, error) {
	method := "POST"

	bodyAtributes := WorkingUrlData(national_id, country, entityType, user_authorized)

	payload := strings.NewReader(bodyAtributes.Encode())

	client := &http.Client{}
	req, err := http.NewRequest(method, urlTruoraAPI, payload)

	if err != nil {
		return "", fmt.Errorf("Error creating the application http %w", err)
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	req.Header.Add("Truora-API-Key", apiKey)
	req.Header.Add("Content-Type", contentType)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("Error making the HTTP application %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("Error reading the Body request %w", err)
	}

	var checkRequest models.TruoraPostResponse

	err = json.Unmarshal(body, &checkRequest)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("Error decoding the JSON code %w", err)
	}

	checkId := checkRequest.Check.CheckID

	return checkId, nil

}

// WorkingUrlData convert an URL type the attributes extracted from the API.
func WorkingUrlData(national_id, country, entityType string, user_authorized bool) url.Values {
	bodyAtributes := url.Values{}
	bodyAtributes.Set("national_id", national_id)
	bodyAtributes.Set("country", country)
	bodyAtributes.Set("type", entityType)
	bodyAtributes.Set("user_authorized", strconv.FormatBool(user_authorized))

	return bodyAtributes

}

func GetScoreTruoraAPI(checkId string) (int, error) {
	urlConsult := urlTruoraAPI + "/" + checkId

	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, urlConsult, payload)

	if err != nil {
		return -1, fmt.Errorf("Error creating the application http %w", err)
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	req.Header.Add("Truora-API-Key", apiKey)
	req.Header.Add("Content-Type", contentType)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return -1, fmt.Errorf("Error making the HTTP application %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return -1, fmt.Errorf("Error reading the Body request %w", err)
	}

	var checkRequest models.TruoraGetResponse

	err = json.Unmarshal(body, &checkRequest)
	if err != nil {
		fmt.Println(err)
		return -1, fmt.Errorf("Error decoding the JSON code %w", err)
	}

	checkScore := checkRequest.Check.Score

	return checkScore, nil

}

func GetTruoraResponse(national_id, country, entityType string, user_authorized bool) (bool, error) {
	checkID, err := PostRequestTruora(national_id, country, entityType, user_authorized)
	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("Error in the Post to the API %w", err)
	}
	time.Sleep(timeOut)

	score, err := GetScoreTruoraAPI(checkID)
	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("Error in the GET to the API %w", err)
	}

	return score == 1, nil
}
