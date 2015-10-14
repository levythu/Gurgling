# Gurgling
An extremely-light framework for Golang to build restful API and Website.

## Quick Start
### Install
	go get github.com/levythu/gurgling
### Setup a HTTP server
	package main
	
	import (
	    . "github.com/levythu/gurgling"
	    "net/http"
	)
	
	func main() {
		// Create a gate router.
		const mountPoint="/"
	    var router=GetRouter(mountPoint)
	
		// Mount one handler
	    router.Get("/", func(req Request, res Response) {
	        res.Send("Hello, World!")
	    })
	
		// Mount the gate to net/http and run the server
	    http.Handle(mountPoint, router)
	    fmt.Println("Running...")
	    http.ListenAndServe(":8080", nil)
	}

