

// Copyright 2017 George Loo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


/*

PartyCracker.go
Copyright 2017 George Loo


*/
package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	// "image"
	"image/color"
	"log"
	"os"
	"math"
	"path/filepath"
    "math/rand"
    "time"

)

const (
	screenwidth = 800
	screenheight = 480
	datafolder = "data"
	sampleRate   = 44100
	kCrackerX = 425
	kCrackerY = 300
	kleftside = 700
	krightside = 711
)

type crackerData struct {
	pointedDir int 
	x, y float64   
	cx, cy int // centre x y 
	w, h int
	image *ebiten.Image
	stretched *ebiten.Image
	explode1, explode2, explode3 *ebiten.Image
	leftSideclick, rightSideClick bool
	timerStart bool
	timercount int 
	leftarrow, rightarrow *ebiten.Image
	side int  // which side the arrow will point
}

type soundData struct {
	mute bool
	//audioContext    *audio.Context
	audioPlayer     *audio.Player
	soundArr       []audio.Player

}


var (
	canChangeFullscreen bool
	aCracker crackerData
	sound soundData
	sound0, sound1 int
	audioContext    *audio.Context
	soundloop *audio.Player
	randnum int

)



func initprog() {

	aCracker.init("cracker2.png", 0, kCrackerX,kCrackerY)
	sound.init()
	sound0 = sound.load("sound0.wav")
	sound1 = sound.load("sound1.wav")
	//ebiten.SetFullscreen(true)
	soundloop = loadloop("sound1.wav")
	soundloop.Play()
	soundloop.Pause()

	rand.Seed( time.Now().UnixNano())

}

func loadloop(fn string) *audio.Player {

	wavF, err := ebitenutil.OpenFile(filepath.Join(datafolder, fn))
	if err != nil {
		log.Fatal(err)
	}

	wavS, err := wav.Decode(audioContext, wavF)
	if err != nil {
		log.Fatal(err)
	}

	s := audio.NewInfiniteLoop(wavS, wavS.Size())

	player, err := audio.NewPlayer(audioContext, s)
	if err != nil {
		log.Fatal(err)
	}
	return player
}

/*
only need one audio context, or so I think...
*/
func (s *soundData) init() {
	const sampleRate  = 44100
	var err error
	s.mute = false
	audioContext, err = audio.NewContext(sampleRate)
	if err != nil {
		log.Fatal(err)
	}

}

func (s *soundData) load(fn string) int {
	var err error
	
	//var audioPlayer     *audio.Player
	
	f, err := os.Open(filepath.Join(datafolder, fn))
	if err != nil {
		log.Fatal(err)
	}

	d, err := wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}

	s.audioPlayer, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}
	s.soundArr = append(s.soundArr, *s.audioPlayer)
	i := len(s.soundArr) - 1
	return i // index to the sound

}


func (s *soundData) play(idx int) error {
	//var err error

	if s.mute {
		return nil
	}
	ap := s.soundArr[idx]
	if !ap.IsPlaying() {
		//fmt.Print("sound or not?\n")
		ap.Rewind()
		err := ap.Play()
		if err != nil {
			panic(err)
		}
	}

	if err := audioContext.Update(); err != nil {
		fmt.Print(" !!!!!!!!!!!! SOUND ERROR \n")
		return err
	}
	return nil 
}


