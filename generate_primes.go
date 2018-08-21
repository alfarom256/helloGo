package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//method to generate primes
//Params
//max_number : int : number to count primes up to
//primes : []int : array of primes up to max_number / 2
func primeGen(max_number int, primes []int) []int {
	//if max_number is odd, increment by one
	if max_number%2 == 1 {
		max_number++
	}
	fmt.Printf("testing up to %d\n", (max_number))
	//for every odd number from 3 to max_number/2
	for i := 3; i < max_number; i += 2 {
		if isPrime(i, primes) {
			fmt.Printf("%d is prime!\n", i)
			primes = append(primes, i)
		}
	}
	return primes
}

//func isPrime(number int, primes []int) (is_prime bool) {
func isPrime(number int, primes []int) bool {
	max_prime := (number + 1) / 2
	for _, j := range primes {
		if j > max_prime {
			break
		} else if number%j == 0 {
			return false
		}
	}
	return true
}

func main() {
	/*
		The basic premise to (roughly, and with no sieves) generate primes:
		For every odd number from Zero to NUMBER,
		Test every prime less than half of NUMBER,
		Add that number to the list of prime.

		We only test up to PRIME_1 < ceil(NUMBER+1 / 2) because every
		NUMBER > PRIME_X >= NUMBER+1 / 2

		results in a float between 1f and 2f

		We only test if NUMBER % PRIME_X != 0
		because if the modulus is 0, NUMBER is factorable by PRIME_X and (NUMBER/PRIME_X)
	*/

	//declare an int
	//max_num := 0

	fmt.Print("Enter number to search up to: ")

	//create a new a stream buffer reading obj to read from STDIN
	reader := bufio.NewReader(os.Stdin)
	//reading from the input buffer, delimited by a newline
	text, _ := reader.ReadString('\n')
	//convert the string from STDIN into an integer
	//before we do this, remember to remove '\n'
	text = strings.TrimSuffix(text, "\n")
	max_num, err := strconv.Atoi(text)

	if err == nil {
		fmt.Printf("You entered %v\n", max_num)
		primes := []int{2, 3}
		primes = primeGen(max_num, primes)
		fmt.Printf("Primes: %v", primes)

	} else {
		fmt.Print("Can't understand, %v is type %v", max_num, reflect.TypeOf(max_num))
	}

}
