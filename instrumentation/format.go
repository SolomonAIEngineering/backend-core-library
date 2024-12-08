// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package instrumentation // import "github.com/SolomonAIEngineering/backend-core-library/instrumentation"

import "fmt"

type MetricName struct {
	ServiceName     string
	OperationName   string
	IsDistributedTx bool
	IsDatabaseTx    bool
	IsError         bool
}

type MetricType string

const (
	Count         MetricType = "count"
	ErrorCount    MetricType = "error_count"
	Latency       MetricType = "latency"
	MetricSummary MetricType = "summary"
)

var (
	OperationLatencyMetric MetricType = "service.operation.latency"
	OperationStatusMetric  MetricType = "service.operation.status"
	RpcStatusMetric        MetricType = "service.rpc.status"
)

// Format formats metric name based on metric type.
func format(m *MetricName, metricType MetricType) *string {
	metric := m.ServiceName
	if m.IsDistributedTx {
		metric = fmt.Sprintf("%s.dtx", metric)
	}

	if m.IsDatabaseTx {
		metric = fmt.Sprintf("%s.db", metric)
	}

	metric = fmt.Sprintf("%s.op.%s", metric, m.OperationName)
	if m.IsError {
		metric = fmt.Sprintf("%s.error", metric)
	}

	metric = appenSuffix(metric, metricType)

	return &metric
}

// appenSuffix appends suffix to metric name.
func appenSuffix(metricName string, metricType MetricType) string {
	return fmt.Sprintf("%s.%s", metricName, metricType)
}
