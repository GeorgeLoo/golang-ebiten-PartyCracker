

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
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	// "image"
	"image/color"
	//"log"
	//"os"
	//"math"
	//"path/filepath"
)

const (
	screenwidth = 800
	screenheight = 480
	datafolder = "data"
	sampleRate   = 44100
	clockwise = 200
	anticlockwise = 202
	notRotating = 204
	kLunarLanderHeight = 50
	kMoonOrbitalSpeed = 1600
	kCommandModule = 801
	kLunarModule = 802
)

var (
	canChangeFullscreen bool
)


func update(screen *ebiten.Image) error {

	if ebiten.IsRunningSlowly() {
		return nil
		
	}

	screen.Fill(color.NRGBA{255, 255, 0, 0xff})

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
	scale := 1.0
	// Initialize Ebiten, and loop the update() function
	if err := ebiten.Run(update, screenwidth, screenheight, scale, "Party Cracker Simulator 0.0 by George Loo"); err != nil {
		panic(err)
	}
	fmt.Printf("Party Cracker Program ended -----------------\n")

}

