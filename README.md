# Huffman Tree
This is a poc for huffman lossless compression for text / byte arrays

# Example
```go
    // Ideally, this is better for larger strings eg. the one in the examples/example.go file
    import (
       encoder "github.com/obbap1/huffman/encoder"
       decoder "github.com/obbap1/huffman/decoder"
    )
    input := "hello"
    encoderHuffman := encoder.Init(input)
    compressedStr, huffmanNode, stats := encoderHuffman.Do() // compressedStr should be a string of bits. for example "1100101"
    // decode
    decoderHuffman := decoder.Init(compressedStr, huffmanNode)
    originalStr, stats := decoderHuffman.Do()
    // originalStr should be the same thing with the input "hello"
    fmt.Printf("%s, %s and %s", compressedStr, originalStr, input)

```

# stats 
The stats is a struct with keys <br>
`OriginalByteSize`: The original number of bytes before compression <br>
`CompressedByteSize`: The number of bytes after compression<br>
`PercentageSaved`: The percentage saved after compression <br>

# Examples
There are examples in the `examples` folder. 

# Tests
You can run tests with `go test ./...`. 
You can also use [act](https://github.com/nektos/act) to run the tests from github actions