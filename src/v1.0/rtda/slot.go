package rtda

import "v1.0/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
