package g

import (
	"os"
	logs "github.com/Sirupsen/logrus"
)
var Root string

func InitRootDir() {
	var err error
	Root, err = os.Getwd()
	if err != nil {
		logs.Fatalln("getwd fail:", err)
	}
}