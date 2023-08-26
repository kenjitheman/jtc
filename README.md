<h2 align="center">json-to-csv simple sync/async converter tool | library</h2>

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

## project structure:

```

```

## installation

- use go get:

```
go get github.com/kenjitheman/jtoc
```

## usage

```
jtoc.Convert(jsonFilePath string, outputCsvFileName string)

jtoc.AsyncConvert(jsonFilePath string, outputCsvFileName string, numWorkers int)
```

- numWorkers - num of concurrent goroutines | change it if you want/need more or
  less

- example:

```
jtoc.Convert("../man.json", "../man.csv")

jtoc.AsyncConvert("../man.json", "../man1.csv", 8)
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
