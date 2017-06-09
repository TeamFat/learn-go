package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//go run upload.go
//chrome open http://127.0.0.1:9090/upload

//处理文件上传我们需要调用req.ParseMultipartForm，里面的参数表示maxMemory， 调用ParseMultipartForm之后，上传的文件存储在maxMemory大小的内存里面，如果文件大小超过了maxMemory， 那么剩下的部分将存储在系统的临时文件中。
//我们可以通过req.FormFile获取上面的文件句柄，然后实例中使用了io.Copy来存储文件
//获取其他非文件字段信息的时候就不需要调用reqe.ParseForm，因为在需要的时候Go自动会去调用。 而且ParseMultipartForm调用一次之后，后面再次调用不会再有效果。
//token 唯一性

// 处理/upload 逻辑
func upload(res http.ResponseWriter, req *http.Request) {
	fmt.Println("method:", req.Method) //获取请求的方法

	if req.Method == "GET" {

		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		data := struct {
			Token string
		}{
			Token: token,
		}
		t.Execute(res, data)

	} else {

		// 方法一：
		// 此文件上传处理是将表单数据缓存到本地之后再处理
		req.ParseMultipartForm(32 << 20)
		file, handler, err := req.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(res, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		/*
		   // 方法二：
		   // 下面方法是一种更优的解决方案，不会去写临时文件
		   reader, err := req.MultipartReader()
		   if err != nil {
		       http.Error(res, err.Error(), http.StatusInternalServerError)
		       return
		   }

		   for {
		       part, err := reader.NextPart()
		       if err == io.EOF {
		           break
		       }

		       //if part.FileName() is empty, skip this iteration.
		       if part.FileName() == "" {
		           continue
		       }
		       dst, err := os.Create("/home/" + part.FileName())
		       defer dst.Close()

		       if err != nil {
		           http.Error(res, err.Error(), http.StatusInternalServerError)
		           return
		       }

		       if _, err := io.Copy(dst, part); err != nil {
		           http.Error(res, err.Error(), http.StatusInternalServerError)
		           return
		       }
		   }
		*/

		fmt.Fprintf(res, "success")
	}
}

func main() {
	http.HandleFunc("/upload", upload)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
