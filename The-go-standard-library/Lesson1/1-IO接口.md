### Go中的IO包

io 包为I/O原语提供基本的接口，包装了这些原语的已有实现。

在 io 包中最重要的是两个接口：Reader 和 Writer 接口。你只要满足这两个接口，就可以使用 IO 包的功能。

### **Reader 使用**

利用 `Reader` 可以容易进行流式数据传输。`Reader` 方法内部是被循环调用，每次迭代，它会从数据源读取一块数据放入缓冲区 `p` （即 Read 的参数 p）中，直到返回 `io.EOF` 错误时停止。

通过 `string.NewReader(string)` 创建一个字符串读取器。

```go
func main() {
	reader := strings.NewReader("This is Golang language.")
	p := make([]byte, 8)
	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(p[:n]))
	}
}
```

结果：

```
This is 
Golang l
anguage.
EOF: 0
```

### **Writer**

表示一个编写器，它从缓冲区读取数据，并将数据写入目标资源。

```go
func main() {
	proverbs := []string{"GET", "POST", "PUT", "DELETE", "PATCH", ""}
	var writer bytes.Buffer
	for _, p := range proverbs {
		n, err := writer.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
		fmt.Println(writer.String())
	}
}
```

### os.File

表示本地系统上的文件。它实现了 `io.Reader` 和 `io.Writer` ，因此可以在任何 io 上下文中使用。也可以读取文件并打印其内容。

```go
func main() {
    file, err := os.Create("./hello.txt")
    defer file.Close()
  	p := make([]byte, 6)
  	// file 类型实现了 io.Writer
  	n, err := file.Write([]byte(p))
}
```

可以用作读取器来从本地文件系统读取文件的内容

```go
func main() {
    file, err := os.Open("./proverbs.txt")
  	defer file.Close()
  	p := make([]byte, 6)
  	n, err := file.Read(p)
}
```

### ioutil

可以使用函数 `ReadFile` 将文件内容加载到 `[]byte` 中

```go
func main() {
  bytes, err := ioutil.ReadFile("./hello.txt") //1.16开始已经弃用，直接os.ReadFile
	file, _ := os.ReadFile("./hello.text")
	fmt.Print(string(file))
}
```

还有其他类型如PipeReader 和 PipeWriter 类型等