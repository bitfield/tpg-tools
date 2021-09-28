package findgo

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func Files(fsys fs.FS) (count int) {
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".go" {
			count++
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func FilesFromArgs(args []string) (count int) {
	if len(args) < 1 {
		return 0
	}
	for _, arg := range args {
		fsys := os.DirFS(arg)
		count += Files(fsys)
	}
	return count
}
