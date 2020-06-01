package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"os"
	"strings"
	"time"
)

func main() {
	text, err := readJwt()
	if err != nil {
		panic(err)
	}

	token, err := parseToken(text)
	if err != nil {
		panic(err)
	}
	claims := token.Claims.(jwt.MapClaims)
	convertDateKeys(claims)

	header, err := convertToJson(token.Header)
	if err != nil {
		panic(err)
	}
	fmt.Println(header)

	payload, err := convertToJson(claims)
	if err != nil {
		panic(err)
	}
	fmt.Println(payload)
}

// Parses the given string to a jwt.Token
func parseToken(text string) (*jwt.Token, error) {
	token, err := jwt.Parse(text, nil)
	if err != nil {
		detailed := err.(*jwt.ValidationError)
		// We don't care about validating the JWT. but the JWT doesn't like that we pass nil as keyFunc and throws
		// ValidationErrorUnverifiable errors. We just ignore it here.
		if detailed.Errors != jwt.ValidationErrorUnverifiable {
			return nil, err
		}
	}
	return token, nil
}

// Convert known date keys (exp, nbf, iat) to a human-readable format
func convertDateKeys(claims jwt.MapClaims) {
	convertibleKeys := []string{"exp", "nbf", "iat"}
	for _, k := range convertibleKeys {
		if val, ok := claims[k]; ok {
			claims[k] = time.Unix(int64(val.(float64)), 0).String()
		}
	}
}

// Converts the given claims to JSON
func convertToJson(aMap map[string]interface{}) (string, error) {
	payload, err := json.MarshalIndent(aMap, "", "  ")
	if err != nil {
		return "", err
	}
	return string(payload), nil
}

// Reads the JWT from stdin or from commandline
func readJwt() (string, error) {
	flag.Parse()
	if len(flag.Args()) == 0 {
		// Read from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("JWT: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		return strings.TrimSuffix(text, "\n"), nil
	} else {
		// Use first provided arg
		return flag.Args()[0], nil
	}
}
