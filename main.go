package main

import (
	"bufio"
	"github.com/kataras/iris"
	"io"
	"os"
	"time"
)

func main() {
	app := iris.New()
	// 从 "./views" 文件夹加载所以的模板
	// 其中扩展名为“.html”并解析它们
	// 使用标准的`html / template`包。

	file := "D:/Go/src/code/Project/Chat"

	app.RegisterView(iris.HTML(file, ".html"))
	// 方法:    GET
	// 资源:  http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		//绑定数据
		//ctx.ViewData("message", "Hello world!")
		// 渲染视图文件: ./views/hello.html
		ctx.View("index.html")
	})
	// 方法:    GET
	//资源:  http://localhost:8080/user/42
	app.Get("/{user:string}/{message:string}", func(ctx iris.Context) {
		customer := ctx.Params().GetString("user")
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		word := timeStr + " " + customer + ": " + ctx.Params().GetString("message") + "&#13;"
		if(ctx.Params().GetString("message") == "" || ctx.Params().GetString("user") == ""){
			return
		}
		filename := file + "/text.txt"
		fromfile1, _ := os.Open(filename)
		br1 := bufio.NewReader(fromfile1)
		str1 := ""
		for i := 0; ; i++{
			a, _, c := br1.ReadLine()
			if c == io.EOF {
				break
			}
			str1 = string(a) + str1
		}
		str1 = word + str1
		_ = fromfile1.Close()
		fromfile, _ := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE, os.ModePerm)
		writein := []byte(str1)
		_, _ = fromfile.Write(writein)
		_ = fromfile.Close()
		fromfile0, _ := os.Open(filename)
		br := bufio.NewReader(fromfile0)
		str := ""
		for i := 0; ; i++{
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			str = string(a) + str
		}
		_, _ = ctx.Writef(str)
		_ = fromfile0.Close()
	})

	app.Get("/show", func(ctx iris.Context) {
		filename := file + "/text.txt"
		fromfile0, _ := os.Open(filename)
		br := bufio.NewReader(fromfile0)
		str := ""
		for i := 0; ; i++{
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			str = string(a) + str
		}
		ctx.Writef(str)
	})
	//启动服务
	app.Run(iris.Addr(":8080"))
}