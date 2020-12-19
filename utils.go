package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"os"
	"strings"
)

func ExpandTilde(input string) string{
	home := os.Getenv("HOME")
	return strings.Replace(input, "~", home, 1)
}

func CreateUUID() string {
	return uuid.NewV4().String()
}

func FileExist(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}

	return false
}

func RandomString() string {
	var output strings.Builder

	charSet := "abcdedfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length := 20

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}

	return output.String()
}

func GenerateRandomTempFilename() string {
	return fmt.Sprintf("/tmp/%s", RandomString())
}

func NilStringPointer() *string{
	return nil
}

func StringPointer(input string) *string{
	return &input
}