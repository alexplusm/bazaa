package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"archive/zip"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

const (
	mediaTempDir = "media_temp"
	mediaRoot    = "media_root"
)

// UploadFiles controller
func UploadFiles(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		fmt.Println("file:", file.Filename, file.Size)

		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		createDir(mediaTempDir)

		// update file name with timestamp
		ts := time.Now().Unix()
		filename := addTimestampToFilename(filepath.Base(file.Filename), ts)
		fmt.Println(filename)

		fp := filepath.Join(mediaTempDir, filename)

		// Destination
		dst, err := os.Create(fp)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err := io.Copy(dst, src); err != nil {
			return err
		}
	}

	// ? request to MessageQueue: unzip archive and fill db
	unzipFiles()

	return c.String(http.StatusOK, "OK")
}

// todo: move to fileutils
func createDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}

// todo: move ?
// add timestamp to file
func addTimestampToFilename(filename string, unixTS int64) string {
	fileParts := strings.Split(filename, ".")
	length := len(fileParts)
	expansion := fileParts[length-1]

	filePartsNew := make([]string, 0, length+1)

	filePartsNew = append(filePartsNew, fileParts[:length-1]...)
	filePartsNew = append(filePartsNew, fmt.Sprint(unixTS))
	filePartsNew = append(filePartsNew, expansion)

	return strings.Join(filePartsNew[:len(filePartsNew)-1], "-") + "." + expansion
}

func unzipFiles() error {
	dir, err := os.Open(mediaTempDir)
	if err != nil {
		return err
	}
	filesInfo, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	createDir(mediaRoot)

	for _, fileInfo := range filesInfo {
		fmt.Println("to unzip", fileInfo.Name(), fileInfo.Size())

		Unzip(filepath.Join(mediaTempDir, fileInfo.Name()), mediaRoot)
	}

	return nil
}

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

// Unzip will decompress a zip archived file,
// copying all files and folders
// within the zip file (parameter 1)
// to an output directory (parameter 2).
func Unzip(src string, destination string) ([]string, error) {

	// a variable that will store any
	//file names available in a array of strings
	var filenames []string

	// OpenReader will open the Zip file
	// specified by name and return a ReadCloser
	// Readcloser closes the Zip file,
	// rendering it unusable for I/O
	// It returns two values:
	// 1. a pointer value to ReadCloser
	// 2. an error message (if any)
	r, err := zip.OpenReader(src)

	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		fmt.Println("FileInfo().Name()", f.FileInfo().Name(), "f.Name", f.Name)

		// this loop will run until there are
		// files in the source directory & will
		// keep storing the filenames and then
		// extracts into destination folder until an err arises

		// Store "path/filename" for returning and using later on
		fpath := filepath.Join(destination, f.Name)

		id, err := uuid.NewRandom()
		if err != nil {
			return filenames, err
		}
		fmt.Println("id", id)

		// Checking for any invalid file paths
		if !strings.HasPrefix(fpath, filepath.Clean(destination)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s is an illegal filepath", fpath)
		}

		// the filename that is accessed is now appended
		// into the filenames string array with its path
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Creating a new Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Creating the files in the target directory
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		// The created file will be stored in
		// outFile with permissions to write &/or truncate
		outFile, err := os.OpenFile(fpath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			f.Mode())

		// again if there is any error this block
		// will be executed and process
		// will return to main function
		if err != nil {
			// with filenames gathered so far
			// and err message
			return filenames, err
		}

		rc, err := f.Open()

		// again if there is any error this block
		// will be executed and process
		// will return to main function
		if err != nil {
			// with filenames gathered so far
			// and err message back to main function
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer so that
		// it closes the outfile before the loop
		// moves to the next iteration. this kinda
		// saves an iteration of memory & time in
		// the worst case scenario.
		outFile.Close()
		rc.Close()

		// again if there is any error this block
		// will be executed and process
		// will return to main function
		if err != nil {
			// with filenames gathered so far
			// and err message back to main function
			return filenames, err
		}
	}

	// Finally after every file has been appended
	// into the filenames string[] and all the
	// files have been extracted into the
	// target directory, we return filenames
	// and nil as error value as the process executed
	// successfully without any errors*
	// *only if it reaches until here.
	return filenames, nil
}
