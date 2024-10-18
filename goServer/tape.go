package poker

import (
	"io"
	"os"
)

type tape struct {
	// os.File type has a Truncate method which effectively empties the file
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}
