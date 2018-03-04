package dependant

import (
	"github.com/sirupsen/logrus"
)

func Exported(fs logrus.Fields) {
	logrus.WithFields(fs).Info("updated_dependant")
}
