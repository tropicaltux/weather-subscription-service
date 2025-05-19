# API Documentation

This directory contains the OpenAPI specification for the Weather Subscription Service API and configuration for code generation.

## Overview

The OpenAPI specification (`openapi.yaml`) has been converted from the [original Swagger specification](https://github.com/mykhailo-hrynko/se-school-5/blob/c05946703852b277e9d6dcb63ffd06fd1e06da5f/swagger.yaml).

This conversion was done using [swagger2openapi](https://www.npmjs.com/package/swagger2openapi) to implement an API-first approach using the [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) library.

A standardized error response schema has been added to provide consistent error handling across all API endpoints.

## Viewing the API Documentation

You can view the API documentation in several ways:

### Using Docker Compose

To start the Swagger UI locally, run the following command **from the project root directory**:

```bash
docker compose up swagger-ui
```

Or to run it in detached mode:

```bash
docker compose up swagger-ui -d
```

The documentation will then be available at: [http://localhost:8080](http://localhost:8080)

## Code Generation

The repository includes a configuration file for [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) to generate server code from the OpenAPI specification. See the [configuration file](oapi-codegen.yaml) for details.

### Generating API Code

#### Using Command Line

To generate server code from the OpenAPI specification, run:

```bash
oapi-codegen -config api/oapi-codegen.yaml api/openapi.yaml
```

#### Using VS Code Tasks

1. Press `Ctrl+Shift+P` (or `Cmd+Shift+P` on macOS)
2. Select "Tasks: Run Task"
3. Choose "Generate API Server"

### Generated Code Location

The generated server code is located in the [internal/api](../internal/http/api) directory. 