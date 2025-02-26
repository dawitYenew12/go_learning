	package main

	import (
		"fmt"
		"reflect"
		"unsafe"
	)

	type authenticationInfo struct {
		username string
		password string
	}

	// create the method below
	func (authI authenticationInfo) getBasicAuth() string {
		return fmt.Sprintf("Authorization: Basic %s:%s", authI.username, authI.password)
	}

	func main() {
		auth := authenticationInfo{
			username: "admin",
			password: "admin",
		}

		fmt.Println("Size of struct:", unsafe.Sizeof(auth))
		typ := reflect.TypeOf(auth).Size()
		fmt.Println(typ)
		totalStringSize := len(auth.username) + len(auth.password)
		fmt.Println("Total size including string content:", unsafe.Sizeof(auth)+uintptr(totalStringSize))
		fmt.Println(auth.getBasicAuth())
	}