package files

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Alexplusm/bazaa/projects/go-db/configs"
)

// INFO: service files with this prefix appear after unpacking
const serviceUnzipFilePrefix = "._"

var extentionsWhiteList = [3]string{"jpg", "png", "jpeg"}

// IsInvalidImageFileName check image filename
func IsInvalidImageFileName(name string) bool {
	return strings.HasPrefix(name, serviceUnzipFilePrefix) ||
		!hasAllowableImageExtension(name)
}

func hasAllowableImageExtension(name string) bool {
	var availableExtention = false

	for _, extension := range extentionsWhiteList {
		if getExtension(name) == extension {
			availableExtention = true
		}
	}
	return availableExtention
}

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

// FileParsingResult TODO: rename!
type FileParsingResult struct {
	filename string
	category categoryType
}

// UnzipImages unzip images
func UnzipImages(filenames []string) ([]FileParsingResult, error) {
	return unzipFiles(configs.MediaTempDir, configs.MediaRoot, filenames)
}

func unzipFiles(srcPath string, destPath string, filenames []string) ([]FileParsingResult, error) {
	dir, err := os.Open(srcPath)
	if err != nil {
		return nil, err // todo: nil?
	}
	filesInfo, err := dir.Readdir(-1)
	if err != nil {
		return nil, err // todo: nil?
	}

	result := make([]FileParsingResult, 0, 500)

	for _, fileInfo := range filesInfo {
		fmt.Println("zip archive", fileInfo.Name(), fileInfo.Size())

		res, err := unzip(filepath.Join(srcPath, fileInfo.Name()), destPath)

		result = append(result, res...)

		if err != nil {
			fmt.Println("EEEEE", err)
		}

		fmt.Println("res", len(res)) // todo: error
	}

	return result, nil
}

func unzip(src string, destination string) ([]FileParsingResult, error) {
	/*
		INFO:
		source: https://www.geeksforgeeks.org/how-to-uncompress-a-file-in-golang/

		TODO:
		Функция заточена для фоток - мб обобщить?
		+ есть бизнесс логика
	*/

	var parsingResults []FileParsingResult

	reader, err := zip.OpenReader(src)
	if err != nil {
		return parsingResults, err
	}
	defer reader.Close()

	rootFolder := destination

	for _, f := range reader.File {
		fname := f.FileInfo().Name()
		fpath := filepath.Join(rootFolder, fname)

		fmt.Println("fname", fname, "| fpath", fpath)

		// INFO: skip nested dirs and invalid files
		if f.FileInfo().IsDir() || IsInvalidImageFileName(fname) {
			continue
		}

		// TODO: after testing refactor with switch
		// TODO: bissness logic -> need remove
		withViolation := strings.HasSuffix(f.Name, filepath.Join(withViolationDir, fname))
		noViolation := strings.HasSuffix(f.Name, filepath.Join(noViolationDir, fname))

		var result FileParsingResult

		if withViolation {
			result = FileParsingResult{fname, withViolationCategory}
		} else if noViolation {
			result = FileParsingResult{fname, noViolationCategory}
		} else {
			result = FileParsingResult{fname, undefinedCategory}
		}

		fmt.Println(result)

		// Creating the files in the target directory
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return parsingResults, err
		}

		// The created file will be stored in outFile with permissions to read
		// todo: нужен ли  os.O_TRUNC пермишен ?
		outFile, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE, f.Mode())
		if err != nil {
			// todo: log error and continue?
			return parsingResults, err
		}

		rc, err := f.Open()
		if err != nil {
			// todo: log error and continue
			return parsingResults, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer so that it closes the outfile
		// before the loop moves to the next iteration.
		outFile.Close()
		rc.Close()

		if err != nil {
			// todo: log error and continue
			return parsingResults, err
		}

		parsingResults = append(parsingResults, result)
	}

	return parsingResults, nil
}
