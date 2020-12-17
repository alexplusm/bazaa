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

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
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
func (service *FileService) UnzipArchives(
	archivesPath []string, dstPath string,
) ([]bo.ImageParsingResult, error) {
	result := make([]bo.ImageParsingResult, 0, 500)

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

// TODO: rename ::: todo: remove
func (service *FileService) GetArchivesPaths(archivesNames []string, srcPath string) ([]string, error) {
	dir, err := os.Open(srcPath)
	if err != nil {
		return nil, err // TODO: build error
	}
	filesInfo, err := dir.Readdir(-1)
	if err != nil {
		return nil, err // TODO: build error
	}

	archivesPaths := make([]string, 0, len(archivesNames))

	for _, fileInfo := range filesInfo {
		for _, archName := range archivesNames {
			if archName == fileInfo.Name() {
				archivesPaths = append(archivesPaths, filepath.Join(srcPath, fileInfo.Name()))
				break
			}
		}
	}

	return archivesPaths, nil
}

// TODO: remove
//func (service *FileService) UnzipImages(filenames []string) ([]bo.ImageParsingResult, error) {
//	return unzipFiles(consts.MediaTempDir, consts.MediaRoot, filenames)
//}

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

//func unzipFiles(srcPath string, destPath string, filenames []string) ([]bo.ImageParsingResult, error) {
//	// TODO: unused filenames !!!
//	dir, err := os.Open(srcPath)
//	if err != nil {
//		return nil, err
//	}
//	// dir.
//	filesInfo, err := dir.Readdir(-1)
//	if err != nil {
//		return nil, err
//	}
//
//	result := make([]bo.ImageParsingResult, 0, 500)
//
//	for _, fileInfo := range filesInfo {
//		res, err := unzip(filepath.Join(srcPath, fileInfo.Name()), destPath)
//		result = append(result, res...)
//
//		if err != nil {
//			return nil, fmt.Errorf("unzip files: %v", err)
//		}
//	}
//	log.Info("unzip archive: count files =", len(result))
//
//	return result, nil
//}

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
			result = bo.ImageParsingResult{Filename: fname, Category: WithViolationCategory}
		} else if noViolation {
			result = bo.ImageParsingResult{Filename: fname, Category: NoViolationCategory}
		} else {
			result = bo.ImageParsingResult{Filename: fname, Category: UndefinedCategory}
		}

		// Creating the files in the target directory
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return parsingResults, err
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
		// before the loop moves to the next iteration.
		err = outFile.Close()
		err = rc.Close()

		if err != nil {
			log.Error("unzip: error while file closing: %v", err)
			continue
		}

		parsingResults = append(parsingResults, result)
	}

	return parsingResults, nil
}
