package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/alecthomas/kingpin"
)

var (
	src    = kingpin.Flag("src", "input Directory or zip file").Short('s').ExistingFileOrDir()
	gofile = kingpin.Flag("gofile", "generated go file name").Default("static.go").Short('g').String()
	dest   = kingpin.Flag("dest", "output Directory when unzip").Default(".").Short('d').String()
)

func main() {

	kingpin.CommandLine.Help = `Generate Go code from any file or Directory.`
	kingpin.Version("0.1.0")
	kingpin.Parse()

	if *src == "" {
		kingpin.Usage()
		return
	}

	info, err := os.Lstat(*src)
	if err != nil {

		log.Println(err)
		kingpin.Usage()
		return
	}

	var data []byte
	if info.IsDir() {
		data, err = Zip(*src)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		data, err = ioutil.ReadFile(*src)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if err = gen(data, *gofile, *dest); err != nil {
		log.Println(err)
		return
	}

}

func gen(data []byte, gofile, destDir string) error {

	outfile, err := os.Create(gofile)
	if err != nil {
		return err
	}
	defer outfile.Close()

	fmt.Fprintf(outfile, `package main
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
var destDir="%v"

	`, destDir)

	outfile.WriteString(`

var zipdata =[]byte{`)

	sw := &ByteWriter{Writer: outfile}
	sw.Write(data)
	outfile.WriteString("}")

	// outfile.WriteString(s)
	outfile.WriteString(`


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
	`)

	return nil
}

var (
	newline    = []byte{'\n'}
	dataindent = []byte{'\t', '\t'}
	space      = []byte{' '}
)

//ByteWriter Byte Writer struct
type ByteWriter struct {
	io.Writer
	c int
}

func (w *ByteWriter) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}

	for n = range p {
		if w.c%12 == 0 {
			w.Writer.Write(newline)
			w.Writer.Write(dataindent)
			w.c = 0
		} else {
			w.Writer.Write(space)
		}

		fmt.Fprintf(w.Writer, "0x%02x,", p[n])
		w.c++
	}
	n++
	return
}

// Zip  directory recursively and return  byte array
func Zip(input string) (data []byte, err error) {

	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	//defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		// fmt.Printf("Crawling: %#v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		// data, err := ioutil.ReadFile(path)
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		f, err := w.Create(path)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, file)

		if err != nil {
			return err
		}
		return nil
	}

	err = filepath.Walk(input, walker)
	if err != nil {
		return nil, err
	}

	//should close it manually,
	//could  not use defer w.Close()
	//or get wrong zip
	if err = w.Close(); err != nil {
		return nil, err
	}

	data = buf.Bytes()
	return
}
