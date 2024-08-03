package kmeans

import (
	"testing"
)

func BenchmarkElbowRandClusters(b *testing.B) {
	countClusrters := 20
	res := make(map[float64]float64, countClusrters)
	dots := GenerateDots(1000)
	dots = CleanEmissions(dots)

	b.StartTimer()

	for i := 1; i < countClusrters; i++ {
		centroids := SelectFirstCentroids(dots, i)
		clusters := FindClusters(dots, centroids)

		res[float64(i)] = SumPowDistance(clusters)
	}

	b.StopTimer()

	PaintElbow(res, 4)
}

func BenchmarkKmeansPlus(b *testing.B) {
	countClusters := 10
	dots := GenerateDots(10000)
	dots = CleanEmissions(dots)

	b.StartTimer()

	centroids := SelectFirstCentroidsPlus(dots, countClusters)
	clusters := FindClusters(dots, centroids)

	b.StopTimer()

	PaintClusters(clusters, 10)
}

func BenchmarkKmeans(b *testing.B) {
	countClusters := 10
	dots := GenerateDots(10000)
	dots = CleanEmissions(dots)

	centroids := SelectFirstCentroids(dots, countClusters)

	b.StartTimer()

	clusters := FindClusters(dots, centroids)

	b.StopTimer()

	PaintClusters(clusters, 10)
}
