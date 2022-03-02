package bundles

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
)

var bundlePartName *regexp.Regexp

func findBundleParts(dir fs.FS) (names []string, err error) {
	dirEntries, err := fs.ReadDir(dir, "/")

	if err != nil {
		err = fmt.Errorf("Error reading dir for bundle parts: %w", err)
		return
	}

	for _, entry := range dirEntries {
		if bundlePartName.MatchString(entry.Name()) {
			names = append(names, entry.Name())
		}
	}

	return
}

func ConcatBundlePieces(dir fs.FS) error {
	bundlePieceNames, err := findBundleParts(dir)

	if err != nil {
		return err
	} else if len(bundlePieceNames) == 0 {
		return fmt.Errorf("No bundle pieces found")
	}

	target, err := os.OpenFile("crcbundle", os.O_CREATE, 0)

	for _, fileName := range bundlePieceNames {
		file, err := dir.Open(fileName)
		if err != nil {
			return fmt.Errorf("Error opening %v for reading: %w", fileName, err)
		}
		defer file.Close()

	}

}

func init() {
	bundlePartName = regexp.MustCompile(`CrcBundlePart\d`)
}
