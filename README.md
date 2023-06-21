# SIMD Implementation for ARM NEON in Golang

This repository provides a collection of SIMD (Single Instruction, Multiple Data) implementations specifically designed for ARM NEON architecture in Golang. SIMD allows parallel processing of data using the same instruction applied to multiple data elements, resulting in significant performance improvements for certain types of algorithms.

## Future Plans

We are actively working on expanding the SIMD implementation to support x86 architecture as well. The upcoming x86 implementation will provide similar SIMD functionalities for parallel processing on x86-based systems.

## Features

- ARM NEON SIMD implementation for various data processing operations.
- Optimized code to take advantage of the parallel processing capabilities of the ARM NEON architecture.
- Examples and usage instructions for each SIMD operation.

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

	"github.com/alivanz/go-simd/arm/neon"
)

func main() {
	var a, b neon.Int8X8
	for i := 0; i < 8; i++ {
		a[i] = neon.Int8(i)
		b[i] = neon.Int8(i * i)
	}
	log.Printf("a = %+v", b)
	log.Printf("b = %+v", a)
	log.Printf("add = %+v", neon.VaddlS8(a, b))
	log.Printf("mul = %+v", neon.VmullS8(a, b))
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

For any questions or feedback regarding this repository, please feel free to contact me at alivan1627@gmai.com