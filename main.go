package main

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Response struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

type ErrorResponse struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPerfect(n int) bool {
	sum := 1
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n && n != 1
}

func isArmstrong(n int) bool {
	sum, temp, digits := 0, n, 0
	for temp > 0 {
		digits++
		temp /= 10
	}
	temp = n
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func classifyNumber(w http.ResponseWriter, r *http.Request) {
	numberStr := r.URL.Query().Get("number")
	if numberStr == "" {
		json.NewEncoder(w).Encode(ErrorResponse{Number: "", Error: true})
		return
	}

	n, err := strconv.Atoi(numberStr)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Number: numberStr, Error: true})
		return
	}

	properties := []string{}
	if n%2 != 0 {
		properties = append(properties, "odd")
	} else {
		properties = append(properties, "even")
	}
	if isArmstrong(n) {
		properties = append(properties, "armstrong")
	}
	
	response := Response{
		Number:     n,
		IsPrime:    isPrime(n),
		IsPerfect:  isPerfect(n),
		Properties: properties,
		DigitSum:   digitSum(n),
		FunFact:    strconv.Itoa(n) + " is an interesting number!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/classify-number", classifyNumber).Methods("GET")

	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":8000", handler)
}
