# Delivery.Graph
## Микросервис по поиску кратчайших путей

Проект по практике студентов М8О 206 Б-21 группы МАИ 

Общий проект https://github.com/mai-806

Запуск

`redis-server`

`go run main.go`

---
### Алгоритмы

![Astar](https://mchow.com/posts/r-and-astar-cats-to-dogs/)

Для поиска в графе используется модифицированный алгоритм A*. Модификация заключается в возможности работы с частично загруженным графом (в нашем случае чанки).