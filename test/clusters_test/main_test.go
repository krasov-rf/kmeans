package test

import (
	"testing"

	cluster "github.com/kras/kMeans/pkg/cluster_sizes"
)

func BenchmarkElbowRandClusters(b *testing.B) {
	countClusrters := 20
	res := make(map[float64]float64, countClusrters)
	sizes := cluster.GenerateSizes(1000)
	sizes = cluster.CleanEmissions(sizes)

	b.StartTimer()

	for i := 1; i < countClusrters; i++ {
		centroids := cluster.SelectFirstCentroids(sizes, i)
		clusters := cluster.FindClusters(sizes, centroids)

		res[float64(i)] = cluster.SumPowDistance(clusters)
	}

	b.StopTimer()

	cluster.PaintElbow(res, 4)
}

func BenchmarkFindRandClusters(b *testing.B) {
	countClusrters := 10
	sizes := cluster.GenerateSizes(1000)
	sizes = cluster.CleanEmissions(sizes)

	b.StartTimer()

	centroids := cluster.SelectFirstCentroids(sizes, countClusrters)
	clusters := cluster.FindClusters(sizes, centroids)

	b.StopTimer()

	cluster.PaintClusters(clusters, 10)
}
