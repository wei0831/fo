package utli

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

type transaction struct {
	action, oldBase, oldFile, newBase, newFile string
}

type actiontaken struct {
	Action    string `json:"action"`
	From      string `json:"from"`
	To        string `json:"to"`
	CreatedAt int64  `json:"createdAt"`
}

func (t *transaction) commit(wet bool) string {
	switch t.action {
	case "MV":
		oldPath := path.Join(t.oldBase, t.oldFile)
		newPath := path.Join(t.newBase, t.newFile)
		newDir, newFile := filepath.Split(newPath)
		newFileExt := filepath.Ext(newFile)
		newFileName := newFile[0 : len(newFile)-len(newFileExt)]

		// Check duplication
		target := path.Join(t.newBase, t.newFile)
		i := 1
		for {
			_, err := os.Stat(target)
			if err != nil {
				break
			}
			attemptFileName := fmt.Sprintf("%s-Copy(%d)%s", newFileName, i, newFileExt)
			target = path.Join(newDir, attemptFileName)
			i++
		}
		dupOrNot := ""
		if i > 1 {
			dupOrNot = "[DUP]"
		}

		// Commit changes
		if wet {
			// Make new folder ?
			if err := os.MkdirAll(newDir, 0755); err != nil {
				return fmt.Sprintf("[ERR] %s\n", err)
			}
			// Move
			if err := os.Rename(oldPath, target); err != nil {
				return fmt.Sprintf("[ERR] %s\n", err)
			}
			// Log
			out, err := json.Marshal(&actiontaken{"MV", string(oldPath), string(target), time.Now().UnixNano()})
			if err != nil {
				panic(err)
			}
			log.info(string(out))
		}

		t.oldFile = oldPath[len(t.oldBase):]
		t.newFile = target[len(t.newBase):]
		return fmt.Sprintf("%s[MV] \"%s\" -> \"%s\"\n", dupOrNot, t.oldFile, t.newFile)

	case "RMDIR":
		removePath := path.Join(t.oldBase, t.oldFile)
		if wet {
			if err := os.Remove(removePath); err != nil {
				return fmt.Sprintf("[ERR] %s\n", err)
			}
			// Log
			out, err := json.Marshal(&actiontaken{"RMDIR", removePath, "", time.Now().UnixNano()})
			if err != nil {
				panic(err)
			}
			log.info(string(out))
		}
		return fmt.Sprintf("[RMDIR] \"%s\"\n", removePath)
	case "MKDIR":
		folderPath := path.Join(t.oldBase, t.oldFile)
		if wet {
			if err := os.MkdirAll(folderPath, 0755); err != nil {
				return fmt.Sprintf("[ERR] %s\n", err)
			}
			// Log
			out, err := json.Marshal(&actiontaken{"MKDIR", folderPath, "", time.Now().UnixNano()})
			if err != nil {
				panic(err)
			}
			log.info(string(out))
		}
		return fmt.Sprintf("[MKDIR] \"%s\"\n", folderPath)
	}

	return fmt.Sprintf("[ERR] invliad action: %s\n", t.action)
}
