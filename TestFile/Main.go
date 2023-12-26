package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendRequest(url string, payload map[string]interface{}) (string, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func parseResponse(response string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(response), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {

	url := "https://krisha.kz/"

	payload := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
		"key3": true,
	}

	response, err := sendRequest(url, payload)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}

	fmt.Println("Ответ от сервера:", response)

	parsedData, err := parseResponse(response)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	fmt.Println("Разобранные данные:", parsedData)
}
