package main
 import "github.com/tigthor/gopract/bytestrings"
 func main() {
 err := bytestrings.WorkWithBuffer()
 if err != nil {
 panic(err)
 }
 // each of these print to stdout
 bytestrings.SearchString()
 bytestrings.ModifyString()
 bytestrings.StringReader()
 }
