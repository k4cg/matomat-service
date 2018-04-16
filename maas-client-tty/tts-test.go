package main

import (
	"github.com/hegedustibor/htgo-tts"
)

func main() {
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak("So long and thanks for the fish")
}
