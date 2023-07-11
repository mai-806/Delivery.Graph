package handlers

import (
	"encoding/json"
	"graph/lstruct"
	"net/http"
	"time"
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
	rub := float64(cost * GetTimeValue(time.Now().String()) * OkladPerHour(cost) / 3600)
	if path != nil {
		response := lstruct.PathInfoResponse{
			CourierID: pathRequest.Courier.ID,
			Path:      path,
			Time:      int(cost),
			Cost:      rub,
		}
		SendJSONResponse(w, http.StatusOK, response)
	} else {
		response := lstruct.ErrorResponse{
			Message: "Not reachable from point destination",
		}
		SendJSONResponse(w, http.StatusNotAcceptable, response)
	}

}
