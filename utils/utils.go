// Description: Generic utils functions
// Author: Pixie79
// ============================================================================
// package utils

package utils

import (
	// "bytes"
	"encoding/base64"
	"encoding/binary"

	// "encoding/gob"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	LogLevel string
)

func init() {
	LogLevel = GetEnvDefault("LOG_LEVEL", "INFO")
}

// Print prints the given message with the specified log level.
//
// Parameters:
//   - level: the log level to use (e.g. "INFO", "ERROR").
//   - msg: the message to be printed.
func Print(level, msg string) {
	switch strings.ToUpper(LogLevel) {
	case "WARNING":
		fmt.Printf("%s: %s\n", strings.ToUpper(level), msg)
	case "ERROR":
		fmt.Printf("%s: %s\n", strings.ToUpper(level), msg)
	case "DEBUG":
		fmt.Printf("%s: %s\n", strings.ToUpper(level), msg)
	case "INFO":
		fmt.Printf("%s: %s\n", strings.ToUpper(level), msg)
	}

}

// GetEnvDefault retrieves the value of the environment variable specified by the key.
// If the environment variable does not exist, it returns the default value.
//
// Parameters:
// - key: the name of the environment variable to retrieve.
// - defaultVal: the value to return if the environment variable does not exist.
//
// Return:
// - string: the value of the environment variable or the default value.
func GetEnvDefault(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// Die exits the program after logging an error message.
//
// It takes a message as a parameter and does not return anything.
func Die(msg string) {
	panic(msg)
}

// MaybeDie is a function that checks if an error exists and calls the Die function with a formatted error message if it does.
//
// Parameters:
// - err: the error to check.
// - msg: the message to include in the error message.
func MaybeDie(err error, msg string) {
	if err != nil {
		Die(fmt.Sprintf("%s: %+v", msg, err))
	}
}

// GetEnvOrDie returns the value of the specified environment variable or exits the program.
//
// It takes a key string as a parameter and returns a string value.
func GetEnvOrDie(key string) string {
	value, set := os.LookupEnv(key)
	if !set {
		Die(fmt.Sprintf("%s: environment variable not set", key))
	}
	return value
}

// InBetween checks if a number is within a given range.
//
// Parameters:
//   - i: the number to check
//   - min: the minimum range value (inclusive)
//   - max: the maximum range value (inclusive)
//
// Returns:
//   - bool: true if the number is within the range, false otherwise.
func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

// ChunkBy splits a slice of items into smaller chunks of a specified size.
//
// items: The slice of items to be split.
// chunkSize: The size of each chunk.
// [][]T: A slice of slices, where each slice represents a chunk of items.
func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

// B64DecodeMsg decodes a base64 encoded key and returns a subset of the key starting from the specified offset.
//
// Parameters:
//   - b64Key: The base64 encoded key to be decoded.
//   - offsetF: An optional integer representing the offset from which to start the subset of the key. If not provided, it defaults to 7.
//
// Returns:
//   - []byte: The subset of the key starting from the specified offset.
//   - error: An error if the decoding or subset operation fails.
func B64DecodeMsg(b64Key string, offsetF ...int) ([]byte, error) {
	offset := 7
	if len(offsetF) > 0 {
		offset = offsetF[0]
	}

	key, err := base64.StdEncoding.DecodeString(b64Key)
	if err != nil {
		return nil, err
	}

	result := key[offset:]
	return result, nil
}

// Contains checks if a string is present in a slice of strings.
//
// Parameters:
// - s: the slice of strings to search in.
// - str: the string to search for.
//
// Returns:
// - bool: true if the string is found, false otherwise.
func Contains(s []string, str string) bool {
	for _, v := range s {
		if strings.EqualFold(v, str) {
			return true
		}
	}
	return false
}

// DifferenceInSlices Returns
// missing from List1 but in list 2
// missing from List2 but in list 1
// common in both
func DifferenceInSlices(l1, l2 []string) ([]string, []string, []string) {
	var missingL1, missingL2, common []string
	sort.Strings(l1)
	sort.Strings(l2)
	for _, v := range l1 {
		if !Contains(l2, v) {
			missingL2 = append(missingL2, v)
		}
	}
	for _, v := range l2 {
		if !Contains(l1, v) {
			missingL1 = append(missingL1, v)
		}
	}
	for _, v := range l1 {
		if Contains(l2, v) {
			common = append(common, v)
		}
	}
	return missingL1, missingL2, common
}

// TimePtr takes a time.Time parameter and returns the same time.Time value.
//
// t: a time.Time parameter.
// Returns: a time.Time value.
func TimePtr(t time.Time) time.Time {
	return t
}

// // CreateBytes encodes the given data to bytes using gob encoding.
// //
// // data: the data to be encoded
// // []byte: the encoded data as a byte slice
// func CreateBytes(data any) []byte {
// 	var envBuffer bytes.Buffer
// 	encData := gob.NewEncoder(&envBuffer)
// 	err := encData.Encode(data)
// 	MaybeDie(err, "encoding to bytes failed")
// 	return envBuffer.Bytes()
// }

// CreateKey generates a key for encryption.
//
// key: The byte array used to generate the key.
// Returns: The generated key.
func CreateKey(key []byte) []byte {
	// If key is empty, use hostname as key
	if len(key) < 1 {
		Die("No key provided: try []byte(Hostname)")
		return []byte{}
	} else {
		return key
	}
}

// intToBytes converts an integer to a byte array.
//
// It takes an integer as a parameter and returns a byte array.
func intToBytes(n int) []byte {
	byteArray := make([]byte, 4)
	binary.BigEndian.PutUint32(byteArray, uint32(n))
	return byteArray
}

// addZeroToStart adds a zero byte at the start of the given byte array.
//
// byteArray: the input byte array
// Returns: the modified byte array with a zero byte added at the start
func addZeroToStart(byteArray []byte) []byte {
	return append([]byte{0}, byteArray...)
}

// EncodedBuffer returns the encoded buffer of an integer.
//
// It takes an integer as input and returns a byte slice that represents the encoded buffer.
func EncodedBuffer(i int) []byte {
	return addZeroToStart(intToBytes(i))
}
