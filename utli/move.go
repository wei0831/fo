package utli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// Movematches TODO
func Movematches(dir, to, find, exclude string, mode int, wet bool) {
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
	for _, item := range workDir {

		isDir := item.IsDir()
		fileName := item.Name()

		// Handle File
		if mode == FileAndFolder ||
			(!isDir && mode == FileOnly) ||
			(isDir && mode == FolderOnly) {

			// If match and not excluded
			if matchName.MatchString(fileName) &&
				((exclude == "") || (exclude != "" && !excludeName.MatchString(fileName))) {
				toDo = append(toDo, transaction{
					"MV",
					filepath.ToSlash(dirBase),
					filepath.ToSlash(fileName),
					filepath.ToSlash(toBase),
					filepath.ToSlash(fileName)})
			}
		}
	}

	// Start doing work
	startMsg := fmt.Sprintf("[CMD][Mode%d] %s Find \"%s\" From \"%s\" -> To \"%s\"\n", mode, "move", find, filepath.ToSlash(dir), filepath.ToSlash(to))
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
