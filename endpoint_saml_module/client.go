// Code generated by apic. DO NOT EDIT.

package endpoint_saml_module

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go"
)

type Client struct {
	apiClient *ngrok.Client
}

func NewClient(apiClient *ngrok.Client) *Client {
	return &Client{apiClient: apiClient}
}

func (c *Client) Replace(
	ctx context.Context,
	arg *ngrok.EndpointSAMLReplace,
) (*ngrok.EndpointSAML, error) {
	var res ngrok.EndpointSAML
	var path bytes.Buffer
	if err := template.Must(template.New("replace_path").Parse("/endpoint_configurations/{{ .ID }}/saml")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	bodyArg = arg.Module

	if err := c.apiClient.Do(ctx, "PUT", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) Get(
	ctx context.Context,
	id string,

) (*ngrok.EndpointSAML, error) {
	arg := &ngrok.Item{ID: id}
	var res ngrok.EndpointSAML
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/endpoint_configurations/{{ .ID }}/saml")).Execute(&path, arg); err != nil {
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

func (c *Client) Delete(
	ctx context.Context,
	id string,

) error {
	arg := &ngrok.Item{ID: id}
	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/endpoint_configurations/{{ .ID }}/saml")).Execute(&path, arg); err != nil {
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
