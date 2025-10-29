package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

//===== Lab #1 =====

func lab1_AssignmentOperators() {
	println("======  Lab #1  ======")
	var a, b int
	fmt.Print("Enter first interger(a): ")
	fmt.Scan(&a)
	fmt.Print("Enter second interger(b): ")
	fmt.Scan(&b)

	fmt.Println("\nIdentify value: ")
	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Printf("==== Operators ====")
	c := a
	fmt.Printf("\nOperation1 c = a => c = %d\n", c)
	c += b
	fmt.Printf("Operation2 c += b => c = %d\n", c)
	c -= b
	fmt.Printf("Operation3 c -= b => c = %d\n", c)
	c *= b
	fmt.Printf("Operation4 c *= b => c = %d\n", c)
	c /= b
	fmt.Printf("Operation5 c /= b => c = %d\n", c)
	c %= b
	fmt.Printf("Operation6 c %%= b => c = %d\n", c)
}

//===== lab #2 LogicalOperators =====

func lab2_AssignmentLogical_Operators() {
	println("=====  Lab #2  ======")
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	fmt.Println("==== Logic operators ====")
	fmt.Println("Both Positive (num1 > 0 && num2 > 0): ", (num1 > 0) && (num2 > 0))
	fmt.Println("One greater than other(num1 > num2 || num2 > num1): ", (num1 > num2) || (num2 > num1))
	fmt.Println("Non Equal (!(num1 == num2)): ", !(num1 == num2))
}

//===== lab #3 Bitwise and Assignment Opr =====

func myAND(a, b int) int        { return a & b }
func myOR(a, b int) int         { return a | b }
func myXOR(a, b int) int        { return a ^ b }
func myNOT(a, b int) (int, int) { return ^a, ^b }
func lab3_AssignmentBitwise_and_Assignment_Opr() {
	println("======  Lab #3  ======")
	var a, b int
	fmt.Print("Enter first number(a): ")
	fmt.Scan(&a)
	fmt.Print("Enter second number(b): ")
	fmt.Scan(&b)

	fmt.Println("===== Bitwise and Assignment Operators =====")
	fmt.Println("a & b = ", myAND(a, b))
	fmt.Println("a | b = ", myOR(a, b))
	fmt.Println("a ^ b = ", myXOR(a, b))
	notA, notB := myNOT(a, b)
	fmt.Println("^a = ", notA)
	fmt.Println("^b = ", notB)

}

//===== lab #4 Mini Calculator =====

func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }

func div(a, b int) (float64, string) {
	if b == 0 {
		return 0, "Error: division by zero"
	}
	return float64(a) / float64(b), ""
}

func mod(a, b int) (int, string) {
	if b == 0 {
		return 0, "Error: modulus by zero"
	}
	return a % b, ""
}

func lab4_AssignmentMiniCalculator() {
	for {
		fmt.Println("==== Lab #4 ====")
		fmt.Println("===== Mini Calculator =====")
		fmt.Println("1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit")
		fmt.Print("Choose: ")

		var choice, a, b int
		fmt.Scan(&choice)

		if choice == 6 {
			fmt.Println("Exiting calculator...")
			break
		}

		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)

		switch choice {
		case 1:
			fmt.Println("Result:", add(a, b))
		case 2:
			fmt.Println("Result:", sub(a, b))
		case 3:
			fmt.Println("Result:", mul(a, b))
		case 4:
			result, err := div(a, b)
			if err != "" {
				fmt.Println(err)
			} else {
				fmt.Println("Result:", result)
			}
		case 5:
			result, err := mod(a, b)
			if err != "" {
				fmt.Println(err)
			} else {
				fmt.Println("Result:", result)
			}
		default:
			fmt.Println("Invalid choice")
		}
	}
}

//===== lab #5 Binary Hex and Base64 Encoding =====

func toBinary(s string) string {
	bin := ""
	for _, c := range s {
		bin += fmt.Sprintf("%08b ", c)
	}
	return bin
}

func lab5_AssignmentBinary_Hex_and_Base64_Encoding() {
	fmt.Println("==== Lab #5 ====")
	fmt.Println("=== Binary_Hex_and_Base64_Encoding ===")
	var input string
	fmt.Print("Enter text: ")
	fmt.Scan(&input)

	fmt.Println("\nBinary:", toBinary(input))
	fmt.Println("Hexadecimal:", hex.EncodeToString([]byte(input)))
	fmt.Println("Base64:", base64.StdEncoding.EncodeToString([]byte(input)))
}

//===== lab #6 XOR Encryption Decryption =====

func xorEncrypt(text string, key byte) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key
	}
	return string(result)
}

func lab6_AssignmentXOR_Encryption_Decryption() {
	fmt.Println("======  Lab #6  ======")
	fmt.Println("=== XOR_Encryption_Decryption ===")
	var text string
	var key byte

	fmt.Print("Enter text: ")
	fmt.Scan(&text)

	fmt.Print("Enter key (0-255): ")
	fmt.Scan(&key)

	encrypted := xorEncrypt(text, key)
	fmt.Println("\nEncrypted: ", encrypted)

	decrypted := xorEncrypt(encrypted, key)
	fmt.Println("Decrypted: ", decrypted)
}

func main() {
	println("-----------------------------")
	println("----   Assignment_Menu   ----")
	println("-----------------------------")
	for {
		fmt.Println("\n====== Lab List ======")
		fmt.Println("1/ Lab #1 - Operators")
		fmt.Println("2/ Lab #2 - Logical Operators")
		fmt.Println("3/ Lab #3 - Bitwise and Assignment Opr")
		fmt.Println("4/ Lab #4 - Mini_Calculator")
		fmt.Println("5/ Lab #5 - Binary_Hex_and_Base64/Encoding")
		fmt.Println("6/ Lab #6 - XOR_Encryption_Decryption")
		fmt.Println("7/ Exit")
		fmt.Print("Choose option you want: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			lab1_AssignmentOperators()
		case 2:
			lab2_AssignmentLogical_Operators()
		case 3:
			lab3_AssignmentBitwise_and_Assignment_Opr()
		case 4:
			lab4_AssignmentMiniCalculator()
		case 5:
			lab5_AssignmentBinary_Hex_and_Base64_Encoding()
		case 6:
			lab6_AssignmentXOR_Encryption_Decryption()
		case 7:
			fmt.Println("Bye Bye")
			return
		default:
			fmt.Println("Invalid!")
		}
	}
}
