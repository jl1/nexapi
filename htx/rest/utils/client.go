/*
 * Copyright (c) 2023, LinstoHu
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/go-playground/validator"
	goquery "github.com/google/go-querystring/query"
)

type HTXClient struct {
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	baseURL     string
	key, secret string
	signVersion string
}

type HTXClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL     string `validate:"required"`
	Key         string
	Secret      string
	SignVersion string
}

func NewHTXRestClient(cfg *HTXClientCfg) (*HTXClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := HTXClient{
		debug:       cfg.Debug,
		logger:      cfg.Logger,
		baseURL:     cfg.BaseURL,
		key:         cfg.Key,
		secret:      cfg.Secret,
		signVersion: cfg.SignVersion,
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	return &cli, nil
}

func (htx *HTXClient) GetDebug() bool {
	return htx.debug
}

func (htx *HTXClient) GetBaseURL() string {
	return htx.baseURL
}

func (htx *HTXClient) GetKey() string {
	return htx.key
}

func (htx *HTXClient) GetSecret() string {
	return htx.secret
}

func (htx *HTXClient) GetHeaders() (map[string]string, error) {
	return map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}, nil
}

func (htx *HTXClient) GenSignatureValues(req HTTPRequest) (url.Values, error) {
	parameters := url.Values{}

	if req.QueryParams != nil {
		q, err := goquery.Values(req.QueryParams)
		if err != nil {
			return nil, err
		}
		parameters = q
	}

	parameters.Add("AccessKeyId", htx.key)
	parameters.Add("SignatureMethod", "HmacSHA256")
	parameters.Add("SignatureVersion", htx.signVersion)
	parameters.Add("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05"))

	return parameters, nil
}

func (htx *HTXClient) NormalizeRequestContent(req HTTPRequest, parameters url.Values) (string, error) {
	if req.Method == "" || req.BaseURL == "" || req.Path == "" || parameters.Encode() == "" {
		return "", fmt.Errorf("gen signature error: method(%s), baseurl(%s), path(%s) and parameters(%s) should not be empty",
			req.Method, req.BaseURL, req.Path, parameters.Encode())
	}

	url, err := url.Parse(req.BaseURL + req.Path)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(req.Method)
	sb.WriteString("\n")
	sb.WriteString(url.Host)
	sb.WriteString("\n")
	sb.WriteString(req.Path)
	sb.WriteString("\n")
	sb.WriteString(parameters.Encode())

	return sb.String(), nil
}

// sign makes a signature by sha256.
func (htx *HTXClient) Sign(plain []byte) string {
	hm := hmac.New(sha256.New, []byte(htx.secret))
	hm.Write(plain)
	return base64.StdEncoding.EncodeToString(hm.Sum(nil))
}

func (htx *HTXClient) SendHTTPRequest(ctx context.Context, req HTTPRequest) (*HTTPResponse, error) {
	client := http.Client{}

	var body io.Reader
	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonBody)
	}

	url, err := url.Parse(req.BaseURL + req.Path)
	if err != nil {
		return nil, err
	}

	if req.Query != nil {
		url.RawQuery = req.Query.Encode()
	}

	request, err := http.NewRequestWithContext(ctx, req.Method, url.String(), body)
	if err != nil {
		return nil, err
	}

	for k, v := range req.Headers {
		request.Header.Set(k, v)
	}

	if htx.GetDebug() {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}

		htx.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if htx.GetDebug() {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		htx.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	return NewResponse(&req, resp, nil), nil
}
