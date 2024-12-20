// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package authn // import "github.com/SolomonAIEngineering/backend-core-library/third-party/authn"

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	jose "gopkg.in/square/go-jose.v2"
)

type internalClient struct {
	client   *retryablehttp.Client
	baseURL  *url.URL
	username string
	password string
	origin   string
}

const (
	remove                = "DELETE"
	get                   = "GET"
	patch                 = "PATCH"
	post                  = "POST"
	put                   = "PUT"
	defaultRetryWaitMin   = 5 * time.Millisecond
	defaultRetryWaitMax   = 10 * time.Millisecond
	defaultRetryMax       = 5
	defaultRequestTimeout = 1 * time.Second
)

func newInternalClient(base, username, password, origin string, retryConfig *RetryConfig) (*internalClient, error) {
	// ensure that base ends with a '/', so ResolveReference() will work as desired
	if base[len(base)-1] != '/' {
		base = base + "/"
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = retryConfig.MaxRetries
	retryClient.RetryWaitMin = retryConfig.MinRetryWaitTime
	retryClient.RetryWaitMax = retryConfig.MaxRetryWaitTime
	retryClient.HTTPClient.Timeout = retryConfig.RequestTimeout

	return &internalClient{
		client:   retryClient,
		baseURL:  baseURL,
		username: username,
		password: password,
		origin:   origin,
	}, nil
}

// TODO: test coverage.
func (ic *internalClient) Key(kid string) ([]jose.JSONWebKey, error) {
	resp, err := http.Get(ic.absoluteURL("jwks"))
	if err != nil {
		return []jose.JSONWebKey{}, err
	}
	defer resp.Body.Close()

	if !isStatusSuccess(resp.StatusCode) {
		return []jose.JSONWebKey{}, fmt.Errorf("received %d from %s", resp.StatusCode, ic.absoluteURL("jwks"))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []jose.JSONWebKey{}, err
	}

	jwks := &jose.JSONWebKeySet{}

	err = json.Unmarshal(bodyBytes, jwks)
	if err != nil {
		return []jose.JSONWebKey{}, err
	}
	return jwks.Key(kid), nil
}

// GetAccount gets the account details for the specified account id.
func (ic *internalClient) GetAccount(id string) (*Account, error) {
	resp, err := ic.doWithAuth(get, "accounts/"+id, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := struct {
		Result Account `json:"result"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data.Result, nil
}

// Update updates the account with the specified id.
func (ic *internalClient) Update(id, username string) error {
	form := url.Values{}
	form.Add("username", username)

	_, err := ic.doWithAuth(patch, "accounts/"+id, strings.NewReader(form.Encode()))
	return err
}

// Signup signs up a user account and returns the jwt token associated with the record
// TODO: unit test.
func (ic *internalClient) Signup(username, password string) (string, error) {
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)

	response, err := ic.doWithAuth(post, "accounts", strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	data := struct {
		Result struct {
			Token string `json:"id_token"`
		} `json:"result"`
	}{}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return "", err
	}

	return data.Result.Token, nil
}

// Logout revokes the established session and refresh token by the application with the authentication service
// TODO: unit test.
func (ic *internalClient) Logout() error {
	_, err := ic.doWithAuth(remove, "session", nil)
	return err
}

// Login logs a user into the backend system
// TODO: unit test.
func (ic *internalClient) Login(username, password string) (string, error) {
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)

	response, err := ic.doWithAuth(post, "session", strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	data := struct {
		Result struct {
			Token string `json:"id_token"`
		} `json:"result"`
	}{}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return "", err
	}

	return data.Result.Token, nil
}

// LockAccount locks the account with the specified id.
func (ic *internalClient) LockAccount(id string) error {
	_, err := ic.doWithAuth(patch, "accounts/"+id+"/lock", nil)
	return err
}

// UnlockAccount unlocks the account with the specified id.
func (ic *internalClient) UnlockAccount(id string) error {
	_, err := ic.doWithAuth(patch, "accounts/"+id+"/unlock", nil)
	return err
}

// ArchiveAccount archives the account with the specified id.
func (ic *internalClient) ArchiveAccount(id string) error {
	_, err := ic.doWithAuth(remove, "accounts/"+id, nil)
	return err
}

// ImportAccount imports an existing account.
func (ic *internalClient) ImportAccount(username, password string, locked bool) (int, error) {
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)
	form.Add("locked", strconv.FormatBool(locked))

	resp, err := ic.doWithAuth(post, "accounts/import", strings.NewReader(form.Encode()))
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	data := struct {
		Result struct {
			ID int `json:"id"`
		} `json:"result"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return -1, err
	}

	return data.Result.ID, err
}

// ExpirePassword expires the users current sessions and flags the account for a required password change on next login.
func (ic *internalClient) ExpirePassword(id string) error {
	_, err := ic.doWithAuth(patch, "accounts/"+id+"/expire_password", nil)
	return err
}

// RequestPasswordReset initiates a password reset request.
func (ic *internalClient) RequestPasswordReset(username string) error {
	baseURI := "password/reset"
	req, err := retryablehttp.NewRequest(get, ic.absoluteURL(baseURI), nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()          // Get a copy of the query values.
	q.Add("username", username)   // Add a new value to the set.
	req.URL.RawQuery = q.Encode() // Encode and assign back to the original query.

	_, err = ic.doWithNoAuthAndRequest(req)
	return err
}

// ResetPassword changes a password given token.
func (ic *internalClient) ResetPassword(password, token string) (string, error) {
	form := url.Values{}
	form.Add("password", password)
	form.Add("token", token)

	response, err := ic.doWithNoAuth(post, "password", strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	data := struct {
		Result struct {
			Token string `json:"id_token"`
		} `json:"result"`
	}{}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return "", err
	}

	return data.Result.Token, nil
}

// ChangePassword changes a password given token.
func (ic *internalClient) ChangePassword(newPassword, currentPassword string) (string, error) {
	form := url.Values{}
	form.Add("password", newPassword)
	form.Add("currentPassword", currentPassword)

	response, err := ic.doWithNoAuth(post, "password", strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	data := struct {
		Result struct {
			Token string `json:"id_token"`
		} `json:"result"`
	}{}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return "", err
	}

	return data.Result.Token, nil
}

// ServiceStats returns the raw request from the /stats endpoint.
func (ic *internalClient) ServiceStats() (*http.Response, error) {
	return ic.doWithAuth(get, "stats", nil)
}

// ServerStats returns the raw request from the /metrics endpoint.
func (ic *internalClient) ServerStats() (*http.Response, error) {
	return ic.doWithAuth(get, "metrics", nil)
}

func (ic *internalClient) absoluteURL(path string) string {
	return ic.baseURL.ResolveReference(&url.URL{Path: path}).String()
}

// unused. this will eventually execute private admin actions.
func (ic *internalClient) get(path string, dest interface{}) (int, error) {
	resp, err := http.Get(ic.absoluteURL(path))
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	err = json.Unmarshal(bodyBytes, dest)
	if err != nil {
		return resp.StatusCode, err
	}
	return resp.StatusCode, nil
}

func (ic *internalClient) doWithNoAuthAndRequest(req *retryablehttp.Request) (*http.Response, error) {
	req.Header.Set("Origin", ic.origin)
	resp, err := ic.client.Do(req)
	if err != nil {
		return nil, err
	}

	if !isStatusSuccess(resp.StatusCode) {
		// try to parse the error response
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, fmt.Errorf("received %d from %s", resp.StatusCode, ic.absoluteURL(req.URL.Path))
		}

		errResp.StatusCode = resp.StatusCode
		errResp.URL = ic.absoluteURL(req.URL.Path)
		return nil, &errResp
	}

	return resp, nil
}

func (ic *internalClient) doWithAuth(verb string, path string, body io.Reader) (*http.Response, error) {
	req, err := retryablehttp.NewRequest(verb, ic.absoluteURL(path), body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(ic.username, ic.password)
	req.Header.Set("Origin", ic.origin)

	if verb == post || verb == patch || verb == put {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := ic.client.Do(req)
	if err != nil {
		return nil, err
	}
	if !isStatusSuccess(resp.StatusCode) {
		// try to parse the error response
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, fmt.Errorf("received %d from %s", resp.StatusCode, ic.absoluteURL(path))
		}

		errResp.StatusCode = resp.StatusCode
		errResp.URL = ic.absoluteURL(path)
		return nil, &errResp
	}
	return resp, nil
}

func (ic *internalClient) doWithNoAuth(verb string, path string, body io.Reader) (*http.Response, error) {
	req, err := retryablehttp.NewRequest(verb, ic.absoluteURL(path), body)
	if err != nil {
		return nil, err
	}

	// origin necessary for this
	req.Header.Set("Origin", ic.origin)
	if verb == post || verb == patch || verb == put {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := ic.client.Do(req)
	if err != nil {
		return nil, err
	}
	if !isStatusSuccess(resp.StatusCode) {
		// try to parse the error response
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, fmt.Errorf("received %d from %s", resp.StatusCode, ic.absoluteURL(path))
		}

		errResp.StatusCode = resp.StatusCode
		errResp.URL = ic.absoluteURL(path)
		return nil, &errResp
	}
	return resp, nil
}

func isStatusSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
