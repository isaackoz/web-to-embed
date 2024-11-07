package convert

import (
	"bytes"
	"compress/gzip"
	"mime"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Asset struct {
	Path           string
	NormalizedName string
	MimeType       string
	Size           int64
	Contents       []byte
}

// GetAssetsFromDir reads all files in the specified directory path,
// and converts each file to an Asset struct containing metadata and contents,
// and returns a slice of these assets. It walks to directory recursively.
//
// Parameters:
//   - dirPath: The path to a directory containing files to convert
//
// Returns:
//   - []Asset: A slice of Asset structs with file data
//   - error: An error if the dir can not be read
func GetAssetsFromDir(dirPath string) ([]Asset, error) {
	var assets []Asset
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		mimeType := mime.TypeByExtension(filepath.Ext(path))
		if mimeType == "" {
			mimeType = "application/octet-stream" // Default to binary if unknown
		}

		compressedData, compressedSize, err := CompressFileContent(contents)

		if err != nil {
			return err
		}
		relativePath, err := filepath.Rel(dirPath, path)
		if err != nil {
			return err
		}

		relativePath = strings.ReplaceAll(relativePath, `\`, `/`)

		normalizedName := regexp.MustCompile(`[^0-9a-zA-Z]`).ReplaceAllString(relativePath, "_")

		asset := Asset{
			Path:           relativePath,
			NormalizedName: normalizedName,
			MimeType:       mimeType,
			Size:           compressedSize,
			Contents:       compressedData,
		}

		assets = append(assets, asset)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return assets, nil
}

// CompressFileContent takes in a []byte and compresses it by gzipping it. Since all browsers support gzip, it makes sense
// to gzip them beforehand to save space, especially in an embedded environment where every byte counts.
//
// Parameters:
//   - contents: The original file contents as a byte slice.
//
// Returns:
//   - []byte: The compressed file contents.
//   - int64: The size of the compressed contents.
//   - error: An error if compression fails.
func CompressFileContent(contents []byte) ([]byte, int64, error) {
	var compressedContents bytes.Buffer
	writer := gzip.NewWriter(&compressedContents)

	_, err := writer.Write(contents)
	if err != nil {
		return nil, 0, err
	}

	err = writer.Close()
	if err != nil {
		return nil, 0, err
	}

	compressedData := compressedContents.Bytes()
	compressedSize := int64(len(compressedData))

	return compressedData, compressedSize, nil
}

func GetTotalSizeOfAssets(assets []Asset) int64 {
	var totalSize int64
	for _, asset := range assets {
		totalSize += asset.Size
	}
	return totalSize
}
