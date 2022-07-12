package shutil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Copy(src, dst string) error {
	srcStat, err := os.Stat(src); if err != nil { return err }

	// src is a file, just copy it
	if !srcStat.IsDir() {
		// if dst is a directory, change dst to include target file name
		if strings.HasSuffix(dst, "/") {
			dst = filepath.Join(dst, filepath.Base(src))
		}
		srcFile, err := os.Open(src); if err != nil { return err }
		dstFile, err := os.Create(dst); if err != nil { return err }
		if io.Copy(dstFile, srcFile) ; err != nil { return err }
	} else {
		os.Mkdir(filepath.Join(dst, filepath.Base(src)), 0755)
		files, err:= os.ReadDir(src); if err != nil { return err }
		for _, file := range files{
			fmt.Println(file.Name())
			// if file.IsDir() {
			// 	Copy(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
			// }
			Copy(filepath.Join(src, file.Name()), filepath.Join(dst, filepath.Join(filepath.Base(src), file.Name())))
		}
	}
	return nil
}
