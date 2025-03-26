package vaultwarden

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Menfre01/vaultwarden-go/pkg/types"
	"github.com/Menfre01/vaultwarden-go/pkg/utils"
)

type ClientInterface interface {
	ListOrganizations(search string) (*types.BaseResponse[types.OrganizationResponse], error)
	ListCollectionsWithOrg(organizationId, search string) (*types.BaseResponse[types.CollectionResponse], error)
	ListItems(organizationId, collectionId, search string) (*types.BaseResponse[types.ItemResponse], error)
	CreateItem(item *types.Item) (*types.BaseResponse[types.ItemResponse], error)
	EditItem(id string, item *types.Item) (*types.BaseResponse[types.ItemResponse], error)
}

func NewClient(server string, httpCli *http.Client) ClientInterface {
	return &client{
		server:  server,
		httpCli: httpCli,
	}
}

type client struct {
	server  string
	httpCli *http.Client
}

func (c *client) makeRequest(method string, path string, headers map[string]string, bodyBytes []byte) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.server, path)
	var buf io.Reader
	if len(bodyBytes) > 0 {
		buf = bytes.NewBuffer(bodyBytes)
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return req, nil
}

func (c *client) ListOrganizations(search string) (*types.BaseResponse[types.OrganizationResponse], error) {
	path := "/list/object/organizations"
	if search != "" {
		path = fmt.Sprintf("%s?search=%s", path, search)
	}
	req, err := c.makeRequest(http.MethodGet, path, map[string]string{
		"Content-Type": "application/json",
	}, nil)
	if err != nil {
		return nil, err
	}
	var ret types.BaseResponse[types.OrganizationResponse]
	return &ret, utils.Send(c.httpCli, req, &ret)
}

func (c *client) ListCollectionsWithOrg(organizationId, search string) (*types.BaseResponse[types.CollectionResponse], error) {
	path := fmt.Sprintf("/list/object/org-collections?organizationId=%s", organizationId)
	if search != "" {
		path = fmt.Sprintf("%s&search=%s", path, search)
	}
	req, err := c.makeRequest(http.MethodGet, path, map[string]string{
		"Content-Type": "application/json",
	}, nil)
	if err != nil {
		return nil, err
	}
	var ret types.BaseResponse[types.CollectionResponse]
	return &ret, utils.Send(c.httpCli, req, &ret)
}

func (c *client) ListItems(organizationId, collectionId, search string) (*types.BaseResponse[types.ItemResponse], error) {
	path := fmt.Sprintf("/list/object/items?organizationId=%s&collectionId=%s", organizationId, collectionId)
	if search != "" {
		path = fmt.Sprintf("%s&search=%s", path, search)
	}
	req, err := c.makeRequest(http.MethodGet, path, map[string]string{
		"Content-Type": "application/json",
	}, nil)
	if err != nil {
		return nil, err
	}
	var ret types.BaseResponse[types.ItemResponse]
	return &ret, utils.Send(c.httpCli, req, &ret)
}

func (c *client) CreateItem(item *types.Item) (*types.BaseResponse[types.ItemResponse], error) {
	path := "/object/item"
	bodyBytes, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	req, err := c.makeRequest(http.MethodPost, path, map[string]string{
		"Content-Type": "application/json",
	}, bodyBytes)
	if err != nil {
		return nil, err
	}
	var ret types.BaseResponse[types.ItemResponse]
	return &ret, utils.Send(c.httpCli, req, &ret)
}

func (c *client) EditItem(id string, item *types.Item) (*types.BaseResponse[types.ItemResponse], error) {
	path := fmt.Sprintf("/object/item/%s", id)
	bodyBytes, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	req, err := c.makeRequest(http.MethodPut, path, map[string]string{
		"Content-Type": "application/json",
	}, bodyBytes)
	if err != nil {

		return nil, err
	}
	var ret types.BaseResponse[types.ItemResponse]
	return &ret, utils.Send(c.httpCli, req, &ret)
}
