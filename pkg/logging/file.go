package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s",LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s",LogSaveName,time.Now().Format(TimeFormat),LogFileExt)
	return fmt.Sprintf("%s%s",prefixPath,suffixPath)
}

func openLogFile(filePath string) *os.File {
	_,err := os.Stat(filePath)
	switch  {
	case os.IsNotExist(err):  //文件或者目录不存在 返回一个bool
		mkDir()
	case os.IsPermission(err):  //反回一个bool 得知权限是否满足
		log.Fatalf("Permission :%v",err)
	}
	// os.OpenFile调用文件,支持传入文件名称,指定的模式调用文件,文件权限
	handle,err:=os.OpenFile(filePath,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err!=nil{
		log.Fatalf("Fail to OpenFile :%v",err)
	}
	return handle
}

func mkDir()  {
	dir,_:=os.Getwd()  //返回与当前目录对应的根路径名
	//os.MkdirAll创建对应的目录以及所需的子目录,若成功则返回nil,否则返回error
	//os.ModePerm const定义ModePerm FileMode = 0777
	err:=os.MkdirAll(dir+"/"+getLogFilePath(),os.ModePerm)
	if err!=nil {
		panic(err)
	}
}




