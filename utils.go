/*
utils.go

Utilitiy functions for wikionastick
*/
package main

import (
	"bytes"
	log "github.com/sirupsen/logrus"
)

func hasExtension(fname string, ext string) bool {

	// If period is not included, add it
	if ext[0:1] != "." {
		var concatBuffer bytes.Buffer
		concatBuffer.Write([]byte("."))
		concatBuffer.Write([]byte(ext))
		ext = concatBuffer.String()
	}

	// Get the length of the provided extension
	ext_length := len(ext)

	// Impossible
	if len(fname) < ext_length {
		return false
	} else {
		if fname[len(fname)-ext_length:] != ext {
			log.Debug("Filename does not have this extension")
			return false
		}
	}
	return true
}