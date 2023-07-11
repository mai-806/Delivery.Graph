package database

import (
	"fmt"
	"graph/lstruct"
	"os"
)

func FindCollisions(x int, y int) error {
	var vertices lstruct.Vertices
	GetVerticesRedis(x+1, y, &vertices)
	GetVerticesRedis(x, y+1, &vertices)
	GetVerticesRedis(x-1, y, &vertices)
	GetVerticesRedis(x, y-1, &vertices)
	curChunk := lstruct.Chunk{X: x, Y: y}
	file, err := os.Create("./database/online/osm/ongrid.txt")

	for id := range vertices {
		vertex := vertices[id]
		if len(vertex.Chunks) > 1 {
			for chunk := range vertex.Chunks {
				if vertex.Chunks[chunk] == curChunk {
					str := fmt.Sprintf("%d %f %f", id, vertex.X, vertex.Y)
					for vChunk := range vertex.Chunks {
						str += fmt.Sprintf("%d %d", vertex.Chunks[vChunk].X, vertex.Chunks[vChunk].Y)
					}
					_, err = file.WriteString(str)
				}
			}
		}
	}
	return err
}
