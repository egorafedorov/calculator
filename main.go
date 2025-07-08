package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("*** Calculator app ***")
	for {
		choiceUserOperation, errorChoiceUserOperation := inputChoiceUserOperation()
		if errorChoiceUserOperation != nil {
			fmt.Println("Error! Choose the correct type of operation")
			continue
		}
		userData, errorUserData := inputUserData()
		if errorUserData != nil {
			fmt.Println("Error! Cannot use letters")
			continue
		}
		userDataTotal := stringNumbersInInt(userData)
		result := calculationData(choiceUserOperation, userDataTotal)
		fmt.Printf("Result: %v\n", result)
		repeatCalculation := repeatCalculation()
		if !repeatCalculation {
			break
		}
	}
}

func inputChoiceUserOperation() (string, error) {
	var choiceUserOperation string
	fmt.Print("Select the type of operation for calculation (AVG/SUM/MED): ")
	fmt.Scan(&choiceUserOperation)
	if choiceUserOperation == "AVG" || choiceUserOperation == "SUM" || choiceUserOperation == "MED" {
		return choiceUserOperation, nil
	}
	return "", errors.New("INVALID_DATA")
}

func inputUserData() (string, error) {
	var userData string
	fmt.Print("Enter numbers separated by commas without spaces for calculation: ")
	fmt.Scan(&userData)
	for _, r := range userData {
		if unicode.IsLetter(r) {
			return "", errors.New("INVALID_DATA")
		}
	}

	if userData == "" || len(userData) == 0 {
		return "", errors.New("INVALID_DATA")
	}

	return userData, nil
}

func stringNumbersInInt(userData string) []int {
	userDataStringSplit := strings.Split(userData, ",")
	var userDataInt []int
	for _, s := range userDataStringSplit {
		num, err := strconv.Atoi(s)
		if err == nil {
			userDataInt = append(userDataInt, num)
		}
	}
	return userDataInt
}

func calculationData(userChoice string, userDataInt []int) int {
	var resultCalculation int
	switch {
	case userChoice == "AVG":
		c := len(userDataInt)
		for i := 0; i < c; i++ {
			resultCalculation += (userDataInt[i])
		}
		resultCalculation = resultCalculation / c
	case userChoice == "SUM":
		for _, value := range userDataInt {
			resultCalculation = resultCalculation + value
		}
	case userChoice == "MED":
		sort.Ints(userDataInt)
		l := len(userDataInt)
		if l == 0 {
			return 0
		} else if l%2 == 0 {
			resultCalculation = (userDataInt[l/2-1] + userDataInt[l/2]) / 2
		} else {
			resultCalculation = userDataInt[l/2]
		}
	}
	return resultCalculation
}

func repeatCalculation() bool {
	var userChoice string
	fmt.Print("Would you like to repeat the calculation? (Y/N): ")
	fmt.Scan(&userChoice)
	if userChoice == "Y" || userChoice == "y" {
		return true
	} else {
		return false
	}
}
