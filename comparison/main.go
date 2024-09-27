package main

import "fmt";

func main() {
	var name1 = "Lee";
	var name2 = "Adama";

	var result bool = name1 == name2;

	fmt.Println(result);

	var valueOne = 1;
	var valueTwo = 2;

	var resultOne = valueOne > 0
	var resultTwo = valueTwo > 1

	var oneTwo = resultOne && resultTwo;

	fmt.Println(oneTwo);
	
	var valueTrue = true;
	var valueFalse = false;

	var resultTrue = valueTrue || valueFalse;

	fmt.Println(resultTrue);
}

