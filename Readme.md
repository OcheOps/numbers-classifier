# Number Classification API

## Overview
The **Number Classification API** is a simple web service built in Golang that classifies numbers based on mathematical properties. It returns JSON responses containing properties such as primality, perfection, digit sum, and Armstrong status.

## Features
- **Check if a number is prime**
- **Check if a number is perfect**
- **Identify Armstrong numbers**
- **Return the sum of digits**
- **Provide a fun fact about the number**
- **CORS-enabled** for cross-origin requests
- **JSON-formatted responses**

## API Endpoint

### GET `/api/classify-number?number={value}`

#### Request Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `number`  | int  | Yes      | The integer number to classify |

#### Response Format (200 OK)
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```

#### Error Response (400 Bad Request)
```json
{
    "number": "invalid_input",
    "error": true
}
```

## Installation & Running Locally

### Prerequisites
- Go 1.18+
- Git

### Steps to Run Locally
```sh
git clone https://github.com/your-repo/number-classifier.git
cd number-classifier
go mod tidy
go run main.go
```

API will be available at `http://localhost:8000`

## Deployment (AWS EC2)
### Steps:
1. **SSH into your EC2 instance**
   ```sh
   ssh -i your-key.pem ubuntu@your-ec2-public-ip
   ```
2. **Install Golang on EC2**
   ```sh
   sudo apt update && sudo apt install -y golang
   ```
3. **Transfer Code to EC2**
   ```sh
   scp -i your-key.pem -r number-classifier ubuntu@your-ec2-public-ip:~
   ```
4. **Run the API on EC2**
   ```sh
   cd number-classifier
go mod tidy
go run main.go
   ```

API will be accessible at `http://your-ec2-public-ip:8000/api/classify-number?number=371`

## License
This project is licensed under the MIT License.

