// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package instrumentation

import "testing"

func TestFormat(t *testing.T) {
	// create a test metric name
	metricName := &MetricName{
		ServiceName:     "test_service",
		OperationName:   "test_operation",
		IsDistributedTx: true,
		IsDatabaseTx:    true,
		IsError:         true,
	}

	// test formatting different metric types
	countMetric := format(metricName, Count)
	if *countMetric != "test_service.dtx.db.op.test_operation.error.count" {
		t.Errorf("format failed, expected: test_service.dtx.db.op.test_operation.error.count, got: %s", *countMetric)
	}

	errorCountMetric := format(metricName, ErrorCount)
	if *errorCountMetric != "test_service.dtx.db.op.test_operation.error.error_count" {
		t.Errorf("format failed, expected: test_service.dtx.db.op.test_operation.error.error_count, got: %s", *errorCountMetric)
	}

	latencyMetric := format(metricName, Latency)
	if *latencyMetric != "test_service.dtx.db.op.test_operation.error.latency" {
		t.Errorf("format failed, expected: test_service.dtx.db.op.test_operation.error.latency, got: %s", *latencyMetric)
	}

	summaryMetric := format(metricName, MetricSummary)
	if *summaryMetric != "test_service.dtx.db.op.test_operation.error.summary" {
		t.Errorf("format failed, expected: test_service.dtx.db.op.test_operation.error.summary, got: %s", *summaryMetric)
	}
}

func TestAppenSuffix(t *testing.T) {
	// test appending suffix to metric name
	metricName := appenSuffix("test_metric", Count)
	if metricName != "test_metric.count" {
		t.Errorf("appenSuffix failed, expected: test_metric.count, got: %s", metricName)
	}
}
