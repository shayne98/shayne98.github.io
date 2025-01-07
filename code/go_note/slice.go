package main
import (
	"fmt"
	"unsafe"
)
func main() {
	arr :=[10]int{1,2,3,4,5,6,7,8,9,10}
	slice1:= arr[1:4]
	slice2:= arr[7:]
	fmt.Printf("slice1 len: %v cap: %v elem %v\n",len(slice1),cap(slice1),slice1)
	fmt.Printf("slice2 len: %v cap: %v elem %v\n",len(slice2),cap(slice2),slice2)
	
	// slice1 len: 3 cap: 9 elem [2 3 4]
	// slice2 len: 3 cap: 3 elem [8 9 10]
	// 可以通过append 来延长slice1的长度，同时修改了底层数组
	slice1 = append(slice1,-1)
	fmt.Printf("slice1 len: %v cap: %v elem %v arr %v\n",len(slice1),cap(slice1),slice1,arr)
	// slice1 len: 4 cap: 9 elem [2 3 4 -1] arr [1 2 3 4 -1 6 7 8 9 10]
	// 若给slice2 append操作会超出cap，此时会触发拷贝并开辟新的底层数组,此时修改元素不会影响arr
	ptr := unsafe.Pointer(&slice2[0])
    fmt.Printf("The starting address of the slice2 before append: %p\n", ptr)
	slice2 = append(slice2,-1)
	slice2[0] = -1
	ptr = unsafe.Pointer(&slice2[0])
    fmt.Printf("The starting address of the slice2 after append: %p\n", ptr)
	// append操作后len 和cap均有变化 
	fmt.Printf("slice2 len: %v cap: %v elem %v arr %v\n",len(slice2),cap(slice2),slice2,arr)
	// slice2 len: 4 cap: 6 elem [-1 9 10 -1] arr [1 2 3 4 -1 6 7 8 9 10] 
	

}
