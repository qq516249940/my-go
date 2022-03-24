package main

import (
	//"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	diff_time = 3600 * 24 * 7
)

func main() {
	//flag.Parse()
	//root := flag.Arg(0)
	root := "/home/deployer/zhijing/data/logs"
	getFilelist(root)
}

func getFilelist(path string) {
	now_time := time.Now().Unix() //当前时间，使用Unix时间戳
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			// fmt.printf(f)
			fmt.Printf("有文件错误,继续执行 %v \r\n",err)
			return nil
		}

		if f.IsDir() {
			//忽略目录
			fmt.Printf("this is a dir: %v !\r\n", path)
			return nil
		   }

		if err != nil {
			fmt.Errorf(" no such file or directory: %v\n", err)
			return nil
		}   

		file_time := f.ModTime().Unix()
		
			// fmt.Println("file_time",file_time)
			// fmt.Println("now_time",now_time)
			// fmt.Println("---", diff_time)
			// fmt.Println(now_time - file_time)
		
		if (now_time - file_time) > diff_time { //判断文件是否超过7天
			fmt.Printf("Delete file %v !\r\n", path)
			fmt.Printf(f.Name())

			if err := os.Remove(path); err != nil {
				fmt.Println(err)
			}
			// os.RemoveAll(path)
		} //else {
		//println(path)
		//}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\r\n", err)
	}
}