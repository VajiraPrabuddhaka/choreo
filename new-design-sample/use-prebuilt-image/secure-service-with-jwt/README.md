# Secure Greeter Service with JWT Authentication

This example demonstrates how to deploy a sample greeter service and leverage OpenChoreo's API management capabilities to implement JWT authentication, OAuth2 security, rate limiting, and CORS policies.

## Overview

This sample showcases OpenChoreo's security features and demonstrates the separation between Platform Engineer (PE) and Developer resources:

- **Platform Engineer**: Defines secure API classes with JWT authentication, OAuth2 scopes, rate limiting, CORS policies, and circuit breaker configurations
- **Developer**: Creates components, workloads, and services that automatically inherit these security policies

### Key Security Features Demonstrated

- **JWT Authentication**: Auth0-based JWT token validation with JWKS endpoint
- **OAuth2 Integration**: Scope-based authorization with "greet" scope requirement  
- **Rate Limiting**: Configurable request limits (5 requests/minute for authenticated, 5 requests/hour for public)
- **CORS Policies**: Cross-origin resource sharing with specific allowed origins and headers
- **Circuit Breaker**: Connection limits and fault tolerance

## Pre-requisites

- Kubernetes cluster with OpenChoreo installed
- The `kubectl` CLI tool installed

## File Structure

```
secure-service-with-jwt/
├── platform-classes.yaml          # PE resources (ServiceClass, APIClass)
├── greeter-service-with-jwt.yaml  # Developer resources (ComponentV2, Workload, Service, API)
└── README.md                      # This guide
```

## Step 1: Platform Engineer Setup

First, the Platform Engineer deploys the class templates that define deployment policies:

```bash
kubectl apply -f platform-classes.yaml
```

This creates:

- **ServiceClass** (`go-service-standard`): Defines resource limits, replicas, service templates, and includes a metrics sidecar container
- **APIClass** (`greeter-service-rest-api-standard`): Defines comprehensive security policies including:
  - JWT authentication with External IDP integration
  - OAuth2 scopes and token validation
  - Rate limiting at API level
  - CORS policies for cross-origin requests
  - Circuit breaker for fault tolerance

## Step 2: Deploy Developer Resources

Deploy the greeter service application:

```bash
kubectl apply -f greeter-service-with-jwt.yaml
```

This creates:

- **ComponentV2**: Component metadata identifying this as a Service type
- **Workload**: Container configuration for the greeter service with:
  - Pre-built container image (`ghcr.io/openchoreo/samples/greeter-service:latest`)
  - REST API endpoint on port 9090
  - OpenAPI 3.0 specification for the reading list API
  - Environment variables for GitHub integration
- **Service**: Runtime service configuration that:
  - References the `go-service-standard` ServiceClass
  - Defines a REST API using the `greeter-service-rest-api-standard` APIClass
  - Exposes the service at both Organization and Public levels
  - Maps the backend service port 9090 to the `/greeter` path

## Step 3: Expose the API Gateway

Port forward the OpenChoreo gateway service to access it locally:

```bash
kubectl port-forward -n choreo-system svc/external-gateway 8443:443 &
```

## Step 4: Test the Service

### Test Without Authentication (Public Endpoint)

Test the public endpoint (limited to 5 requests/minute):

```bash
curl -k https://development.choreoapis.localhost:8443/default/greeter-service/greeter/greet
```

### Test With JWT Authentication

For authenticated requests (5 requests/minute), you'll need a valid JWT token from the Auth0 endpoint:

```bash
# First, obtain a JWT token (replace with your Auth0 credentials)
TOKEN=$(curl -s -X POST "https://dev-tfsf6412a2bn011a.us.auth0.com/oauth/token" \
  -H "Content-Type: application/json" \
  -d '{
    "client_id": "YOUR_CLIENT_ID",
    "client_secret": "YOUR_CLIENT_SECRET", 
    "audience": "openchoreo:greeter:service",
    "grant_type": "client_credentials",
    "scope": "greet"
  }' | jq -r '.access_token')

# Use the token to make authenticated requests
curl -k -H "Authorization: Bearer $TOKEN" \
  https://development.choreoapis.localhost:8443/default/greeter-service/greeter/greet
```

> [!TIP]
> #### Verification
>
> You should receive a successful response from the greeter service:
> ```
> Hello, Stranger!
> ```
>
> **Rate Limiting**: Notice that public requests are limited to 5/hour while authenticated requests allow 5/minute.
> **CORS**: The API supports cross-origin requests from `https://example.com` and `https://app.example.com`.
> **Circuit Breaker**: The service has a maximum of 2 concurrent connections for fault tolerance.

## Clean Up

Remove all resources:

```bash
# Remove developer resources
kubectl delete -f greeter-service-with-jwt.yaml

# Remove platform classes (optional, as they're shared)
kubectl delete -f platform-classes.yaml
```
