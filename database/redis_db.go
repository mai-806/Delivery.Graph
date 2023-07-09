package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"graph/lstruct"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func ConnectRedisDB() error {
	// Создание клиента Redis
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Адрес и порт Redis сервера
		Password: "",               // Пароль (если требуется)
		DB:       0,                // Индекс базы данных
	})
	// Проверка соединения с Redis
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	log.Println("Connected to Redis:", pong)

	return nil
}

func GetRedis(key string) (string, error) {
	// Получение значения из Redis по ключу
	result, err := client.Get(context.Background(), key).Result()
	return result, err
}

func SetRedis(key string, value string, expiration time.Duration) error {
	// Установка значения в Redis с указанием срока жизни
	err := client.Set(context.Background(), key, value, expiration).Err()
	return err
}

func AddEdgesToRedis(x int, y int, data lstruct.Edges) error {
	key := fmt.Sprintf("%d%d", x, y)
	hash := make(map[string]interface{})
	SelectRedis(1)
	for k, v := range data {
		EdgesBytes, err := json.Marshal(v)
		if err != nil {
			return err
		}

		hash[fmt.Sprintf("%d", k)] = string(EdgesBytes)
	}

	err := client.HMSet(context.Background(), key, hash).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetEdgesRedis(x int, y int, data *lstruct.Edges) error {
	key := fmt.Sprintf("%d%d", x, y)
	SelectRedis(1)
	result, err := client.HGetAll(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if len(result) == 0 {
		return errors.New("Key not found")
	}

	for k, v := range result {
		keyInt, err := strconv.Atoi(k)
		if err != nil {
			return err
		}
		var final_id map[int]float64
		err = json.Unmarshal([]byte(v), &final_id)
		if err != nil {
			return err
		}
		for key_fin, value_fin := range final_id {
			if (*data)[keyInt] == nil {
				(*data)[keyInt] = make(map[int]float64)
			}

			(*data)[keyInt][key_fin] = value_fin
		}

	}

	return nil
}

func AddVerticesToRedis(x int, y int, data lstruct.Vertices) error {
	key := fmt.Sprintf("%d%d", x, y)
	hash := make(map[string]interface{})

	for k, v := range data {
		vertexBytes, err := json.Marshal(v)
		if err != nil {
			return err
		}

		hash[fmt.Sprintf("%d", k)] = string(vertexBytes)
	}
	SelectRedis(0)
	err := client.HMSet(context.Background(), key, hash).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetVerticesRedis(x int, y int, data *lstruct.Vertices) error {
	key := fmt.Sprintf("%d%d", x, y)
	SelectRedis(0)
	result, err := client.HGetAll(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if len(result) == 0 {
		return errors.New("Key not found")
	}

	for k, v := range result {
		keyInt, err := strconv.Atoi(k)
		if err != nil {
			return err
		}
		var vertex lstruct.Vertex
		err = json.Unmarshal([]byte(v), &vertex)
		if err != nil {
			return err
		}

		(*data)[keyInt] = vertex
	}

	return nil
}

func SelectRedis(index int) error {
	// Установка таблицы в Redis
	err := client.Do(context.Background(), "SELECT", index).Err()
	return err
}

func EraseAllTablesRedis() error {
	// Удаление всех таблиц
	err := client.FlushAll(context.Background()).Err()
	return err
}

func GetChunk(point lstruct.Coordinate) lstruct.Chunk {
	//ЗАГЛУШКА
	return lstruct.Chunk{X: 0, Y: 0}
}

func CloseRedisDB() error {
	if client == nil {
		return nil // Проверяем, что клиент Redis был инициализирован
	}
	// Закрытие соединения с Redis
	err := client.Close()
	if err != nil {
		return err
	}
	log.Println("Redis connection closed")
	return nil
}
