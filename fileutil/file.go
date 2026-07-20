package fileutil

import (
	"archive/zip"
	"bufio"
	"io"
	"io/fs"
	"os"
	"sync"
)

type FileReader struct {
	*bufio.Reader
	file   *os.File
	offset int64
}

func NewFileReader(path string) (*FileReader, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *FileReader) ReadLine() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (f *FileReader) Offset() int64 { _ = "STUB: not implemented"; return 0 }

func (f *FileReader) SeekOffset(offset int64) error { _ = "STUB: not implemented"; return nil }

func (f *FileReader) Close() error { _ = "STUB: not implemented"; return nil }

func IsExist(path string) bool { _ = "STUB: not implemented"; return false }

func CreateFile(path string) bool { _ = "STUB: not implemented"; return false }

func CreateDir(absPath string) error { _ = "STUB: not implemented"; return nil }

func CopyDir(srcPath string, dstPath string) error { _ = "STUB: not implemented"; return nil }

func IsDir(path string) bool { _ = "STUB: not implemented"; return false }

func RemoveFile(path string, onDelete ...func(path string)) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoveDir(path string, onDelete ...func(path string)) error {
	_ = "STUB: not implemented"
	return nil
}

func CopyFile(srcPath string, dstPath string) error { _ = "STUB: not implemented"; return nil }

func ClearFile(path string) error { _ = "STUB: not implemented"; return nil }

func ReadFileToString(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ReadFileByLine(path string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ListFileNames(path string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func IsZipFile(filepath string) bool { _ = "STUB: not implemented"; return false }

func Zip(path string, destPath string) error { _ = "STUB: not implemented"; return nil }

func zipFile(filePath string, destPath string) error { _ = "STUB: not implemented"; return nil }

func zipFolder(folderPath string, destPath string) error { _ = "STUB: not implemented"; return nil }

func addFileToArchive1(fpath string, archive *zip.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func addFileToArchive2(w *zip.Writer, basePath, baseInZip string) error {
	_ = "STUB: not implemented"
	return nil
}

func UnZip(zipFile string, destPath string) error { _ = "STUB: not implemented"; return nil }

func ZipAppendEntry(fpath string, destPath string) error { _ = "STUB: not implemented"; return nil }

func safeFilepathJoin(path1, path2 string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func IsLink(path string) bool { _ = "STUB: not implemented"; return false }

func FileMode(path string) (fs.FileMode, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileMode), nil
}

func MiMeType(file any) string { _ = "STUB: not implemented"; return "" }

func CurrentPath() string { _ = "STUB: not implemented"; return "" }

func FileSize(path string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func DirSize(path string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func MTime(filepath string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func Sha(filepath string, shaType ...int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ReadCsvFile(filepath string, delimiter ...rune) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WriteCsvFile(filepath string, records [][]string, append bool, delimiter ...rune) error {
	_ = "STUB: not implemented"
	return nil
}

func WriteStringToFile(filepath string, content string, append bool) error {
	_ = "STUB: not implemented"
	return nil
}

func WriteBytesToFile(filepath string, content []byte) error { _ = "STUB: not implemented"; return nil }

func ReadFile(path string) (reader io.ReadCloser, closeFn func(), err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil, nil
}

func escapeCSVField(field string, delimiter rune) string { _ = "STUB: not implemented"; return "" }

func WriteMapsToCsv(filepath string, records []map[string]any, appendToExistingFile bool, delimiter rune,
	headers ...[]string) error {
	_ = "STUB: not implemented"
	return nil
}

func isCsvSupportedType(v interface{}) bool { _ = "STUB: not implemented"; return false }

func ChunkRead(file *os.File, offset int64, size int, bufPool *sync.Pool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParallelChunkRead(filePath string, linesCh chan<- []string, chunkSizeMB, maxGoroutine int) error {
	_ = "STUB: not implemented"
	return nil
}
