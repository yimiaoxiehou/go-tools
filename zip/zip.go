package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ZipDirContent zips the contents of a directory and saves the result to a destination file
func ZipDirContent(destination, sourceDir string) error {
	// Create the destination zip file
	zipFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create a zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the source directory and add its contents to the zip file
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Generate the path for the file in the zip
		zipPath := strings.TrimPrefix(path, sourceDir)
		if info.IsDir() {
			zipPath += "/"
		}

		// Skip if the file is not a regular file
		if !info.Mode().IsRegular() {
			return nil
		}

		// Create a header for the file to be added to the zip
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = zipPath

		// Create the file in the zip
		file, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// Open the source file
		source, err := os.Open(path)
		if err != nil {
			return err
		}
		defer source.Close()

		// Copy the source file to the zip file
		_, err = io.Copy(file, source)
		if err != nil {
			return err
		}

		return nil
	})
}

// Unzip extracts the contents of a zip file to a destination directory
// sourceZip: the path to the zip file to be extracted
// destDir: the destination directory where the contents will be extracted
func Unzip(sourceZip, destDir string) error {
	// Open the zip file for reading
	zr, err := zip.OpenReader(sourceZip)
	if err != nil {
		return err
	}
	defer zr.Close()

	// Create the destination directory if it doesn't exist
	if destDir != "" {
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return err
		}
	}

	// Iterate through each file in the zip and extract its contents
	for _, file := range zr.File {
		path := filepath.Join(destDir, file.Name)

		// Create directories if the file is a directory
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			continue
		}

		// Open the file from the zip
		fr, err := file.Open()
		if err != nil {
			return err
		}
		defer fr.Close()

		// Create and open the corresponding file in the destination directory
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer fw.Close()

		// Copy the contents from the zip file to the destination file
		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}
	}
	return nil
}
