# Go Data Types
Basic Go datatypes
- bool
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64
- byte (alias for uint8)
- rune (alias for uint32)
- float32, float64
- complex64, complex128

## Runes
- Unicode is a superset of ASCII
- Unicode contains all the characters that are present in today's writing system
    - These include all dialects, accents, control characters (space, tab, etc)
- Each character is assigned a standard number called a `Unicode Code Point`
- In Go, the Unicode Code Point is called a `rune`  