# Client Go SDK 1.0.0

A Go SDK for Client.

- API version: v1
- SDK version: 1.0.0

The public Abbey API. Used for integrating with Abbey and building interfaces to extend the Abbey platform. See https://docs.abbey.io for more information.

## Table of Contents

- [Authentication](#authentication)
- [Services](#services)

## Authentication

### Access Token

The client API uses a access token as a form of authentication.

The access token can be set when initializing the SDK like this:

```go

```

Or at a later stage:

```go

```

## Services

### GrantKitsService

#### ListGrantKits

List Grant Kits

#### CreateGrantKit

Create a Grant Kit

#### ValidateGrantKit

Validate a Grant Kit

#### GetGrantKitById

Retrieve a Grant Kit by ID

#### UpdateGrantKit

Update a Grant Kit

#### DeleteGrantKit

Delete a Grant Kit

### IdentitiesService

#### ListEnrichedIdentities

List all Identities with enriched metadata

#### CreateIdentity

Create an Identity

#### GetIdentity

Retrieve an Identity

#### UpdateIdentity

Update an Identity

#### DeleteIdentity

Delete an Identity

### DemoService

#### GetDemo

Get Demo

#### CreateDemo

Create Demo Access

#### DeleteDemo

Delete Demo Access
