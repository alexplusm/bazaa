package services

import (
	"archive/zip"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/fileutils"
)

type FileService struct{}

func (service *FileService) CopyFiles(files []*multipart.FileHeader, copyPath string) ([]string, error) {
	filenames := make([]string, 0)

	for _, file := range files {
		fmt.Println("file:", file.Filename, file.Size)

		// Source
		src, err := file.Open()
		if err != nil {
			return filenames, err
		}

		filename := filepath.Base(file.Filename)
		fp := filepath.Join(copyPath, filename)

		// Destination
		dst, err := os.Create(fp)
		if err != nil {
			return filenames, err
		}

		// Copy
		if _, err := io.Copy(dst, src); err != nil {
			return filenames, err
		}
		filenames = append(filenames, filename)

		if err = src.Close(); err != nil {
			return filenames, err
		}
		if err = dst.Close(); err != nil {
			return filenames, err
		}

	}

	return filenames, nil
}

func (service *FileService) UnzipImages(filenames []string) ([]bo.ImageParsingResult, error) {
	return unzipFiles(consts.MediaTempDir, consts.MediaRoot, filenames)
}

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
		if fileutils.GetExtension(name) == extension {
			availableExtention = true
		}
	}
	return availableExtention
}

const (
	withViolationDir = "withViolation"
	noViolationDir   = "noViolation"
)

// TODO: Должны совпадать с опциями игры: придется синхронизиться руками
// question: "Есть ли нарушение?"
// [{option: 0, value: "Есть нарушение"}, {option: 1, value: "Нет нарушения"}]
// TODO:  bo/image.go | сделать структуру, и вынести бизнес логику туда!
const (
	WithViolationCategory = "0"
	NoViolationCategory   = "1"
	UndefinedCategory     = "undefined"
)

func unzipFiles(srcPath string, destPath string, filenames []string) ([]bo.ImageParsingResult, error) {
	dir, err := os.Open(srcPath)
	if err != nil {
		return nil, err // todo: nil?
	}
	filesInfo, err := dir.Readdir(-1)
	if err != nil {
		return nil, err // todo: nil?
	}

	result := make([]bo.ImageParsingResult, 0, 500)

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

func unzip(src string, destination string) ([]bo.ImageParsingResult, error) {
	/*
		INFO:
		source: https://www.geeksforgeeks.org/how-to-uncompress-a-file-in-golang/

		TODO:
		Функция заточена для фоток - мб обобщить?
		+ есть бизнесс логика
	*/

	var parsingResults []bo.ImageParsingResult

	reader, err := zip.OpenReader(src)
	if err != nil {
		return parsingResults, err
	}
	defer reader.Close()

	rootFolder := destination

	for _, f := range reader.File {
		fname := f.FileInfo().Name()
		fpath := filepath.Join(rootFolder, fname)

		// INFO: skip nested dirs and invalid files
		if f.FileInfo().IsDir() || IsInvalidImageFileName(fname) {
			continue
		}

		// TODO: after testing refactor with switch
		// TODO: business logic -> need remove
		withViolation := strings.HasSuffix(f.Name, filepath.Join(withViolationDir, fname))
		noViolation := strings.HasSuffix(f.Name, filepath.Join(noViolationDir, fname))

		var result bo.ImageParsingResult

		if withViolation {
			result = bo.ImageParsingResult{fname, WithViolationCategory}
		} else if noViolation {
			result = bo.ImageParsingResult{fname, NoViolationCategory}
		} else {
			result = bo.ImageParsingResult{fname, UndefinedCategory}
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
