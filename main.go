package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

const (
	O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
	O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)

func main() {
	songs := "#EXTM3U"
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path as a command-line argument.")
		return
	}
	for i := 1; i < len(os.Args); i = i + 1 {
		filePath := os.Args[i]
		fmt.Printf("File path: %s\n", filePath)
		songs = songs + "\n#EXTINF:,\n" + filePath
	}
	var songlistname string
	fmt.Printf("欢迎使用“Woshijiaobenxiaozi”牌M3U8歌单生成器！\n请输入歌单名称：")
	fmt.Scanf("%s", &songlistname)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前工作目录失败：%v\n", err)
		return
	}
	CreateFile(currentDir + "\\" + songlistname + ".m3u8")
	WriteStringFile(currentDir+"\\"+songlistname+".m3u8", songs)
	time.Sleep(time.Second * 10)
}

// CreateFile 创建文件
func CreateFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("创建成功！file:%s\n", file.Name())
	}
}

// WriteStringFile 写入字符串
func WriteStringFile(name string, songs string) {
	file, err := os.OpenFile(name, O_RDWR, 0775) // 以读写模式打开文件，并且打开时清空文件
	if err != nil {
		fmt.Printf("err:%v\n", nil)
	}
	file.WriteString(songs)
}
