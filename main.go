package main

import (
	"fmt"
	gim "github.com/ozankasikci/go-image-merge"
	"image/color"
	"image/png"
	"log"
	"os"
)

type config struct {
	Badan [][]string
	Kaki [][]string
	Tangan [][]string
	Kepala [][]string
	Mata [][]string
	Mulut [][]string
	Prop1 [][]string
	Prop2 [][]string
	Prop3 [][]string
}

func main() {
	runConfig := config{
		Badan:  [][]string{{"Badan_A_color", "Badan_A_line"}},
		Kaki:   [][]string{{"Kaki_A_color", "Kaki_A_line"}},
		Tangan: [][]string{{"Tangan_A_color", "Tangan_A_line"}},
		Kepala: [][]string{{"Kepala_A_color", "Kepala_A_line"}},
		Mata:   [][]string{{"Mata_A_color", "Mata_A_line"}, {"Mata_B_color", "Mata_B_line"}}, // every %1
		Mulut:  [][]string{{"Mulut_A_color", "Mulut_A_line"}, {"Mulut_B_line"}}, // every %2
		Prop1:  [][]string{{}, {"Prop1_A_color", "Prop1_A_line"}, {"Prop1_A_color2", "Prop1_A_line"}}, // every %4
		Prop2:  [][]string{{}, {"Prop2_A_color", "Prop2_A_line"}}, // every %12
		Prop3:  [][]string{{}, {"Prop3_A_color", "Prop3_A_line"}, {"Prop3_B_color", "Prop3_B_line"}}, // every %24
	}

	for i := 0; i < 72; i++ {
		currentImages := []string{}
		currentImages = append(currentImages, runConfig.Badan[0]...)
		currentImages = append(currentImages, runConfig.Kaki[0]...)
		currentImages = append(currentImages, runConfig.Tangan[0]...)
		currentImages = append(currentImages, runConfig.Kepala[0]...)
		currentImages = append(currentImages, runConfig.Mata[i%2]...)
		currentImages = append(currentImages, runConfig.Mulut[(i/2)%2]...)
		currentImages = append(currentImages, runConfig.Prop1[(i/4)%3]...)
		currentImages = append(currentImages, runConfig.Prop2[(i/12)%2]...)
		currentImages = append(currentImages, runConfig.Prop3[(i/24)%3]...)

		if (i/4)%3 == 0 && ((i/24)%3 != 0 || (i/12)%2 != 0) {
			continue
		}

		generateImage(currentImages, fmt.Sprintf("#%03d", i))
		log.Println(fmt.Sprintf("done #%03d", i))
	}
}

func generateImage(images []string, destFilename string) {
	insideGrids := []*gim.Grid{}
	for _, image := range images {
		insideGrids = append(insideGrids, &gim.Grid{ImageFilePath: fmt.Sprintf("./images/%s.png", image)})
	}

	grids := []*gim.Grid{
		{
			ImageFilePath: "./images/BG_A.png",
			BackgroundColor: color.RGBA{R: 255, G: 255, B: 255, A: 0},
			Grids: insideGrids,
		},
	}

	rgba, err := gim.New(grids, 1, 1).Merge()
	if nil != err {
		panic(err)
	}

	f, _ := os.Create(fmt.Sprintf("./results/%s.png", destFilename))
	err = png.Encode(f, rgba)
	if nil != err {
		panic(err)
	}
}
