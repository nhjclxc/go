package main

import (
	"fmt"
	"go/types"
	"strconv"
	"sync"
)

/*
*
熟练掌握go里面的所有基本数据类型的相关特性
*/
func main() {

	//// go中各种数据类型的转化必须显示的进行，注意：go不能和Java一样进行自动转化
	//
	//var i int = -66
	//var f32 float32 = float32(i)
	//var f64 float64 = float64(f32)
	//var u uint8 = uint8(i)
	//fmt.Println(i, f32, f64, u)

	//test54_01()

	//test54_02()

	//test54_03()

	//test54_04()

	//test54_05()

	test54_06()

}

// 在 Go 语言中，make 和 new 都用于内存分配，但它们的用途和适用场景完全不同。
func test54_06() {
	/*
		1. make 和 new 的核心区别
		关键点			make						new
		作用			初始化并返回非零值			仅分配内存，返回零值
		适用于		slice、map、channel		指针（*T）
		返回类型		返回具体类型				返回指向类型的指针（*T）
		是否初始化	是（可以直接使用）			否（只是分配空间，需要手动初始化）
	*/

	// 2. make 的用法
	//make 主要用于创建并初始化以下三种引用类型：
	//切片（slice）
	//映射（map）
	//通道（channel）
	//这些类型本质上是带有底层数据结构的引用类型，make 负责分配内存并初始化相关的内部数据结构（如 slice 的底层数组、map 的哈希表等）。

	// 3. new 的用法
	//new 只用于分配内存，它返回的是一个指向类型的指针，但不会进行初始化。
	//📌 new 创建指针
	// new(int) 分配了一块 int 类型的内存，并返回指针 *int
	// new 只会分配零值内存，不会初始化复杂的结构
	p := new(int)   // 分配一个 int 类型的内存，返回 *int
	fmt.Println(p)  // 0xc000014090 (指针地址)
	fmt.Println(*p) // 0 （默认零值）

	// 📌 new vs make
	var p1 = new([]int)     // ❌ 这里的 p1 是 *([]int)，指向 nil，这是一个指针
	var p2 = make([]int, 5) // ✅ 这里的 p2 是 []int，已分配空间，这是一个切片

	fmt.Println(p1 == nil)  // false (p1 是一个指向 nil slice 的指针)
	fmt.Println(*p1 == nil) // true (p1 指向的 slice 仍然是 nil)
	fmt.Println(p2 == nil)  // false (p2 是已初始化的切片)

	/*
		4. make vs new 的使用场景
		✅ 使用 make 的场景（初始化 slice、map、channel）

		需要创建 slice（并分配底层数组）
		需要创建 map（并分配哈希表）
		需要创建 channel（并分配缓冲区）
		✅ 使用 new 的场景（创建指针）

		需要创建基础类型的指针
		需要创建结构体指针
		减少拷贝（指针可以避免值传递）
	*/

	// 5. new 的高级用法
	//📌 new 用于创建结构体指针
	type Person struct {
		Name string
		Age  int
	}

	var p50 *Person = new(Person)
	fmt.Println(p50)
	p5 := new(Person) // 返回 *Person 指针
	fmt.Println(p5)
	fmt.Println(p5.Name) // ""
	fmt.Println(p5.Age)  // 0

	// 📌 new 避免值拷贝
	//在 Go 语言中，值传递会导致数据拷贝，使用 new 返回指针可以避免拷贝：
	//func createInt() *int {
	//	return new(int) // 返回指针
	//}
	//6. 总结
	//操作			适用类型				返回类型		是否初始化		适用场景
	//new(T)		任何类型				*T（指针）	否（默认零值）	需要指针，避免值拷贝
	//make(T, ...)	slice、map、channel	T（非指针）	是				创建 slice、map、channel
	//🚀 make vs new 选择指南
	//✅ 如果要创建 slice、map、channel，并能直接使用 → make
	//✅ 如果只需要内存分配，并返回指针 → new
	//✅ 如果是 struct 结构体，一般用字面量 &T{} 代替 new(T)

	// 7. 推荐使用 &T{} 代替 new(T)
	//Go 更推荐使用 &T{} 而不是 new(T)：
	//✅ new(T) 适用于避免拷贝
	//✅ &T{} 更适用于初始化结构体
	p71 := new(Person) // 传统方式
	p72 := &Person{}   // 推荐方式
	fmt.Println(p71)
	fmt.Println(p72)

}

