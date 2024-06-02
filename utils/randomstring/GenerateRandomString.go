package randomstring

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomString() string {
	// generate random numerical number with length 9 digits
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(999999999)
	b := fmt.Sprintf("%09d", a)

	// get current date
	currentTime := time.Now()
	// with minutes format
	// currentTime.Format("2006-01-02 15:04:05")

	dateString := currentTime.Format("2006-01-02-15:04:05")

	// remove space
	result := fmt.Sprintf("%s-%s-%s", "order", dateString, b)

	return result
}
