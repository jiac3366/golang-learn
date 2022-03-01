package mocks

import "fmt"

type Fucker struct {
	Content string
}

func (f *Fucker) Get(url string) string {
	fmt.Println("Get", url)
	return f.Content
}

// 当struct Fucker同时实现了Get()和Put(), 他就是RetriverAndPoster类型了

func (f *Fucker) Put(url string, form map[string]string) string {
	f.Content = form["contents"] // 试图去修改Content成员变量，但如果不传pointer receiver不能达到目的
	return "OK"
}

type Lover struct {
	LoverContent string
}

func (f *Lover) Get(url string) string {
	return f.LoverContent
}
