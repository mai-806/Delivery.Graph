package lstruct

import "github.com/go-playground/validator/v10"

type PathRequest struct {
	Courier       Courier    `json:"courier" validate:"required"`
	EndCoordinate Coordinate `json:"end_coordinate" validate:"required"`
}

type PathRequestV2 struct {
	Courier       Courier    `json:"courier" validate:"required"`
	EndCoordinate Coordinate `json:"end_coordinate" validate:"required"`
	Time          string     `json:"time" validate:"required"`
}

type PathMultipleStartRequest struct {
	Couriers      []Courier  `json:"couriers"`
	EndCoordinate Coordinate `json:"end_coordinate"`
}

func ValidatePath(pathRequest PathRequest) error {
	validate := validator.New()

	err := validate.Struct(pathRequest)
	if err != nil {
		return err
	}

	return nil
}

func ValidatePathV2(pathRequest PathRequestV2) error {
	validate := validator.New()

	err := validate.Struct(pathRequest)
	if err != nil {
		return err
	}

	return nil
}

func ValidatePathMultiple(pathMultipleStartRequest PathMultipleStartRequest) error {
	validate := validator.New()

	err := validate.Struct(pathMultipleStartRequest)
	if err != nil {
		return err
	}

	return nil
}
