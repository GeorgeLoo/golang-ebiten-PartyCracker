

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
	//"github.com/hajimehoshi/ebiten/audio"
	//"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	// "image"
	"image/color"
	"log"
	//"os"
	"math"
	"path/filepath"
)

const (
	screenwidth = 800
	screenheight = 480
	datafolder = "data"
	sampleRate   = 44100

)

type crackerData struct {
	pointedDir int 
	x, y float64   
	cx, cy int // centre x y 
	w, h int
	image *ebiten.Image
	stretched *ebiten.Image

}

var (
	canChangeFullscreen bool
	aCracker crackerData
)



func initprog() {

	aCracker.init("cracker2.png", 0, 425,300)
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

func (l *crackerData) init(picFilename string,
						  pointedDir int,
						  x float64,
						  y float64  ) {

	l.image = readimg(picFilename)
	l.pointedDir = pointedDir
	l.x = x
	l.y = y
}					

func (l *crackerData) draw(screen *ebiten.Image) {
	w, h := l.image.Size()
	//fmt.Printf("w %d h %d \n",w,h)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Reset()
	opts.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	opts.GeoM.Rotate(float64(l.pointedDir % 360) * 2 * math.Pi / 360)
	//opts.GeoM.Scale( 1.0, 1.0 )
	opts.GeoM.Scale( 1.0, 1.0 )
	opts.GeoM.Translate(l.x, l.y)

	screen.DrawImage(l.image, opts)

}



func update(screen *ebiten.Image) error {

	if ebiten.IsRunningSlowly() {
		return nil
		
	}

	screen.Fill(color.NRGBA{255, 255, 0, 0xff})  // yellow

	aCracker.draw(screen)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if mx < 50 && my < 50 {
			togglFullscreen()
		}

		if mx > (screenwidth-50) && my < 50 {
			togglFullscreen()
		}
	}

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

