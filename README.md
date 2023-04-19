# pwdGen

This is a command-line password generator written in [Go](https://golang.com). It generates secure and random passwords of a specified length and complexity level.

## Usage
```bash
$ go run main.go -l [length] -c [complexity]
```

## Flags
```bash
-l [length] - Length of the password to be generated
-c [complexity] - Complexity level of the password to be generated
```

## Complexity Levels
```bash
[1] - Letters (lowercase and uppercase) & numbers
[2] - Letters (lowercase and uppercase), numbers and special characters
[3] - Letters (lowercase and uppercase), numbers and all special characters
```