package cred_checkers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/hugin-and-munin/cred-checker/internal/model"
)

var ErrNotFound = errors.New("company not found")
var ErrTooManyResults = errors.New("too many results found")

type gosuslugiCredsChecker struct {
	httpClient *http.Client
}

func NewGosuslugiCredsChecker(httpClient *http.Client) *gosuslugiCredsChecker {
	return &gosuslugiCredsChecker{
		httpClient: httpClient,
	}
}

func (c *gosuslugiCredsChecker) SearchCompany(ctx context.Context, inn string) (*model.Company, error) {
	jsonData, err := json.Marshal(buildRequestForInn(inn))
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(jsonData)

	req, err := http.NewRequest(
		http.MethodPost,
		"https://www.gosuslugi.ru/api/nsi-suggest/v1/MOB_ORGS/",
		bodyReader,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	jsonResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("call gosuslugi: status code: %d, status: %s; response body: %s", resp.StatusCode, resp.Status, string(jsonResponse))
	}

	response := ResponseBody{}
	err = json.Unmarshal(jsonResponse, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w, response body: %s", err, string(jsonResponse))
	}

	if response.Total == 0 {
		return nil, fmt.Errorf("%w: inn: %s", ErrNotFound, inn)
	}

	if response.Total > 1 {
		return nil, fmt.Errorf("%w: inn: %s", ErrTooManyResults, inn)
	}

	result := &model.Company{
		Name:           response.Items[0].Name,
		Inn:            inn,
		HasCredentials: true,
	}

	return result, nil
}

func buildRequestForInn(inn string) RequestBody {
	return RequestBody{
		Filter: Filter{
			Union: Union{
				UnionKind: "AND",
				Subs: []Subs{
					{
						Simple: Simple{
							AttributeName: "inn",
							Condition:     "EQUALS",
							Value: Value{
								AsString: inn,
							},
						},
					},
				},
			},
		},
		TreeFiltering:      "ONELEVEL",
		PageNum:            1,
		PageSize:           100,
		ParentRefItemValue: "",
		SelectAttributes:   []string{"*"},
		Tx:                 "",
	}
}
