package services

import (
	"example/go-calculator/internal"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var cacheDuration, _ = time.ParseDuration("20s")

type CalculationResponse struct {
	Action 	string 		`json:"action"`
	X 		float64 	`json:"x"`
	Y 		float64		`json:"y"`
	Answer 	float64		`json:"answer"`
	Cached 	bool		`json:"cached"`
}

func Add(x float64, y float64) (*CalculationResponse, error) {

	cacheKey := fmt.Sprintf("ADD|%f|%f", x, y)
	usedValueFromCache := true

	client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })

	var result float64;

	cacheResult, err := client.Get(cacheKey).Result()
    if err != nil {
		log.Println("Info: ", fmt.Sprintf("Key %s is not available in cache", cacheKey))
		usedValueFromCache = false
    }

	if !usedValueFromCache {
		result = x + y

		err := client.Set(cacheKey, result, cacheDuration).Err()
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

	client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })

	var result float64;

	cacheResult, err := client.Get(cacheKey).Result()
    if err != nil {
		log.Println("Info: ", fmt.Sprintf("Key %s is not available in cache", cacheKey))
		usedValueFromCache = false
    }

	if !usedValueFromCache {
		result = x - y
		cacheDuration, _ := time.ParseDuration("20s")
		err := client.Set(cacheKey, result, cacheDuration).Err()
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

	client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })

	var result float64;

	cacheResult, err := client.Get(cacheKey).Result()
    if err != nil {
		log.Println("Info: ", fmt.Sprintf("Key %s is not available in cache", cacheKey))
		usedValueFromCache = false
    }

	if !usedValueFromCache {
		result = x * y
		cacheDuration, _ := time.ParseDuration("20s")
		err := client.Set(cacheKey, result, cacheDuration).Err()
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

	client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })

	var result float64;

	cacheResult, err := client.Get(cacheKey).Result()
    if err != nil {
		log.Println("Info: ", fmt.Sprintf("Key %s is not available in cache", cacheKey))
		usedValueFromCache = false
    }

	if !usedValueFromCache {
		result = x / y
		cacheDuration, _ := time.ParseDuration("20s")
		err := client.Set(cacheKey, result, cacheDuration).Err()
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