// Go 语言中的空值处理与 nil 使用详解 🚀
// 在 Go 语言中，空值（Zero Value） 是指变量在未初始化时的默认值，而 nil 是 Go 语言中一个特殊的 预定义标识符，用于表示引用类型的“空”状态。
// 在 Go 语言中，只有某些特定类型 可以取 nil，使用 nil 需要谨慎，避免 nil 相关的运行时错误（如 panic）。
func test54_05() {
	/*
	   	1. 什么是 nil？
	      在 Go 中，nil 主要用于表示某些引用类型的空值，比如：
	   	   指针 (*T)
	   	   切片 ([]T)
	   	   映射 (map[T]T)
	   	   通道 (chan T)
	   	   接口 (interface{})
	   	   函数 (func(...) ...)

	      这些类型的默认零值都是 nil，但它们的行为可能不同。例如：
	   	   nil map 不能写入
	   	   nil slice 可以读取但长度为 0
	   	   nil channel 可能会导致 Goroutine 永久阻塞
	   	   nil interface 可能会导致 panic
	*/

	// 2. nil 的使用场景
	//（1）指针
	//指针在 Go 里默认是 nil
	//常见场景：检测指针是否为 nil、传递 nil 指针给函数
	var ptr *int
	fmt.Println(ptr)
	if ptr == nil {
		fmt.Println("指针遍历ptr为nil，不能解引用")
	}
	println("--------------------------------------")

	// （2）切片（slice）
	// 切片的默认零值是 nil，但它可以被安全读取：
	var sli2 []int           // nil切片
	fmt.Println(sli2 == nil) // true
	fmt.Println(len(sli2))   // 0
	fmt.Println(cap(sli2))   // 0
	//fmt.Println(sli2[0])     // panic: runtime error: index out of range [0] with length 0

	// ⚠️ nil 切片和空切片 ([]int{}) 是不同的！
	var sli22 []int = []int{} // 空切片
	fmt.Println(sli22 == nil) // true
	fmt.Println(len(sli22))   // 0
	fmt.Println(cap(sli22))   // 0
	//fmt.Println(sli22[0])     // panic: runtime error: index out of range [0] with length 0

	// ✅ 最佳实践：
	//使用 nil 切片可以减少内存分配
	//但返回空切片比 nil 更好，以避免 nil 引发的 panic

	println("--------------------------------------")

	// （3）映射（map）
	//未初始化的 map 是 nil，不能直接写入：
	var map31 map[string]int
	fmt.Println(map31 == nil)
	//map31["apple"] = 10 // ❌ panic: assignment to entry in nil map
	fmt.Println(types.Nil{})

	// 以上的解决方法1，创建一个静态的
	var map32 map[string]int = map[string]int{}
	fmt.Println(map32 == nil)
	map32["apple"] = 10
	fmt.Println(map32)

	// 以上的解决方法2：使用make创建
	var map33 map[string]int = make(map[string]int)
	fmt.Println(map33 == nil)
	map33["apple"] = 101
	fmt.Println(map33)

	println("--------------------------------------")

	// https://chatgpt.com/c/67d22576-391c-8012-9f79-b30a4c487e3c
	// （4）通道（channel）

	// （5）接口（interface）

	/*
			4. nil 使用的最佳实践
			✅ 什么时候使用 nil？
				未初始化的引用类型变量（避免 panic）
				判断值是否存在
				返回 nil 以表示“无值”
				在 select 语句中控制 channel

			✅ 什么时候避免 nil？
				返回 slice 或 map 时，建议返回空值而不是 nil
				结构体接收器方法避免 nil 指针调用
				Goroutine 处理 nil 通道，避免死锁
				避免 interface{} 为空但 != nil 的问题
			5. 总结
		类型					默认值 (nil)		访问 nil 的行为
		指针 (*T)			nil				不能解引用
		切片 ([]T)			nil				可读但 len=0，不能 append
		映射 (map[T]T)		nil				可读但不能写入
		通道 (chan T)		nil				读写会阻塞
		接口 (interface{})	nil				nil 但可能持有非 nil 值
		掌握 nil 的使用方式，能够有效避免 panic，编写更健壮的 Go 代码！🚀
	*/
}

func getSlice(useNil bool) []int {
	if useNil {
		return nil // 返回 nil
	}
	return []int{} // 返回空切片
}

