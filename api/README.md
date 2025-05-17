# API Documentation

This directory contains the OpenAPI specification for the Weather Subscription Service API.

## Overview

The OpenAPI specification (`openapi.yaml`) has been converted from the [original Swagger specification](https://github.com/mykhailo-hrynko/se-school-5/blob/c05946703852b277e9d6dcb63ffd06fd1e06da5f/swagger.yaml).

This conversion was done using [swagger2openapi](https://www.npmjs.com/package/swagger2openapi) to implement an API-first approach using the [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) library.

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

### Online Documentation

The API documentation is also available online at: [https://weather.devcontainer.click/doc/api](https://weather.devcontainer.click/doc/api) 