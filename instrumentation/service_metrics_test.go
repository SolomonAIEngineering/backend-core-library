// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package instrumentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServiceBaseMetrics(t *testing.T) {
	serviceName := "test-service"
	metrics, err := newServiceBaseMetrics(&serviceName)

	assert.NoError(t, err)
	assert.NotNil(t, metrics)
	assert.Equal(t, metrics.RequestCountMetric.MetricName, "test-service.grpc.request.count")
	assert.Equal(t, metrics.RequestLatencyMetric.MetricName, "test-service.grpc.request.latency")
	assert.Equal(t, metrics.ErrorCountMetric.MetricName, "test-service.error.count")
	assert.Equal(t, metrics.RequestStatusSummaryMetric.MetricName, "test-service.grpc.request.summary")
	assert.Equal(t, metrics.DbOperationCounter.MetricName, "test-service.db.operation.counter")
	assert.Equal(t, metrics.DbOperationLatency.MetricName, "test-service.db.operation.latency")
	assert.Equal(t, metrics.GrpcRequestCounter.MetricName, "test-service.grpc.request.counter")
	assert.Equal(t, metrics.GrpcRequestLatency.MetricName, "test-service.grpc.request.latency")
}
