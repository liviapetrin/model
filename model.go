// Copyright 2018 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY Type, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"time"
)

const (
	// RetryCount is used in a few hard-coded retries.
	RetryCount = 3
	// LevelError represents a log of 'error' level.
	LevelError = byte(0)
	// LevelWarn represents a log of 'warn' level.
	LevelWarn = byte(1)
	// LevelInfo represents a log of 'info' level.
	LevelInfo = byte(2)
	// LevelDebug represents a log of 'debug' level.
	LevelDebug = byte(3)
	// TransportPackageTypeLog represents a package of type 'log'.
	TransportPackageTypeLog = byte(0)
	// TransportPackageTypeHiPriLog represents a package of type 'high priority log'.
	TransportPackageTypeHiPriLog = byte(1)
	// TransportPackageTypeHealhcheck represents a package of type 'healthcheck'.
	TransportPackageTypeHealhcheck = byte(2)
	// LogTypeLog represents a log of type 'log'.
	LogTypeLog = byte(0)
	// LogTypeAudit represents a log of type 'audit'.
	LogTypeAudit = byte(1)
	// ContextTypeID holds the context type ID.
	ContextTypeID = "t"
	// CorrelationIDField holds the correlation id field name.
	CorrelationIDField = "CorrelationID"
	// DeliveryMethodClientSpecified represents a client specified log delivery method.
	DeliveryMethodClientSpecified = byte(0)
	// DeliveryMethodRoundRobin represents a round robin log delivery method.
	DeliveryMethodRoundRobin = byte(1)
)

// TransportPackage holds data being transported to the server.
// ID: sequential number.
// Type: One of "TransportPackageType*".
// Data: Package specific reference to the concrete oject.
// Payload: Data variable serialized.
// RetryCount: Number of retries executed on this package.
type TransportPackage struct {
	ID         uint64
	Type       byte
	Data       interface{}
	Payload    []byte
	RetryCount byte
}

// CorrelationData contains common data related to correlated logs.
type CorrelationData struct {
	CorrelationID string
	Name          string
	Custom        map[string]interface{}
}

// LogData holds log data.
// Timestamp: Log timestamp.
// Level: One of "Level*".
// Type: One of "LogType*".
// Weight: Log weight.
// Message: Log Message.
// Context: User specific log context data.
// Error: Log error data.
// ContextMap: Context object serialized into a map.
// Linked: Linked LogData objects.
type LogData struct {
	Timestamp       time.Time
	Level           byte
	Type            byte
	Weight          int
	Message         string
	Error           error
	ContextMap      []interface{}
	CorrelationData *CorrelationData
	ContextMaps     map[string][]string // todo: remove and check how to pass to workers this info.
}

// LogGroup holds a collection of log data and its common data.
// CorrelationData: Logs correlation data.
// Logs: List of logs beloging to this group.
// TODO: have a common props map here with all common props values.
type LogGroup struct {
	CorrelationData *CorrelationData
	Logs            []*LogData
}

// LoggedData holds log data that is sent to the logging systems.
type LoggedData struct {
	Type    byte                   `json:"Type,omitempty"`
	Weight  int                    `json:"Weight,omitempty"`
	Message string                 `json:"Message,omitempty"`
	Error   error                  `json:"Error,omitempty"`
	Context map[string]interface{} `json:"Context,omitempty"`
}

