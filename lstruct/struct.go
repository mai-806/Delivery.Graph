package lstruct

import (
	"github.com/go-playground/validator/v10"
)

type Coordinate struct {
	Lon float64 `json:"lon" validate:"gte=-180,lte=180"`
	Lat float64 `json:"lat" validate:"gte=-90,lte=90"`
}

type Courier struct {
	ID       int        `json:"id" validate:"gte=0"`
	Position Coordinate `json:"position" validate:"required"`
}

type CourierPointID struct {
	ID      int
	PointID int
}

type Vertices map[int]Vertex

type Vertex struct {
	X      float64
	Y      float64
	Chunks []Chunk
}

type Chunk struct {
	X int
	Y int
}

type Edges map[int]map[int]float64

func ValidateCoordinate(coordinate Coordinate) error {
	validate := validator.New()

	err := validate.Struct(coordinate)
	if err != nil {
		return err
	}

	return nil
}
