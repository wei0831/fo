package utli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// Folderout check given "dir" directory and empty out the child folders
// move files in the child folders to "to" directory
// then delete those empty folders
func Folderout(dir string, to string, wet bool) {
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
		os.Exit(1)
	}

	// Check toDir
	if _, err := os.Stat(to); os.IsNotExist(err) {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Start checking files in the dir
	parentDir, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Organize todo
	var toDo []transaction
	for _, parentItem := range parentDir {
		if parentItem.IsDir() {
			todoPath := path.Join(dir, parentItem.Name())
			childDir, err := ioutil.ReadDir(todoPath)
			if err != nil {
				fmt.Printf("[ERR] %s\n", err)
				os.Exit(1)
			}
			for _, childItem := range childDir {
				oldPath := path.Join(parentItem.Name(), childItem.Name())
				newPath := childItem.Name()
				toDo = append(toDo, transaction{
					"MV",
					filepath.ToSlash(dirBase),
					filepath.ToSlash(oldPath),
					filepath.ToSlash(toBase),
					filepath.ToSlash(newPath)})
			}
			toDo = append(toDo, transaction{"RMDIR",
				filepath.ToSlash(dirBase),
				filepath.ToSlash(parentItem.Name()),
				"",
				""})
		}
	}

	// Start doing work
	startMsg := fmt.Sprintf("[CMD] %s From \"%s\" -> To \"%s\"\n", "folderout", filepath.ToSlash(dir), filepath.ToSlash(to))
	if wet {
		startMsg = "[WET]" + startMsg
		log.info("#" + startMsg)
	}
	fmt.Printf(startMsg)

	// Commit transactions
	for _, i := range toDo {
		fmt.Print(i.commit(wet))
	}
}
