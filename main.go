package main

import (
	"github.com/RadhiFadlillah/boom/internal/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.BoomCmd().Execute(); err != nil {
		logrus.Fatalln(err)
	}
}
