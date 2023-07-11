package database

import (
	"bufio"
	"fmt"
	"graph/lstruct"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var files []string

func Visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Ошибка при посещении пути: %v\n", err)
		return nil
	}
	if info.IsDir() {
		return nil
	}
	files = append(files, path)
	return nil
}

func FileName(path string) string {
	_, filename := filepath.Split(path)
	return filename[:len(filename)-4]
}

func ParseStructure(structure string) (string, int, int, error) {
	// Разделяем структуру по разделителю "_"
	structureParts := strings.Split(structure, "_")

	// Проверяем, что структура содержит 3 части
	if len(structureParts) != 3 {
		return "", 0, 0, fmt.Errorf("неправильная структура")
	}

	// Получаем char
	char := structureParts[0]

	// Получаем первое int значение
	int1, err := strconv.Atoi(structureParts[1])
	if err != nil {
		return "", 0, 0, fmt.Errorf("неправильное значение int: %v", err)
	}

	// Получаем второе int значение
	int2, err := strconv.Atoi(structureParts[2])
	if err != nil {
		return "", 0, 0, fmt.Errorf("неправильное значение int: %v", err)
	}

	return char, int1, int2, nil
}

func parseFileEdges(filename string) (lstruct.Edges, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	var data lstruct.Edges = make(map[int]map[int]float64)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 3 {
			return nil, fmt.Errorf("incorrect number of fields in line")
		}

		idStart, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}
		if data[idStart] == nil {
			data[idStart] = make(map[int]float64)
		}

		idFinish, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}

		weight, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			return nil, err
		}
		data[idStart][idFinish] = weight
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func parseFileVertices(filename string) (lstruct.Vertices, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	var data lstruct.Vertices = make(map[int]lstruct.Vertex)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		id, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}

		x, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			return nil, err
		}

		y, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return nil, err
		}

		var chunks []lstruct.Chunk
		for j := 3; j < len(fields); j += 2 {
			X, err := strconv.Atoi(fields[j])
			if err != nil {
				return nil, err
			}

			Y, err := strconv.Atoi(fields[j+1])
			if err != nil {
				return nil, err
			}

			chunks = append(chunks, lstruct.Chunk{X: X, Y: Y})
		}

		data[id] = lstruct.Vertex{X: x, Y: y, Chunks: chunks}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func LoadFromTextToRedis(path string) error {
	root := path
	err := filepath.Walk(root, Visit)
	if err != nil {
		return err
	}

	for _, file := range files {
		filename := FileName(file)
		typeFile, x, y, err := ParseStructure(filename)
		if err != nil {
			return err
		}

		if typeFile == "E" {
			data, err := parseFileEdges(file)
			if err != nil {
				return err
			}
			AddEdgesToRedis(x, y, data)
		} else {
			data, err := parseFileVertices(file)
			if err != nil {
				return err
			}
			AddVerticesToRedis(x, y, data)
		}
	}

	return err
}
