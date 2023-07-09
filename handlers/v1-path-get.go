package handlers

import (
	"encoding/json"
	"graph/lstruct"
	"net/http"
)

func GetV1Path(w http.ResponseWriter, r *http.Request) {
	var pathRequest lstruct.PathRequest
	err := json.NewDecoder(r.Body).Decode(&pathRequest)
	if err != nil {
		errorResponse := lstruct.ErrorResponse{
			Message: "Bad Input",
		}
		SendJSONResponse(w, http.StatusBadRequest, errorResponse)
		return
	}

	err = lstruct.ValidatePath(pathRequest)
	if err != nil {
		errorResponse := lstruct.ErrorResponse{
			Message: err.Error(),
		}
		SendJSONResponse(w, http.StatusBadRequest, errorResponse)
		return
	}

	vertices := lstruct.Vertices{}
	edges := lstruct.Edges{}
	chunks := map[lstruct.Chunk]bool{}

	path, cost := findPath(pathRequest.Courier.Position, pathRequest.EndCoordinate, &vertices, &edges, &chunks)

	// Создание и отправка ответа
	if path != nil {
		response := lstruct.PathInfoResponse{
			CourierID: pathRequest.Courier.ID,
			Path:      path,
			Time:      int(cost),
			Cost:      cost,
		}
		SendJSONResponse(w, http.StatusOK, response)
	} else {
		response := lstruct.ErrorResponse{
			Message: "Not reachable from point destination",
		}
		SendJSONResponse(w, http.StatusNotAcceptable, response)
	}

}
