Any integer type will work with %c, but the rune alias indicates that the number 960 represents a character.  
Rather than memorize Unicode code points, Go provides a character literal. Just enclose a character in single quotes 'A'. If no type is specified, Go will infer a rune, so the following three lines are equivalent:  
```go
grade := 'A'
var grade = 'A'
var grade rune = 'A'
```
Though character literals can also be used with the 'byte' alias, it is intended for binary data. A byte is an alias for the uint8 type. 
Byte can be used for English characters defined by ASCII, an older 128-character subset of Unicode.  
  
Strings in Go are immutable, as they are in Python, Java, and JavaScript. Unlike strings in Ruby and character arrays in C, you can’t modify a string in Go