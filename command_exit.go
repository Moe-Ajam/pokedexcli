package main

import "os"

func callbackExit(cfg *config, param string) error {
	os.Exit(0)
	return nil
}
