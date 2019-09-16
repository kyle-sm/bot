package main

import "fmt"

func HandleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}