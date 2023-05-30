package main

import (
	"fmt"
	"time"

	"github.com/Zumari/labora-wallet/pruebas/services"
)

const (
	urlTruoraAPI = "https://api.checks.truora.com/v1/checks"
	contentType  = "application/x-www-form-urlencoded"
	timeOut      = 5 * time.Second
)

func main() {

	// score, err := services.GetTruoraResponse("74909799", "PE", "person", true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(err)

	// fmt.Println(score)

	checkID, err := services.PostRequestTruora("74909799", "PE", "person", true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(checkID, err)

	time.Sleep(timeOut)

	var score int
	score, err = services.GetScoreTruoraAPI(checkID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(score, err)

}
