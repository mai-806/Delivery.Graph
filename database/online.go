package database

import (
	"fmt"
	"graph/lstruct"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func createURL(bbox []float64) string {
	//  float64 в строки
	bboxStr := make([]string, len(bbox))
	for i, val := range bbox {
		bboxStr[i] = fmt.Sprintf("%f", val)
	}

	// Формируем строку URL
	url := fmt.Sprintf("https://overpass-api.de/api/map?bbox=%s", strings.Join(bboxStr, ","))

	return url
}

func LoadChunkOnline(chunk lstruct.Chunk, point lstruct.Coordinate, width_chunk float64, height_chunk float64, max_id int) error {
	var xmin, xmax, ymin, ymax float64
	var epsilon float64
	epsilon = 0.0001
	xmin = point.Lon + width_chunk*float64(chunk.X) - epsilon
	xmax = point.Lon + width_chunk*float64(chunk.X+1) + epsilon
	ymin = point.Lat + height_chunk*float64(chunk.Y) - epsilon
	ymax = point.Lat + height_chunk*float64(chunk.Y+1) + epsilon
	bbox := []float64{xmin, xmax, ymin, ymax}
	fileURL := createURL(bbox)
	//качает файл по ссылке
	resp, err := http.Get(fileURL)
	if err != nil {
		fmt.Printf("Ошибка при отправке GET-запроса: %v\n", err)
		return err
	}
	//defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка при получении файла. Код состояния: %v\n", resp.StatusCode)
		return err
	}

	file, err := os.Create("map")
	if err != nil {
		fmt.Printf("Ошибка при создании файла: %v\n", err)
		return err
	}
	//defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Ошибка при записи данных в файл: %v\n", err)
		return err
	}

	fmt.Println("Файл успешно скачан и сохранен.")

	// Первый скрипт

	pythonKey := "-f"
	pythonFile := "run.py"
	openmapFile := "..\\..\\map"

	pythonInterpreter := "python"

	pythonArgs := []string{pythonFile, pythonKey, openmapFile}
	directory := "online"
	err = os.Chdir(directory)
	if err != nil {
		log.Fatal(err)
	}
	directory = "osm"
	err = os.Chdir(directory)
	if err != nil {
		log.Fatal(err)
	}

	cmdPython := exec.Command(pythonInterpreter, pythonArgs...)

	cmdPython.Stdin = os.Stdin
	cmdPython.Stdout = os.Stdout
	cmdPython.Stderr = os.Stderr

	err = cmdPython.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Программа 1 на Python успешно выполнена.")
	// Егор
	FindCollisions(chunk.X, chunk.Y)
	//Второй скрипт

	pythonArgs = []string{"script.py", "chunk.x", "chunk.y", "point.Lon", "point.Lat", "width_chunk", "max_id"} //аргументы для второго скрипта

	cmdPython = exec.Command(pythonInterpreter, pythonArgs...)

	cmdPython.Stdin = os.Stdin
	cmdPython.Stdout = os.Stdout
	cmdPython.Stderr = os.Stderr

	err = cmdPython.Run() //python script.py chunk.x chunk.y point.Lon point.Lat width_chunk max_id
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Программа 2  на Python успешно выполнена.")

	//Скрипт Егора

	path := "./database/online/osm/chunks"
	err = LoadFromTextToRedis(path)
	if err != nil {
		return err
	}
	// Удаление файлов
	err = os.Remove("..pypgr")
	if err != nil {
		return err
	}

	err = os.Remove("..pypgr_names")
	if err != nil {
		return err
	}

	// Удаление папки
	err = os.RemoveAll("chunks")
	if err != nil {
		return err
	}

	// Переход на папку выше
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	// Удаление файла
	err = os.Remove("map")
	if err != nil {
		return err
	}
	return nil
}
