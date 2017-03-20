# Luminescence: Lifeband Backend (nwHacks 2017)

Golang backend server which manages operations concerning data retrieval and storage.

## Setup

To setup the dependencies, run:
```
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
go get -u github.com/lib/pq
go install
```

To build the server, run:
```
go build
```

Run the built executable:
```
./backend
```
