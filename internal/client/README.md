# Client Go SDK 1.0.0

A Go SDK for Client.

- API version: v1
- SDK version: 1.0.0

The public Abbey API. Used for integrating with Abbey and building interfaces to extend the Abbey platform.
See https://docs.abbey.io for more information.

## Table of Contents

- [Authentication](#authentication)
- [Services](#services)

## Authentication

### Access Token

The client API uses a access token as a form of authentication.

The access token can be set when initializing the SDK like this:

```go
// Constructor initialization
```

Or at a later stage:

```go
// Setter initialization
```

## Services

### GrantKitsService

#### ListGrantKits

#### CreateGrantKit

#### GetGrantKitById

#### UpdateGrantKit

#### DeleteGrantKit

### IdentitiesService

#### ListEnrichedIdentities

#### CreateIdentity

#### GetIdentity

#### UpdateIdentity

#### DeleteIdentity

### DemoService

#### GetDemo

#### CreateDemo

#### DeleteDemo
