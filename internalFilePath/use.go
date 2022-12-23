package internalFilePath

import (
	"fmt"
	"os"
	"path/filepath"
)

//学习filepath包，兼容各操作系统的文件路径

func Use() {
	//返回所给路径的绝对路径
	path, _ := filepath.Abs("./1.txt")
	fmt.Println(path) // /Users/xx/goModProject/output/goTools/1.txt

	//返回路径最后一个元素
	fmt.Println(filepath.Base("./1.txt")) // 1.txt
	//如果路径为空字符串，返回.
	fmt.Println(filepath.Base(""))
	//如果路径只有斜线，返回/
	fmt.Println(filepath.Base("///"))

	//返回等价的最短路径
	//1.用一个斜线替换多个斜线
	//2.清除当前路径.
	//3.清除内部的..和他前面的元素
	//4.以/..开头的，变成/
	fmt.Println(filepath.Clean("C:/a/b/../c")) // C:/a/c
	fmt.Println(filepath.Clean("./1.txt"))     // 1.txt

	//返回路径最后一个元素的目录
	//路径为空则返回.
	fmt.Println(filepath.Dir("./a/b/c"))  // a/b
	fmt.Println(filepath.Dir("C:/a/b/c")) // C:/a/b

	//返回链接文件的实际路径
	path2, _ := filepath.EvalSymlinks("1.lnk")
	fmt.Println(path2)

	//返回路径中的扩展名
	//如果没有点，返回空
	fmt.Println(filepath.Ext("./a/b/c/d.jpg")) // .jpg

	//将路径中的/替换为路径分隔符
	fmt.Println(filepath.FromSlash("./a/b/c")) // ./a/b/c

	//返回所有匹配的文件
	match, _ := filepath.Glob("./*.go")
	fmt.Println(match) // [main.go]

	//判断路径是不是绝对路径
	fmt.Println(filepath.IsAbs("./a/b/c")) // false
	fmt.Println(filepath.IsAbs("C:/a/b/c"))

	//连接路径，返回已经clean过的路径
	fmt.Println(filepath.Join("C:/a", "/b", "/c")) // C:/a/b/c

	//匹配文件名，完全匹配则返回true
	fmt.Println(filepath.Match("*", "a"))
	fmt.Println(filepath.Match("*", "C:/a/b/c"))
	fmt.Println(filepath.Match("\\b", "b"))

	//返回以basepath为基准的相对路径
	path3, _ := filepath.Rel("C:/a/b", "C:/a/b/c/d/../e")
	fmt.Println(path3)

	//将路径使用路径列表分隔符分开，见os.PathListSeparator
	//linux下默认为:，windows下为;
	fmt.Println(filepath.SplitList("C:/windows;C:/windows/system"))

	//分割路径中的目录与文件
	dir, file := filepath.Split("C:/a/b/c/d.jpg")
	fmt.Println(dir, file) // C:/a/b/c/   d.jpg

	//将路径分隔符使用/替换
	fmt.Println(filepath.ToSlash("C:/a/b")) // C:/a/b

	//返回分区名
	fmt.Println(filepath.VolumeName("C:/a/b/c"))

	//遍历指定目录下所有文件
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
