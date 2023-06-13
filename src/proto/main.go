package main

import (
	cmds "Go-test/src/proto/cmd"
	"Go-test/src/proto/protobuf/Data"
	"fmt"
	"io/ioutil"
	// "proto/protobuf/Data"
	"strings"

	"google.golang.org/protobuf"
)

const (
	GO_PROTO_OUT    = `D:\proto\go_proto`
	PROTO_PATH      = `D:\proto\bog2_meta\proto` // 要遍历的文件夹
	PROTO_INIT_PATH = `D:\proto\bog2_meta`       // 要遍历的文件夹
)

func main() {
	animator := Data.AnimatorTransitionDatas{}

	// result := cmds.Command("dir")
	// fmt.Println(result)
	// 要遍历的文件夹
	// dir := `D:\proto\bog2_meta\proto`

	// 遍历的文件夹
	// 参数：要遍历的文件夹，层级（默认：0）
	findDir(PROTO_PATH, 0)

}

// 遍历的文件夹
func findDir(dir string, num int) {
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// 遍历这个文件夹
	for _, fi := range fileinfo {

		// 重复输出制表符，模拟层级结构
		print(strings.Repeat("\t", num))

		// 判断是不是目录
		if fi.IsDir() {
			// println(`目录：`, fi.Name())
			findDir(dir+`/`+fi.Name(), num+1)
		} else {
			// println(`文件：`, fi.Name())
			proto_file := dir + `/` + fi.Name()
			// fmt.Println("文件：", proto_file)
			proto_cmd := "protoc --proto_path=\"" + PROTO_INIT_PATH + "\" --go_out=\"" + GO_PROTO_OUT + "\" " + proto_file
			fmt.Println("文件：", proto_cmd)
			result := cmds.Command(proto_cmd)
			fmt.Println(result)
		}
	}
}
