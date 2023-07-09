package handlers

import (
	"encoding/json"
	"graph/lstruct"
	"net/http"
)

func GetV1PointIsAvailable(w http.ResponseWriter, r *http.Request) {
	// Чтение и декодирование тела запроса в формате JSON
	var coordinate lstruct.Coordinate
	err := json.NewDecoder(r.Body).Decode(&coordinate)
	if err != nil {
		errorResponse := lstruct.ErrorResponse{
			Message: "Bad Input",
		}
		SendJSONResponse(w, http.StatusBadRequest, errorResponse)
		return
	}

	err = lstruct.ValidateCoordinate(coordinate)
	if err != nil {
		errorResponse := lstruct.ErrorResponse{
			Message: err.Error(),
		}
		SendJSONResponse(w, http.StatusBadRequest, errorResponse)
		return
	}

	// Проверка принадлежности координат диапазону
	available := coordinate.Lon >= -90 && coordinate.Lon <= 90 && coordinate.Lat >= -90 && coordinate.Lat <= 90

	// Создание и отправка ответа
	response := lstruct.PointAvailableResponse{
		Available: available,
	}

	SendJSONResponse(w, http.StatusOK, response)
}