// Go 语言中的 map 详解及高阶用法 🚀
// 在 Go 语言中，map（映射）是一种内置的数据结构，用于存储键值对（key-value）。
// 它类似于其他语言中的 哈希表（HashMap） 或 字典（Dictionary），具有 O(1) 的查找和插入时间复杂度，是高效的数据存储方式之一。
func test54_04() {
	// 1. map 的基础概念
	//（1）什么是 map？
	//map 是一种无序的 key-value 对集合，其键必须是可比较的类型（如 string、int、bool、float、interface{} 等），
	//但不能是 slice、map 或 function（因为这些类型不可比较）。

	// (2)创建一个map
	// 直接创建
	map11 := map[int]string{
		1: "12345",
		2: "qaz",
		3: "wsx", // 注意：最后一个
	}
	fmt.Println(map11)

	// 使用make创建
	map12 := make(map[int]string)
	map12[1] = "qqq"
	map12[1] = "aaa"
	map12[2] = "qqq"
	map12[3] = "zzz"
	fmt.Println(map12)

	// 先声明变量，在赋值数据
	var map13 map[int]string = map[int]string{}
	fmt.Println(map13)
	map13[11] = "qwertyu"
	fmt.Println(map13)
	map13 = nil // 置空map之后就不能在使用了
	fmt.Println(map13)
	//map13[22] = "qwerqwertyutyu" // panic: assignment to entry in nil map

	//2. map 的基本操作
	// （1）增删改查
	map21 := make(map[int]string)
	map21[1] = "qqq"
	map21[2] = "aaa"
	map21[3] = "zzz"
	map21[4] = "xxx"
	map21[5] = "sss"
	map21[6] = "www"
	fmt.Println(map21)

	// b) 获取值
	fmt.Println(map21[5])

	// c) 判断键是否存在
	val, err := map21[5]
	if err {
		fmt.Println("数据获取成功！", val)
	} else {
		fmt.Println("数据获取失败失败！")
	}
	val2, err2 := map21[555]
	if err2 {
		fmt.Println("数据获取成功！", val2)
	} else {
		fmt.Println("数据获取失败失败！")
	}

	// d) 删除键
	fmt.Println(map21)
	delete(map21, 2)
	fmt.Println(map21)

	// e) 遍历 map
	// ⚠️ map 是无序的，遍历结果是随机的！⚠️
	for key, val := range map21 {
		fmt.Printf(" key = %d, val = %s \n", key, val)
	}
	println("------------------------------------------")

	// 3. map 的高阶用法
	//（1）map 作为函数参数
	map31 := make(map[int]string)
	map31[1] = "qqq"
	map31[2] = "aaa"
	map31[3] = "zzz"
	map31[4] = "xxx"
	map31[5] = "sss"
	map31[6] = "www"
	fmt.Println(map31)
	// ⚠️ map 作为函数参数时，是 "引用传递"，即函数内部修改会影响外部！
	mapFunc1(map31)
	fmt.Println(map31)
	println("------------------------------------------")

	// （2）map 的嵌套
	map32 := make(map[int]map[int]string)
	map32[1] = map[int]string{
		11: "1-1",
		12: "1-2",
		13: "1-3",
	}
	map32[2] = map[int]string{
		21: "2-1",
		22: "2-2",
	}
	map32[3] = map[int]string{
		31: "3-1",
	}
	map32[4] = map[int]string{}
	// map[1:map[11:1-1 12:1-2 13:1-3] 2:map[21:2-1 22:2-2] 3:map[31:3-1] 4:map[]]
	fmt.Println(map32)
	println("------------------------------------------")
	map32[4][41] = "4-1" // 不建议使用这种方法，因为map32[4]可能为nil，建议先拿出来看看map32[4]有没有在给他赋值
	// map32[5][41] = "5-1" // panic: assignment to entry in nil map
	val3, err := map32[4]
	if err {
		val3[42] = "4-2"
	}
	fmt.Println(map32)
	println("------------------------------------------")

	// 遍历嵌套map
	for key1, innerMap := range map32 {
		fmt.Printf("外层 key = %d \n", key1)
		for key2, val := range innerMap {
			fmt.Printf("\t 内层 key = %d， val = %s \n", key2, val)
		}
		println()
	}
	println("------------------------------------------")

	// （3）map 与 slice 结合
	// a) map[string][]int，map的value中存[]int
	var map33 = make(map[string][]int)
	fmt.Println(map33)
	map33["sli1"] = []int{1, 2, 3, 4, 5}
	map33["sli2"] = []int{21, 22, 232, 24, 25}
	fmt.Println(map33)
	println("------------------------------------------")
	// b) []map[string]，slice中存map[string]

	sli33 := []map[int]string{map11, map12, map13}
	fmt.Println(sli33)
	println("------------------------------------------")

	// 由于map是线程不安全的方法，因此需要引入线程安全的映射 sync.Map
	//（4）并发安全的 sync.Map
	//Go 的 map 不是线程安全的，多个 Goroutine 并发读写 map 可能会导致 fatal error: concurrent map writes。
	test54_04_syncMap()
	println("------------------------------------------")

}

