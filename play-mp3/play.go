package main

import (
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
)

var (
	err    error
	file   *os.File
	dec    *minimp3.Decoder
	c      *oto.Context
	player *oto.Player
)

func play(arguments []string) {

	_ = arguments

	if file, err = os.Open("./" + arguments[0]); err != nil {
		log.Fatal(err)
	}
	if dec, err = minimp3.NewDecoder(file); err != nil {
		log.Fatal(err)
	}
	started := dec.Started()
	<-started

	log.Printf("Convert audio sample rate: %d, channels: %d\n", dec.SampleRate, dec.Channels)

	if c == nil {
		c, err = oto.NewContext(dec.SampleRate, dec.Channels, 2, 1024)
		if err != nil {
			log.Fatal(err)
		}
	}

	player = c.NewPlayer()

	var waitForPlayOver = new(sync.WaitGroup)
	waitForPlayOver.Add(1)

	go func() {
		for {
			var data = make([]byte, 1024)
			_, err := dec.Read(data)
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}
			_, err = player.Write(data)
			if err != nil {
				break
			}
		}
		log.Println("over play.")
		waitForPlayOver.Done()
	}()
	waitForPlayOver.Wait()

	<-time.After(time.Second)
	dec.Close()
	err = player.Close()
	if err != nil {
		log.Fatal(err)
	}
	//err = c.Close()
}

/*

package main

import (
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
)

func main() {
	var err error
	var file *os.File
	var dec *minimp3.Decoder
	var player *oto.Player

	if file, err = os.Open("./talk.mp3"); err != nil {
		log.Fatal(err)
	}
	if dec, err = minimp3.NewDecoder(file); err != nil {
		log.Fatal(err)
	}
	started := dec.Started()
	<-started

	log.Printf("Convert audio sample rate: %d, channels: %d\n", dec.SampleRate, dec.Channels)

	c, err := oto.NewContext(dec.SampleRate, dec.Channels, 2, 1024)
	if err != nil {
		log.Fatal(err)
	}

	player = c.NewPlayer()

	var waitForPlayOver = new(sync.WaitGroup)
	waitForPlayOver.Add(1)

	go func() {
		for {
			var data = make([]byte, 1024)
			_, err := dec.Read(data)
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}
			player.Write(data)
		}
		log.Println("over play.")
		waitForPlayOver.Done()
	}()
	waitForPlayOver.Wait()

	<-time.After(time.Second)
	dec.Close()
	player.Close()
}


*/
