package cmd

import (
	"errors"
)

// Run dispatches subcommands by name.
func Run(name string) error {
	switch name {
	case "importExcel":
		return importLetters()
	case "importAudio":
		return importAudio()
	case "importWord":
		return importWord()
	case "importWordAudio":
		return importWordAudio()
	case "importGroupPhonics":
		return importGroupPhonics()
	default:
		return errors.New("unknown command: " + name)
	}
}
