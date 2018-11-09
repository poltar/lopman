package fileactions

//Handle extraction, installation, and logging of fetched archives

import (
	"path/filepath"

	_ "golang.org/x/build/tarutil"

	_ "github.com/h2non/filetype"
)

//Extract archive and return true if success, false if failure
func extract(path string) bool {
	abspath, er := filepath.Abs(path)

	_ = abspath

	if er != nil {
		return false
	}

	return true
}

//Install installs the archive by extracting, moving to install directory, and adding to database, return true if success, false if failure
func Install(path string) bool {
	return true
}

//Add extracted archive to database, return true if success, false if failure
func addToDatabase(path string) bool {
	return true
}

//Move extracted archive to install directory, return true if success, false if failure
func move(installpath string) bool {
	return true
}
