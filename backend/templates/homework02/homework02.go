package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ============ 指针练习 ============

// 1. 定义一个函数，接收一个整数指针作为参数，在函数内部将该指针指向的值增加10
func AddTen(num *int) {
	// 用*来获取指针指向的值，然后加10
	*num = *num + 10
}

// 2. 接收一个整数切片的指针，将切片中的每个元素乘以2
func MultiplyByTwo(nums *[]int) {
	// 遍历切片中的每个元素
	for i := 0; i < len(*nums); i++ {
		// 把每个元素乘以2
		(*nums)[i] = (*nums)[i] * 2
	}
}

// ============ Goroutine练习 ============

// 3. 使用协程打印奇数和偶数
func PrintOddEven() {
	// 创建一个等待组，用来等待所有协程完成
	var wg sync.WaitGroup

	// 添加2个协程到等待组
	wg.Add(2)

	// 启动第一个协程，打印奇数
	go func() {
		// 协程结束时通知等待组
		defer wg.Done()
		for i := 1; i <= 10; i = i + 2 {
			fmt.Printf("奇数: %d\n", i)
		}
	}()

	// 启动第二个协程，打印偶数
	go func() {
		// 协程结束时通知等待组
		defer wg.Done()
		for i := 2; i <= 10; i = i + 2 {
			fmt.Printf("偶数: %d\n", i)
		}
	}()

	// 等待所有协程完成
	wg.Wait()
}

// 4. 任务调度器，并发执行任务并统计执行时间
func TaskScheduler(tasks []func()) {
	// 创建等待组
	var wg sync.WaitGroup

	// 给每个任务启动一个协程
	for i, task := range tasks {
		wg.Add(1)

		// 启动协程执行任务
		go func(taskNum int, t func()) {
			defer wg.Done()

			// 记录开始时间
			startTime := time.Now()

			// 执行任务
			t()

			// 计算执行时间
			duration := time.Since(startTime)
			fmt.Printf("任务 %d 执行完成，耗时: %v\n", taskNum, duration)
		}(i+1, task)
	}

	// 等待所有任务完成
	wg.Wait()
}

// ============ 面向对象练习 ============

// 定义Shape接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 5. 定义Rectangle结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle的Area方法
func (r Rectangle) Area() float64 {
	// 矩形面积 = 长 × 宽
	return r.Width * r.Height
}

// Rectangle的Perimeter方法
func (r Rectangle) Perimeter() float64 {
	// 矩形周长 = 2 × (长 + 宽)
	return 2 * (r.Width + r.Height)
}

// 定义Circle结构体
type Circle struct {
	Radius float64
}

// Circle的Area方法
func (c Circle) Area() float64 {
	// 圆形面积 = π × 半径²
	pi := 3.14159
	return pi * c.Radius * c.Radius
}

// Circle的Perimeter方法
func (c Circle) Perimeter() float64 {
	// 圆形周长 = 2 × π × 半径
	pi := 3.14159
	return 2 * pi * c.Radius
}

// 6. 使用组合的方式创建Person和Employee结构体
type Person struct {
	Name string
	Age  int
}

// Employee结构体组合了Person
type Employee struct {
	Person     // 匿名字段，组合Person
	EmployeeID string
}

// Employee的PrintInfo方法
func (e Employee) PrintInfo() {
	fmt.Printf("员工信息:\n")
	fmt.Printf("  姓名: %s\n", e.Name)
	fmt.Printf("  年龄: %d\n", e.Age)
	fmt.Printf("  员工ID: %s\n", e.EmployeeID)
}

// ============ Channel练习 ============

// 7. 使用通道实现两个协程之间的通信
func ChannelCommunication() {
	// 创建一个通道，用来传递整数
	ch := make(chan int)

	// 创建等待组
	var wg sync.WaitGroup
	wg.Add(2)

	// 生产者协程：生成1到10的整数
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			// 把数字发送到通道
			ch <- i
		}
		// 关闭通道，表示不再发送数据
		close(ch)
	}()

	// 消费者协程：从通道接收整数并打印
	go func() {
		defer wg.Done()
		// 从通道接收数据，直到通道关闭
		for num := range ch {
			fmt.Printf("接收到: %d\n", num)
		}
	}()

	// 等待所有协程完成
	wg.Wait()
}

