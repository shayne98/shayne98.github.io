---
layout: post
title: "Go notes"
date:   2024-7-9
tags: [Go]
comments: true
author: Shayne
---
#### 关于slice的初始化

```go
var a []int// 这样情况下会赋零值，对于slice而言，是nil
a := make(string[],0)
a := string[]{}//以上两种初始化方式都是非空，同理map以及chan
//Badcase
ins, err := rpc.GetInsuranceInfoList(ctx, append(make([]string, 0), strconv.FormatInt(parentOrderId, 10)), orderInfo.UserId)

//Correct

ins, err := rpc.GetInsuranceInfoList(ctx, []string{strconv.FormatInt(parentOrderId, 10)},
```

注意，对于nil的slice而言，除了不能索引其他例如遍历等操作均不会报panic

```go
/ 一个为nil的slice，除了不能索引外，其他的操作都是可以的

// Note: 如果这个slice是个指针，不适用这里的规则

var a []int  

fmt.Printf("len(a):%d, cap(a):%d, a==nil:%v\n", len(a),cap(a), a == nil) //0 0 true

for _, v := range a{// 不会panic

        fmt.Println(v) 

}

aa := a[0:0]     // 也不会panic，只要索引都是0



// nil的map，我们可以简单把它看成是一个只读的map

var b map[string]string

if val, ok := b["notexist"];ok{// 不会panic

        fmt.Println(val)

}

for k, v := range b{// 不会panic

        fmt.Println(k,v)

}

delete(b, "foo") // 也不会panic

fmt.Printf("len(b):%d, b==nil:%v\n", len(b), b == nil) // 0 true
```

#### 基本类型的传参

go中所有类型的传参均为值传递，但slice和map、chan这类指针类型的变量又有不同

在传入slice的函数本地更改slice 中的值并不会有所变化，而map以及chan会变化

```go
a := []string{}
apendMe(a)
fmt.Println(a)//[]
}
func apendMe(strs []string) {
	strs = append(strs, "a")//

```

结构体中若含有指针类型的变量，传输结构体作为参数时，结构体中的变量只拷贝非指针变量，指针变量指向的内容是公用的

#### for循环中的起go rountine

若gorutine 使用了for的局部变量，等价于在闭包中的局部变量，是个引用。

闭包即，一个函数返回另一个函数，返回的那个函数使用了父级函数中的局部变量，且该返回函数的调用在外部

```go
for _, i := range []int{1, 2, 3} {
		go func(x int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(3 * time.Second)
// 输出 333
```

同时注意，for循环中的临时变量并不是每次迭代生成新的变量，遍历的对象以及每次遍历得到的元素都是一个副本

```go
For key, element = range aContainer {...}
```

关于上面for循环有几个点：

1. 实际遍历的aContainer是原始值的一个副本
2. element是遍历到的元素的原始值的一个副本
3. key和element整个循环都是同一个变量，而不是每次迭代都生成新变量

所以针对比较大的element，遍历直接用下标来访问会比较好

#### init函数的使用

功能： 主要为了初始化客户端、缓存以及加载一些固定的配置项，使用时需要注意：

1. 初始化客户端、结构体这种操作时，为方便后续单元测试的Mock，应该单独抽成一个方法
2. init函数不能依赖于环境变量以及先后初始化顺序，避免io call

![20240710112431](https://raw.githubusercontent.com/shayne98/Figure-Bed/main/blog/20240710112431.png)

#### panic的使用方式

使用defer配合recover来完成对于panic的处理

```go
func main() {
	fmt.Println("Enter function")
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recovered from panic")
		}
		fmt.Println("Exit panic defer function")
	}()
	panic("Panic statement")
	fmt.Println("Enter function")
}
```
