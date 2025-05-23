# SIMD Implementation in Golang

This repository contains an implementation of SIMD (Single Instruction, Multiple Data) operations in Go, specifically targeting ARM NEON architecture. The goal is to provide optimized parallel processing capabilities for certain computational tasks.

## Future Plans

We are actively working on expanding the SIMD implementation to support x86 architecture as well. The upcoming x86 implementation will provide similar SIMD functionalities for parallel processing on x86-based systems.

## Hacks

When we call a C function through CGO, there are some overheads due to Go design.
In general, avoiding CGO would be a good idea.
But I found a hack, instead of relying on CGO, we can utilize `linkname` directive to call C code, bypass CGO, and get better performance.

```
goos: darwin
goarch: arm64
pkg: github.com/alivanz/go-simd/arm/neon
BenchmarkMultRef-8                131395              9168 ns/op
BenchmarkMultSimd-8               598742              1954 ns/op
BenchmarkMultSimdBypass-8         605554              1959 ns/op
BenchmarkMultSimdFull-8          1816879               661.3 ns/op
BenchmarkMultSimdCgo-8             13020             92213 ns/op
PASS
```

```
goos: darwin
goarch: arm64
pkg: github.com/alivanz/go-simd/arm/neon
cpu: Apple M2
BenchmarkVmulqF32N-8                8848            124616 ns/op        33657.86 MB/s       1422 B/op          0 allocs/op
BenchmarkVmulqF32C-8                2256            528683 ns/op        7933.49 MB/s        5577 B/op          0 allocs/op
BenchmarkVmulqF32Ref-8              3630            327995 ns/op        12787.69 MB/s       3466 B/op          0 allocs/op
PASS
ok      github.com/alivanz/go-simd/arm/neon     5.793s
```

The floating-point multiplication benchmarks demonstrate significant performance differences between implementations:

- `VmulqF32N` (Native): Achieves the highest throughput at 33.6 GB/s with minimal memory allocation (1422 B/op). This implementation leverages direct SIMD instructions for optimal performance.
- `VmulqF32C` (C): Shows the lowest performance at 7.9 GB/s with higher memory allocation (5577 B/op), likely due to the overhead of CGO calls and memory management.
- `VmulqF32Ref` (Reference): Performs at 12.8 GB/s with moderate memory usage (3466 B/op), serving as a baseline for comparison.

These results highlight the importance of using native SIMD implementations over CGO-based solutions for performance-critical applications. The native implementation is approximately 2.6x faster than the reference implementation, while the C implementation is about 1.6x slower than the reference.

## Features

- SIMD operations for ARM NEON architecture.
- High-performance parallel processing for specific tasks.
- Utilizes the power of SIMD instructions to process multiple data elements simultaneously.
- Supports a range of data types, including integers and floating-point numbers.
- Modular design for easy integration into existing projects.
- Well-documented code for understanding and extending the implementation.

## Roadmap

- [x] Implement SIMD operations for ARM NEON architecture.
- [ ] Add support for x86 architecture.
- [ ] Expand SIMD operations for additional data types.
- [ ] Optimize performance for specific use cases.
- [ ] Develop comprehensive test suite for validation.

## Usage

To use the SIMD implementations in your project, follow these steps:

1. Import the required package in your Go code:

```go
import "github.com/alivanz/go-simd"
```

2. Use the SIMD functions in your code as needed. Example:

```go
package main

import (
	"log"

	"github.com/alivanz/go-simd/arm"
	"github.com/alivanz/go-simd/arm/neon"
)

func main() {
	var a, b arm.Int8X8
	var add, mul arm.Int16X8
	for i := 0; i < 8; i++ {
		a[i] = arm.Int8(i)
		b[i] = arm.Int8(i * i)
	}
	log.Printf("a = %+v", a)
	log.Printf("b = %+v", b)
	neon.VaddlS8(&add, &a, &b)
	neon.VmullS8(&mul, &a, &b)
	log.Printf("add = %+v", add)
	log.Printf("mul = %+v", mul)
}

```

## Supported Operations

Only ARM Neon supported, for now.

Refer to the documentation in each respective file for more details on how to use each operation.

## Contributing

Contributions to this project are welcome. To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with descriptive messages.
4. Push your changes to your forked repository.
5. Submit a pull request to the main repository.

Please ensure that your code follows the existing code style and includes appropriate tests.

## Acknowledgments

- The ARM NEON architecture documentation for providing valuable insights into SIMD programming techniques.
- The open-source community for their contributions and inspiration.

## Contact

For any questions or feedback regarding this repository, please feel free to contact me at [alivan1627@gmail.com](mailto:alivan1627@gmail.com)