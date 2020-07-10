package main

import (
	"errors"
	"log"
	"net/url"
	"syscall/js"
)

func readBrowserURL() (*url.URL, error) {

	g := js.Global()
	if !g.Truthy() {
		return nil, errors.New("not in browser (js) environment")
	}

	locstr := js.Global().Get("window").Get("location").Call("toString").String()

	log.Printf("locstr = %v", locstr)
	u, err := url.Parse(locstr)
	if err != nil {

		log.Printf("Error parsing url")
		return u, err
	}

	return u, nil
}

func sessionStorageSet(k, v string) {
	js.Global().Get("localStorage").Call("setItem", k, v)
}

func sessionStorageGet(k string) js.Value {
	return js.Global().Get("localStorage").Call("getItem", k)
}

func sessionStorageDelete(k string) {
	js.Global().Get("localStorage").Call("removeItem", k)
}

func sessionStorageClear() {
	js.Global().Get("localStorage").Call("clear")
}
