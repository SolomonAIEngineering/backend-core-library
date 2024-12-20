// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package authn // import "github.com/SolomonAIEngineering/backend-core-library/third-party/authn"

import (
	"errors"
	"net/url"
	"time"

	jwt "gopkg.in/square/go-jose.v2/jwt"
)

var (
	// ErrNoKey is returned when no key is found in the keychain.
	ErrNoKey = errors.New("no keys found")
)

// A JWT Claims extractor (JWTClaimsExtractor) implementation
// which extracts claims from Authn idToken.
type idTokenVerifier struct {
	audience  jwt.Audience
	keychain  JWKProvider
	issuerURL *url.URL
}

// NewIDTokenVerifier creates a new idTokenVerifier object by using keychain as the JWK provider
// Claims are verified against the values specified in config.
func NewIDTokenVerifier(issuer, audience string, keychain JWKProvider) (JWTClaimsExtractor, error) {
	return newIDTokenVerifierWithAudiences(issuer, jwt.Audience{audience}, keychain)
}

// newIDTokenVerifierWithAudiences creates a new idTokenVerifier object by using keychain as the JWT provider
// Claims are verified against issuer and the set of audiences.
func newIDTokenVerifierWithAudiences(issuer string, audiences jwt.Audience, keychain JWKProvider) (*idTokenVerifier, error) {
	issuerURL, err := url.Parse(issuer)
	if err != nil {
		return nil, err
	}

	return &idTokenVerifier{
		audience:  audiences,
		keychain:  keychain,
		issuerURL: issuerURL,
	}, nil
}

// Gets verified claims from an Authn idToken.
func (verifier *idTokenVerifier) GetVerifiedClaims(idToken string) (*jwt.Claims, error) {
	var err error

	claims, err := verifier.claims(idToken)
	if err != nil {
		return nil, err
	}

	err = verifier.verify(claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// Gets claims object from an idToken using the key from keychain
// Key from keychain is fetched using KeyID found in idToken's header.
func (verifier *idTokenVerifier) claims(idToken string) (*jwt.Claims, error) {
	var err error

	idJwt, err := jwt.ParseSigned(idToken)
	if err != nil {
		return nil, err
	}

	headers := idJwt.Headers
	if len(headers) != 1 {
		return nil, errors.New("multi-signature JWT not supported or missing headers information")
	}
	keyID := headers[0].KeyID
	keys, err := verifier.keychain.Key(keyID)
	if err != nil {
		return nil, err
	}
	if len(keys) == 0 {
		return nil, ErrNoKey
	}
	key := keys[0]

	claims := &jwt.Claims{}
	err = idJwt.Claims(key, claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// Verify the claims against the configured values.
func (verifier *idTokenVerifier) verify(claims *jwt.Claims) error {
	// Validate rest of the claims
	var err = claims.Validate(jwt.Expected{
		Issuer:   verifier.issuerURL.String(),
		Time:     time.Now(),
		Audience: verifier.audience,
	})
	if err != nil {
		return err
	}

	return nil
}
