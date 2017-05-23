package main

//AngryProfessor struct
type AngryProfessor struct {
	K                   int
	StudentsArrivalTime []int
}

//Result for the AngryProfessor
func (prof *AngryProfessor) Result() string {
	studentsOnTime := 0
	for _, t := range prof.StudentsArrivalTime {
		if t <= 0 {
			studentsOnTime++
		}
	}
	if studentsOnTime < prof.K {
		return "YES"
	}
	return "NO"
}

func main() {

}
