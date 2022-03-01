package main

import (
	"fmt"
	"programming/interface_program/mock"
)

const url = "https://www.imooc.com"

type Retriver interface {
	Get(url string) string
}

type Poster interface {
	Put(url string, form map[string]string) string
}

//问题起源：如何让别人传进来的r强制实现Get()呢？——让别人实现我的接口
func download(r Retriver) {
	fmt.Println(r.Get("https://baidu.com"))
}

func post(p Poster) {
	p.Put(url, map[string]string{
		"method": "POST",
	})
}

// RetriverAndPoster 接口的组合
type RetriverAndPoster interface {
	Retriver
	Poster
}

func session(s RetriverAndPoster) string {

	s.Put(url, map[string]string{
		"contents": "another new contents",
	})
	return s.Get(url)

}

func inspect_type(r Retriver) {
	// 传进来的不是引用 是类型(%T) 和 值或指针(%v)
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) { // switch r.(type)
	case *mocks.Fucker:
		fmt.Println("Contents", v.Content)

	case *mocks.Lover:
		fmt.Println("Contents", v.LoverContent)
	}
}

func main() {
	var r Retriver
	// 为什么Fucker的Get()如果改成从(f Fucker)改成(f *Fucker) 初始化赋值要从mocks变成&mocks呢？
	// 因为 Get()要接收一个指针的类型，又因为调用Get()的就是r 意味着传进去方法里面的就是r，所以r是一个指针变量，故加上“&”

	r = &mocks.Fucker{Content: "Fucker"}
	test_r := mocks.Fucker{Content: "Fucker+Put()"}
	inspect_type(r)
	r = &mocks.Lover{LoverContent: "Lover"}

	inspect_type(r)

	//判断类型：1 switch 2 type assertion
	if value, ok := r.(*mocks.Fucker); ok {
		fmt.Println(value.Content) // 如果不是mocks.Fucker类型又不加判断：main.Retriver is *mocks.Lover, not mocks.Fucker
	} else {
		fmt.Println("not mocks.Fucker")
	}

	// 测试接口组合
	fmt.Println("test RetriverAndPoster")
	fmt.Println(session(&test_r)) // 这里不能传入r, 因为r被声明为Retriver类型 只实现了Get()

	// 测试多态
	fmt.Println("test download()1")
	download(r)

	fmt.Println("test download()2")
	download(&test_r)
}
