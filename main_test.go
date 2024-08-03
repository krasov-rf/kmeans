package kmeans

import (
	"testing"
)

func BenchmarkElbowRandClusters(b *testing.B) {
	countClusrters := 20
	res := make(map[float64]float64, countClusrters)
	Dots := GenerateDots(1000)
	Dots = CleanEmissions(Dots)

	b.StartTimer()

	for i := 1; i < countClusrters; i++ {
		centroids := SelectFirstCentroids(Dots, i)
		clusters := FindClusters(Dots, centroids)

		res[float64(i)] = SumPowDistance(clusters)
	}

	b.StopTimer()

	PaintElbow(res, 4)
}

func BenchmarkKmeansPlus(b *testing.B) {
	countClusters := 10
	Dots := GenerateDots(10000)
	Dots = CleanEmissions(Dots)

	b.StartTimer()

	centroids := SelectFirstCentroidsPlus(Dots, countClusters)
	clusters := FindClusters(Dots, centroids)

	b.StopTimer()

	PaintClusters(clusters, 10)
}

func BenchmarkKmeans(b *testing.B) {
	countClusters := 10
	Dots := GenerateDots(10000)
	Dots = CleanEmissions(Dots)

	centroids := SelectFirstCentroids(Dots, countClusters)

	b.StartTimer()

	clusters := FindClusters(Dots, centroids)

	b.StopTimer()

	PaintClusters(clusters, 10)
}
