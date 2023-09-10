// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package instrumentation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewRequestCountMetric(t *testing.T) {
	serviceName := "test-service"
	expected := &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.count", serviceName),
		ServiceName: serviceName,
		Help:        "Tracks the number of request serviced by the service partitioned by name and status code",
		Subsystem:   Subsystem(GrpcSubSystem),
		Namespace:   RequestNamespace,
	}

	result := newRequestCountMetric(&serviceName)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestNewRequestCountMetric() failed, expected: %+v, result: %+v", result, expected)
	}
}

func TestNewRequestLatencyMetric(t *testing.T) {
	serviceName := "test-service"
	expected := &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.latency", serviceName),
		ServiceName: serviceName,
		Help:        "Tracks the latency associated with various requests partitioned by service name, target name, status code, and latency",
		Subsystem:   Subsystem(GrpcSubSystem),
		Namespace:   RequestNamespace,
	}

	result := newRequestLatencyMetric(&serviceName)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestNewRequestLatencyMetric() failed, expected: %+v, result: %+v", result, expected)
	}
}

func TestNewErrorCountMetric(t *testing.T) {
	serviceName := "test-service"
	expected := &Metric{
		MetricName:  fmt.Sprintf("%s.error.count", serviceName),
		ServiceName: serviceName,
		Help:        "Tracks the number of errors encountered by the service",
		Subsystem:   Subsystem(ErrorSubSystem),
		Namespace:   ServiceNamespace,
	}

	result := newErrorCountMetric(&serviceName)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestNewErrorCountMetric() iled, expected: %+v, result: %+v", result, expected)
	}
}

func TestNewRequestStatusSummaryMetric(t *testing.T) {
	serviceName := "test-service"
	expected := &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.summary", serviceName),
		ServiceName: serviceName,
		Help:        "Tracks the status of all requests serviced by the service",
		Subsystem:   Subsystem(RequestNamespace),
		Namespace:   RequestNamespace,
	}

	result := newRequestStatusSummaryMetric(&serviceName)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestNewRequestStatusSummaryMric() failed, expected: %+v, result: %+v", result, expected)
	}
}

func TestNewDbOperationCounter(t *testing.T) {
	serviceName := "test-service"
	expected := &Metric{
		MetricName:  fmt.Sprintf("%s.db.operation.counter", serviceName),
		ServiceName: serviceName,
		Help:        "Tracks the number of db tx processed by the service",
		Subsystem:   Subsystem(DbSubSystem),
		Namespace:   DatabaseNamespace,
	}

	result := newDbOperationCounter(serviceName)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestNewDbOperationCounter failed, expected: %+v, result: %+v", result, expected)
	}
}
func TestNewDbOperationLatency(t *testing.T) {
	serviceName := "test-service"
	expected := &Metric{
		MetricName:  fmt.Sprintf("%s.db.operation.latency", serviceName),
		ServiceName: serviceName,
		Help:        "Tracks the latency of all db tx performed by the service.",
		Subsystem:   Subsystem(DbSubSystem),
		Namespace:   DatabaseNamespace,
	}

	result := newDbOperationLatency(serviceName)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestNewDbOperationLatency failed, expected: %+v, result: %+v", result, expected)
	}
}

func TestNewGrpcRequestCounter(t *testing.T) {
	serviceName := "test-service"
	expected := &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.counter", serviceName),
		ServiceName: serviceName,
		Help:        "Tracks the number of grpc requests processed by the service. Partitioned by status code and operation",
		Subsystem:   Subsystem(GrpcSubSystem),
		Namespace:   RequestNamespace,
	}

	result := newGrpcRequestCounter(serviceName)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestNewGrpcRequestCounter failed, expected: %+v, result: %+v", result, expected)
	}
}

func TestNewGrpcRequestLatency(t *testing.T) {
	serviceName := "test-service"
	expected := &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.latency", serviceName),
		ServiceName: serviceName,
		Help:        "Tracks the latency of all outgoing grpc requests initiated by the service. Partitioned by status code and operation",
		Subsystem:   Subsystem(GrpcSubSystem),
		Namespace:   RequestNamespace,
	}

	result := newGrpcRequestLatency(serviceName)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestNewGrpcRequestLatency failed, expected: %+v, result: %+v", result, expected)
	}
}