// 声明一个全局变量sync.Map

var syncMap sync.Map

func test54_04_syncMap() {
	// 添加数据
	syncMap.Store("key1", 123)
	syncMap.Store("key2", 456)
	syncMap.Store("key3", 789)

	// 获取数据
	val, err := syncMap.Load("key1")
	if err {
		fmt.Println("syncMap[key1] = ", val)
	}

	// 遍历sync.Map
	syncMap.Range(func(key, value any) bool {
		fmt.Printf("syncMap[\"%s\"] = %d \n", key, value)
		return true
	})

	// 4. map 的性能优化
	//（1）预分配容量
	//如果事先知道大概的 map 大小，可以用 make(map[T]T, capacity) 预分配，减少扩容开销：
	map4 := make(map[string]int, 1000) // 预分配 1000 个键值对
	//（2）避免 map 频繁扩容
	//map 在元素数量增长时，可能会触发 rehash（重新分配内存），影响性能。若 map 频繁增删，可考虑 定期清空 或 重新创建 map：
	map4 = make(map[string]int, len(map4)) // 重新分配内存
	// 3）清空 map
	//Go 没有 clear(map) 方法，只能重新创建一个相同大小的map在赋值给该变量，或者 手动删除所有键（性能较差）通过遍历map来调用delete(map, key)
	map4 = make(map[string]int, len(map4)) // 直接重新创建

	fmt.Println(map4)

	//5. map 常见陷阱
	//（1）读取 nil map

}

func mapFunc1(m map[int]string) {
	for key, val := range m {
		m[key] = m[key] + "-" + strconv.Itoa(key)
		fmt.Printf(" key = %d, val = %s \n", key, val)
	}
}

