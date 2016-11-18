# IAM Sample Client in Go

## Prerequisites

- Docker

or

- Go

In addition to the aforementioned development dependencies, you will need the following:

- Client ID / Secret
- Client authorized to access api.cimpress.io
- Permissions configured in IAM

## Getting started

First and foremost, populate a `.env` file with client ID and client secret:

```
$ cat > .env <<EOF
CLIENT_ID=<client-id>
CLIENT_SECRET=<client-secret>
EOF
```

### Using Docker

```
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.6 bash -c make
```

### Using locally installed Go

Install godotenv dependency:

```
$ go get github.com/joho/godotenv
```

Run the code:

```
$ go run main.go
```
