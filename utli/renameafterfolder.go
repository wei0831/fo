package utli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

// RenameAfterFolder TODO
func RenameAfterFolder(dir, find, exclude string, wet bool) {
	// Get absolute path
	dirBase, _ := filepath.Abs(dir)

	// Check dir
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Start checking files in the dir
	workDir, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}

	// Organize todo
	matchName := regexp.MustCompile(find)
	excludeName := regexp.MustCompile(exclude)

	var toDo []transaction
	for _, parentItem := range workDir {
		if parentItem.IsDir() {
			folderName := parentItem.Name()
			// If match and not excluded
			if matchName.MatchString(folderName) &&
				((exclude == "") || (exclude != "" && !excludeName.MatchString(folderName))) {
				todoPath := path.Join(dirBase, folderName)
				childDir, err := ioutil.ReadDir(todoPath)
				if err != nil {
					fmt.Printf("[ERR] %s\n", err)
					os.Exit(1)
				}
				for _, childItem := range childDir {
					fileExt := filepath.Ext(childItem.Name())
					// Take care only files
					if !childItem.IsDir() {
						toDo = append(toDo, transaction{
							"MV",
							filepath.ToSlash(""),
							filepath.ToSlash(path.Join(todoPath, childItem.Name())),
							filepath.ToSlash(""),
							filepath.ToSlash(path.Join(todoPath, folderName+fileExt))})
					}
				}
			}
		}
	}

	// Start doing work
	startMsg := fmt.Sprintf("[CMD] %s Find \"%s\" In \"%s\"\n", "renameafterfolder", find, filepath.ToSlash(dir))
	if wet {
		startMsg = "[WET]" + startMsg
		log.info("#" + startMsg)
	} else {
		startMsg = "[DRY]" + startMsg
	}
	fmt.Printf(startMsg)

	// Commit transactions
	for _, i := range toDo {
		fmt.Print(i.commit(wet))
	}
}
