package kmeans

import (
	"image/color"
	"log"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandFloats(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func GenerateDots(count_dot int) []Dot {
	Dots := []Dot{}

	for i := count_dot; i != 0; i-- {
		Dots = append(Dots, Dot{X: RandFloats(1, 1000), Y: RandFloats(1, 1000)})
	}
	return Dots
}

func PaintClusters(clusters map[Dot][]Dot, inch font.Length) {
	p := plot.New()
	for _, c := range clusters {
		var xys plotter.XYs

		for _, i := range c {
			xys = append(xys, plotter.XY{X: i.X, Y: i.Y})
		}

		s, err := plotter.NewScatter(xys)
		if err != nil {
			log.Fatalf("Ошибка при добавлении точек: %v", err)
		}
		s.Color = color.RGBA{
			R: uint8(RandInt(0, 255)),
			G: uint8(RandInt(0, 255)),
			B: uint8(RandInt(0, 255)),
			A: 255,
		}
		p.Add(s)
	}

	if err := p.Save(inch*vg.Inch, inch*vg.Inch, "points.png"); err != nil {
		log.Fatalf("Ошибка при сохранении файла: %v", err)
	}
}

// нарисовать локоть
func PaintElbow(dots map[float64]float64, inch font.Length) {
	var xys plotter.XYs
	p := plot.New()

	for x, y := range dots {
		xys = append(xys, plotter.XY{X: x, Y: y})

		s, err := plotter.NewScatter(xys)
		if err != nil {
			log.Fatalf("Ошибка при добавлении точек: %v", err)
		}
		p.Add(s)
	}

	if err := p.Save(inch*vg.Inch, inch*vg.Inch, "elbow.png"); err != nil {
		log.Fatalf("Ошибка при сохранении файла: %v", err)
	}
}
