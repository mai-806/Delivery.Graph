package lstruct

import "github.com/go-playground/validator/v10"

type PathRequest struct {
	Courier       Courier    `json:"courier" validate:"required"`
	EndCoordinate Coordinate `json:"end_coordinate" validate:"required"`
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

func ValidatePathMultiple(pathMultipleStartRequest PathMultipleStartRequest) error {
	validate := validator.New()

	err := validate.Struct(pathMultipleStartRequest)
	if err != nil {
		return err
	}

	return nil
}
