package judge0

import (
	"net/http"
	"fmt"
	"encoding/json"
	)


type Client struct {
	authProvider AuthProvider
	httpClient   *http.Client
	baseURL      string
}






//returns a slice of all supported languages by judge0
func (c *Client) GetLanguages() ([]Language, error){
	
	
	data, err := c.doRequest(c.baseURL+"/languages", "GET", nil)

	if err!=nil{
		return nil, fmt.Errorf("failed to get languages: %w", err)
	}

	var languages []Language
	err = json.Unmarshal(data, &languages)
	if err!=nil{
		return nil, fmt.Errorf("failed to serialize/Unmarshal languages: %w", err)
	}

	return languages, nil

}

//returns a slice of all the judge0 status codes
func (c *Client) GetStatuses() ([]Status, error){

	data, err := c.doRequest(c.baseURL+"/statuses", "GET", nil)

	if err!=nil{
		return nil, fmt.Errorf("failed to get statuses: %w", err)
	}

	
	var statuses []Status
	err = json.Unmarshal(data, &statuses)
	if err!=nil{
		return nil, fmt.Errorf("failed to serialize/Unmarshal statuses: %w", err)
	}

	return statuses, nil

}