package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Response struct {
	Number     float64  `json:"number"`
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
	n = int(math.Abs(float64(n))) // Ensure negative numbers are handled correctly
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func getFunFact(n float64) string {
	url := fmt.Sprintf("http://numbersapi.com/%v", n)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("Could not fetch fact for %v", n)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error reading fact for %v", n)
	}
	return string(body)
}

func classifyNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Ensure JSON response
	numberStr := r.URL.Query().Get("number")
	if numberStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Number: "", Error: true})
		return
	}

	n, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Number: numberStr, Error: true})
		return
	}

	properties := []string{}
	if int(n)%2 != 0 {
		properties = append(properties, "odd")
	} else {
		properties = append(properties, "even")
	}
	if n == float64(int(n)) && isArmstrong(int(n)) {
		properties = append(properties, "armstrong")
	}

	response := Response{
		Number:     n,
		IsPrime:    isPrime(int(n)),
		IsPerfect:  isPerfect(int(n)),
		Properties: properties,
		DigitSum:   digitSum(int(n)),
		FunFact:    getFunFact(n),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/classify-number", classifyNumber).Methods("GET")

	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":8000", handler)
}

