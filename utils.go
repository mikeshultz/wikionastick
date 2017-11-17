/*
utils.go

Utilitiy functions for wikionastick
*/
package main

import (
	"bytes"
	log "github.com/sirupsen/logrus"
)

func HasExtension(fname string, ext string) bool {

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
			return false
		}
	}
	return true
}

func LogLevelTranslate(level string) log.Level {
	switch level {
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "error":
		return log.ErrorLevel
	default: 
		return log.WarnLevel
	}
}