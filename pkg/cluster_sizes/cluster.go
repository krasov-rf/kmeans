package cluster

import (
	"math"
	"math/rand"
)

type Size struct {
	X float64
	Y float64
}

func CalcCentroid(dots []Size) Size {
	var x, y float64
	l := float64(len(dots))
	for _, s := range dots {
		x += s.X
		y += s.Y
	}
	return Size{
		X: x / l,
		Y: y / l,
	}
}

func Wcss(clusters map[Size][]Size) []float64 {
	var wcss float64
	var deviations, deviationsPercentages []float64

	for centroid, dots := range clusters {
		var sum float64
		for _, dot := range dots {
			sum += math.Pow(Distance(centroid, dot), 2)
		}
		wcss += sum
		deviations = append(deviations, sum)
	}

	for _, d := range deviations {
		deviationsPercentages = append(deviationsPercentages, (d/wcss)*100)
	}
	return deviationsPercentages
}

// формула расстояния
func Distance(oneDot, twoDot Size) float64 {
	return math.Sqrt(math.Pow(twoDot.X-oneDot.X, 2) + math.Pow(twoDot.Y-oneDot.Y, 2))
}

// калибровка центроид
func NewCentroids(dots, centroids []Size) map[Size][]Size {
	countCentroids := len(centroids)

	newClusters := make(map[Size][]Size, countCentroids)

	for _, dot := range dots {
		near := struct {
			Cluster Size
			D       float64
		}{
			D: -1,
		}
		for _, centroid := range centroids {
			d := Distance(centroid, dot)
			if near.D < 0 || d < near.D {
				near.D = d
				near.Cluster = centroid
			}
		}
		newClusters[near.Cluster] = append(newClusters[near.Cluster], dot)
	}

	newCentroids := make(map[Size][]Size, countCentroids)

	for _, cluster := range newClusters {
		newCentroids[CalcCentroid(cluster)] = cluster
	}

	return newCentroids
}

// ищем минимальную дистанцию до центроида
func FindMinDistance(dot Size, kDots []Size) float64 {
	minDistance := math.MaxFloat64
	for k := range kDots {
		distance := math.Pow(Distance(kDots[k], dot), 2)
		if minDistance > distance {
			minDistance = distance
		}
	}
	return minDistance
}

// выбирает начальные равноудаленные центроиды
func SelectFirstCentroids(dots []Size, countCentroids int) []Size {
	var kDots []Size
	dotsCount := len(dots)

	kDots = append(kDots, dots[RandInt(0, dotsCount)])

	for len(kDots) < countCentroids {
		var sumAllDistance float64
		minDistancesForDots := make([]float64, dotsCount)

		for d := range dots {
			minDistance := FindMinDistance(dots[d], kDots)
			minDistancesForDots[d] = minDistance
			sumAllDistance += sumAllDistance
		}

		var sumCenter float64
		rnd := rand.Float64() * sumAllDistance
		for dot, distance := range minDistancesForDots {
			sumCenter += distance
			if sumCenter > rnd {
				kDots = append(kDots, dots[dot])
				break
			}
		}
	}
	return kDots
}

func FindClusters(sizes, centroids []Size) map[Size][]Size {
	var okWcss bool
	var lastWcss, currentWcss []float64
	var res map[Size][]Size

	for {
		res = NewCentroids(sizes, centroids)
		currentWcss = Wcss(res)
		if len(lastWcss) != 0 {
			for i, w := range lastWcss {
				if currentWcss[i] != w {
					okWcss = true
					continue
				}
				okWcss = false
			}
			if !okWcss {
				break
			}
		}

		lastWcss = currentWcss
		centroids = []Size{}
		for cluster := range res {
			centroids = append(centroids, cluster)
		}
	}

	return res
}

// сумма квадрата дистанций до точки
func SumPowDistance(clusters map[Size][]Size) float64 {
	var sum float64
	for centroid, dots := range clusters {
		for _, dot := range dots {
			sum += math.Pow(Distance(centroid, dot), 2)
		}
	}
	return sum
}

func GetDistances2Dot(dots []Size, targetDot Size) []float64 {
	var res []float64
	for _, dot := range dots {
		res = append(res, Distance(dot, targetDot))
	}
	return res
}

// среднее значение
func GetMean(ditstancesToDots []float64) float64 {
	var sum float64
	for _, d := range ditstancesToDots {
		sum += d
	}
	return sum / float64(len(ditstancesToDots))
}

// дисперсия
func GetVariance(ditstancesToDots []float64, mean float64) float64 {
	var res float64
	for _, d := range ditstancesToDots {
		res += math.Pow(d-mean, 2)
	}
	return res / float64(len(ditstancesToDots)-1)
}

// стандартное отклонение
func GetStandardDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}

// z оценка
func zScore(ditstanceToDot, mean, sd float64) float64 {
	return (ditstanceToDot - mean) / sd
}

// чистим входные данные от шума
func CleanEmissions(dots []Size) []Size {
	var res []Size
	startDot := Size{
		X: 0, Y: 0,
	}

	ds := GetDistances2Dot(dots, startDot)
	mean := GetMean(ds)
	variance := GetVariance(ds, mean)
	sd := GetStandardDeviation(variance)

	for _, dot := range dots {
		score := zScore(Distance(dot, startDot), mean, sd)
		if -4 < score && score < 4 {
			res = append(res, dot)
		}
	}
	return res
}
