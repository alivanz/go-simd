package neon

/*
#include <arm_neon.h>
*/
import "C"

type Poly8 uint8
type Poly16 uint16
type Poly64 uint64
type Poly128 C.poly128_t
