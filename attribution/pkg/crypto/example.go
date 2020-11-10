package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/colinmarc/hdfs"
	"os"
	"strings"
)

func main() {

	testLine()
}

func splitTest() {
	s := string("111")
	split := strings.Split(s, ",")
	fmt.Println(split)
}

func testHdfs() {
	// 这个就是hdfs的地址
	//client, err := hdfs.New("localhost:9000")
	client, err := hdfs.NewClient(hdfs.ClientOptions{
		Addresses: []string{"localhost:9000"},
		User:      "hadoop",
	})
	if err != nil {
		fmt.Println("fail to connect to hdfs")
		return
	}
	defer client.Close()

	filePath := "/test/json2.txt"
	//writer, err := client.Create(filePath)
	//if err != nil {
	//	fmt.Printf("fail to open reader for writing, err[%v]\n", err)
	//	return
	//}
	// 这里一定要记得关闭！！！
	//defer writer.Close()

	// => Abominable are the tumblers into which he pours his poison.
	//myMap := make(map[string]string)
	//myMap["123"] = "fff"
	//marshal, _ := json.Marshal(myMap)
	//_, _ = writer.Write(marshal)

	// 这个就是要打开的文件
	reader, _ := client.Open(filePath)
	defer reader.Close()
	buf := make([]byte, 128)
	length, _ := reader.ReadAt(buf, 0)
	fmt.Println(string(buf))

	result := make(map[string]interface{})
	// 这里要使用指定长度的切片，否则会转化失败
	err = json.Unmarshal(buf[0:length], &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func testLine() {
	file, err := os.Open("C:\\Users\\v_zefanliu\\Desktop\\json.txt")
	if err != nil {
		fmt.Printf("fail to open file, err[%v]", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "\t")
		fmt.Println(split)
	}
	arr := []string{"aaa\n", "bbb\n"}
	fmt.Println(arr)
}
