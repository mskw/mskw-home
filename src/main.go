package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"os"
	"os/signal"
	"syscall"
)

const pidFile = "/www/pid/mskw-home"

func main() {
	if IsExist(pidFile) {
		panic(fmt.Errorf("another server process exists"))
	}
	writePidFile()
	setUpSignalHandler()

	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/welcome", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	app.Get("/ping", func(ctx context.Context) {
		ctx.WriteString("pong")
	})

	app.Run(iris.Addr(":8848"))

}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func setUpSignalHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出", s)
				os.Remove(pidFile)
				os.Exit(0)
			}
		}
	}()
}

func writePidFile() {
	file, err := os.OpenFile(pidFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Errorf("create pid file error: %v", err))
	}
	_, err = file.WriteString(fmt.Sprintf("%v", os.Getpid()))
	if err != nil {
		panic(fmt.Errorf("create pid file error: %v", err))
	}
	err = file.Close()
	if err != nil {
		panic(fmt.Errorf("create pid file error: %v", err))
	}
}
