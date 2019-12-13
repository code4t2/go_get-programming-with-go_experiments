To cater to situations where integers are too small, floating-point too imprecise, or another numeric type would be more suitable  
**Constants** can be declared with a type, just like variables. And just like variables, a uint64 constant can’t possibly contain a number like 24 quintillion:
```go
const distance uint64 = 24000000000000000000          // (1)
```
    (1) Constant 24000000000000000000 overflows uint64  
For **variables**, Go uses _type inference_ to determine the type, and in the case of 24 quintillion, overflows the int type.  
**Constants** are different. Rather than infer a type, constants can be untyped. The following line doesn’t cause an overflow error:
```go
const distance = 24000000000000000000
```
Constants are declared with the `const` keyword, but **every literal value in your program is a constant too**. That means unusually sized numbers can be used directly. **Calculations on constants and literals are performed during compilation rather than while the program is running**. The Go compiler is written in Go. Under the hood, untyped numeric constants are backed by the big package. 
There’s a caveat to constants of unusual size. Though the Go compiler utilizes the big package for untyped numeric constants, `const`s and `big.Int` values aren’t interchangeable.