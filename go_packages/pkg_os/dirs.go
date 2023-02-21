package pkg_os

import (
	"fmt"
	"os"
)

// GetWorkdir 获得工作目录
func GetWorkdir() {
	curDir, _ := os.Getwd()
	fmt.Println(curDir)
	os.Chdir("/Users/maggie/Golang/go_packages/pkg_os")
	curDir, _ = os.Getwd()
	fmt.Println(curDir)

	tmpDir := os.TempDir()
	fmt.Println(tmpDir)
}

// CreateFile 创建文件
// 等价于 OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
func CreateFile() {
	os.Chdir("go_packages/pkg_os")
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Create file failed - ", err)
	} else {
		fmt.Println(f.Name())
	}
}

// CreateDir 创建文件夹 - Mkdir / MkdirAll
func CreateDir() {
	os.Chdir("go_packages/pkg_os")
	// 创建单个目录
	err := os.Mkdir("folder1", os.ModePerm)
	if err != nil {
		fmt.Println("Create folder failed - ", err)
	} else {
		fmt.Println("Create folder successfully")
	}

	// 递归创建级联目录
	err = os.MkdirAll("folder1/folder2/folder3", os.ModePerm)
	if err != nil {
		fmt.Println("Create folder failed - ", err)
	} else {
		fmt.Println("Create folder successfully")
	}
}

// RemoveDir 删除目录或文件
func RemoveDir() {
	os.Chdir("go_packages/pkg_os")
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Println("Remove failed - ", err)
	} else {
		fmt.Println("Remove successfully")
	}

	// 递归删除级联目录
	err = os.RemoveAll("folder1")
	if err != nil {
		fmt.Println("Remove failed - ", err)
	} else {
		fmt.Println("Remove successfully")
	}
}

// Rename_ 重命名
func Rename_() {
	os.Chdir("go_packages/pkg_os")
	err := os.Rename("test.txt", "new.txt")
	if err != nil {
		fmt.Println("Rename failed")
	} else {
		fmt.Println("Rename successfully")
	}
}

func WriteFile() {
	os.Chdir("go_packages")
	err := os.WriteFile("new.txt", []byte("哈哈哈哈"), os.ModePerm)
	if err != nil {
		fmt.Println("Write file failed", err)
	} else {
		file, _ := os.ReadFile("new.txt")
		fmt.Println(string(file))
	}
}
