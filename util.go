package keygen

import (
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

func writeFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0700)
}

func encodePemBlock(out io.Writer, block *pem.Block) error {
	if err := pem.Encode(out, block); err != nil {
		return fmt.Errorf("could not write pem block due to %s ", err.Error())
	}
	return nil
}
