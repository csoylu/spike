//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package backend

import (
	"context"

	"github.com/spiffe/spike-sdk-go/api/entity/data"

	"github.com/spiffe/spike/pkg/store"
)

type DatabaseConfigKey string

const (
	KeyDataDir                DatabaseConfigKey = "data_dir"
	KeyDatabaseFile           DatabaseConfigKey = "database_file"
	KeyJournalMode            DatabaseConfigKey = "journal_mode"
	KeyBusyTimeoutMs          DatabaseConfigKey = "busy_timeout_ms"
	KeyMaxOpenConns           DatabaseConfigKey = "max_open_conns"
	KeyMaxIdleConns           DatabaseConfigKey = "max_idle_conns"
	KeyConnMaxLifetimeSeconds DatabaseConfigKey = "conn_max_lifetime_seconds"
)

// Backend defines the interface for secret storage and management backends
type Backend interface {
	// Initialize initializes the backend
	Initialize(ctx context.Context) error
	// Close closes the backend
	Close(ctx context.Context) error

	// StoreSecret stores a secret at the specified path
	StoreSecret(ctx context.Context, path string, secret store.Secret) error
	// LoadSecret loads a secret from the specified path
	LoadSecret(ctx context.Context, path string) (*store.Secret, error)

	// StorePolicy stores a policy object in the backend storage.
	StorePolicy(ctx context.Context, policy data.Policy) error

	// LoadPolicy retrieves a policy by its ID from the backend storage.
	// It returns the policy object and an error, if any.
	LoadPolicy(ctx context.Context, id string) (*data.Policy, error)

	// DeletePolicy removes a policy object identified by the given Id from
	// storage.
	// ctx is the context for managing cancellations and timeouts.
	// id is the identifier of the policy to delete.
	// Returns an error, if the operation fails.
	DeletePolicy(ctx context.Context, id string) error

	// StoreKeyRecovery stores the encrypted key recovery data blob.
	// The data includes the root key and its Shamir shards.
	StoreKeyRecovery(ctx context.Context, meta store.KeyRecoveryData) error

	// LoadKeyRecovery retrieves the encrypted key recovery data blob.
	// Returns the nonce and encrypted data containing the root key and its
	// Shamir shards.
	LoadKeyRecovery(ctx context.Context) (meta store.KeyRecoveryData, err error)
}

// Config holds configuration for backend initialization
type Config struct {
	// Common configuration fields
	EncryptionKey string
	Location      string // Could be a file path, S3 bucket, etc.

	// Backend-specific configuration
	Options map[DatabaseConfigKey]any
}

// Factory creates a new backend instance
type Factory func(cfg Config) (Backend, error)
