# IAM Sample Client in Go

## Prerequisites

You will need the following:

- Client ID / Secret
- Client authorized to access api.cimpress.io
- Permissions configured in IAM

## Getting started

Install godotenv dependency:

```
$ go get github.com/joho/godotenv
```

Populate a `.env` file with client ID and client secret

```
$ cat > .env <<EOF
CLIENT_ID=<client-id>
CLIENT_SECRET=<client-secret>
EOF
```

Run the code:

```
$ go run main.go
```
