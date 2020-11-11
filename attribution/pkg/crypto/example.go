package main

import (
	"fmt"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/conf"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/uitl"
	"math/big"
)

func main() {
	testHdfs()
}

//func testHdfs() {
//	// 这个就是hdfs的地址
//	//client, err := hdfs.New("localhost:9000")
//	client, err := hdfs.NewClient(hdfs.ClientOptions{
//		Addresses: []string{"localhost:9000"},
//		User:      "hadoop",
//	})
//	if err != nil {
//		fmt.Println("fail to connect to hdfs")
//		return
//	}
//	defer client.Close()
//
//	filePath := "/test/json2.txt"
//	//writer, err := client.Create(filePath)
//	//if err != nil {
//	//	fmt.Printf("fail to open reader for writing, err[%v]\n", err)
//	//	return
//	//}
//	// 这里一定要记得关闭！！！
//	//defer writer.Close()
//
//	// => Abominable are the tumblers into which he pours his poison.
//	//myMap := make(map[string]string)
//	//myMap["123"] = "fff"
//	//marshal, _ := json.Marshal(myMap)
//	//_, _ = writer.Write(marshal)
//
//	// 这个就是要打开的文件
//	reader, _ := client.Open(filePath)
//	defer reader.Close()
//	buf := make([]byte, 128)
//	length, _ := reader.ReadAt(buf, 0)
//	fmt.Println(string(buf))
//
//	result := make(map[string]interface{})
//	// 这里要使用指定长度的切片，否则会转化失败
//	err = json.Unmarshal(buf[0:length], &result)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(result)
//}

func testHdfs() {
	hdfsConf := conf.NewHdfsConf("/crypto", "v_zefanliu", []string{"localhost:9000"})
	manager := uitl.NewHdfsKeyManager(hdfsConf)
	err := manager.StorageEncryptKey("123456", big.NewInt(123124))
	if err != nil {
		fmt.Println("fail to storage key", err)
		return
	}
	key, err := manager.GetEncryptKey("123456")
	if err != nil {
		fmt.Println("fail to get key", err)
		return
	}
	fmt.Println("key:", key)
}