func readimg(fn string) *ebiten.Image {
	var err error
	var fname string
	fname = filepath.Join(datafolder, fn)
	img, _, err := ebitenutil.NewImageFromFile(
		fname,
		ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	return img

}

func (l *crackerData) CrackerTimer(screen *ebiten.Image) {
	

	screen.Fill(color.NRGBA{255, 255, 0, 0xff})  // yellow
	l.draw(screen, l.image)

	if !l.timerStart {
		return
	}
	l.timercount++
	if l.timercount == 6 {
		//l.draw(screen,l.explode1)
		//sound.play(sound0)

	}
	if l.timercount == 1 {
		l.draw(screen,l.explode1)
		sound.play(sound0)
	}

	if l.timercount > 1 {
		l.draw(screen,l.explode1)
	}

	if l.timercount == 16 {
		sound.play(sound1)
	}

	if l.timercount > 16 {
		l.draw(screen,l.explode2)
	}

	if l.timercount > 32 {
		l.draw(screen,l.explode3)
	}

	if l.timercount == 60 {

		randnum = rand.Intn(100)+1
		if randnum > 50 {
			fmt.Print("RIGHT SIDE \n")
			l.draw(screen,l.rightarrow)
			l.side = krightside
		} else {
			fmt.Print("LEFT SIDE \n")
			l.side = kleftside
		}

	}

	if l.timercount > 120 {
		if l.side == kleftside {
			l.draw(screen,l.leftarrow)
		} else {
			l.draw(screen,l.rightarrow)
		}
	}


	if l.timercount > 300 {
		fmt.Print("timer stopped\n")

		l.timerStart = false
		l.leftSideclick = false
		l.rightSideClick = false

	}

}

func (l *crackerData) CheckForCrackerPull() {

	if !l.timerStart && l.leftSideclick && l.rightSideClick {
		//l.leftSideclick = false
		//l.rightSideClick = false
		//sound.play(sound0)
		fmt.Print("***********BOOMZ \n")
		l.timercount = 0
		l.timerStart = true
	}
}

func (l *crackerData) init(picFilename string,
						  pointedDir int,
						  x float64,
						  y float64  ) {

	l.image = readimg(picFilename)
	l.pointedDir = pointedDir
	l.x = x
	l.y = y
	l.leftSideclick = false 
	l.rightSideClick = false
	l.timerStart = false
	l.explode1 = readimg("cracker3.png")
	l.explode2 = readimg("cracker4.png")
	l.explode3 = readimg("cracker5.png")

	l.leftarrow = readimg("leftarrow.png")
	l.rightarrow = readimg("rightarrow.png")

}					

func (l *crackerData) draw(screen *ebiten.Image, image *ebiten.Image) {
	w, h := image.Size()
	//fmt.Printf("w %d h %d \n",w,h)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Reset()
	opts.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	opts.GeoM.Rotate(float64(l.pointedDir % 360) * 2 * math.Pi / 360)
	//opts.GeoM.Scale( 1.0, 1.0 )
	opts.GeoM.Scale( 1.0, 1.0 )
	opts.GeoM.Translate(l.x, l.y)

	screen.DrawImage(image, opts)

}

func approx(n,l,h int) bool {

	// if math.Abs(float64(n)-float64(m)) < 10 {
	// 	return true
	// }
	if n > l && n < h {
		return true
	}

	// if n > 355 && m == 0 {
	// 	return true //handle 0 case 
	// }
	return false 
}


func update(screen *ebiten.Image) error {

	if ebiten.IsRunningSlowly() {
		return nil
		
	}

	//screen.Fill(color.NRGBA{255, 255, 0, 0xff})  // yellow

	//aCracker.draw(screen,aCracker.image)
	aCracker.CrackerTimer(screen)


	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if mx < 50 && my < 50 {
			togglFullscreen()
		}

		if mx > (screenwidth-50) && my < 50 {  // top right hand corner
			togglFullscreen()
		}

		h := 100

		if approx(my,screenheight/2-h,screenheight/2+h) &&
			approx(mx,0,screenwidth/4) {
				//sound.play(sound0)
				aCracker.leftSideclick = true
				fmt.Print("***********leftSideclick \n")
			}

		if approx(my,screenheight/2-h,screenheight/2+h) &&
			approx(mx,screenwidth-screenwidth/4,screenwidth) {
				//sound.play(sound0)
				aCracker.rightSideClick = true
				fmt.Print("***********rightSideClick \n")
			}

	}

	if ebiten.IsKeyPressed(ebiten.KeyU) {
		sound.play(sound0)
	}

	aCracker.CheckForCrackerPull()

	return nil
}

func togglFullscreen() {
	//if ebiten.IsKeyPressed(ebiten.KeyF) {
	if canChangeFullscreen {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
		canChangeFullscreen = false
		
	} else {
		canChangeFullscreen = true
	}
}

func main() {

	initprog()
	scale := 1.0
	// Initialize Ebiten, and loop the update() function
	if err := ebiten.Run(update, screenwidth, screenheight, scale, "Party Cracker Simulator 0.0 by George Loo"); err != nil {
		panic(err)
	}
	fmt.Printf("Party Cracker Program ended -----------------\n")

}

