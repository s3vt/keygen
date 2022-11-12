package keygen

import "io"

type Keymaker interface {
	MakeKeys(reader io.Reader) error
	PrintKeys(printPrivate bool) error
	WriteKeysToFile(filename string) error
}
