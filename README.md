# 汉明距离（Hamming Distance）

# 一、安装
```bash
go get -u github.com/golang-infrastructure/go-hamming-distance
```

# 二、代码示例  

计算Int类型的汉明距离：

```go
a := 1
b := 3
distance := hamming_distance.Int(a, b)
t.Log(distance)
// Output:
// 1
```

计算String类型的汉明距离：

```go
s1 := "abcd"
s2 := "abcf"
distance := hamming_distance.String(s1, s2)
// distance, err := hamming_distance.StringE(s1, s2)
t.Log(distance)
// Output:
// 1
```

# 三、汉明距离是什么及实现方式

TODO 2022-12-4 23:55:26 

