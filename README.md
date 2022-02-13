---
stack: GO
lang: all
---

# Multi Projects in Golang
to be..

## go build
```
go build main.go
```

## godotenv
```
go get github.com/joho/godotenv
```

## get version of module - analyzing a package to see if it's go module compatible
go list -m -version pkg-name
- example:
```
go list -m -versions github.com/dgrijalva/jwt-go
```

## Got a problem with lunch? GOPATH should be set to
```
export GOPATH=$GOROOT
unset GOROOT
```

##  GO111MODULE set "on" or "off"
```
go env -w GO111MODULE=off
```