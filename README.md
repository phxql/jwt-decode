# jwt-decode
Small tool to decode JWTs, written in Go

## Features

* Prints out the payload in JSON formatting
* Converts `exp`, `nbf` and `iat` to human readable format

## Building

```
go build
```

## Running

Build it, then: 

```
// Will ask on stdin for token
jwt-decode

// Or provide jwt on commandline
jwt-decode $JWT
```

## License

[GPLv3](LICENSE)