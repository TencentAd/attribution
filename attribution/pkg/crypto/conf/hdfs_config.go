package conf

import "flag"

var fileName = flag.String("file_name", "test", "")
var user = flag.String("user", "hadoop", "")

// 该类是用于hdfs的配置类
type HdfsConf struct {
	NodeAddress []string
	FileName    string
	User        string
}

// 该函数运行用户使用传参的方法创建一个HDFS的配置对象
func NewHdfsConf(fileName string, user string, NodeAddresses []string) *HdfsConf {
	return &HdfsConf{
		NodeAddress: NodeAddresses,
		FileName:    fileName,
		User:        user,
	}
}

// 该方法允许用户使用命令行参数的形式输入配置参数
func (h *HdfsConf) InitHdfsConf() {
	var hdfsAddrs addrs
	flag.Var(&hdfsAddrs, "hdfs_addresses", "")
	flag.Parse()
	h.FileName = *fileName
	h.User = *user
	h.NodeAddress = hdfsAddrs
}
