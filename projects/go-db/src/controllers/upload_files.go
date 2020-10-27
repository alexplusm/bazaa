package controllers

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Alexplusm/bazaa/projects/go-db/src/configs"
	"github.com/Alexplusm/bazaa/projects/go-db/src/utils/files"
	"github.com/labstack/echo"
)

const (
	undefinedDir     = "undefined"
	withViolationDir = "withViolation"
	noViolationDir   = "noViolation"
)

type categoryType int8

const (
	withViolationCategory categoryType = iota
	noViolationCategory
	undefinedCategory
)

type fileParsingResult struct {
	filename string
	category categoryType
}

// UploadFiles controller
func UploadFiles(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	/*
	* Although the field is called "files"
	* expected only one file - zip archive
	* The array is made for the future.
	 */
	files := form.File["files"]

	for _, file := range files {
		fmt.Println("file:", file.Filename, file.Size)

		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		filename := filepath.Base(file.Filename)
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

	unzipFiles()

	return c.String(http.StatusOK, "OK")
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

	for _, fileInfo := range filesInfo {
		fmt.Println("zip archive", fileInfo.Name(), fileInfo.Size())

		r, _ := unzip(filepath.Join(configs.MediaTempDir, fileInfo.Name()), configs.MediaRoot)
		fmt.Println("RESULT", r, len(r)) // todo: error
	}

	return nil
}

// unzip will decompress a zip archived file
func unzip(src string, destination string) ([]fileParsingResult, error) {
	/*
		source: https://www.geeksforgeeks.org/how-to-uncompress-a-file-in-golang/
	*/

	var parsingResults []fileParsingResult

	reader, err := zip.OpenReader(src)
	if err != nil {
		return parsingResults, err
	}
	defer reader.Close()

	rootFolder := destination

	for _, f := range reader.File {
		fname := f.FileInfo().Name()
		fpath := filepath.Join(rootFolder, fname)

		// for nested dirs
		if f.FileInfo().IsDir() || files.IsInvalidImageFileName(fname) {
			continue
		}

		withViolation := strings.HasSuffix(f.Name, filepath.Join(withViolationDir, fname))
		noViolation := strings.HasSuffix(f.Name, filepath.Join(noViolationDir, fname))

		var result fileParsingResult

		// TODO: after testing refactor with switch
		if withViolation {
			result = fileParsingResult{fname, withViolationCategory}
		} else if noViolation {
			result = fileParsingResult{fname, noViolationCategory}
		} else {
			result = fileParsingResult{fname, undefinedCategory}
		}

		fmt.Println(result)

		parsingResults = append(parsingResults, result)

		// Creating the files in the target directory
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return parsingResults, err
		}

		// The created file will be stored in
		// outFile with permissions to write &/or truncate
		outFile, err := os.OpenFile(fpath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			f.Mode())
		if err != nil {
			return parsingResults, err
		}

		rc, err := f.Open()
		if err != nil {
			return parsingResults, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer so that it closes the outfile
		// before the loop moves to the next iteration.
		outFile.Close()
		rc.Close()

		if err != nil {
			return parsingResults, err
		}
	}

	return parsingResults, nil
}
