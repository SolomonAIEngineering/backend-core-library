// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package authn // import "github.com/SolomonAIEngineering/backend-core-library/third-party/authn"

import (
	"errors"
	"net/http"
	"time"

	jwt "gopkg.in/square/go-jose.v2/jwt"
)

// TODO: jose/jwt references are all over the place. Refactor possible?

// ErrInvalidOptions is returned by SubjectFrom if invalid options are used.
var ErrInvalidOptions = errors.New("invalid options for SubjectFrom")

// Client provides JWT verification for ID tokens generated by the AuthN server. In the future it
// will also implement the server's private APIs (aka admin actions).
type Client struct {
	config   Config
	iclient  *internalClient
	kchain   *keychainCache
	verifier JWTClaimsExtractor
}

// RetryConfig provides a mechanism by which clients can configure http retries parameters.
type RetryConfig struct {
	MaxRetries       int
	MinRetryWaitTime time.Duration
	MaxRetryWaitTime time.Duration
	RequestTimeout   time.Duration
}

var (
	defaultRetryConfigs = RetryConfig{
		MaxRetries:       defaultRetryMax,
		MinRetryWaitTime: defaultRetryWaitMin,
		MaxRetryWaitTime: defaultRetryWaitMax,
		RequestTimeout:   defaultRequestTimeout,
	}
)

// New returns an initialized and configured Client.
func New(config Config, origin string, retryConfig *RetryConfig) (*Client, error) {
	var err error
	config.setDefaults()

	ac := Client{}

	ac.config = config

	ac.iclient, err = newInternalClient(config.PrivateBaseURL, config.Username, config.Password, origin, retryConfig)
	if err != nil {
		return nil, err
	}

	ac.kchain = newKeychainCache(time.Duration(config.KeychainTTL)*time.Minute, ac.iclient)
	ac.verifier, err = NewIDTokenVerifier(config.Issuer, config.Audience, ac.kchain)
	if err != nil {
		return nil, err
	}

	return &ac, nil
}

// SubjectFrom will return the subject inside the given idToken if and only if the token is a valid
// JWT that passes all verification requirements. The returned value is the AuthN server's account
// ID and should be used as a unique foreign key in your users data.
//
// If the JWT does not verify, the returned error will explain why. This is for debugging purposes.
func (ac *Client) SubjectFrom(idToken string) (string, error) {
	return ac.subjectFromVerifier(idToken, ac.verifier)
}

// SubjectFromWithAudience works like SubjectFrom but allows specifying a different
// JWT audience.
func (ac *Client) SubjectFromWithAudience(idToken string, audience jwt.Audience) (string, error) {
	verifier, err := newIDTokenVerifierWithAudiences(ac.config.Issuer, audience, ac.kchain)
	if err != nil {
		return "", err
	}

	return ac.subjectFromVerifier(idToken, verifier)
}

// ClaimsFrom will return all verified claims inside the given idToken
// if and only if the token is a valid JWT that passes all
// verification requirements. If the JWT does not verify, the returned
// error will explain why. This is for debugging purposes.
func (ac *Client) ClaimsFrom(idToken string) (*jwt.Claims, error) {
	return ac.claimsFromVerifier(idToken, ac.verifier)
}

// ClaimsFromWithAudience works like ClaimsFrom but allows
// specifying a different JWT audience.
func (ac *Client) ClaimsFromWithAudience(idToken string, audience jwt.Audience) (*jwt.Claims, error) {
	verifier, err := newIDTokenVerifierWithAudiences(ac.config.Issuer, audience, ac.kchain)
	if err != nil {
		return nil, err
	}
	return ac.claimsFromVerifier(idToken, verifier)
}

func (ac *Client) subjectFromVerifier(idToken string, verifier JWTClaimsExtractor) (string, error) {
	claims, err := verifier.GetVerifiedClaims(idToken)
	if err != nil {
		return "", err
	}
	return claims.Subject, nil
}

func (ac *Client) claimsFromVerifier(idToken string, verifier JWTClaimsExtractor) (*jwt.Claims, error) {
	claims, err := verifier.GetVerifiedClaims(idToken)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// GetAccount gets the account with the associated id.
func (ac *Client) GetAccount(id string) (*Account, error) { //Should this be a string or an int?
	return ac.iclient.GetAccount(id)
}

// Update updates the account with the associated id.
func (ac *Client) Update(id, username string) error {
	return ac.iclient.Update(id, username)
}

// LockAccount locks the account with the associated id.
func (ac *Client) LockAccount(id string) error {
	return ac.iclient.LockAccount(id)
}

// UnlockAccount unlocks the account with the associated id.
func (ac *Client) UnlockAccount(id string) error {
	return ac.iclient.UnlockAccount(id)
}

// ArchiveAccount archives the account with the associated id.
func (ac *Client) ArchiveAccount(id string) error {
	return ac.iclient.ArchiveAccount(id)
}

// ImportAccount imports an account with the provided information, returns the imported account id.
func (ac *Client) ImportAccount(username, password string, locked bool) (int, error) {
	return ac.iclient.ImportAccount(username, password, locked)
}

// ExpirePassword expires the password of the account with the associated id.
func (ac *Client) ExpirePassword(id string) error {
	return ac.iclient.ExpirePassword(id)
}

// LoginAccount attempts to log in the account with the input credentials and returns a jwt token.
func (ac *Client) LoginAccount(username, password string) (string, error) {
	return ac.iclient.Login(username, password)
}

// SignupAccount attempts to sign up the account with the input credentials and returns a jwt token.
func (ac *Client) SignupAccount(username, password string) (string, error) {
	return ac.iclient.Signup(username, password)
}

// LogOutAccount logs a user out of the systems by revoking all associated tokens to the account.
func (ac *Client) LogOutAccount() error {
	return ac.iclient.Logout()
}

// ServiceStats gets the http response object from calling the service stats endpoint.
func (ac *Client) ServiceStats() (*http.Response, error) {
	return ac.iclient.ServiceStats()
}

// ServerStats gets the http response object from calling the server stats endpoint.
func (ac *Client) ServerStats() (*http.Response, error) {
	return ac.iclient.ServerStats()
}

// RequestPasswordReset initiates a password reset request.
func (ac *Client) RequestPasswordReset(username string) error {
	return ac.iclient.RequestPasswordReset(username)
}

// ResetPassword resets a password based on the provided token.
func (ac *Client) ResetPassword(password, token string) (string, error) {
	return ac.iclient.ResetPassword(password, token)
}

// ChangePassword attempts to change a password while authenticated.
func (ac *Client) ChangePassword(neswPassword, oldPassword string) (string, error) {
	return ac.iclient.ChangePassword(neswPassword, oldPassword)
}

// DefaultClient can be initialized by Configure and used by SubjectFrom.
var DefaultClient *Client

func defaultClient() *Client {
	if DefaultClient == nil {
		panic("Please initialize DefaultClient using Configure")
	}
	return DefaultClient
}

// Configure initializes the default AuthN client with the given config. This is necessary to
// use lib.SubjectFrom without keeping a reference to your own AuthN client.
func Configure(config Config, origin string) error {
	client, err := New(config, origin, &defaultRetryConfigs)
	if err != nil {
		return err
	}
	DefaultClient = client
	return nil
}

// SubjectFrom will use the the client configured by Configure to extract a subject from the
// given idToken.
func SubjectFrom(idToken string) (string, error) {
	return defaultClient().SubjectFrom(idToken)
}
