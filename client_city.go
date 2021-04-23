package izi_client

import (
	"net/http"
)

type cityClient struct {
	cli *Client
}

func (c cityClient) GetById(id string, p SearchBaseRequestParams) (*CityCompactForm, error) {
	p.form = compactForm
	resp := &CityCompactForm{}
	req := internalRequest{
		endpoint:            citiesByIdUrl + id,
		method:              "GetById",
		withRequest:         nil,
		withResponse:        resp,
		withQueryParams:     p.toQueryParam(),
		acceptedStatusCodes: []int{http.StatusOK},
		apiName:             "City",
	}
	err := c.cli.executeRequest(req)
	return resp, err
}

func (c cityClient) Search(p SearchPaginationBaseRequestParams) ([]CityCompactForm, error) {
	p.form = compactForm
	var resp []CityCompactForm
	req := internalRequest{
		endpoint:            citiesUrl,
		method:              "Search",
		withRequest:         nil,
		withResponse:        &resp,
		withQueryParams:     p.toQueryParam(),
		acceptedStatusCodes: []int{http.StatusOK},
		apiName:             "City",
	}
	err := c.cli.executeRequest(req)
	return resp, err
}

func (c cityClient) GetByIdFullForm(id string, p SearchBaseRequestParams) (*CityFullForm, error) {
	p.form = fullForm
	resp := &CityFullForm{}
	req := internalRequest{
		endpoint:            citiesByIdUrl + id,
		method:              "GetByIdFullForm",
		withRequest:         nil,
		withResponse:        resp,
		withQueryParams:     p.toQueryParam(),
		acceptedStatusCodes: []int{http.StatusOK},
		apiName:             "City",
	}
	err := c.cli.executeRequest(req)
	return resp, err
}

func (c cityClient) SearchFullForm(p SearchPaginationBaseRequestParams) ([]CityFullForm, error) {
	p.form = fullForm
	var resp []CityFullForm
	req := internalRequest{
		endpoint:            citiesUrl,
		method:              "SearchFullForm",
		withRequest:         nil,
		withResponse:        &resp,
		withQueryParams:     p.toQueryParam(),
		acceptedStatusCodes: []int{http.StatusOK},
		apiName:             "City",
	}
	err := c.cli.executeRequest(req)
	return resp, err
}
