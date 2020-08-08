package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func record(arguments []string) {

	_ = arguments

	binary, lookErr := exec.LookPath("/usr/local/bin/rec")
	if lookErr != nil {
		fmt.Printf("not found: %v\n", lookErr)
		return
	}

	//filename := "recording.mp3"
	filename := "recording.flac"

	var cmd *exec.Cmd
	// 24000 44100
	cmd = exec.Command(binary, filename, "channels", "2", "rate", "44100", "trim", "0", arguments[0])

	//cmd = exec.Command(binary, "recording.mp3", "channels", "1", "rate", "24000", "trim", "0", arguments[0])

	/*

	 File Size: 33.5k     Bit Rate: 148k
	  Encoding: FLAC          Info: Processed by SoX
	  Channels: 1 @ 16-bit
	Samplerate: 16000Hz
	Replaygain: off
	  Duration: 00:00:01.81

	*/

	//if len(arguments) == 1 {
	//} else {
	//	switch len(arguments[1:]) {
	//	case 1:
	//		cmd = exec.Command(binary, arguments[1])
	//	case 2:
	//		cmd = exec.Command(binary, arguments[1], arguments[2])
	//	case 3:
	//		cmd = exec.Command(binary, arguments[1], arguments[2], arguments[3])
	//	case 4:
	//		cmd = exec.Command(binary, arguments[1], arguments[2], arguments[3], arguments[4])
	//	}
	//}

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Env = os.Environ()
	err = cmd.Run()
	if err != nil {
		fmt.Printf("could not run", err)
		return
	}

	audioContent, err := ioutil.ReadFile("recording.mp3")
	if err != nil {
		fmt.Printf("could not read file: %v\n", err)
		return
	}

	// Append EOF to get back to the prompt
	audioContent = append(audioContent, byte(0))

	// The audioContent is binary.
	err = ioutil.WriteFile(filename, audioContent, 0644)
	if err != nil {
		fmt.Printf("could not write file: %v\n", err)
		return
	}
	fmt.Printf("Audio content written to file: %v\n", filename)
}
