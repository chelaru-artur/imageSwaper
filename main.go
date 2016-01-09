package main

import (
	"image"
	"image/draw"
	"image/png"
	"image/jpeg"
	"log"
	"os"
	"os/exec"
	"math/rand"
	"flag"
)

func main() {
    fileName :=  flag.String("filename", "*.jpg", "filename with extension")
    step := flag.Int("square-size", 50, "size of each square")
    flag.Parse()
    log.Println(*fileName)
	inputFile, _ := os.Open(*fileName)
	defer inputFile.Close()
	inputImg, _ := jpeg.Decode(inputFile)
	bounds := inputImg.Bounds()
	rectangles, points := GetImageParts(inputImg.Bounds(), *step)
	m := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y)) //*NRGBA (image.Image interface)
	for index, rect := range rectangles {
		rand.Seed(int64(index))
		rnd := rand.Intn(len(points))
		point := points[rnd]
		points = append(points[:rnd], points[rnd+1:]...)
		draw.Draw(m, rect, inputImg, point, draw.Src)
	} 
	w, _ := os.Create("new.png")
	defer w.Close()
	png.Encode(w, m) //Encode writes the Image m to w in PNG format.
	Show(w.Name())
}

// show  a specified file by Preview.app for OS X(darwin)
func Show(name string) {
	command := "xdg-open"
	cmd := exec.Command(command, name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func GetImageParts(bounds image.Rectangle, step int) ([]image.Rectangle, []image.Point) {
	rectangles := []image.Rectangle{}
	points := []image.Point{}
	for x := 0; x < bounds.Max.X; x++ {
	 		for y := 0 ; y < bounds.Max.Y; y++ {
	 			if x % step == 0 {	
	 				if y % step == 0 {
	 				maxX := x + step
	 				maxY := y + step
							if maxX <= bounds.Max.X && maxY <= bounds.Max.Y{
									 				newRect := image.Rect(x, y, maxX, maxY)
	 				newPoint := image.Point{x, y}
	 				rectangles = append(rectangles, newRect)
	 				points =  append(points, newPoint)	
							}
	 			}
	 		}	
	 	}
	}
	return rectangles, points
}