// ClientConfig holds client logging configuration.
// Enabled: true if logging is enabled; false otherwise.
// AppName: Name of the logging app.
// Level: Logging level. One of "LogType*".
// Endpoint: Server endpoint.
// NumberOfConnections: Number of connections the client will keep with the server.
// NumberOfHiPriConnections: Number of high priority connections the client will keep with the server.
// NumberOfBackupConnections: Number of backup connections the client will keep with the server.
// NumberOfHiPriBackupConnections: Number of high priority connections the client will keep with the server.
// ConnectionResetInterval: Interval which the client will reset its connection to the server.
// ChannelSize: Size of the channel used to temporarily hold messages that will be sent to the server.
// OverflowChannelSize: Size of the overflow channel used to temporarily store messages in emergency scenarios.
// OverflowChannelLoggingLevel: Level of the messages that will be stored in the overflow channel.
// HipriLoggingLevel: Level of the nessages that will be sent over the high priority connections.
// HipriChannelSize: Size of the channel used to temporarily hold high priority messages that will be sent to the server.
// TargetMessageBatchSize: Number of messages that should be batched together before being sent to the server.
// SendBatchLogsInterval: The maximum interval which the batched log messages should be sent to the server.
// CommonLabels: Key-value data pairs that should be attached to every log message.
// ServerConfigName: Server configuration name the client should default to.
// HealthCheckInterval: Interval which the client will send a health check command to the server.
// HealthCheckFailureThreshold: Number of failed send healh check commands for a connection to be deemed unhealthy.
// UserRequestTimout: Used to estimate requests that timed out on clients. This value is used to set the 'timedout'
//   field in the request tracking log entry.
// ConnectionShutdownTimout: Maximum time to wait for the logs to drain during shutdown for each connection.
type ClientConfig struct {
	Enabled                        bool              `json:"enabled"`
	AppName                        string            `json:"appName"`
	Level                          byte              `json:"level"`
	Endpoint                       string            `json:"endpoint"`
	NumberOfConnections            int               `json:"numberOfConnections"`
	NumberOfHiPriConnections       int               `json:"numberOfHiPriConnections"`
	NumberOfBackupConnections      int               `json:"numberOfBackupConnections"`
	NumberOfHiPriBackupConnections int               `json:"numberOfHiPriBackupConnections"`
	ConnectionResetInterval        time.Duration     `json:"connectionResetInterval"`
	ChannelSize                    int               `json:"channelSize"`
	OverflowChannelSize            int               `json:"overflowChannelSize"`
	OverflowChannelLoggingLevel    byte              `json:"overflowChannelLoggingLevel"`
	HipriLoggingLevel              byte              `json:"hipriLoggingLevel"`
	HipriChannelSize               int               `json:"hipriChannelSize"`
	TargetMessageBatchSize         int               `json:"targetMessageBatchSize"`
	SendBatchLogsInterval          time.Duration     `json:"sendBatchLogsInterval"`
	CommonLabels                   map[string]string `json:"commonLabels"`
	ServerConfigGroup              string            `json:"serverConfigGroup"`
	ServerConfigName               string            `json:"serverConfigName"`
	HealthCheckInterval            time.Duration     `json:"healthCheckInterval"`
	HealthCheckFailureThreshold    int               `json:"healthCheckFailureThreshold"`
	RequestTrackingTimout          int               `json:"requestTrackingTimout"`
	ConnectionShutdownTimout       time.Duration     `json:"connectionShutdownTimout"`
	ProjectID                      string            `json:"ProjectID"`           // TODO: remove. Here just for direct logging tests.
	CredentialsFilePath            string            `json:"CredentialsFilePath"` // TODO: remove. Here just for direct logging tests.
}

// ServerConfigs ... TODO
// ServicePort holds the server port.
// ShutdownTimeout contains the timeout to shutdown the server.
// ReadTimeout holds the read timeout.
// WriteTimeout holds the write timeout.
// Logging contains the logging configs.
type ServerConfigs struct {
	ServicePort     int
	ShutdownTimeout string
	ReadTimeout     string
	WriteTimeout    string
	Logging         *ServerLoggingConfigs
}

// ServerLoggingConfigs ... TODO
type ServerLoggingConfigs struct {
	Enabled                bool
	DefaultConfigGroupName string
	DefaultConfigName      string
	DeliveryMethod         byte
	Configs                []*ServerLoggingConfig
}

// ServerLoggingConfig ... TODO
type ServerLoggingConfig struct {
	Group               string
	Name                string
	ProjectID           string
	CredentialsFilePath string
	Level               byte
	NumberOfWorkers     int
	MessagesChannelSize int
	ShutdownTimeout     time.Duration
}

// OpenConnectionDataRequest holds open connection request data.
// ClientID: Client provided ID.
// IsHiPri: true if the requesting connection should be high priority; false otherwise.
// ConfigName: Default server config name used for the connection.
// CommonLabels: Key-value data pairs that should be attached to every log message for this connection.
// ContextMaps: Key-value data pairs containing the maps for each context object.
type OpenConnectionDataRequest struct {
	ClientID      string
	IsHiPri       bool
	ClientConfigs *ClientConfig
	ContextMaps   map[string][]string
}

// OpenConnectionDataResponse holds open connection response data.
// ConnectionID: Server provided unique connecton ID.
// StreamingEndpoint: Server provided streaming endpoint the client should use to start the streaming connection.
type OpenConnectionDataResponse struct {
	ConnectionID      string
	StreamingEndpoint string
}

// ListConnectionResponse holds a list of connections response data.
// ClientID: Client provided unique client ID.
// ConnectionID: Server provided unique connecton ID.
type ListConnectionResponse struct {
	ClientID     string
	ConnectionID string
}

// GetConnectionResponse holds connection response data.
// IsActive: true if the connection is active; false otherwise.
// ClientID: Client provided unique client ID.
// ConnectionID: Server provided unique connecton ID.
// StreamingEndpoint: Server provided streaming endpoint the client should use to start the streaming connection.
// IsHiPri: true if the requesting connection should be high priority; false otherwise.
// ClientConfigs: Holds client logging configuration.
type GetConnectionResponse struct {
	IsActive          bool
	ClientID          string
	ConnectionID      string
	StreamingEndpoint string
	IsHiPri           bool
	ClientConfigs     *ClientConfig
	LastReceivedTime  string
}

// PostConnectionRequest holds post connection request data.
// IsActive: true if the connection is active; false otherwise.
// ClientConfigs: Holds client logging configuration.
type PostConnectionRequest struct {
	IsActive      bool
	ClientConfigs *ClientConfig
}

// PostConnectionResponse holds post connection response data.
// IsActive: true if the connection is active; false otherwise.
// ClientConfigs: Holds client logging configuration.
type PostConnectionResponse struct {
	IsActive      bool
	ClientConfigs *ClientConfig
}
