// 1) число знак число =
// 2) число знак число знак число с приоритетом
// 3) число знак число знак скобка число знак число скобка
// 4) проверка на корректность скобок и знаков

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func count(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "*":
		return num1 * num2, nil
	case "/":
		return num1 / num2, nil
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	default:
		return 0, fmt.Errorf("error")
	}

}

func Calc(expression string) (float64, error) {
	list := strings.Split(expression, "")
	res := 0.0
	var stackNum []float64
	var stackSign []string
	for _, v := range list {
		fmt.Println(stackNum, stackSign, res)
		if number, err := strconv.ParseFloat(v, 64); err == nil {
			stackNum = append(stackNum, number)
		} else {
			n := len(stackNum) - 1
			fmt.Println(n)

			switch v {

			case "*", "/":
				fmt.Println(v == "*", n == -1)

				if n == 0 {

				} else if stackSign[0] == "+" || stackSign[0] == "-" {
					result, err := count(stackNum[n-1], stackNum[n], stackSign[0])
					if err != nil {
						return 0, err
					}

					res = result
					stackSign = stackSign[:n]
					stackNum = stackNum[:n-1]
					stackNum = append(stackNum, result)

				} else {
					result, err := count(stackNum[0], stackNum[1], stackSign[0])
					if err != nil {
						return 0, err
					}
					res = result
					stackSign = stackSign[1:]
					stackNum = stackNum[2:]
					stackNum = append(stackNum, result)

				}
				stackSign = append(stackSign, v)

			case "+", "-":
				if n == 0 {
				} else {

					result, err := count(stackNum[0], stackNum[1], stackSign[0])
					if err != nil {
						return 0, err
					}

					res = result
					stackSign = stackSign[1:]
					stackNum = stackNum[2:]
					stackNum = append(stackNum, result)
				}
				stackSign = append(stackSign, v)

			}
		}
	}
	if len(stackSign) != 0 {
		fmt.Println(stackNum, stackSign)
		result, err := count(stackNum[0], stackNum[1], stackSign[0])
		if err != nil {
			return 0, err
		}

		res = result
	}
	fmt.Println(stackNum, stackSign)
	return res, nil
}

// num = list[0]
// sign = list[1]
// for i := 0; i < len(list); i++ {
// 	if fmt.Sprintf("%T", list[i]) == "int" {
//         num = list[i]
// 	}else{
//         sign = list[i]
//     }
// num1, err := strconv.ParseFloat(list[0], 64)
// if err != nil {
// 	return 0, errors.New("error")
// }
// num2, err := strconv.ParseFloat(list[2], 64)
// if err != nil {
// 	return 0, errors.New("error")
// }
// switch sign {
// case "*":
// 	return num1 * num2, nil
// case "/":
// 	return num1 / num2, nil
// case "+":
// 	return num1 + num2, nil
// case "-":
// 	return num1 - num2, nil
// default:
// 	return 0, errors.New("error")
// }
// if fmt.Sprintf("%T", list[i]) == "int" && i+2 < len(list) &&
// 	(list[i+2] == "(" || fmt.Sprintf("%T", list[i]) == "int") {

// }

func main() {
	fmt.Println(Calc("2+2*2"))
    fmt.Println(Calc("2*2+2"))
}