// Go 语言中的 Slice 详解及高阶用法 🚀
// 在 Go 语言中，slice（切片）是基于 array（数组）的动态、可变长度的数据结构，提供了更灵活的数组操作能力。
// 理解 slice 是写好 Go 代码的关键，尤其是涉及内存管理、性能优化、高效数据操作等方面。
func test54_03() {

	/*
		1. Slice 的基础概念
		（1）什么是 Slice？
			slice 是一个 引用类型，本质上是对 底层数组的一个视图，它包含以下三部分：
				指向底层数组的指针
				切片的长度（len）
				切片的容量（cap）
	*/

	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sli1 := arr[0:8]
	fmt.Println(sli1)
	fmt.Println(len(sli1))
	fmt.Println(cap(sli1))

	// (2)、slice的创建方式
	// 1.2.1、直接使用 []T{} 语法
	sli121 := []int{11, 22, 33}
	fmt.Println(sli121)

	// 1.2.2、通过数组创建
	var arr122 []int = []int{1, 2, 3, 4, 5, 6}
	sli122 := arr122[0:3]
	fmt.Println(sli122)

	// 1.2.3、通过 make 创建
	var sli123 []int = make([]int, 6, 9)
	fmt.Println(sli123)
	_ = append(sli123, 666) // 可以看得到必须将append新创建的sli赋值才能生效，意思很明显append不改变原切片
	fmt.Println(sli123)
	sli123 = append(sli123, 999)
	fmt.Println(sli123)

	println("--------------------------------------")

	//2. Slice 的高级用法
	//（1）切片扩容
	//切片容量 cap 确定了最大可用空间，如果 append() 超出 cap，Go 运行时会创建新的底层数组，并将原切片的数据复制过去。
	s := []int{1, 2, 3}
	fmt.Printf("len = %d, cap = %d \n", len(s), cap(s)) // len = 3, cap = 3

	s = append(s, 4, 5, 6)                              // 超出原始容量，触发扩容
	fmt.Println(s)                                      // [1 2 3 4 5 6]
	fmt.Printf("len = %d, cap = %d \n", len(s), cap(s)) // len = 6, cap = 6
	s = append(s, 7, 8, 9)                              // 超出原始容量，触发扩容
	fmt.Println(s)                                      // [1 2 3 4 5 6 7 8 9]]
	fmt.Printf("len = %d, cap = %d \n", len(s), cap(s)) // len = 9, cap = 12

	/*（2）扩容机制
	Go 采用指数扩容策略（大致是 1.25x~2x 扩展），核心规则：
		长度 ≤ 1024 时，按 2 倍 扩展
		长度 > 1024 时，每次增长 1.25 倍
		触发扩容时，会创建新数组，拷贝旧数据到新数组
	*/

	//（3）切片共享底层数组
	//多个切片可能共享同一块底层数组，修改其中一个会影响其他切片：
	arr23 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr23)
	sli231 := arr23[0:8]
	sli232 := arr23[1:6] // 共享一个底层数组
	fmt.Println(sli231)
	fmt.Println(sli232)
	// 修改原数组，看输出
	arr23[2] = 666
	fmt.Println(arr23)  // 修改为666
	fmt.Println(sli231) // 修改为666
	fmt.Println(sli232) // 修改为666
	// 修改某一个切片数据，看输出
	sli231[3] = 999
	fmt.Println(arr23)  // 修改为 999
	fmt.Println(sli231) // 修改为 999
	fmt.Println(sli232) // 修改为 999
	// 以上表明，slice切片是共享了同一个底层数组，sli只是一个指向数组的指针
	// 以上在其他语言（如：Java或c）里面被称为浅拷贝

	// ✅ 解决以上浅拷贝的方法：使用 copy() 复制切片，避免共享底层数组：

	sli233 := make([]int, len(sli232))
	copy(sli233, sli232)
	fmt.Println(sli233)
	sli233[1] = 888
	fmt.Println(sli233)
	fmt.Println(arr23)  // 未修改
	fmt.Println(sli231) // 未修改
	fmt.Println(sli232) // 未修改

	sli233t := deepCopy(arr23, 1, 5)
	fmt.Println("deepCopy", sli233t)
	println("---------------------------------------")

	// 4）删除元素
	//Go 没有提供直接删除元素的方法，但可以用 append() 变通处理：
	//a) 删除索引 i 处的元素
	fmt.Println(sli231)
	i := 2
	// 前面的sli231[0:i]，表示取sli2310到i-1位置的数据
	// 后面的sli231[i+1:len(sli231)]，表示取i+1以后的所有数据，其中三个点...表示解压缩切片，即将[1,2,3]的数据转化为1,2,3。这个类似于js的ES6的写法
	sli231 = append(sli231[0:i], sli231[i+1:len(sli231)]...)
	fmt.Println(sli231)
	sli231 = append(sli231[:i], sli231[i+1:]...) // 简化版本
	fmt.Println(sli231)
	println("---------------------------------------")

	// b) 删除切片开头的元素，头删法
	fmt.Println(sli231)
	sli231 = sli231[1:]
	fmt.Println(sli231)
	println("---------------------------------------")

	// c) 删除切片末尾的元素，尾删法
	fmt.Println(sli231)
	sli231 = sli231[:len(sli231)-1]
	fmt.Println(sli231)
	println("---------------------------------------")

	/*
		（5）切片的性能优化
			a) 预分配容量，减少扩容
				如果已知最终长度，使用 make() 预分配内存，避免 append() 触发多次扩容
			b) 使用 copy() 避免修改原数据
				如果数据共享底层数组，建议使用 copy() 创建新切片，避免副作用。
			c) 避免 append() 引起的底层数组变化
				如果需要多个切片指向不同的内存，尽量提前 copy()，或直接创建新 slice。
	*/
}

// 深拷贝切片
// src: 原数据，可以是数组或切片
// startIndex: 开始拷贝的索引
// endIndex: 结束拷贝的索引
// res: 深拷贝返回的切片
func deepCopy(src []int, startIndex int, endIndex int) (res []int) {
	res = make([]int, endIndex-startIndex)
	copy(res, src[startIndex:endIndex])
	return
}

