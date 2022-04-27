# tram-cli project
Go CLI app to fetch data from Zaragoza's Tram and write them to CSV

# Usage

## Show help

```shell
go run ./cmd/tram-cli/main.go help
```

## Help of main command

```shell
go run ./cmd/tram-cli/main.go tram-stops help
```

## Tram stops status to csv

### To default output file
```shell
go run ./cmd/tram-cli/main.go tram-stops
```

### To custom output file
```shell
go run ./cmd/tram-cli/main.go tram-stops -o my_output.csv
```