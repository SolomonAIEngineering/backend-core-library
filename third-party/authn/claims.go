package authn // import "github.com/SimifiniiCTO/simfiny-core-lib/third-party/authn"

import "gopkg.in/square/go-jose.v2/jwt"

// Claims represents the claims in an Authn idToken.
type Claims struct {
	// The time before which the JWT MUST NOT be accepted for processing.
	AuthTime *jwt.NumericDate `json:"auth_time"`
	jwt.Claims
}
