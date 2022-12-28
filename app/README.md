This directory is the implementation of an app that listens to http requests, moves the robot to a specific location, and stores
the information in a database.

## Directory structure

```shell
├── README.md
├── logger
│   ├── logger.go
├── server
│   ├── http.go
│   └── http_test.go
├── storage
│     ├── client.go
│     ├── migrations
│     │   ├── 20221218110026_create_executions_table.go
│     │   └── main.go
│     └── models.go
└── svc
    └── service.go
```

The app is divided in three components: `server`, `storage`, `svc`.

* The `server` is responsible for instantiating a http server and handling requests.
* The `storage` is responsible for persisting information in a storage backend. By default, Postgres is used.
* The `svc` is responsible for the business logic of the app. The `svc` is the bridge between the `server` and the `storage`.
