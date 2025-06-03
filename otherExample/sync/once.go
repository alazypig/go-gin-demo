package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image
var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":   loadIcon("left.png"),
		"right":  loadIcon("right.png"),
		"top":    loadIcon("top.png"),
		"bottom": loadIcon("bottom.png"),
	}
}

func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func main() {

}
