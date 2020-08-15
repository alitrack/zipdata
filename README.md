# zipdata

* Generate go code from directory(zip it recursively) or zip file.
* The generated code does not need this library.
* The generated code can automatically unzip the zipdata to the given folder.

## Installation

To install the command line program, use the following:

```bash
go get -u github.com/alitrack/zipdata
```

## Usage

``` bash
# zip directory named static  and generate code named static.go
$ zipdata -s static -g static.go -d .

# or zip static.zip file  and generate code named static.go
$ zipdata -s static.zip -g static.go -d .
```

## Sample generated code

```Go
package main
import (
 "archive/zip"
 "bytes"
 "fmt"
 "io"
 "os"
 "path/filepath"
)

func init() {
 err := UnzipBytes(zipdata, destDir)

 if err != nil {
  fmt.Println(err)
  return
 }
}

func main(){
 fmt.Println(destDir)
}
var destDir="."


var zipdata =[]byte{
  0x50, 0x4b, 0x03, 0x04, 0x14, 0x00, 0x00, 0x00, 0x08, 0x00, 0xdb, 0x82,
  0x0f, 0x51, 0x7d, 0xfd, 0xd6, 0x5c, 0x02, 0x06, 0x00, 0x00, 0x03, 0x10,
  0x00, 0x00, 0x0a, 0x00, 0x1c, 0x00, 0x7a, 0x69, 0x70, 0x64, 0x61, 0x74,
  0x61, 0x2e, 0x67, 0x6f, 0x55, 0x54, 0x09, 0x00, 0x03, 0x5d, 0x9b, 0x37,
  0x5f, 0x8f, 0x9b, 0x37, 0x5f, 0x75, 0x78, 0x0b, 0x00, 0x01, 0x04, 0xf5,
  0x01, 0x00, 0x00, 0x04, 0x14, 0x00, 0x00, 0x00, 0xa5, 0x57, 0x6d, 0x73,
  0xd3, 0x46, 0x10, 0xfe, 0x2c, 0xfd, 0x8a, 0xad, 0x3a, 0x69, 0xa4, 0x60,
  0xe4, 0xc0, 0xc7, 0x30, 0xfe, 0x00, 0x21, 0xa1, 0xe9, 0x40, 0xc2, 0x04,
  0x28, 0xd3, 0x02, 0x83, 0x2f, 0xf6, 0xc9, 0xbe, 0x89, 0x7c, 0xa7, 0x39,
  0x9d, 0x62, 0x4c, 0x9a, 0xff, 0xde, 0xdd, 0x5b, 0xbd, 0xf9, 0x25, 0xd0,
  0x4c, 0xf3, 0x21, 0x3a, 0xeb, 0xf6, 0xf6, 0xe5, 0xd9, 0xdd, 0xe7, 0x56,
  0x85, 0x98, 0x5c, 0x8b, 0x99, 0x84, 0x85, 0x50, 0x3a, 0x0c, 0xd5, 0xa2,
  0x30, 0xd6, 0x41, 0x1c, 0x06, 0x91, 0xb0, 0x93, 0xb9, 0xba, 0x91, 0xc3,
  0xef, 0xaa, 0x88, 0xf0, 0xe7, 0xd5, 0xca, 0xc9, 0x92, 0x16, 0xd9, 0xc2,
  0xd1, 0x43, 0x19, 0xfe, 0x3f, 0x54, 0xa6, 0x72, 0x2a, 0xa7, 0x1f, 0xb9,
  0x99, 0xd1, 0xc3, 0x78, 0xb1, 0x42, 0xb8, 0xf9, 0x30, 0x53, 0xb9, 0xa4,
  0x45, 0x14, 0xe2, 0x9b, 0x99, 0x72, 0xf3, 0xea, 0x2a, 0x9d, 0x98, 0xc5,
  0x50, 0xe4, 0x72, 0xe2, 0xe6, 0x66, 0x21, 0xca, 0xe1, 0xb5, 0xd2, 0xb3,
  0x42, 0xe9, 0x28, 0x4c, 0xc2, 0xf0, 0x46, 0x58, 0x32, 0x5d, 0xda, 0x09,
  0xe0, 0xdf, 0x08, 0xea, 0xbd, 0xf4, 0x34, 0x17, 0xb3, 0x38, 0xc2, 0xd7,
  0xd1, 0x00, 0x22, 0xa5, 0x8b, 0xca, 0xc1, 0x4b, 0x65, 0x51, 0x85, 0xb1,
  0x2b, 0x30, 0x16, 0xd0, 0x45, 0x20, 0x53, 0x51, 0x92, 0xbe, 0x9b, 0xa3,
  0xff, 0xf1, 0x7e, 0xb9, 0x9f, 0xa4, 0x27, 0xdf, 0x54, 0xe9, 0x50, 0xc3,
  0x29, 0xee, 0x5c, 0x58, 0x3c, 0x10, 0x27, 0x61, 0x30, 0x33, 0x24, 0xb8,
  0xa5, 0x9b, 0x5f, 0x93, 0xfa, 0x99, 0xd4, 0xd2, 0x0a, 0x27, 0xa7, 0x30,
  0x33, 0x5e, 0x29, 0x68, 0xb1, 0x20, 0xcd, 0x2f, 0x65, 0x26, 0xaa, 0xdc,
  0xa1, 0x1f, 0x4e, 0x38, 0x35, 0x49, 0x67, 0xa6, 0x33, 0x37, 0x43, 0x73,
  0xef, 0x9c, 0x45, 0x95, 0x64, 0x63, 0x2a, 0x4b, 0xb7, 0xc3, 0x7f, 0x7a,
  0x4d, 0x16, 0x10, 0xaf, 0xf5, 0x08, 0x96, 0x73, 0xa9, 0xa1, 0xd2, 0x04,
  0x74, 0xcf, 0x4c, 0xda, 0xa9, 0x9f, 0xf6, 0xd5, 0x23, 0x4e, 0x59, 0xa5,
  0x27, 0x3e, 0x61, 0x71, 0x02, 0xb7, 0x08, 0x6d, 0x63, 0xe7, 0xd8, 0x2c,
  0x16, 0x42, 0x4f, 0x5f, 0x2b, 0x2d, 0xd3, 0xdf, 0x65, 0x5e, 0xa0, 0x0b,
  0xe3, 0x57, 0x75, 0x3c, 0xf0, 0xca, 0xc0, 0xc4, 0x4c, 0x25, 0x64, 0xd6,
  0x2c, 0x40, 0xe8, 0x15, 0x07, 0x87, 0xe8, 0xb5, 0x8e, 0xa4, 0xe3, 0x4e,
  0xd5, 0x9f, 0xd2, 0x96, 0xca, 0xe8, 0x38, 0x3a, 0x4c, 0x9f, 0x44, 0x49,
  0xf7, 0xfe, 0xad, 0xb0, 0xa5, 0x44, 0x2f, 0xc2, 0x40, 0x65, 0x70, 0x40,
  0x99, 0x1a, 0x8d, 0x20, 0x8a, 0xd0, 0x8d, 0xa0, 0x95, 0xf9, 0x50, 0x62,
  0x41, 0x11, 0x10, 0x81, 0x95, 0xae, 0xb2, 0x3a, 0x0c, 0xee, 0x48, 0x5e,
  0x67, 0x66, 0x00, 0xd2, 0x5a, 0x38, 0x1a, 0x81, 0x29, 0xd3, 0xd7, 0x84,
  0x63, 0x4c, 0x2a, 0x12, 0xaf, 0x8c, 0x76, 0x7e, 0x19, 0x81, 0x56, 0xb9,
  0x8f, 0x29, 0xc0, 0x6a, 0x4a, 0xdf, 0x62, 0xcc, 0x2e, 0xd7, 0x31, 0xee,
  0x25, 0x3f, 0x35, 0x40, 0xd5, 0x33, 0x15, 0x4e, 0xc0, 0xa7, 0x2f, 0x54,
  0xac, 0x5e, 0x29, 0x19, 0x4d, 0xcf, 0x4a, 0x9f, 0x7c, 0xef, 0x23, 0x09,
  0xb0, 0x17, 0x23, 0xf8, 0x5b, 0x15, 0x8d, 0xfd, 0x2d, 0x07, 0x82, 0x9d,
  0x0e, 0xb4, 0xf6, 0xd0, 0x60, 0x70, 0x07, 0x32, 0x2f, 0xe5, 0x96, 0x56,
  0xee, 0x87, 0xf4, 0x52, 0x8a, 0x29, 0x55, 0xde, 0xff, 0xb2, 0x10, 0xb6,
  0xc8, 0x8c, 0x00, 0x0b, 0x33, 0x66, 0x3b, 0x07, 0x5c, 0xad, 0xb8, 0xa0,
  0x92, 0x4a, 0x9e, 0x6d, 0xea, 0xdd, 0xa1, 0xb6, 0x87, 0xd3, 0x5d, 0x5d,
  0x3f, 0x8d, 0xbe, 0x1a, 0xae, 0x01, 0x34, 0x5a, 0x49, 0x29, 0x02, 0x06,
  0xa5, 0xaf, 0xb8, 0x84, 0xb4, 0x63, 0x95, 0x50, 0x4e, 0x30, 0x32, 0x16,
  0xe9, 0xb2, 0x78, 0x6c, 0x25, 0x16, 0x57, 0xcc, 0x67, 0x77, 0x24, 0xb2,
  0x31, 0x4d, 0x6f, 0xc9, 0x3c, 0x36, 0x47, 0x26, 0x2d, 0xd4, 0x9a, 0xd2,
  0xe3, 0xdc, 0xd4, 0xe5, 0x84, 0xb4, 0x92, 0x9e, 0x16, 0xe4, 0x75, 0x16,
  0xb7, 0x76, 0xc6, 0x45, 0x9f, 0x9d, 0x1e, 0x4a, 0x4e, 0x3b, 0x89, 0xa8,
  0x69, 0x1f, 0xa5, 0x95, 0xe3, 0x9a, 0xa8, 0x83, 0xf9, 0x40, 0x0d, 0xf8,
  0x82, 0x14, 0xc5, 0xb8, 0x60, 0xa4, 0x6b, 0x28, 0x92, 0x70, 0x57, 0x60,
  0xe4, 0xf1, 0x7d, 0x30, 0x07, 0x2d, 0xcc, 0xdc, 0xa6, 0xb7, 0xe1, 0x9a,
  0x78, 0xab, 0xf7, 0xce, 0x73, 0x5e, 0xfd, 0x73, 0x14, 0xed, 0xdd, 0x10,
  0x53, 0x8e, 0xd7, 0x0c, 0x37, 0x50, 0x7d, 0xb4, 0xca, 0xc9, 0x9a, 0x05,
  0xc6, 0xcc, 0x95, 0xb5, 0x9f, 0x30, 0xe2, 0x14, 0xde, 0x8e, 0x49, 0xbe,
  0x5c, 0x52, 0x34, 0xbf, 0x51, 0x24, 0xfe, 0x88, 0xbd, 0xe5, 0xc7, 0x51,
  0x03, 0xfa, 0x1d, 0xc9, 0xb0, 0x3a, 0x5f, 0x00, 0xc9, 0x6e, 0x1b, 0xd1,
  0x5d, 0x44, 0xea, 0x86, 0x43, 0xd8, 0xb5, 0x5b, 0xde, 0x73, 0x0a, 0x3d,
  0xe3, 0xb8, 0x3d, 0x9f, 0x11, 0x92, 0xd4, 0x08, 0x98, 0xf1, 0x03, 0x5c,
  0xa6, 0xbc, 0xbe, 0xbf, 0xc4, 0x82, 0x0c, 0x1f, 0x5f, 0x07, 0x90, 0x51,
  0x08, 0x56, 0x68, 0xcc, 0x7c, 0xab, 0x21, 0xa5, 0x6e, 0x62, 0xe0, 0x29,
  0x93, 0x24, 0xd1, 0x64, 0x35, 0xfd, 0xc3, 0xa8, 0x16, 0x54, 0x3c, 0x9d,
  0x9e, 0x23, 0x65, 0xd7, 0x3d, 0x97, 0xf9, 0x73, 0x67, 0xc8, 0x03, 0x71,
  0xb2, 0xc6, 0x04, 0x01, 0x56, 0xef, 0x9b, 0xeb, 0xa9, 0xb2, 0xcf, 0xf3,
  0x3c, 0xf6, 0x2a, 0x07, 0x54, 0xd0, 0x6f, 0x90, 0x23, 0xdf, 0x4a, 0xbb,
  0xa0, 0xe3, 0xbd, 0x0e, 0x6f, 0xf2, 0xcf, 0x55, 0xdf, 0x9d, 0x6b, 0x3c,
  0x20, 0xbd, 0x5e, 0x49, 0xb2, 0xae, 0x65, 0xab, 0x39, 0xd7, 0x1b, 0x22,
  0xf0, 0x3d, 0x8e, 0xda, 0xf5, 0x69, 0xbf, 0xb1, 0xb2, 0xf4, 0xa2, 0xc0,
  0xf6, 0x4c, 0x7a, 0x86, 0x7f, 0xa8, 0x02, 0xff, 0x71, 0x5f, 0xb1, 0x9e,
  0x5e, 0x5b, 0x05, 0x94, 0xa5, 0xd3, 0x8d, 0xa6, 0x25, 0xe5, 0x9e, 0x9b,
  0xba, 0xb0, 0x2f, 0xbe, 0x7e, 0xbc, 0xbc, 0x38, 0x7f, 0xfd, 0xd7, 0x3f,
  0x7e, 0x7d, 0x7c, 0x79, 0xf2, 0xfc, 0xfd, 0x09, 0xaf, 0xdf, 0x5f, 0x7e,
  0x38, 0x3f, 0x26, 0x50, 0x29, 0xa6, 0x38, 0x79, 0xb8, 0x4f, 0xb5, 0x03,
  0x6b, 0x4e, 0x7d, 0xed, 0xd8, 0x12, 0x6f, 0xad, 0x62, 0x15, 0xb7, 0x5e,
  0x72, 0x04, 0x0f, 0xb0, 0x52, 0xf3, 0x64, 0xbd, 0x81, 0x92, 0xd4, 0x7a,
  0xc3, 0x61, 0xd7, 0xce, 0x5c, 0x8a, 0x7c, 0xef, 0x51, 0x9b, 0x7c, 0xfa,
  0xe2, 0xaf, 0x09, 0x2e, 0xd3, 0x5e, 0xd7, 0xaf, 0x91, 0xe1, 0xfd, 0x24,
  0xe8, 0x41, 0xf4, 0x84, 0x93, 0x9e, 0xcb, 0x25, 0xd7, 0x66, 0xdd, 0x47,
  0x61, 0xd0, 0x56, 0x6b, 0x0b, 0x37, 0x95, 0x7d, 0x27, 0x87, 0xef, 0x6d,
  0xfa, 0x4e, 0x7d, 0xf7, 0x40, 0xfe, 0x9c, 0x2d, 0xdb, 0xb0, 0x36, 0xba,
  0xa9, 0x47, 0x10, 0x18, 0xbe, 0x6f, 0xfb, 0xf5, 0xf8, 0xeb, 0x49, 0x4a,
  0xcb, 0x65, 0x8e, 0xc3, 0x00, 0x4f, 0x53, 0x35, 0x49, 0xec, 0x7f, 0xd6,
  0xfb, 0xc4, 0xc3, 0xe8, 0xb1, 0xd2, 0x53, 0xa9, 0x5d, 0x7f, 0xcb, 0xed,
  0x0f, 0x80, 0xfe, 0x13, 0x47, 0x20, 0xf7, 0xfa, 0x93, 0xfd, 0xb3, 0x80,
  0x3b, 0x09, 0xc1, 0xdb, 0xd1, 0x0b, 0xd0, 0x12, 0xea, 0x35, 0xc2, 0x55,
  0x4d, 0x5c, 0xe8, 0x56, 0x85, 0x84, 0x9e, 0x08, 0xbf, 0xa6, 0x00, 0x31,
  0xdf, 0xfc, 0x2e, 0x0c, 0x88, 0x83, 0x5d, 0xcb, 0x93, 0xf1, 0x12, 0x0e,
  0xba, 0x13, 0x09, 0x2b, 0x8c, 0x8b, 0xda, 0x72, 0x02, 0xb1, 0x26, 0x71,
  0x86, 0xd5, 0x27, 0xc3, 0x77, 0x31, 0x02, 0x98, 0x63, 0xa3, 0x14, 0x09,
  0x8d, 0x21, 0x87, 0x3d, 0x04, 0x19, 0x3d, 0xa2, 0x13, 0x0d, 0x0d, 0x95,
  0x14, 0x7e, 0x1f, 0x8f, 0x2c, 0xd3, 0xc9, 0xde, 0x93, 0xa7, 0xdd, 0x91,
  0xa0, 0xa6, 0x43, 0x5b, 0xb3, 0x62, 0x0d, 0x5b, 0xb2, 0x63, 0xab, 0x83,
  0xad, 0xde, 0xc5, 0x01, 0x08, 0x0e, 0x37, 0xb8, 0x62, 0xe3, 0x8c, 0x47,
  0xd2, 0xf3, 0x49, 0x58, 0x5f, 0x1b, 0xcd, 0x45, 0xd7, 0x08, 0xe2, 0x54,
  0x78, 0xf8, 0x6d, 0xef, 0xf0, 0xe9, 0xb7, 0x01, 0x0e, 0x88, 0xc5, 0x27,
  0xfd, 0x85, 0xa4, 0x51, 0xf7, 0xa3, 0x47, 0xfe, 0xce, 0xd4, 0xf4, 0xac,
  0xe3, 0xf2, 0xc5, 0x4d, 0xa3, 0x0b, 0xc0, 0xb4, 0x9d, 0x20, 0xf1, 0x59,
  0xe1, 0xbc, 0x76, 0x23, 0xf3, 0x15, 0x8e, 0x76, 0x53, 0xa8, 0x8b, 0xc1,
  0x97, 0x29, 0x08, 0x6b, 0xc5, 0x8a, 0x41, 0xa6, 0x89, 0x87, 0xe7, 0xe7,
  0xa6, 0xae, 0xd7, 0xab, 0x7e, 0x0d, 0xdc, 0x30, 0xb8, 0xaa, 0x3c, 0x11,
  0x23, 0x1c, 0x31, 0x17, 0xfc, 0x8b, 0x2a, 0xc3, 0x8e, 0xa6, 0x72, 0x5b,
  0xf6, 0x6a, 0x9b, 0x63, 0x88, 0x51, 0x9a, 0x2f, 0x0c, 0x6e, 0xfb, 0x65,
  0xaf, 0xe1, 0x97, 0x22, 0xbf, 0x96, 0xcc, 0x6c, 0xe8, 0x46, 0xec, 0xf9,
  0x9b, 0x3d, 0x18, 0xf8, 0x11, 0x8d, 0x18, 0xa8, 0xa1, 0xe9, 0x35, 0x2f,
  0xda, 0x8b, 0x81, 0xee, 0xa1, 0xf6, 0x06, 0xcd, 0xe2, 0xe8, 0xd8, 0x0a,
  0xca, 0xd1, 0xec, 0x08, 0xf6, 0x7e, 0xbd, 0xf9, 0xac, 0x09, 0x36, 0x22,
  0xdf, 0xdd, 0x73, 0xd6, 0x1a, 0x69, 0xdc, 0xb1, 0xcc, 0xd6, 0x64, 0xd8,
  0xef, 0x20, 0x96, 0x42, 0x93, 0xdd, 0x60, 0x77, 0xb4, 0x3d, 0xd9, 0x35,
  0x16, 0x37, 0x67, 0x22, 0xcf, 0xdd, 0x0f, 0x71, 0x87, 0x01, 0xdb, 0x18,
  0x88, 0xc8, 0xfc, 0x89, 0x2e, 0x2b, 0x2b, 0xc1, 0xcd, 0x85, 0xa3, 0x59,
  0xc8, 0xcd, 0xc7, 0xa0, 0x4a, 0xd0, 0xc6, 0x81, 0xb8, 0x2a, 0x4d, 0x5e,
  0x39, 0xf9, 0x0c, 0x14, 0x26, 0x73, 0x6e, 0xaa, 0x7c, 0xea, 0xdf, 0xe3,
  0x50, 0x8d, 0x23, 0xd2, 0x12, 0x3f, 0xb5, 0x20, 0x1a, 0x46, 0x29, 0xab,
  0x79, 0x3f, 0xc7, 0x53, 0xa5, 0x56, 0x45, 0x21, 0x1d, 0xcc, 0x05, 0x3e,
  0x74, 0x09, 0xce, 0xc0, 0xd2, 0xd8, 0x6b, 0xb8, 0x92, 0x13, 0x51, 0x61,
  0xe5, 0x9e, 0xc1, 0xd4, 0xe8, 0x7d, 0x07, 0xb8, 0xe6, 0x53, 0x8d, 0x09,
  0x8f, 0x6c, 0x39, 0x80, 0x2b, 0x2c, 0x1b, 0xc9, 0x0e, 0xad, 0x4c, 0x65,
  0xb1, 0xc8, 0x44, 0xfe, 0x18, 0x55, 0xa0, 0x65, 0xfa, 0xb0, 0xe0, 0x43,
  0x0e, 0x5b, 0xad, 0xc4, 0xb6, 0x5b, 0x94, 0xfe, 0x18, 0xf5, 0xac, 0x01,
  0x41, 0x95, 0xf2, 0xd8, 0x1a, 0xf4, 0xcf, 0xca, 0x1c, 0xbf, 0x9e, 0x6e,
  0x58, 0x29, 0xb9, 0x97, 0xb5, 0xc8, 0x2d, 0x9b, 0x61, 0xf2, 0x21, 0xd0,
  0x6d, 0x5d, 0x24, 0xa8, 0x8f, 0x47, 0xd1, 0xff, 0xa8, 0xa0, 0x9f, 0x75,
  0xea, 0x4f, 0xd6, 0xd6, 0xde, 0xeb, 0x1f, 0xb1, 0x72, 0xb9, 0x65, 0x06,
  0xc0, 0x55, 0xfc, 0xa3, 0x21, 0x17, 0x7f, 0x0e, 0x3a, 0xee, 0x1e, 0x0e,
  0xeb, 0xc4, 0x4c, 0x28, 0xa9, 0x94, 0x28, 0xfc, 0x3a, 0xab, 0x44, 0x9e,
  0xaf, 0x06, 0xb4, 0x39, 0xf1, 0x7b, 0x3e, 0x6b, 0x84, 0xff, 0x66, 0xd7,
  0xa0, 0x04, 0x96, 0xfe, 0x0c, 0x33, 0xb6, 0xb4, 0x46, 0xcf, 0x08, 0xc2,
  0xde, 0xd7, 0x40, 0x2b, 0xb7, 0x35, 0x60, 0xec, 0x74, 0x85, 0xc7, 0x44,
  0x4c, 0x61, 0x96, 0xf2, 0x65, 0x97, 0xf4, 0x08, 0xe5, 0x5f, 0x50, 0x4b,
  0x01, 0x02, 0x1e, 0x03, 0x14, 0x00, 0x00, 0x00, 0x08, 0x00, 0xdb, 0x82,
  0x0f, 0x51, 0x7d, 0xfd, 0xd6, 0x5c, 0x02, 0x06, 0x00, 0x00, 0x03, 0x10,
  0x00, 0x00, 0x0a, 0x00, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00,
  0x00, 0x00, 0xa4, 0x81, 0x00, 0x00, 0x00, 0x00, 0x7a, 0x69, 0x70, 0x64,
  0x61, 0x74, 0x61, 0x2e, 0x67, 0x6f, 0x55, 0x54, 0x05, 0x00, 0x03, 0x5d,
  0x9b, 0x37, 0x5f, 0x75, 0x78, 0x0b, 0x00, 0x01, 0x04, 0xf5, 0x01, 0x00,
  0x00, 0x04, 0x14, 0x00, 0x00, 0x00, 0x50, 0x4b, 0x05, 0x06, 0x00, 0x00,
  0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x50, 0x00, 0x00, 0x00, 0x46, 0x06,
  0x00, 0x00, 0x00, 0x00,}


func unzip(zipReader *zip.Reader, destDir string) error {
 for _, f := range zipReader.File {
  fpath := filepath.Join(destDir, f.Name)
  if f.FileInfo().IsDir() {
   os.MkdirAll(fpath, os.ModePerm)
  } else {
   if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
    return err
   }

   inFile, err := f.Open()
   if err != nil {
    return err
   }
   defer inFile.Close()

   outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
   if err != nil {
    return err
   }
   defer outFile.Close()

   _, err = io.Copy(outFile, inFile)
   if err != nil {
    return err
   }
  }
 }

 return nil
}

//UnzipBytes unzip from byte[] data
func UnzipBytes(data []byte, destDir string) error {

 r := bytes.NewReader(data)

 zipReader, err := zip.NewReader(r, r.Size())

 if err != nil {
  return err
 }

 return unzip(zipReader, destDir)
}
```
