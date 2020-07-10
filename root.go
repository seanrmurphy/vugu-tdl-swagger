package main

import "log"

func (c *Root) GotoHome() {
	log.Printf("Home Button Pressed")
	c.Navigate("/", nil)
}

func (c *Root) GotoLogin() {
	log.Printf("Login Button Pressed")
	c.Navigate("/login", nil)
}

//func (c *Root) BeforeBuild() {
//log.Printf("Before builder...")
//if c.LoggedIn == false {
//c.Navigate("/login", nil)
//}
//}
