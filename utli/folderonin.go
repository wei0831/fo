package utli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// Folderin TODO
func Folderin(dir string, to string, wet bool) {
	// Not provide? Same as dir
	if to == "" {
		to = dir
	}

	// Get absolute path
	dirBase, _ := filepath.Abs(dir)
	toBase, _ := filepath.Abs(to)

	// Check dir
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(2)
	}

	// Check toDir
	if _, err := os.Stat(to); os.IsNotExist(err) {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(2)
	}

	// Start checking files in the dir
	workDir, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(2)
	}

	// Organize todo
	var toDo []transaction
	for _, item := range workDir {
		if !item.IsDir() {
			fileName := item.Name()
			oldPath := fileName
			fileExt := filepath.Ext(fileName)
			fileNameOnly := fileName[0 : len(fileName)-len(fileExt)]
			newPath := path.Join(fileNameOnly, fileName)
			toDo = append(toDo, transaction{
				"MKDIR",
				filepath.ToSlash(toBase),
				filepath.ToSlash(fileNameOnly),
				"",
				""})
			toDo = append(toDo, transaction{
				"MV",
				filepath.ToSlash(dirBase),
				filepath.ToSlash(oldPath),
				filepath.ToSlash(toBase),
				filepath.ToSlash(newPath)})
		}
	}
	// Start doing work
	wetOrNot := ""
	if wet {
		wetOrNot = "[WET]"
	} else {
		wetOrNot = "[DRY]"
	}
	log.init()
	startMsg := fmt.Sprintf("%s[CMD] %s From \"%s\" -> To \"%s\"\n", wetOrNot, "folderin", filepath.ToSlash(dir), filepath.ToSlash(to))
	log.info("#" + startMsg)
	fmt.Printf(startMsg)

	// Commit transactions
	for _, i := range toDo {
		msg := i.commit(wet)
		fmt.Print(msg)
	}
}
