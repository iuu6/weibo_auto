package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	fmt.Println("             _       ")
	fmt.Println("            | |      ")
	fmt.Println("  ____ _____| |_   _ ")
	fmt.Println(" / ___|___  ) | | | |")
	fmt.Println("| |    / __/| | |_| |")
	fmt.Println("|_|   (_____)\\_)__  |")
	fmt.Println("              (____/ ")
	fmt.Println("-------------------------------------------")
	fmt.Println("本软件为热衷绿野荣誉出品，为蓝天彼端工作室编写！")
	fmt.Println("本软件完全免费，软件版本2.1！现开源发行！")
	fmt.Println("使用本软件产生的一切后果我们不负责任！")
	fmt.Println("本项目仅供开发者学习，产生的一切后果与我们无关")
	fmt.Println("无论是用何种方式倒卖本软件都是可耻的行为！！！")
	fmt.Println("无论是用何种方式倒卖本软件都是可耻的行为！！！")
	fmt.Println("无论是用何种方式倒卖本软件都是可耻的行为！！！")
	client := &http.Client{}

	// 1. 指定要读取的文本文件路径
	filePath := "text.txt"
	// 2. 使用ioutil.ReadFile函数读取整个文件内容到一个字节切片([]byte)中
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("无法读取文件:", err)
		return
	}
	fileContent := string(content)
	fmt.Println("读取的文本内容：")
	fmt.Println(fileContent)
	encodedString := url.QueryEscape(fileContent)
	fmt.Println("转换的文本内容：")
	fmt.Println(encodedString)
	// 生成随机字符串作为data中的占位符
	randomString := generateRandomString(10)

	// 第一次请求
	makeRequest(client, randomString, encodedString)
	for {
		// 生成随机字符串作为data中的占位符
		randomString := generateRandomString(10)

		// 随机生成下一次请求的时间间隔（5-10分钟）
		randomWaitTime := generateRandomWaitTime()
		fmt.Printf("下一次请求将在 %d 秒后执行。\n", randomWaitTime)

		// 等待一段时间后再次请求
		time.Sleep(time.Duration(randomWaitTime) * time.Second)
		makeRequest(client, randomString, encodedString)
	}

}

// makeRequest 发送HTTP请求
func makeRequest(client *http.Client, randomString string, encodedString string) {
	originalData := "callback_url=xxxxxxxx" //在这里输入callback_url
	newdata := strings.Replace(originalData, "N9HfC3G0qQwmlcnBQ7M", randomString, -1)
	newdata2 := strings.Replace(newdata, "zOejxvmMtLu1vthF1FU", encodedString, -1)
	var data = strings.NewReader(newdata2)
	req, err := http.NewRequest("POST", "https://api.weibo.cn/2/statuses/send?", data) //在这里补全send
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "api.weibo.cn")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "xxxxx")
	req.Header.Set("X-Sessionid", "xxxxx")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("cronet_rid", "xxxxx")
	req.Header.Set("SNRT", "xxxxx")
	req.Header.Set("x-engine-type", "xxxxx")
	req.Header.Set("X-Log-Uid", "xxxxx")
	req.Header.Set("X-Validator", "xxxxx")
	req.Header.Set("User-Agent", "xxxxx")
	req.Header.Set("Authorization", "xxxxx")
	req.Header.Set("Accept", "*/*")
	// req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "en-US,en")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

// generateRandomString 生成指定长度的随机字符串
func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// generateRandomWaitTime 生成5-10分钟之间的随机等待时间（单位：秒）
func generateRandomWaitTime() int {
	rand.Seed(time.Now().UnixNano())
	min := 5 * 60  // 5分钟的秒数
	max := 10 * 60 // 10分钟的秒数
	return rand.Intn(max-min) + min
}
