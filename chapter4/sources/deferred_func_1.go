package main

import (
	"fmt"
	"os"
	"sync"
)

func writeToFile(fname string, data []byte, mu *sync.Mutex) error {
	mu.Lock()
	f, err := os.OpenFile(fname, os.O_RDWR, 0666)
	if err != nil {
		mu.Unlock()
		return err
	}

	_, err = f.Seek(0, 2)
	if err != nil {
		f.Close()
		mu.Unlock()
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		f.Close()
		mu.Unlock()
		return err
	}

	err = f.Sync()
	if err != nil {
		f.Close()
		mu.Unlock()
		return err
	}

	err = f.Close()
	if err != nil {
		mu.Unlock()
		return err
	}

	mu.Unlock()
	return nil
}

func showFileOperation() {
	var mu sync.Mutex
	err := writeToFile("./foo.txt", []byte("hello, defer!\n"), &mu)
	if err != nil {
		fmt.Println("writeToFile error:", err)
		return
	}
	fmt.Println("writeToFile ok")
}

var mu sync.Mutex

func writeToFile2(fileName string, c []byte) error {
	mu.Lock()
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open error: ", err)
		mu.Unlock()
		return err
	}
	_, err = f.Seek(0, 2)
	if err != nil {
		fmt.Println("Seek error: ", err)
		f.Close()
		mu.Unlock()
		return err
	}
	_, err = f.Write(c)
	if err != nil {
		fmt.Println("write error: ", err)
		f.Close()
		mu.Unlock()
		return err
	}
	err = f.Sync()
	if err != nil {
		fmt.Println("Sync error: ", err)
		f.Close()
		mu.Unlock()
		return err
	}
	err = f.Close()
	if err != nil {
		fmt.Println("write error: ", err)
		mu.Unlock()
		return err
	}

	mu.Unlock()
	return nil
}

func writeToFile3(fileName string, c []byte) error {
	mu.Lock()
	defer mu.Unlock()
	f, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("open error: ", err)
		return err
	}
	defer f.Close()
	_, err = f.Seek(0, 2)
	if err != nil {
		fmt.Println("Seek error: ", err)
		return err
	}
	_, err = f.Write(c)
	if err != nil {
		fmt.Println("write error: ", err)
		return err
	}
	err = f.Sync()
	if err != nil {
		fmt.Println("Sync error: ", err)
		return err
	}
	// err = otherOperation()
	// if err != nil {
	// 	fmt.Println("otherOperation error: ", err)
	// 	return err
	// }
	err = f.Close()
	if err != nil {
		fmt.Println("write error: ", err)
		return err
	}
	return nil
}

func showFileOperation2() {
	err := writeToFile3("/a.txt", []byte("Hello, world"))
	if err != nil {
		fmt.Println("writeToFile2 error: ", err)
		return
	}
	fmt.Println("write ok")
}

func main() {
	showFileOperation2()
}

func copyMap(src map[int]string) map[int]string {
	return make(map[int]string)
}
