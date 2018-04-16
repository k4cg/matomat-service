package main

import (
	"os"
	"github.com/faiface/beep/speaker"
)

func main() {
	f, _ := os.Open("audio/output.mp3")
	s, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)
	peaker.Play(beep.Seq(s, beep.Callback(func() {
	       close(done)
	)))
}
