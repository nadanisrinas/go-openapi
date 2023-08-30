package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"openapi-golang/dto"
)

func FUnmarshalResponse(data []byte) (dto.IResponse, error) {
	var r dto.IResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func RequestPostA(url string, payload string, apiKey string) {
	// initiate client http client
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	if err != nil {
		log.Fatalln(err)
	}
	res, err := client.Do(req)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//body type byte
	//harus di close supaya tidak memory leak
	defer res.Body.Close()
	stringBody := string(body)

	r, errPars := FUnmarshalResponse([]byte(stringBody))
	if errPars != nil {
		fmt.Println(errPars.Error())
		return
	}

	fmt.Println(r.Choices[0].Message.Content)

}
