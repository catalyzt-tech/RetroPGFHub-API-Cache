package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

func GetProjectApi(limit int, offset int, authorizeKey string) (*Response, error) {
	url := fmt.Sprintf("https://vote.optimism.io/api/v1/projects?limit=%d&offset=%d", limit, offset)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		slog.Error("failed to make new request", "error", err)
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authorizeKey))

	res, err := client.Do(req)
	if err != nil {
		slog.Error("failed to client do", "error", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("failed to make io readall", "error", err)
		return nil, err
	}

	var parseRes Response
	err = json.Unmarshal(body, &parseRes)
	if err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %q", body)
		return nil, err
	}

	return &parseRes, nil

}

func GetAllProjects(authorizeKey string) ([]Datum, error) {
	var conditionBreak bool = false
	var limit, offset int = 100, 0

	var datas []Datum

	for !conditionBreak {
		res, err := GetProjectApi(limit, offset, authorizeKey)
		if err != nil {
			return datas, err
		}

		datas = append(datas, res.Data...)

		if res.Meta.HasNext {
			offset += 100
		} else {
			conditionBreak = true
		}

	}

	return datas, nil

}
