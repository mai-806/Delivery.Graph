package handlers

import (
	"encoding/json"
	"graph/database"
	"graph/lstruct"
	"net/http"
)

func GetV1SecretLoadDatabase(w http.ResponseWriter, r *http.Request) {
	var request lstruct.ErrorResponse
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errorResponse := lstruct.ErrorResponse{
			Message: "Bad Input",
		}
		SendJSONResponse(w, http.StatusBadRequest, errorResponse)
		return
	}
	if request.Message != "iamgay" && request.Message != "iamgayestgay" {
		errorResponse := lstruct.ErrorResponse{
			Message: "Wrong message",
		}
		SendJSONResponse(w, http.StatusOK, errorResponse)
		return
	}
	// Создание и отправка ответа
	if request.Message == "iamgayestgay" {
		database.EraseAllTablesRedis()
		response := lstruct.ErrorResponse{
			Message: "All tables cleared successfully",
		}
		SendJSONResponse(w, http.StatusOK, response)
		return
	}
	if request.Message == "iamgay" {
		vertices := lstruct.Vertices{
			0: {X: 0.0, Y: 0.0, Chunks: []lstruct.Chunk{{X: 0, Y: 0}}},
			1: {X: 1.0, Y: 1.0, Chunks: []lstruct.Chunk{{X: 0, Y: 0}}},
			2: {X: 2.0, Y: 2.0, Chunks: []lstruct.Chunk{{X: 0, Y: 0}}},
			3: {X: 3.0, Y: 3.0, Chunks: []lstruct.Chunk{{X: 0, Y: 0}}},
		}

		edges := lstruct.Edges{
			0: {1: 3.0, 2: 5.0},
			1: {2: 3.0, 3: 6.0},
			2: {3: 1.0},
			3: {},
		}
		database.AddEdgesToRedis(0, 0, edges)
		database.AddVerticesToRedis(0, 0, vertices)
		response := lstruct.ErrorResponse{
			Message: "Example loaded successfully",
		}
		SendJSONResponse(w, http.StatusOK, response)
		return
	}

}