// 指针基本知识
func test54_02() {
	//使用 new 创建指针 pointer
	//1. 基本用法
	ptr := new(int)   // 创建一个指向 int 类型的指针
	fmt.Println(ptr)  // 输出类似 0xc0000120a0（指针地址）
	fmt.Println(*ptr) // 输出 0（int 的零值）
	//new(int) 分配了一个 int 类型的内存，并返回该内存的指针。
	//*ptr 获取指针指向的值，默认值为 0。

	//2. new 和 var 的区别
	var x int        // 直接声明变量 x，值为 0
	ptr2 := new(int) // 使用 new 分配内存并返回指针，指向 0

	fmt.Println(x)              // 0
	fmt.Println(&x)             // 0xc00000a138
	fmt.Println(*ptr2)          // 0
	fmt.Println(&ptr2)          // 0xc000064070
	fmt.Println(&(*ptr2))       // 0xc00000a140
	fmt.Println(*(&(*ptr2)))    // 0
	fmt.Println(&(*(&(*ptr2)))) // 0xc00000a140

	//区别：
	//var x int 直接分配变量 x 在栈帧上，类型为 int。
	//new(int)  返回的是指针 *int，分配的内存在堆上（Go 运行时会优化）。

	//3. new 创建结构体指针
	type Person struct {
		Name string
		Age  int
	}

	p := new(Person) // 创建 *Person 指针
	fmt.Println(p)   // 输出 &{ 0}（指向零值的结构体）

	//相当于：
	var p2 *Person = &Person{} // 另一种写法
	fmt.Println(p2)

	//4. new vs & 取地址
	p1 := new(int)   // 使用 new
	p22 := &Person{} // 直接取地址

	fmt.Println(p1, *p1) // 0xc000010250 0
	fmt.Println(p22)     // &{ 0}
	//区别：
	//new(Type) 只分配内存，不初始化（返回指针）。
	//&Type{} 直接初始化结构体（也是返回指针）。

	//5. 什么时候用 new？
	//一般不推荐用 new，因为 &Type{} 更直观，代码更简洁。
	//new 适用于简单类型（int、float64）的指针分配，不常用于复杂结构体。
	//总结
	//✅ new(T) 分配内存并返回指针，默认值是零值。
	//✅ new 适用于创建基本类型和结构体的指针，但不如 &T{} 直观。
	//✅ 在实际开发中，结构体推荐用 &T{}，而不是 new(T)。
}

