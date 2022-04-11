package main

import "fmt"

func rotationalCipher(input string, rotationFactor int) string {
	b := make([]byte, len(input))

	for i, c := range input {
		var value uint8
		switch {
		case 'A' <= c && c <= 'Z':
			value = (input[i] - 'A' + uint8(rotationFactor)) % 26
			b[i] = value + 'A'

		case 'a' <= c && c <= 'z':
			value = (input[i] - 'a' + uint8(rotationFactor)) % 26
			b[i] = value + 'a'

		case '0' <= c && c <= '9':
			value = (input[i] - '0' + uint8(rotationFactor)) % 10
			b[i] = value + '0'

		default:
			b[i] = input[i]
		}
	}

	return string(b)
}

func main() {
	fmt.Println(rotationalCipher("Zebra-493?", 3))
}
