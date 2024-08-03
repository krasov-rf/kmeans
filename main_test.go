package kmeans

import (
	"testing"
)

func BenchmarkElbowRandClusters(b *testing.B) {
	countClusrters := 20
	res := make(map[float64]float64, countClusrters)
	sizes := GenerateSizes(1000)
	sizes = CleanEmissions(sizes)

	b.StartTimer()

	for i := 1; i < countClusrters; i++ {
		centroids := SelectFirstCentroids(sizes, i)
		clusters := FindClusters(sizes, centroids)

		res[float64(i)] = SumPowDistance(clusters)
	}

	b.StopTimer()

	PaintElbow(res, 4)
}

func BenchmarkKmeansPlus(b *testing.B) {
	countClusrters := 10
	sizes := GenerateSizes(10000)
	sizes = CleanEmissions(sizes)

	b.StartTimer()

	centroids := SelectFirstCentroidsPlus(sizes, countClusrters)
	clusters := FindClusters(sizes, centroids)

	b.StopTimer()

	PaintClusters(clusters, 10)
}

func BenchmarkKmeans(b *testing.B) {
	countClusrters := 10
	sizes := GenerateSizes(10000)
	sizes = CleanEmissions(sizes)

	centroids := SelectFirstCentroids(sizes, countClusrters)

	b.StartTimer()

	clusters := FindClusters(sizes, centroids)

	b.StopTimer()

	PaintClusters(clusters, 10)
}
