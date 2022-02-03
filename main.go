package main

import (
	"fmt"
	gim "github.com/ozankasikci/go-image-merge"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

type config struct {
	BGWarna []string
	BGProp []string
	Kursi []string
	Badan []string
	Baju []string
	AksBadan []string
	AksBadanB []string
	Kaki []string
	Tangan []string
	Kepala []string
	Mata []string
	Mulut []string
	AksKepala []string
	Meja []string
	PropKanan []string
	PropKiri []string
}

type chance struct {
	BGWarna []int
	BGProp []int
	Kursi []int
	Badan []int
	Baju []int
	AksBadan []int
	AksBadanB []int
	Kaki []int
	Tangan []int
	Kepala []int
	Mata []int
	Mulut []int
	AksKepala []int
	Meja []int
	PropKanan []int
	PropKiri []int
}

// Prop Kanan 5
// Prop Kiri 5
// Aksesoris Kepala 3
// Mata 6 Wajib
// Mulut 5 Wajib
// Tangan 2 Wajib
// Kepala 1 Wajib
// Meja 3
// Kaki 1 Wajib
// Aksesoris Badan 5
// Baju 4
// Badan 1 Wajib
// Kursi 2
// BG Prop 4
// BG Warna 3+clear

// BG Warna - BG Prop - Kursi - Badan - Baju - Aksesoris Badan - Kaki - Tangan - Kepala - Mulut - Mata - Aksesoris Kepala - Meja - Prop Kanan - Prop Kiri

func main() {
	rand.Seed(time.Now().UnixNano())

	runConfig := config{
		BGWarna: []string{"", "Cat_0049_BG-Warna-1", "Cat_0048_BG-Warna-2"},
		BGProp: []string{"", "Cat_0047_BG-Prop-1", "Cat_0046_BG-Prop-2", "Cat_0045_BG-Prop-3", "Cat_0044_BG-Prop-4"},
		Kursi: []string{"", "Cat_0042_Kursi-2", "Cat_0041_Kursi-4"},
		Badan: []string{"Cat_0040_Badan-1"},
		Baju: []string{"", "Cat_0038_Baju-2", "Cat_0037_Baju-4", "Cat_0036_Baju-6", "Cat_0035_Baju-5"},
		AksBadan: []string{"", "Cat_0034_Aksesoris-Badan-1", "Cat_0033_Aksesoris-Badan-3", "Cat_0032_Aksesoris-Badan-5", "Cat_0031_Aksesoris-Badan-6"},
		AksBadanB: []string{"", "Cat_0030_Aksesoris-Badan-B-1"},
		Kaki: []string{"Cat_0029_Kaki-1"},
		Tangan: []string{"Cat_0024_Tangan-1", "Cat_0023_Tangan-2"},
		Kepala: []string{"Cat_0025_Kepala-1"},
		Mata: []string{"Cat_0017_Mata-1", "Cat_0016_Mata-2", "Cat_0015_Mata-3", "Cat_0014_Mata-4", "Cat_0013_Mata-6", "Cat_0012_Mata-7"},
		Mulut: []string{"Cat_0022_Mulut-1", "Cat_0021_Mulut-2", "Cat_0020_Mulut-3", "Cat_0019_Mulut-4", "Cat_0018_Mulut-5"},
		AksKepala: []string{"", "Cat_0011_Aksesoris-Kepala-2", "Cat_0010_Aksesoris-Kepala-3"},
		Meja: []string{"Cat_0028_Meja-3", "Cat_0027_Meja-2", "Cat_0026_Meja-1"},
		PropKanan: []string{"", "Cat_0004_Prop-Kanan-1", "Cat_0003_Prop-Kanan-2", "Cat_0002_Prop-Kanan-3", "Cat_0001_Prop-Kanan-4", "Cat_0000_Prop-Kanan-5"},
		PropKiri: []string{"", "Cat_0009_Prop-Kiri-1", "Cat_0008_Prop-Kiri-2", "Cat_0007_Prop-Kiri-3", "Cat_0006_Prop-Kiri-4", "Cat_0005_Prop-Kiri-5"},
	}

	bodRunChance := chance{
		BGWarna: []int{5, 40, 100},
		BGProp: []int{10, 20, 40, 70, 100},
		Kursi: []int{10, 55, 100},
		Badan: []int{100},
		Baju: []int{10, 25, 55, 80, 100},
		AksBadan: []int{25, 25, 50, 75, 100},
		AksBadanB: []int{50, 100},
		Kaki: []int{100},
		Tangan: []int{50, 100},
		Kepala: []int{100},
		Mata: []int{15, 30, 45, 60, 75, 100},
		Mulut: []int{15, 30, 45, 60, 100},
		AksKepala: []int{60, 80, 100},
		Meja: []int{30, 60, 100},
		PropKanan: []int{20, 40, 60, 80, 100, 101},
		PropKiri: []int{15, 30, 45, 60, 75, 100},
	}

	managerRunChance := chance{
		BGWarna: []int{20, 60, 100},
		BGProp: []int{30, 40, 60, 90, 100},
		Kursi: []int{30, 70, 100},
		Badan: []int{100},
		Baju: []int{30, 50, 100, 100, 100},
		AksBadan: []int{10, 20, 45, 70, 100},
		AksBadanB: []int{100, 100},
		Kaki: []int{100},
		Tangan: []int{60, 100},
		Kepala: []int{100},
		Mata: []int{25, 40, 55, 70, 85, 100},
		Mulut: []int{20, 40, 60, 100, 100},
		AksKepala: []int{70, 85, 100},
		Meja: []int{35, 70, 100},
		PropKanan: []int{20, 40, 60, 80, 100, 101},
		PropKiri: []int{15, 30, 45, 60, 100, 100},
	}

	staffRunChance := chance{
		BGWarna: []int{40, 70, 100},
		BGProp: []int{40, 60, 80, 100, 100},
		Kursi: []int{50, 100, 100},
		Badan: []int{100},
		Baju: []int{35, 65, 100, 100, 100},
		AksBadan: []int{20, 60, 60, 60, 100},
		AksBadanB: []int{100, 100},
		Kaki: []int{100},
		Tangan: []int{70, 100},
		Kepala: []int{100},
		Mata: []int{15, 30, 45, 60, 100, 100},
		Mulut: []int{35, 60, 85, 100, 100},
		AksKepala: []int{80, 90, 100},
		Meja: []int{40, 80, 100},
		PropKanan: []int{40, 55, 70, 85, 100, 101},
		PropKiri: []int{25, 50, 75, 100, 100, 100},
	}

	internRunChance := chance{
		BGWarna: []int{60, 80, 100},
		BGProp: []int{55, 70, 85, 100, 100},
		Kursi: []int{80, 100, 100},
		Badan: []int{100},
		Baju: []int{50, 100, 100, 100, 100},
		AksBadan: []int{40, 100, 100, 100, 100},
		AksBadanB: []int{100, 100},
		Kaki: []int{100},
		Tangan: []int{90, 100},
		Kepala: []int{100},
		Mata: []int{25, 50, 75, 100, 100, 100},
		Mulut: []int{55, 70, 85, 100, 100},
		AksKepala: []int{90, 95, 100},
		Meja: []int{50, 100, 100},
		PropKanan: []int{60, 70, 80, 90, 100, 101},
		PropKiri: []int{40, 60, 80, 100, 100, 100},
	}

	exists := make(map[int64]bool, 0)
	combinations := [][]int{}
	for i := 0; i < 200; i++ {
		runChance := internRunChance
		if i < 20 {
			runChance = bodRunChance
		} else if i < 50 {
			runChance = managerRunChance
		} else if i < 100 {
			runChance = staffRunChance
		}

		currentCombination := []int{
			getRandomIndex(runChance.BGWarna),
			getRandomIndex(runChance.BGProp),
			getRandomIndex(runChance.Kursi),
			0,
			getRandomIndex(runChance.Baju),
			getRandomIndex(runChance.AksBadan),
			getRandomIndex(runChance.AksBadanB),
			0,
			getRandomIndex(runChance.Tangan),
			0,
			getRandomIndex(runChance.Mata),
			getRandomIndex(runChance.Mulut),
			getRandomIndex(runChance.AksKepala),
			getRandomIndex(runChance.Meja),
			getRandomIndex(runChance.PropKanan),
			getRandomIndex(runChance.PropKiri),
		}

		exist := getExistIndex(currentCombination)
		if _, ok := exists[exist]; ok {
			i--
			continue
		} else {
			exists[exist] = true
		}

		currentImages := []string{}
		currentImages = append(currentImages, runConfig.BGWarna[currentCombination[0]])
		currentImages = append(currentImages, runConfig.BGProp[currentCombination[1]])
		currentImages = append(currentImages, runConfig.Kursi[currentCombination[2]])
		currentImages = append(currentImages, runConfig.Badan[currentCombination[3]])
		currentImages = append(currentImages, runConfig.Baju[currentCombination[4]])
		currentImages = append(currentImages, runConfig.AksBadan[currentCombination[5]])
		currentImages = append(currentImages, runConfig.AksBadanB[currentCombination[6]])
		currentImages = append(currentImages, runConfig.Kaki[currentCombination[7]])
		currentImages = append(currentImages, runConfig.Tangan[currentCombination[8]])
		currentImages = append(currentImages, runConfig.Kepala[currentCombination[9]])
		currentImages = append(currentImages, runConfig.Mata[currentCombination[10]])
		currentImages = append(currentImages, runConfig.Mulut[currentCombination[11]])
		currentImages = append(currentImages, runConfig.AksKepala[currentCombination[12]])
		currentImages = append(currentImages, runConfig.Meja[currentCombination[13]])
		currentImages = append(currentImages, runConfig.PropKanan[currentCombination[14]])
		currentImages = append(currentImages, runConfig.PropKiri[currentCombination[15]])

		generateImage(currentImages, fmt.Sprintf("#%04d", i+1))
		log.Printf("image #%04d generated\n", i+1)

		combinations = append(combinations, currentCombination)
	}
}

func getExistIndex(combination []int) int64 {
	exist := int64(0)
	length := len(combination)
	for i := 0; i < length; i++ {
		exist += int64(combination[i]) * int64(math.Pow(10, float64(14-i)))
	}
	return exist
}

func getRandomIndex(chances []int) int {
	rnd := rand.Intn(100)
	length := len(chances)
	for i := 0; i < length; i++ {
		if rnd <= chances[i] {
			return i
		}
	}
	return 0
}

func generateImage(images []string, destFilename string) {
	insideGrids := []*gim.Grid{}
	for _, image := range images {
		if image == "" {
			continue
		}

		insideGrids = append(insideGrids, &gim.Grid{ImageFilePath: fmt.Sprintf("./images/%s.png", image)})
	}

	grids := []*gim.Grid{
		{
			ImageFilePath: "./images/Cat_0999_BG.png",
			BackgroundColor: color.White,
			Grids: insideGrids,
		},
	}

	rgba, err := gim.New(grids, 1, 1).Merge()
	if nil != err {
		panic(err)
	}

	f, err := os.Create(fmt.Sprintf("./results/%s.png", destFilename))
	if nil != err {
		panic(err)
	}

	err = png.Encode(f, rgba)
	if nil != err {
		panic(err)
	}
}
