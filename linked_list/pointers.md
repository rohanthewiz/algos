Given:
```go
func main() {
	a := 2
	p := &a
	fmt.Println(*p)
}
```

### What is a Pointer?
- A variable holds a value
- A pointer holds the address of the variable
- In general terms: A pointer points to the container of a value

```go
o := struct{ attrib1 int, attrib2 int}
```