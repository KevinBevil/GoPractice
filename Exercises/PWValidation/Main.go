// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {

	const (
		usage    = "Usage: [username] [password]"
		errUser  = "Access denied for %q.\n"
		errPwd   = "Invalid password for %q.\n"
		accessOK = "Access granted to %q.\n"
	)

	// credMap holds usernames/passwords as key/value pairs
	credMap := map[string]string{"jack": "1888", "inanc": "1879"}

	loginValid, credValid := false, false

	// if 2 entries have not been entered, cease program with usage info
	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}
	// Takes in username/password from user - standard in
	username, password := os.Args[1], os.Args[2]

	// For each key/value pair, we compare username/password with the pairs in the vault
	for k, v := range credMap {
		if username == k && password == v {
			credValid = true
			break
		} else if username == k {
			loginValid = true
			break
		}
	}
	// After looping through cred map, if valid cred was found or if only valid login was found, report, else 'Access denied'
	if credValid {
		fmt.Printf(accessOK, username)
	} else if loginValid {
		fmt.Printf(errPwd, username)
	} else {
		fmt.Printf(errUser, username)
	}
}
