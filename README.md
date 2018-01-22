# Luminescence: Lifeband Backend (nwHacks 2017)

[![Build Status](https://travis-ci.org/nwHacks2017/backend.svg?branch=master)](https://travis-ci.org/nwHacks2017/backend)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/nwHacks2017/backend/blob/master/LICENSE)

Golang backend server which manages operations concerning data retrieval and storage.

## Setup

First, an instance of the [Lifeband CockroachDB database](https://github.com/nwHacks2017/database) must already be running.

Navigate to the directory with the source files:
```
cd backend/
```

Setup the dependencies:
```
go get -t ./...
```

Create a configuration file (which will not be tracked by Git):
```
cp config.json.template config.json
```

Fill in the server port, current user, and database details in `config.json`. You will need to retrieve and store the SSL certificate and key of the database user to connect to.

Build the server:
```
go build
```

Run the built executable:
```
./backend
```
