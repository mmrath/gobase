//go:generate go run github.com/UnnoTed/fileb0x web-static.yaml

package main

import "github.com/mmrath/gobase/apps/uaa/cmd"

func main() {
	cmd.Main()
}
