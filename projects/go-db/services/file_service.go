package services

import (
	"archive/zip"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/utils/fileutils"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/logutils"
)

type FileService struct{}

func (service *FileService) SaveFile(file *multipart.FileHeader, dstPath string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("%v SaveFile: %v", logutils.GetStructName(service), err)
	}

	filename := filepath.Base(file.Filename)
	dstFilepath := filepath.Join(dstPath, filename)

	dst, err := os.Create(dstFilepath)
	if err != nil {
		return "", fmt.Errorf("%v SaveFile: %v", logutils.GetStructName(service), err)
	}

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("%v SaveFile: %v", logutils.GetStructName(service), err)
	}

	if err = src.Close(); err != nil {
		return "", fmt.Errorf("%v SaveFile: %v", logutils.GetStructName(service), err)
	}
	if err = dst.Close(); err != nil {
		return "", fmt.Errorf("%v SaveFile: %v", logutils.GetStructName(service), err)
	}

	return dstFilepath, nil
}

func (service *FileService) SaveFiles(files []*multipart.FileHeader, dstPath string) ([]string, error) {
	filesPaths := make([]string, 0)

	for _, file := range files {
		fpath, err := service.SaveFile(file, dstPath)
		if err != nil {
			return filesPaths, err
		}

		filesPaths = append(filesPaths, fpath)
	}

	return filesPaths, nil
}

// TODO: return values || !!!
func (service *FileService) UnzipArchives(archivesPath []string, dstPath string) ([]zip.File, error) {
	result := make([]zip.File, 0, 1024)

	for _, archivePath := range archivesPath {
		// TODO: method
		res, err := unzip(archivePath, dstPath)
		if err != nil {
			return nil, fmt.Errorf("unzip files: %v", err)
		}

		result = append(result, res...)
	}
	log.Info("unzip archive: count files =", len(result))

	return result, nil
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

func unzip(src string, dstPath string) ([]zip.File, error) {
	/*
		INFO:
		source: https://www.geeksforgeeks.org/how-to-uncompress-a-file-in-golang/

		TODO:
		Функция заточена для фоток - мб обобщить?
		+ есть бизнесс логика
	*/
	results := make([]zip.File, 0, 1024)

	reader, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	rootFolder := dstPath

	for _, f := range reader.File {
		fname := f.FileInfo().Name()
		fpath := filepath.Join(rootFolder, fname)

		// INFO: skip nested dirs and invalid files
		if f.FileInfo().IsDir() || IsInvalidImageFileName(fname) { // TODO: funcHandler
			continue
		}

		// Creating the files in the target directory
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return nil, err
		}

		// The created file will be stored in outFile with permissions to read
		outFile, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE, f.Mode())
		if err != nil {
			log.Error("unzip: error while file opening: ", err, fpath)
			continue
		}

		rc, err := f.Open()
		if err != nil {
			log.Error("unzip: error while file opening: ", err, fpath, outFile.Close())
			continue
		}

		_, err = io.Copy(outFile, rc)

		// INFO: Close the file without defer so that it closes the outfile
		// 		before the loop moves to the next iteration.
		err = outFile.Close()
		err = rc.Close()

		if err != nil {
			log.Error("unzip: error while file closing: %v", err)
			continue
		}

		results = append(results, *f)
	}

	return results, nil
}
