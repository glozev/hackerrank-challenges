package main

type UtopianTree struct {
	InitialHeight int
	GrowthCycles  int
}

func (tr *UtopianTree) ResultHeight() int {
	newHeight := tr.InitialHeight
	for i := 0; i < tr.GrowthCycles; i++ {
		if i%2 == 0 {
			newHeight *= 2
		} else {
			newHeight++
		}
	}
	return newHeight
}

func main() {

}
