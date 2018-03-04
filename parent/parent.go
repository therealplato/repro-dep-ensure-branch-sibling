package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/therealplato/repro-dep-ensure-branch-sibling/dependant"
)

func main() {
	dependant.Exported(logrus.Fields{})
	fmt.Println("vim-go")
}
