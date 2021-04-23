package izi_client

import (
	"net/http"
	"strings"
)

type objectClient struct {
	cli *Client
}

func (c objectClient) GetById(id string, p SearchBaseRequestParams) (*ObjectCompactForm, error) {
	p.form = compactForm
	resp := &ObjectCompactForm{}
	req := internalRequest{
		endpoint:            objectByIdUrl + id,
		method:              "GetById",
		withRequest:         nil,
		withResponse:        resp,
		withQueryParams:     p.toQueryParam(),
		acceptedStatusCodes: []int{http.StatusOK},
		apiName:             "Object",
	}
	err := c.cli.executeRequest(req)
	return resp, err
}

func (c objectClient) GetByIds(id []string, p SearchBaseRequestParams) ([]ObjectCompactForm, error) {
	p.form = compactForm
	var resp []ObjectCompactForm
	req := internalRequest{
		endpoint:            objectByIdsUrl + strings.Join(id, ","),
		method:              "GetByIds",
		withRequest:         nil,
		withResponse:        &resp,
		withQueryParams:     p.toQueryParam(),
		acceptedStatusCodes: []int{http.StatusOK},
		apiName:             "Object",
	}
	err := c.cli.executeRequest(req)
	return resp, err
}

func (c objectClient) GetByIdFullForm(id string, p SearchBaseRequestParams) (*ObjectFullForm, error) {
	p.form = fullForm
	resp := &ObjectFullForm{}
	req := internalRequest{
		endpoint:            objectByIdUrl + id,
		method:              "GetByIdFullForm",
		withRequest:         nil,
		withResponse:        resp,
		withQueryParams:     p.toQueryParam(),
		acceptedStatusCodes: []int{http.StatusOK},
		apiName:             "Object",
	}
	err := c.cli.executeRequest(req)
	return resp, err
}

func (c objectClient) GetByIdsFullForm(id []string, p SearchBaseRequestParams) ([]ObjectFullForm, error) {
	p.form = fullForm
	var resp []ObjectFullForm
	req := internalRequest{
		endpoint:            objectByIdUrl + strings.Join(id, ","),
		method:              "GetByIdsFullForm",
		withRequest:         nil,
		withResponse:        &resp,
		withQueryParams:     p.toQueryParam(),
		acceptedStatusCodes: []int{http.StatusOK},
		apiName:             "Object",
	}
	err := c.cli.executeRequest(req)
	return resp, err
}

func (c objectClient) Search(p SearchObjectRequest) (*ObjectCompactForm, error) {
	p.form = compactForm
	resp := &ObjectCompactForm{}
	req := internalRequest{
		endpoint:            objectSearchUrl,
		method:              "Search",
		withRequest:         nil,
		withResponse:        resp,
		withQueryParams:     p.toQueryParam(),
		acceptedStatusCodes: []int{http.StatusOK},
		apiName:             "Object",
	}
	err := c.cli.executeRequest(req)
	return resp, err
}

func (c objectClient) SearchFullForm(p SearchObjectRequest) (*ObjectFullForm, error) {
	p.form = fullForm
	resp := &ObjectFullForm{}
	req := internalRequest{
		endpoint:            objectSearchUrl,
		method:              "SearchFullForm",
		withRequest:         nil,
		withResponse:        resp,
		withQueryParams:     p.toQueryParam(),
		acceptedStatusCodes: []int{http.StatusOK},
		apiName:             "Object",
	}
	err := c.cli.executeRequest(req)
	return resp, err
}
