// Code generated by apic. DO NOT EDIT.

package endpoint_configurations

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v2"
)

// Endpoint Configuration managementAn
// (https://ngrok.com/docs/ngrok-link#api-endpoint-configurations)Endpoint
// Configuration describes
// a ngrok network endpoint instance.Endpoints are your gateway to ngrok features!

type Client struct {
	apiClient *ngrok.Client
}

func NewClient(apiClient *ngrok.Client) *Client {
	return &Client{apiClient: apiClient}
}

// Create a new endpoint configuration
//
// https://ngrok.com/docs/api#api-endpoint-configurations-create
func (c *Client) Create(ctx context.Context, arg *ngrok.EndpointConfigurationCreate) (*ngrok.EndpointConfiguration, error) {
	if arg == nil {
		arg = new(ngrok.EndpointConfigurationCreate)
	}
	var res ngrok.EndpointConfiguration
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/endpoint_configurations")).Execute(&path, arg); err != nil {
		panic(err)
	}
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	bodyArg = arg

	if err := c.apiClient.Do(ctx, "POST", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Delete an endpoint configuration. This operation will fail if the endpoint
// configuration is still referenced by any reserved domain or reserved address.
//
// https://ngrok.com/docs/api#api-endpoint-configurations-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/endpoint_configurations/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()

	if err := c.apiClient.Do(ctx, "DELETE", apiURL, bodyArg, nil); err != nil {
		return err
	}
	return nil
}

// Returns detailed information about an endpoint configuration
//
// https://ngrok.com/docs/api#api-endpoint-configurations-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.EndpointConfiguration, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.EndpointConfiguration
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/endpoint_configurations/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()

	if err := c.apiClient.Do(ctx, "GET", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Returns a list of all endpoint configurations on this account
//
// https://ngrok.com/docs/api#api-endpoint-configurations-list
func (c *Client) list(ctx context.Context, arg *ngrok.Paging) (*ngrok.EndpointConfigurationList, error) {
	if arg == nil {
		arg = new(ngrok.Paging)
	}
	var res ngrok.EndpointConfigurationList
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/endpoint_configurations")).Execute(&path, arg); err != nil {
		panic(err)
	}
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	queryVals := make(url.Values)
	if arg.BeforeID != nil {
		queryVals.Set("before_id", *arg.BeforeID)
	}
	if arg.Limit != nil {
		queryVals.Set("limit", *arg.Limit)
	}
	apiURL.RawQuery = queryVals.Encode()

	if err := c.apiClient.Do(ctx, "GET", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Returns a list of all endpoint configurations on this account
//
// https://ngrok.com/docs/api#api-endpoint-configurations-list
func (c *Client) List(ctx context.Context, paging *ngrok.Paging) *Iter {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	if paging.Limit == nil {
		paging.Limit = ngrok.String("100")
	}
	return &Iter{
		client:     c,
		ctx:        ctx,
		limit:      paging.Limit,
		lastItemID: paging.BeforeID,
		n:          -1,
	}
}

// Iter allows the caller to iterate through a list of values while
// automatically fetching new pages worth of values from the API.
type Iter struct {
	client     *Client
	ctx        context.Context
	n          int
	items      []ngrok.EndpointConfiguration
	err        error
	limit      *string
	lastItemID *string
}

// Next() returns true if there is another value available in the iterator. If it
// returs true it also advances the iterator to that next available item.
func (it *Iter) Next() bool {
	// no more if there is an error
	if it.err != nil {
		return false
	}

	// advance the iterator
	it.n += 1

	// is there an available item?
	if it.n < len(it.items) {
		it.lastItemID = ngrok.String(it.Item().ID)
		return true
	}

	// fetch the next page
	resp, err := it.client.list(it.ctx, &ngrok.Paging{
		BeforeID: it.lastItemID,
		Limit:    it.limit,
	})
	if err != nil {
		it.err = err
		return false
	}

	// page with zero items means there are no more
	if len(resp.EndpointConfigurations) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.EndpointConfigurations
	return it.Next()
}

// Item() returns the EndpointConfiguration currently
// pointed to by the iterator.
func (it *Iter) Item() *ngrok.EndpointConfiguration {
	return &it.items[it.n]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *Iter) Err() error {
	return it.err
}

// Updates an endpoint configuration. If a module is not specified in the update,
// it will not be modified. However, each module configuration that is specified
// will completely replace the existing value. There is no way to delete an
// existing module via this API, instead use the delete module API.
//
// https://ngrok.com/docs/api#api-endpoint-configurations-update
func (c *Client) Update(ctx context.Context, arg *ngrok.EndpointConfigurationUpdate) (*ngrok.EndpointConfiguration, error) {
	if arg == nil {
		arg = new(ngrok.EndpointConfigurationUpdate)
	}
	var res ngrok.EndpointConfiguration
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/endpoint_configurations/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	bodyArg = arg

	if err := c.apiClient.Do(ctx, "PATCH", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
