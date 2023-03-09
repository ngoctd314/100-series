# All necessary concept in Golang

## 1. Memory order guarantees in Go

Many compilers (at compile time) and CPU processors (at run time) often make some optimizations by adjusting the instruction orders, so that the instruction execution orders may differ from the orders presented in code. Instruction ordering is also often called memory ordering.

Surely, instruction reordering can't be arbitrary. The basic requirement for a reordering inside a specified goroutine is the reordering must not be detectable by the goroutine itself if the goroutine doesn't share data with other goroutines. In other words, from the perspective of such a goroutine, it can think its instruction execution order is always the same as the order specified by code, even if instruction reordering really happens inside it.

However, if some goroutines share some data, then instruction reordering happens inside one of these goroutine may be observed by the others goroutines.