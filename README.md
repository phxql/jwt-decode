# ARCHIVED!

See [this tool](https://github.com/mike-engel/jwt-cli) for a nice alternative.

# jwt-decode
Small tool to decode JWTs, written in Go

## Features

* Prints out the header and the payload in JSON formatting
* Converts `exp`, `nbf` and `iat` to human readable format

## Building

```
go build
```

## Running

Build it, then run: 

```
# Will ask on stdin for token
jwt-decode

# Or provide jwt on commandline
jwt-decode $JWT
```

## Example

```
# Run with commandline argument
jwt-decode eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzY290Y2guaW8iLCJleHAiOjEzMDA4MTkzODAsIm5hbWUiOiJDaHJpcyBTZXZpbGxlamEiLCJhZG1pbiI6dHJ1ZX0.03f329983b86f7d9a9f5fef85305880101d5e302afafa20154d094b229f75773

# Output
{
  "alg": "HS256",
  "typ": "JWT"
}
{
  "admin": true,
  "exp": "2011-03-22 19:43:00 +0100 CET",
  "iss": "scotch.io",
  "name": "Chris Sevilleja"
}
```

## License

[GPLv3](LICENSE)
