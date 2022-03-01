package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer,
		//	err.Error(),
		//	http.StatusInternalServerError)
		//return
		// 改用统一错误处理,直接往外面扔error,
		// 所以在HandleFileList --> 自定义1个类型处理error --> func (writer http.ResponseWriter, request *http.Request)
		return err
	}

	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}

// 自定义类型 这个类型能定义的也包括了函数的返回值
// 为了接受一个 [比HandleFunc()第2个参数返回值多了个error] 的func
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 将一个变形的(HandleFunc()第2个参数返回值本来不可以是error，而我们却返回了error)func 变回符合规定的func
// 这里体现了函数式编程，还对错误逻辑做了抽离
func errorWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		code := http.StatusOK
		if err != nil {
			switch {
			case os.IsNotExist(err):
				log.Println(err.Error())
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}

}

func main() {
	//http.HandleFunc("/list/", HandleFileList)
	http.HandleFunc("/list/", errorWrapper(HandleFileList)) // 用errorWrapper做统一错误处理
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
