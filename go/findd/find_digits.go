package main

//FindDigits in N that exactly divide it
type FindDigits struct {
	Numb int
}

//Count the digits in N that exactly devide it
func (dig *FindDigits) Count() int {
	count := 0
	var digits []int
	n := dig.Numb
	for n > 0 {
		dig := n % 10
		if dig != 0 {
			digits = append(digits, dig)
		}
		n /= 10
	}
	for _, v := range digits {
		if dig.Numb%v == 0 {
			count++
		}
	}
	return count
}

func main() {

}
