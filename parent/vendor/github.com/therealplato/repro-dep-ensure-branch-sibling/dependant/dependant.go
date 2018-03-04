package dependant

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Exported(logrus.Fields) {
	fmt.Println("dependant")
}
