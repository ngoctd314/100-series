## 1. What is rune, Why rune alias for int32

In Go, a rune is a Unicode represented a single value. UTF-8 encodes characters from one to four bytes. Hence, up to 32 bits. It's reason why in Go, a rune is an alias of int32.

Something else to hightlight about UTF-8: some people believe that Go strings are always UTF-8, yet this isn't true. A string is a sequence of arbitrary bytes; it's not necessarily based on UTF-8. Hence, when we manipulate a variable that wasn't initialized from a string literal (reading from file system), we can't necessarily expect it uses the UTF-8 encoding.

**Question: Tell me about result of below code, then explain your answer**
```go
func main() {
	s := "hêllo"
	for i, v := range s {
		fmt.Printf("%d: %s", i, string(s[i]))
		fmt.Printf("%d: %s", i, string(v))
	}
}
```
 
**Answer**
- Result

```txt
0: h 0: h
1: Ã 1: ê
3: l 3: l
4: l 4: l
5: o 5: o
```

- Explain

We don't iterate over each rune, instead, we iterate over each starting index of a rune (value of i)

Here are another approach

```go
func main() {
	s := "hêllo"
	runes := []rune(s)
	for i, v := range runes {
		fmt.Printf("%d: %s ", i, string(runes[i]))
		fmt.Printf("%d: %s", i, string(v))
		fmt.Println()
	}
}
```
The only difference is about the position: instead of printing the starting index of the rune's bytes sequence, it prints the run's index directly.

Note that this solution introduces a runtime overhead compared to the previous one. Indeed, converting a string into a slice of rune requires allocating an additional slice and converting the bytes into runes; hence, an O(n) time complexity with n the number of bytes in the string. Therefore, if we want to iterate over all the runes, we should use the first solution.

However, if we want to access the i th rune of a string with the first option, we don’t have access to the rune index but the starting index of a rune in the byte sequence. Hence, we should favor the second option in most cases.

```go
s := "hêllo"
r := []rune(s)[4]
fmt.Printf("%c\n", r) // o
```