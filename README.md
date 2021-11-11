# Maakfile

> This project is inspired by [npm-run-all](https://www.npmjs.com/package/npm-run-all) and [GNU Make](https://www.gnu.org/software/make/manual/make.html)
> 
> I wanted to be able to use a file to run all scripts from... with watch mode, with parallel execution


## Example

```sh
maak init
maak run <command>
maak list
maak [--help -h]
```

**Maakfile.toml**

```toml
[sequential]
# run `maak run get` to run `go get` then `go mod tidy`
get = ["go get", "go mod tidy"]
build = "go build -o build/maak main.go"
		
[parallel]
# run `maak run test` to run below echos and sleep in parallel
# parallel only allows an array of strings
test = ["echo 'this is a test'", "sleep 3", "echo currently sleeping"]
[watch]
# run `maak run watch-build` to run the the build script
# while watching all files ending with `.go` & `.mod`
build = [".go", ".mod"]
# this will watch all files
test = "*"
```

## Init

`maak init`

This will create a **Maakfile.toml** in the current directory

## Run

`maak run build`

This will run the build script documented in the Maakfile

### Without command

Running `maak run` without specifying a command will list all commands and then ask which command to run.

## List

`maak run list`

this will show a nice list of all the availible commands that are available for `maak run `.
