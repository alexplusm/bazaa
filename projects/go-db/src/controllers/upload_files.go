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
	"strconv"

	"github.com/Alexplusm/bazaa/projects/go-db/src/configs"
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

		createDir(configs.MediaTempDir)

		// update file name with timestamp
		ts := time.Now().Unix()
		filename := addTimestampToFilename(filepath.Base(file.Filename), ts)
		fmt.Println(filename)

		fp := filepath.Join(configs.MediaTempDir, filename)

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

// todo: refactor
func getTimestampFromFilename(filename string) int64 {
	fileParts := strings.Split(filename, ".")
	filePartsWithoutExpansion := fileParts[len(fileParts)-2]
	filePartsWithoutExpansion2 := strings.Split(filePartsWithoutExpansion, "-")
	ts := filePartsWithoutExpansion2[len(filePartsWithoutExpansion2)-1]

	tsInt, err := strconv.ParseInt(ts, 10, 0)
	// todo: proccess error
	if err != nil {
		return -1
	}

	return tsInt
}

func unzipFiles() error {
	dir, err := os.Open(configs.MediaTempDir)
	if err != nil {
		return err
	}
	filesInfo, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	createDir(configs.MediaRoot)

	for _, fileInfo := range filesInfo {
		fmt.Println("to unzip", fileInfo.Name(), fileInfo.Size())

		Unzip(filepath.Join(configs.MediaTempDir, fileInfo.Name()), configs.MediaRoot)
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
	reader, err := zip.OpenReader(src)

	if err != nil {
		return filenames, err
	}
	defer reader.Close()

	timestamp := getTimestampFromFilename(src)
	rootFolder := filepath.Join(destination, fmt.Sprint(timestamp))
	os.MkdirAll(rootFolder, os.ModePerm)

	for _, f := range reader.File {
		// this loop will run until there are
		// files in the source directory & will
		// keep storing the filenames and then
		// extracts into destination folder until an err arises

		fname := f.FileInfo().Name()
		fpath := filepath.Join(rootFolder, f.Name)

		if isInvalidFileName(fname) {
			continue
		}

		fmt.Println("fname", fname)

		// TODO: Checking for any invalid file paths
		if !strings.HasPrefix(fpath, filepath.Clean(destination)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s is an illegal filepath", fpath)
		}

		// the filename that is accessed is now appended
		// into the filenames string array with its path
		filenames = append(filenames, fpath)

		// for nested dirs
		if f.FileInfo().IsDir() {
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

func isDSStoreFile(name string) bool {
	return name == ".DS_Store"
}

func isInvalidFileName(name string) bool {
	return strings.HasPrefix(name, "._") || isDSStoreFile(name)
}
