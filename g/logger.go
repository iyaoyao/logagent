package g

import (
	logs "github.com/Sirupsen/logrus"
	"os"
	"time"
)



func checkLogFileIsExist(filename string) bool{
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err){
		exist = false
	}
	return exist
}

func InitLogLevel(level string) {
	switch level {
	case "info":
	logs.SetLevel(logs.InfoLevel)
	case "debug":
	logs.SetLevel(logs.DebugLevel)
	case "warn":
	logs.SetLevel(logs.WarnLevel)
	default:
	logs.Fatal("log conf only allow [info, debug, warn], please check your confguire")
	}
	return
}


func initLogFile(filename string) (f *os.File,err error){
	var file *os.File
	if checkLogFileIsExist(filename){
		file, err = os.OpenFile(filename,os.O_APPEND,0666)
	}else{
		file, err = os.Create(filename)
	}
	return file,err

}

func InitLogger()  (err error) {
	logs.SetFormatter(&logs.JSONFormatter{})
	logFileName  :=  time.Now().Format("2006-01-02") + "-logagent.log"
	f, err := initLogFile(logFileName)
	if err!=nil{
		logs.Error("init log file error:%s",err)
	}
	logs.SetOutput(f)
	return err
}
