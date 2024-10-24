package tools

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	mrand "math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

// ErrTimeout is returned if WaitFor takes longer than 300 second to happen.
var ErrTimeout = errors.New("Timed out")

// WaitFor polls a predicate function once per second to wait for a certain state to arrive.
func WaitFor(predicate func() (bool, error)) error {
	for i := 0; i < 300; i++ {
		time.Sleep(1 * time.Second)

		satisfied, err := predicate()
		if err != nil {
			return err
		}
		if satisfied {
			return nil
		}
	}
	return ErrTimeout
}

// MakeNewPassword generates a new string that's guaranteed to be different than the given one.
func MakeNewPassword(oldPass string) string {
	randomPassword := RandomString("", 16)
	for randomPassword == oldPass {
		randomPassword = RandomString("", 16)
	}
	return randomPassword
}

// RandomString generates a string of given length, but random content.
// All content will be within the ASCII graphic character set.
// (Implementation from Even Shaw's contribution on
// http://stackoverflow.com/questions/12771930/what-is-the-fastest-way-to-generate-a-long-random-string-in-go).
func RandomString(prefix string, n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	_, _ = rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return prefix + string(bytes)
}

// RandomInt will return a random integer between a specified range.
func RandomInt(min, max int) int {
	mrand.Seed(time.Now().Unix())
	return mrand.Intn(max-min) + min
}

// Elide returns the first bit of its input string with a suffix of "..." if it's longer than
// a comfortable 40 characters.
func Elide(value string) string {
	if len(value) > 40 {
		return value[0:37] + "..."
	}
	return value
}

// PrintResource returns a resource as a readable structure
func PrintResource(t *testing.T, resource interface{}) {
	b, _ := json.MarshalIndent(resource, "", "  ")
	t.Logf(string(b))
}

// ExtractNetworkAddress removes the mask from the CIDR block
func ExtractNetworkAddress(cidr string) string {
	parts := strings.Split(cidr, "/")
	if len(parts) != 2 {
		return ""
	}
	return parts[0]
}

// SetLastOctet changes the last octet of the IP address to the specified value
func SetLastOctet(ip string, newLastOctet int) string {
	// Split IP into its components
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return ip
	}

	// Replace the last octet with the new value
	parts[3] = fmt.Sprintf("%d", newLastOctet)

	// Reassemble the IP address
	return strings.Join(parts, ".")
}

// logFatal is a helper function to log fatal errors during the test.
func logFatal(t *testing.T, message string) {
	t.Helper()
	t.Fatal(message)
}

// AssertLengthGreaterThan checks if the length of the provided list is greater than the specified number.
// If the condition fails, it logs a fatal error and fails the test.
func AssertLengthGreaterThan(t *testing.T, list interface{}, threshold int) {
	t.Helper()

	// Use reflection to get the length of the list
	listValue := reflect.ValueOf(list)

	// Ensure the list is a slice or array
	if listValue.Kind() != reflect.Slice && listValue.Kind() != reflect.Array {
		logFatal(t, fmt.Sprintf("expected a slice or array, but got %s", listValue.Kind().String()))
		return
	}

	// Check if the length is greater than the threshold
	listLen := listValue.Len()
	if listLen <= threshold {
		logFatal(t, fmt.Sprintf("expected length to be greater than %d, but got %d", threshold, listLen))
	}
}
