# Bexs

This project aims to develop a Rest API and a CLI interface to consult with the best routes.

# Getting Started

These instructions will provide you a copy of the project that can be run on your local machine for development and testing purposes. Consult deployment item for notes on how to deploy the project on a live system.

# Prerequisites

This package was created with go1.16 and all you need is the standard go library.

# Installing

This is what you need to install the application from the source code:

```bash
    git clone https://github.com/paraizofelipe/bexs
```

To build the docker version you can use the `Makefile`:

```bash
    make dk-build 
```

# Running the tests

Until I finish this README there is not so much Unit tests written.

But I will try to coverage unless 80% of unit tests for this code as soon as possible.

You can run tests like this:

```bash
    make test
```

To run this code locally for test purposes use:

```bash
    make start-api
```

or run the CLI:

```bash
   make start-cli
```

# Deployment

This codebase is cloud-native by design so you can use lots of environments to make this run anywhere you want.

But to make this even easier to you the codebase also provides a Dockerfile.

Deploy with docker:

```bash
    make dk-deploy
```

# API

## api/routes

### GET

Request 

```curl
    curl -X GET "http://localhost:3000/api/routes/?from=GRU&to=CDG"
```

Response

```json
    {
        "routes": [
            {
                "from": "GRU",
                "to": "BRC",
                "price": 10
            },
            {
                "from": "BRC",
                "to": "SCL",
                "price": 5
            },
            {
                "from": "SCL",
                "to": "ORL",
                "price": 20
            },
            {
                "from": "ORL",
                "to": "CDG",
                "price": 5
            }
        ],
        "total_price": 40
    }
```

### POST

Request

```curl
    curl -X POST "http://localhost:3000/api/routes/?from=GRU&to=CDG&price=1"
```

Response

```json
    {
        "msg": "route successfully created"
    }
```
