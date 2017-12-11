package main

import (
	"os"
	"time"
	"bufio"
	"io"
	"os/exec"
	"fmt"
	"net/http"
	"log"
	"math/rand"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/image", getImage)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getImage(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//u := r.Form["url"][0]
	from, err := url.QueryUnescape(r.URL.RawQuery)
	u := strings.Split(from, "=")[1]

	command := "node"
	curl := "--url=" + u
	fmt.Println(curl)
	//output := r.Form["output"][0]
	imageDir := "img/"
	fileName := GetRandomString(22) + ".jpg"
	output := imageDir + fileName
	coutput := "--output=" + output
	js := "page2image.js"
	params := []string{js, curl, coutput}
	b, err2 := execCommand(command, params)
	if (b) {
	}
	if err2 != nil {
		fmt.Println(err2)
		http.Error(w, err2.Error(), 500)
		return
	}

	modtime := time.Now()

	//w.Header().Set("Content-Type", "application/jpg")
	//w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", output))
	w.Header().Add("Content-Disposition", "Attachment")

	content, err := os.Open(output)
	if err != nil {
		fmt.Println(err2)
		//fmt.Println(output, err)
		http.Error(w, err.Error(), 500)
		return
	}
	http.ServeContent(w, r, fileName, modtime, content)
}

func execCommand(commandName string, params []string) (bool, error) {
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	cmd.Run()
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			//break
			fmt.Println(err2)
			return false, err2
		}
		fmt.Println(line)
	}
	cmd.Wait()
	return true, nil
}

//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
