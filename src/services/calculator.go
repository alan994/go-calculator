package services

import (
	"example/go-calculator/internal"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)


type CalculationResponse struct {
	Action 	string 		`json:"action"`
	X 		float64 	`json:"x"`
	Y 		float64		`json:"y"`
	Answer 	float64		`json:"answer"`
	Cached 	bool		`json:"cached"`
}

func getRedisClient() (*redis.Client, error) {
	dbName, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	host := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")

	if len(host) == 0 {
		return nil, internal.NewErrorf("Redis host is undefined")
	}

	return redis.NewClient(&redis.Options{
        Addr: host,
        Password: password,
        DB: dbName,
    }), nil
}

func Add(x float64, y float64) (*CalculationResponse, error) {

	cacheKey := fmt.Sprintf("ADD|%f|%f", x, y)
	usedValueFromCache := true

	client, error := getRedisClient()
	if error != nil{
		return nil, internal.WrapErrorf(error, "Connecting to redis failed")
	}

	var result float64;

	cacheResult, err := client.Get(cacheKey).Result()
    if err != nil {
		log.Println("Info: ", fmt.Sprintf("Key %s is not available in cache", cacheKey))
		usedValueFromCache = false
    }

	if !usedValueFromCache {
		result = x + y

		err := client.Set(cacheKey, result, 1*time.Minute).Err()
		if err != nil {
			log.Println("Error: Failed to save result to cache", err)
		}
	} else {
		parsedCacheValue, err := strconv.ParseFloat(cacheResult, 64)
		if err != nil {
			log.Println("Error: ", "Failed to parse cached value")
			return nil, internal.WrapErrorf(err, "Communication with cache failed")
		}
		result = parsedCacheValue
	}

	var response = new (CalculationResponse)
	response.Action = "Add"
	response.Answer = result
	response.Cached = usedValueFromCache
	response.X = x
	response.Y = y

	return response, nil
}


func Subtract(x float64, y float64) (*CalculationResponse, error) {
	cacheKey := fmt.Sprintf("SUBTRACT|%f|%f", x, y)
	usedValueFromCache := true

	client, error := getRedisClient()
	if error != nil{
		return nil, internal.WrapErrorf(error, "Connecting to redis failed")
	}

	var result float64;

	cacheResult, err := client.Get(cacheKey).Result()
    if err != nil {
		log.Println("Info: ", fmt.Sprintf("Key %s is not available in cache", cacheKey))
		usedValueFromCache = false
    }

	if !usedValueFromCache {
		result = x - y
		err := client.Set(cacheKey, result, 1*time.Minute).Err()
		if err != nil {
			log.Println("Error: Failed to save result to cache", err)
		}
	} else {
		parsedCacheValue, err := strconv.ParseFloat(cacheResult, 64)
		if err != nil {
			log.Println("Error: ", "Failed to parse cached value")
			return nil, internal.WrapErrorf(err, "Communication with cache failed")
		}
		result = parsedCacheValue
	}

	var response = new (CalculationResponse)
	response.Action = "subtract"
	response.Answer = result
	response.Cached = usedValueFromCache
	response.X = x
	response.Y = y

	return response, nil
}

func Multiply(x float64, y float64) (*CalculationResponse, error) {
	cacheKey := fmt.Sprintf("MULTIPLY|%f|%f", x, y)
	usedValueFromCache := true

	client, error := getRedisClient()
	if error != nil{
		return nil, internal.WrapErrorf(error, "Connecting to redis failed")
	}

	var result float64;

	cacheResult, err := client.Get(cacheKey).Result()
    if err != nil {
		log.Println("Info: ", fmt.Sprintf("Key %s is not available in cache", cacheKey))
		usedValueFromCache = false
    }

	if !usedValueFromCache {
		result = x * y
		err := client.Set(cacheKey, result, 1*time.Minute).Err()
		if err != nil {
			log.Println("Error: Failed to save result to cache", err)
		}
	} else {
		parsedCacheValue, err := strconv.ParseFloat(cacheResult, 64)
		if err != nil {
			log.Println("Error: ", "Failed to parse cached value")
			return nil, internal.WrapErrorf(err, "Communication with cache failed")
		}
		result = parsedCacheValue
	}

	var response = new (CalculationResponse)
	response.Action = "multiply"
	response.Answer = result
	response.Cached = usedValueFromCache
	response.X = x
	response.Y = y

	return response, nil
}

func Divide(x float64, y float64) (*CalculationResponse, error) {
	if y == 0{
		log.Println("Error: Y can't be 0")
		return nil, internal.NewErrorf("Y can't be 0")
	}

	cacheKey := fmt.Sprintf("DIVIDE|%f|%f", x, y)
	usedValueFromCache := true

	client, error := getRedisClient()
	if error != nil{
		return nil, internal.WrapErrorf(error, "Connecting to redis failed")
	}

	var result float64;

	cacheResult, err := client.Get(cacheKey).Result()
    if err != nil {
		log.Println("Info: ", fmt.Sprintf("Key %s is not available in cache", cacheKey))
		usedValueFromCache = false
    }

	if !usedValueFromCache {
		result = x / y
		err := client.Set(cacheKey, result, 1*time.Minute).Err()
		if err != nil {
			log.Println("Error: Failed to save result to cache", err)
		}
	} else {
		parsedCacheValue, err := strconv.ParseFloat(cacheResult, 64)
		if err != nil {
			log.Println("Error: ", "Failed to parse cached value")
			return nil, internal.WrapErrorf(err, "Communication with cache failed")
		}
		result = parsedCacheValue
	}

	var response = new (CalculationResponse)
	response.Action = "divide"
	response.Answer = result
	response.Cached = usedValueFromCache
	response.X = x
	response.Y = y

	return response, nil
}