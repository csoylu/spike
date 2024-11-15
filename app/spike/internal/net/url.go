//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"net/url"

	"github.com/spiffe/spike/app/spike/internal/env"
)

// UrlSecretGet returns the URL for getting a secret.
func UrlSecretGet() string {
	u, _ := url.JoinPath(env.NexusApiRoot(), "/v1/secrets?action=get")
	return u
}

// UrlSecretPut returns the URL for putting a secret.
func UrlSecretPut() string {
	u, _ := url.JoinPath(env.NexusApiRoot(), "/v1/secrets")
	return u
}

// UrlSecretDelete returns the URL for deleting a secret.
func UrlSecretDelete() string {
	u, _ := url.JoinPath(env.NexusApiRoot(), "/v1/secrets?action=delete")
	return u
}

// UrlSecretUndelete returns the URL for undeleting a secret.
func UrlSecretUndelete() string {
	u, _ := url.JoinPath(env.NexusApiRoot(), "/v1/secrets?action=undelete")
	return u
}

// UrlSecretList returns the URL for listing secrets.
func UrlSecretList() string {
	u, _ := url.JoinPath(env.NexusApiRoot(), "/v1/secrets?action=list")
	return u
}

// UrlInit returns the URL for initializing SPIKE Nexus.
func UrlInit() string {
	u, _ := url.JoinPath(env.NexusApiRoot(), "/v1/init")
	return u
}

// UrlInitState returns the URL for checking the initialization state of
// SPIKE Nexus.
func UrlInitState() string {
	u, _ := url.JoinPath(env.NexusApiRoot(), "/v1/init?action=check")
	return u
}

// UrlAdminLogin returns the URL for logging in as an admin.
func UrlAdminLogin() string {
	u, _ := url.JoinPath(env.NexusApiRoot(), "/v1/login?action=admin")
	return u
}
