package http

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type HTTPClient struct {
	client *client.Client
}

func newHTTPClient(useTLS bool) (*HTTPClient, error) {
	var options []config.ClientOption
	if useTLS {
		options = append(options, client.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}))
	}
	options = append(options, client.WithDialer(standard.NewDialer()))

	c, err := client.NewClient(options...)
	if err != nil {
		return nil, err
	}
	return &HTTPClient{client: c}, nil
}

func doRequest(ctx context.Context, method, url string, headers map[string]string, body []byte, useTLS bool) (*protocol.Response, error) {
	hc, err := newHTTPClient(useTLS)
	if err != nil {
		return nil, err
	}
	req := protocol.AcquireRequest()
	defer protocol.ReleaseRequest(req)
	res := protocol.AcquireResponse()
	defer protocol.ReleaseResponse(res)

	req.SetRequestURI(url)
	req.Header.SetMethod(method)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	if method == consts.MethodPost && body != nil {
		req.SetBody(body)
	}

	err = hc.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Get(ctx context.Context, url string, headers map[string]string, useTLS bool) (*protocol.Response, error) {
	if url == "" {
		return nil, errors.New("URL cannot be empty")
	}
	return doRequest(ctx, consts.MethodGet, url, headers, nil, useTLS)
}

func Post(ctx context.Context, url string, headers map[string]string, body []byte, useTLS bool) (*protocol.Response, error) {
	if url == "" {
		return nil, errors.New("URL cannot be empty")
	}
	return doRequest(ctx, consts.MethodPost, url, headers, body, useTLS)
}