func test54_01() {
	/*
		1. 布尔类型 (bool)
		取值：true / false
		默认值：false
		不能使用 int 代替 bool（如 1 不能表示 true）
		逻辑运算符：&&（与）、||（或）、!（非）
	*/

	/*
		2. 整型 (int, uint, intX, uintX)
		Go 语言的整型有 有符号 (int) 和 无符号 (uint) 两种：

		类型	位数	范围
		int8	8	-128 ~ 127
		int16	16	-32,768 ~ 32,767
		int32	32	-2,147,483,648 ~ 2,147,483,647
		int64	64	-9,223,372,036,854,775,808 ~ 9,223,372,036,854,775,807
		uint8	8	0 ~ 255
		uint16	16	0 ~ 65,535
		uint32	32	0 ~ 4,294,967,295
		uint64	64	0 ~ 18,446,744,073,709,551,615
		int / uint	取决于系统	32 位系统是 32 位，64 位系统是 64 位
		注意点
		int 和 uint 大小取决于 CPU 架构（32/64 位）。
		整型转换可能导致溢出，如 var x int8 = 128 会报错。
		uintptr 可存储指针地址（与 unsafe.Pointer 相关）。
	*/

	/*
		3. 浮点数 (float32, float64)
		取值范围：
		float32: 约 7 位小数精度
		float64: 约 15 位小数精度
		默认值：0.0
		不能直接与 int 运算（需要转换）
		计算可能有 精度误差（浮点数不能精确表示部分小数）
	*/

	/*
		4. 复数 (complex64, complex128)
		complex64 由 float32 + float32i 组成
		complex128 由 float64 + float64i 组成
		复数的实部和虚部可通过 real() 和 imag() 获取
	*/

	/*
		5. 字符 (rune 和 byte)
		byte（uint8 别名）：存储 ASCII 字符
		rune（int32 别名）：存储 Unicode（UTF-8）字符
		Go 的字符串是 不可变 的
		rune 适用于 中文、emoji 等非 ASCII 字符
	*/
	// Go 的字符串是 不可变 的
	var ch1 byte = 'A'
	var ch2 rune = '你'
	fmt.Printf("%c %d\n", ch1, ch1) // A 65
	fmt.Printf("%c %d\n", ch2, ch2) // 你 20320

	/*
		6. 字符串 (string)
		默认值：""
		不可变（不能直接修改 string 的字符）
		使用 + 进行拼接
		使用 len() 获取字节长度（不是字符数）
		处理 Unicode 需要 rune 或 utf8.RuneCountInString()
	*/

	s := "你好，Go"
	fmt.Println(len(s)) // 输出 1（因为中文占 3 字节） 11 = 3*3 + 2

	r := []rune(s)
	fmt.Println(len(r)) // 输出 5（因为有 5 个字符）

	s2 := "hello"
	bs := []byte(s2) // 转换为可变的 `[]byte`
	bs[0] = 'H'
	s3 := string(bs) // 转回 string
	fmt.Println(s3)  // "Hello"

	r4 := []rune(s2)
	r4[0] = 97
	fmt.Println(r4)
	s5 := string(r4) // 转回 string
	fmt.Println(s5)  // "aello"

	/*
		7. 指针 (*T)
		var p *int = nil // 默认值 nil
		取地址 &，获取指针的值 *
		new() 创建指针
	*/
	a := 10
	p := &a
	fmt.Println(*p) // 10
	*p = 20
	fmt.Println(a) // 20

	//在 Go 语言中，new 关键字可以用于创建 指向零值的变量的指针。它的作用是分配内存并返回指向该内存的指针。
	ptr := new(int)   // 创建一个指向 int 类型的指针
	fmt.Println(ptr)  // 输出类似 0xc0000120a0（指针地址）
	fmt.Println(*ptr) // 输出 0（int 的零值）
	//new(int) 分配了一个 int 类型的内存，并返回该内存的指针。
	//*ptr 获取指针指向的值，默认值为 0。

	/*
								8. 切片 ([]T)
								变长数组，底层基于 array
								len() 获取长度，cap() 获取容量
								append() 动态扩展容量

					数组 vs. 切片
					特性				数组 (array)			切片 (slice)
					长度				固定，不可变			动态增长
					是否引用传递		值传递，拷贝整个数组	引用传递，底层共享数据
					适用场景			适合 小型固定数据		适合 灵活变长数据

				总结
				数组是 Go 的复合数据类型，不是基本数据类型。
				数组长度固定，不同长度的数组是不同的类型。
				数组是值类型，函数传递时会拷贝整个数组（不像切片是引用传递）。
				数组适用于小型、固定大小的数据，大部分情况下应该使用 切片（slice） 代替数组。
				如果你的数据长度是动态变化的，通常应该选择 切片（slice） 而不是 数组（array）。


			切片的本质
			切片由 三部分 组成：

			指向底层数组的指针（ptr）
			切片长度（len）：当前切片包含的元素数
			切片容量（cap）：切片底层数组从 ptr 开始的最大可用元素数

		注意：如果 cap(s) 不足，append() 会创建新的底层数组，原来的切片将不会修改原数组。
	*/
	s11 := []int{1, 2, 3}
	s11 = append(s11, 4)
	fmt.Println(s11) // [1 2 3 4]

	/*
		9. 映射 (map[K]V)
		无序键值对
		make(map[K]V) 初始化
		v, ok := map[key] 判断键是否存在
	*/
	m := make(map[string]int)
	m["age"] = 25
	v, ok := m["age"]
	fmt.Println(v, ok) // 25 true

	/*
		10. 接口 (interface{})
		interface{} 可表示 任意类型
		需要 类型断言 (.(T)) 或 反射 (reflect 包) 处理
	*/
	var x interface{} = "hello"
	s12, ok := x.(string)
	fmt.Println(s12, ok) // hello true

	/*
		类型				关键特性
		bool			逻辑运算，不与 int 互换
		int/uint		整数类型，避免溢出
		float32/float64	可能有精度误差
		complex64/128	复数计算
		byte/rune		处理字符，rune 用于 Unicode
		string			不可变，可用 rune 处理 Unicode
		pointer			变量地址操作
		slice			变长数组
		map				无序键值存储
		interface{}		动态类型
	*/
}