// 8. 使用带缓冲的通道
func BufferedChannel() {
	// 创建一个缓冲大小为10的通道
	ch := make(chan int, 10)

	// 创建等待组
	var wg sync.WaitGroup
	wg.Add(2)

	// 生产者协程：发送100个整数
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		// 发送完毕，关闭通道
		close(ch)
	}()

	// 消费者协程：接收并打印
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("收到: %d\n", num)
		}
	}()

	// 等待所有协程完成
	wg.Wait()
}

// ============ 锁机制练习 ============

// 9. 使用sync.Mutex保护共享计数器
func MutexCounter() int {
	// 共享的计数器
	counter := 0
	// 创建互斥锁
	var mutex sync.Mutex
	// 创建等待组
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// 每个协程递增1000次
			for j := 0; j < 1000; j++ {
				// 加锁
				mutex.Lock()
				// 递增计数器
				counter = counter + 1
				// 解锁
				mutex.Unlock()
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 返回最终的计数器值（应该是10000）
	return counter
}

// 10. 使用原子操作实现无锁计数器
func AtomicCounter() int64 {
	// 使用int64类型的原子计数器
	var counter int64
	// 创建等待组
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// 每个协程递塹1000次
			for j := 0; j < 1000; j++ {
				// 使用原子操作递增
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 返回最终的计数器值（应该是10000）
	return counter
}

// main函数，演示所有功能
func main() {
	fmt.Println("========== 指针练习 ==========")
	fmt.Println("\n=== 1. AddTen - 通过指针修改值 ===")
	num := 5
	fmt.Printf("修改前: %d\n", num)
	AddTen(&num)
	fmt.Printf("修改后: %d\n", num)

	fmt.Println("\n=== 2. MultiplyByTwo - 修改切片元素 ===")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("修改前: %v\n", nums)
	MultiplyByTwo(&nums)
	fmt.Printf("修改后: %v\n", nums)

	fmt.Println("\n========== Goroutine练习 ==========")
	fmt.Println("\n=== 3. PrintOddEven - 并发打印奇偶数 ===")
	PrintOddEven()

	fmt.Println("\n=== 4. TaskScheduler - 任务调度器 ===")
	tasks := []func(){
		func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("任务A完成")
		},
		func() {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("任务B完成")
		},
		func() {
			time.Sleep(150 * time.Millisecond)
			fmt.Println("任务C完成")
		},
	}
	TaskScheduler(tasks)

	fmt.Println("\n========== 面向对象练习 ==========")
	fmt.Println("\n=== 5. Shape接口 - Rectangle和Circle ===")
	rect := Rectangle{Width: 5.0, Height: 3.0}
	fmt.Printf("矩形: 宽=%.1f, 高=%.1f\n", rect.Width, rect.Height)
	fmt.Printf("  面积: %.2f\n", rect.Area())
	fmt.Printf("  周长: %.2f\n", rect.Perimeter())

	circle := Circle{Radius: 4.0}
	fmt.Printf("圆形: 半径=%.1f\n", circle.Radius)
	fmt.Printf("  面积: %.2f\n", circle.Area())
	fmt.Printf("  周长: %.2f\n", circle.Perimeter())

	fmt.Println("\n=== 6. Employee - 组合结构体 ===")
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  28,
		},
		EmployeeID: "EMP001",
	}
	emp.PrintInfo()

	fmt.Println("\n========== Channel练习 ==========")
	fmt.Println("\n=== 7. ChannelCommunication - 通道通信 ===")
	ChannelCommunication()

	fmt.Println("\n=== 8. BufferedChannel - 带缓冲的通道 ===")
	BufferedChannel()

	fmt.Println("\n========== 锁机制练习 ==========")
	fmt.Println("\n=== 9. MutexCounter - 互斥锁计数器 ===")
	result1 := MutexCounter()
	fmt.Printf("最终计数: %d\n", result1)

	fmt.Println("\n=== 10. AtomicCounter - 原子操作计数器 ===")
	result2 := AtomicCounter()
	fmt.Printf("最终计数: %d\n", result2)

	fmt.Println("\n========== 所有练习完成! ==========")
}
