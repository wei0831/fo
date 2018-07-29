package utli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// ReplaceName TODO
func ReplaceName(dir, to, find, replace, exclude string, mode int, wet bool) {
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
	parentDir, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(2)
	}

	// Organize todo
	matchName := regexp.MustCompile(find)
	excludeName := regexp.MustCompile(exclude)

	var toDo []transaction
	for _, item := range parentDir {

		isDir := item.IsDir()
		oldFileName := item.Name()

		// Handle File
		if mode == FileAndFolder ||
			(!isDir && mode == FileOnly) ||
			(isDir && mode == FolderOnly) {

			// If match and not excluded
			if matchName.MatchString(oldFileName) &&
				((exclude == "") || (exclude != "" && !excludeName.MatchString(oldFileName))) {
				newFileName := matchName.ReplaceAllString(oldFileName, replace)
				toDo = append(toDo, transaction{
					"MV",
					filepath.ToSlash(dirBase),
					filepath.ToSlash(oldFileName),
					filepath.ToSlash(toBase),
					filepath.ToSlash(newFileName)})
			}

		}
	}

	// Start doing work
	startMsg := fmt.Sprintf("[CMD][Mode%d] %s Find \"%s\" Replace with \"%s\" From \"%s\" -> To \"%s\"\n", mode, "replacename", find, replace, filepath.ToSlash(dir), filepath.ToSlash(to))
	if wet {
		startMsg = "[WET]" + startMsg
		log.init()
		log.info("#" + startMsg)
	}
	fmt.Printf(startMsg)

	// Commit transactions
	for _, i := range toDo {
		fmt.Print(i.commit(wet))
	}
}
