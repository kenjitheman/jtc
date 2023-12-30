## json-to-csv simple sync/async converter

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

## Project structure:

```go
jtoc
│
├── go.mod
├── jtoc.go
├── LICENSE
└── README.md
```

## Installation

```sh
go get github.com/kenjitheman/jtoc
```

## Usage

```go
jtoc.Convert(jsonFilePath string, outputCsvFileName string)

jtoc.AsyncConvert(jsonFilePath string, outputCsvFileName string, numWorkers int)
```

- NumWorkers - num of concurrent goroutines | change it if you want/need more or
  less.

- Example:

```go
jtoc.Convert("../man.json", "../man.csv")

jtoc.AsyncConvert("../man.json", "../man1.csv", 8)
```

## Contributing

- Pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change.

## license

- [MIT](./LICENSE)
