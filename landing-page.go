package main

import "log"

func (c *LandingPage) BeforeBuild() {
	log.Printf("Before builder...")
	if c.LoggedIn == false {
		c.Navigate("/login", nil)
	}
}
