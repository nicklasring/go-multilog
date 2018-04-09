package multilog

import (
	"io"
	"log"
	"os"
)

func addToMultiWriter(writers ...io.Writer) {
	log.SetOutput(io.MultiWriter(writers...))
}

func Add(files []string) {
	lwriters := make([]io.Writer, 0)
	for _, file := range files {
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
		if err != nil {
			log.Fatal(err)
		}
		lwriters = append(lwriters, f)
	}

	lwriters = append(lwriters, os.Stdout)
	addToMultiWriter(lwriters...)
}
