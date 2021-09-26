package findgo

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func Files(path string) (count int) {
	fsys := os.DirFS(path)
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
