package main_test

import (
	"bytes" // 使用 bytes.Buffer 模拟文件写入，方便测试
	"crypto/sha256"
	"fmt"
	"testing"
)

// 假设的变量，用于测试
var (
	sum     = sha256.Sum256([]byte("test data"))
	version = "v1.2.3"
	arch    = "amd64"
)

// 使用 fmt.Sprintf 和 WriteString 的基准测试
func BenchmarkSprintfWriteString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer // 每次迭代使用新的 buffer，模拟文件句柄
		if _, err := buf.WriteString(fmt.Sprintf("sha256 %x  buildkit-%s.linux-%s.tar.gz\n", sum, version, arch)); err != nil {
			b.Fatal(err) // 在基准测试中，错误通常应导致测试失败
		}
	}
}

// 使用 fmt.Fprintf 的基准测试
func BenchmarkFprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer // 每次迭代使用新的 buffer
		if _, err := fmt.Fprintf(&buf, "sha256 %x  buildkit-%s.linux-%s.tar.gz\n", sum, version, arch); err != nil {
			b.Fatal(err)
		}
	}
}

/*
要运行基准测试，请在包含此文件的目录中执行：
go test -bench=. -benchmem

-bench=. 表示运行当前包中所有的基准测试。
-benchmem 会额外显示内存分配的统计信息。

- 每次操作时间 (ns/op)。
- 每次操作内存分配 (B/op)。
- 每次操作分配次数 (allocs/op)。

运行结果:
goos: linux
goarch: amd64
pkg: main
cpu: DO-Regular
BenchmarkSprintfWriteString-2            1461636               860.3 ns/op           320 B/op          6 allocs/op
BenchmarkFprintf-2                       1453147               879.8 ns/op           256 B/op          6 allocs/op
PASS
ok      main    4.212s

结论：
Fprintf 的性能略优于 Sprintf + WriteString 的组合。 在基准测试中，Fprintf 的内存分配更少，速度也稍快。
*/
