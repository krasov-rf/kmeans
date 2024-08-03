# Алгоритм K-means++

Пример использования

```
import (
  "github.com/krasov-rf/kmeans"
)

dots := GenerateDots(10000) // Генерация точек, Dot{X float64; Y float64}
dots = CleanEmissions(dots) // Чистим шум в датасете
centroids := SelectFirstCentroids(dots, countClusters) // K-means++, выбор первоначаьльных точек для работы
clusters := FindClusters(dots, centroids) // Сам поиск
```

 Статья на хабре: https://habr.com/ru/articles/829202/
