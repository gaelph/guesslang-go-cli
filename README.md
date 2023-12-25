# Guesslang Go CLI

A CLI around [`robherley/guesslang-go`](https://github.com/robherley/guesslang-go)

## Dependencies

As with [`robherley/guesslang-go`](https://github.com/robherley/guesslang-go), you will
need to install [`libtensorflow`](https://www.tensorflow.org/install/lang_c) C API.

On macOS, using Hombrew :
```sh
brew install libtensorflow
```

On linux:
```sh
scripts/install_libtensorflow
```
## Building

Due to the tensoflow C API dependency, you will need to enable CGO and if
necessary, add the includes and link flags:
```
CGO_ENABLE=1
CGO_CFLAGS="-I$(brew --prefix)/include"
CGO_LDFLAGS="-L$(brew --prefix)/lib"
```

The usual `go build` and/or `go install` commands will do.

## Usage

Unlike the original python [yoeo/guesslang](https://github.com/yoeo/guesslang)
this CLI only supports reading input from `stdin`.

Therefore, the only way to use it is as so:
```sh
echo "SELECT * FROM Users LIMIT 10;" | guesslang-go-cli
# Outputs "sql" to stdout
```
