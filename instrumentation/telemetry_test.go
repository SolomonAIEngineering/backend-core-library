package instrumentation

import (
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestClient_NewChildSpan(t *testing.T) {
	// Create a new transaction
	txn := &newrelic.Transaction{}

	// Create a new client
	client := &Client{}

	// Start a new child span
	span := client.NewChildSpan("test", *txn)

	// Assert that the span was started correctly
	if span == nil {
		t.Errorf("Expected a newrelic.Segment, but got nil")
	}

	if span.Name != "test" {
		t.Errorf("Expected span name to be \"test\", but got %s", span.Name)
	}
}

func TestWithRequest(t *testing.T) {
	s := &Client{}
	req := &http.Request{}
	txn := &newrelic.Transaction{}
	reqWithTrace := s.WithRequest(req, *txn)
	if reqWithTrace != nil {
		t.Error("Expected nil request to be returned")
	}
}

func TestClient_StartExternalSegment(t *testing.T) {
	// create a mock transaction
	app := newrelic.Application{}
	txn := app.StartTransaction("test", nil, nil)

	// create a mock HTTP request
	req, err := http.NewRequest("GET", "http://example.com", nil)
	assert.NoError(t, err)

	// create a client
	client := &Client{}

	// start the external segment
	extSeg := client.StartExternalSegment(txn, req)

	// assert that the external segment is not nil
	assert.NotNil(t, extSeg)

	// end the transaction
	txn.End()
}

func TestClient_NewRoundtripper(t *testing.T) {
	// create a mock transaction
	app := newrelic.Application{}
	txn := app.StartTransaction("test", nil, nil)

	// create a client
	client := &Client{}

	// create a round tripper
	roundTripper := client.NewRoundtripper(txn)

	// assert that the round tripper is not nil
	assert.NotNil(t, roundTripper)

	// end the transaction
	txn.End()
}

func TestClient_StartDatastoreSegment(t *testing.T) {
	// create a mock transaction
	app := newrelic.Application{}
	txn := app.StartTransaction("test", nil, nil)

	// create a client
	client := &Client{}

	// start the datastore segment
	dsSeg := client.StartDatastoreSegment(txn, "select")

	// assert that the datastore segment is not nil
	assert.NotNil(t, dsSeg)

	// end the transaction
	txn.End()
}

func TestClient_StartRedisDatastoreSegment(t *testing.T) {
	// create a mock transaction
	app := newrelic.Application{}
	txn := app.StartTransaction("test", nil, nil)

	// create a client
	client := &Client{}

	// start the redis datastore segment
	dsSeg := client.StartRedisDatastoreSegment(txn, "get")

	// assert that the redis datastore segment is not nil
	assert.NotNil(t, dsSeg)

	// end the transaction
	txn.End()
}

func TestClient_StartNosqlDatastoreSegment(t *testing.T) {
	// create a mock transaction
	app := newrelic.Application{}
	txn := app.StartTransaction("test", nil, nil)

	// create a client
	client := &Client{}

	// start the nosql datastore segment
	dsSeg := client.StartNosqlDatastoreSegment(txn, "find", "my_collection")

	// assert that the nosql datastore segment is not nil
	assert.NotNil(t, dsSeg)

	// end the transaction
	txn.End()
}

func TestClient_StartMessageQueueSegment(t *testing.T) {
	// create a mock transaction
	app := newrelic.Application{}
	txn := app.StartTransaction("test", nil, nil)

	// create a client
	client := &Client{}

	// start the message queue segment
	mqSeg := client.StartMessageQueueSegment(txn, "my_queue")

	// assert that the message queue segment is not nil
	assert.NotNil(t, mqSeg)

	// end the transaction
	txn.End()
}

func TestClient_StartSegment(t *testing.T) {
	txn := &newrelic.Transaction{}
	name := "segment-name"
	client := &Client{}
	segment := client.StartSegment(txn, name)

	if segment == nil {
		t.Errorf("Expected StartSegment to return a non-nil *newrelic.Segment object")
	}
}

func TestClient_GetUnaryServerInterceptors(t *testing.T) {
	client := &Client{
		Logger: zap.NewNop(),
		client: &newrelic.Application{},
	}

	interceptors := client.GetUnaryServerInterceptors()
	if len(interceptors) != 1 {
		t.Errorf("Expected GetUnaryServerInterceptors to return an array with 2 elements, but got %d", len(interceptors))
	}
}

func TestClient_GetStreamServerInterceptors(t *testing.T) {
	client := &Client{
		Logger: zap.NewNop(),
		client: &newrelic.Application{},
	}

	interceptors := client.GetStreamServerInterceptors()
	if len(interceptors) != 1 {
		t.Errorf("Expected GetStreamServerInterceptors to return an array with 2 elements, but got %d", len(interceptors))
	}
}

func TestClient_GetUnaryClientInterceptors(t *testing.T) {
	client := &Client{
		Logger: zap.NewNop(),
	}

	interceptors := client.GetUnaryClientInterceptors()
	if len(interceptors) != 1 {
		t.Errorf("Expected GetUnaryClientInterceptors to return an array with 2 elements, but got %d", len(interceptors))
	}
}

func TestClient_GetStreamClientInterceptors(t *testing.T) {
	client := &Client{
		Logger: zap.NewNop(),
	}

	interceptors := client.GetStreamClientInterceptors()
	if len(interceptors) != 1 {
		t.Errorf("Expected GetStreamClientInterceptors to return an array with 2 elements, but got %d", len(interceptors))
	}
}
func TestConfigureNewrelicClient(t *testing.T) {
	logger := zap.NewNop()
	serviceName := "test-service"
	newrelicKey := randomString(40)
	enabled := false

	// create a new client and configure it
	client := &Client{
		ServiceName: serviceName,
		NewrelicKey: newrelicKey,
		Enabled:     enabled,
		Logger:      logger,
	}
	err := client.configureNrClient()

	// assert that there were no errors returned
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// assert that the client was configured correctly
	if client.client == nil {
		t.Error("expected client to be configured, but it was nil")
	}
}
