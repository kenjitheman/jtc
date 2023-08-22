<h2 align="center">json-to-csv simple sync/async converter tool | library</h2>

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

## project structure:

```
├── cmd
│   └── main.go
├── core
│   └── core.go
├── go.mod
└── README.md
```

## installation

- use go get:

```
go get github.com/kenjitheman/json-to-csv
```

## usage

```
Convert(jsonFilePath string, outputCsvFileName string)

AsyncConvert(jsonFilePath string, outputCsvFileName string)
```

- example:

```
core.Convert("../man.json", "../man.csv")

core.AsyncConvert("../man.json", "../man1.csv")
```

- also you can change the number of concurrent goroutines:

```
const numWorkers = 4 // num of concurrent goroutines | change it if you want/need more or less
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
