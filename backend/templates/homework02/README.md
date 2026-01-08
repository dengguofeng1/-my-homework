# Go 作业说明

## homework01 - Go 基础

这个作业包含8个基础算法题目：

1. **只出现一次的数字** - 使用 map 统计每个数字出现的次数
2. **回文数** - 反转数字并比较
3. **有效的括号** - 使用切片模拟栈来匹配括号
4. **最长公共前缀** - 逐个比较字符串
5. **加一** - 处理进位问题
6. **删除有序数组中的重复项** - 双指针法
7. **合并区间** - 先排序再合并
8. **两数之和** - 使用 map 记录已遍历的数字

### 运行测试

```bash
cd backend/templates/homework01
go test -v
```

## homework02 - Go 进阶

这个作业包含10个进阶练习：

### 指针
1. **AddTen** - 通过指针修改变量值
2. **MultiplyByTwo** - 通过切片指针修改切片元素

### Goroutine
3. **PrintOddEven** - 使用协程并发打印奇偶数
4. **TaskScheduler** - 任务调度器，并发执行任务并统计时间

### 面向对象
5. **Shape接口** - Rectangle 和 Circle 实现 Shape 接口
6. **Employee** - 使用组合方式创建结构体

### Channel
7. **ChannelCommunication** - 使用通道进行协程间通信
8. **BufferedChannel** - 使用带缓冲的通道

### 锁机制
9. **MutexCounter** - 使用 sync.Mutex 保护共享数据
10. **AtomicCounter** - 使用原子操作实现无锁计数器

### 运行测试

```bash
cd backend/templates/homework02
go test -v
```

## 代码风格说明

这些作业代码采用初学者友好的风格：
- 详细的中文注释
- 简单直接的实现方式
- 避免使用复杂的语法特性
- 每一步都有清晰的说明
