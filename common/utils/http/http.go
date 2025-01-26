package http

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type HTTPClient struct {
	client *client.Client
}

func newHTTPClient() (*HTTPClient, error) {
	c, err := client.NewClient()
	if err != nil {
		return nil, err
	}
	return &HTTPClient{client: c}, nil
}

func Get(ctx context.Context, url string, headers map[string]string) (*protocol.Response, error) {
	hc, err := newHTTPClient()
	if err != nil {
		return nil, err
	}
	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	req.SetRequestURI(url)
	req.Header.SetMethod(consts.MethodGet)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	err = hc.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Post(ctx context.Context, url string, headers map[string]string, body []byte) (*protocol.Response, error) {
	hc, err := newHTTPClient()
	if err != nil {
		return nil, err
	}
	//var postArgs protocol.Args
	//postArgs.Set("name", "cloudwego") // Set post args
	//status, responseBody, err := hc.client.Post(context.Background(), nil, "http://localhost:8080/hello", &postArgs)
	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	req.SetRequestURI(url)
	req.Header.SetMethod(consts.MethodPost)
	req.SetBody(body)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	err = hc.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
