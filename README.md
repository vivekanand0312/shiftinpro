# ShiftInPRO
### Mysql:
```sh
mysql -uroot -p
show databases;
use vdmnpDB;
```

### BUILD time errors check: 
```sh
1. go build ./...
```
> OR
```sh
2. go vet ./...
```
> OR
```sh
3. go test ./...
```

### RUN: 
```sh
go run ./cmd/server/main.go
```