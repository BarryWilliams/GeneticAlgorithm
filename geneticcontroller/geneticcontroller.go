package geneticcontroller
import (
	"sort"
	"fmt"
)

type Candidate interface {
	Value() float64
	MutateWith(parent2 Candidate) Candidate
}

type Candidates []Candidate

func (slice Candidates) Len() int {
return len(slice)
}

var lessFunc func(slice Candidates, i int, j int) bool

func (slice Candidates) Less(i, j int) bool {

	return lessFunc(slice, i, j)
	//return slice[i].Value() < slice[j].Value();
}

func (slice Candidates) Swap(i, j int) {
slice[i], slice[j] = slice[j], slice[i]
}

func (slice Candidates) Sort(){
	sort.Sort(slice)
}

type GeneticController struct {
	Population Candidates
	sortFunc func(gc *GeneticController)
	Iterations int64
}

func NewGeneticController(population []Candidate, sortOrder string) GeneticController{
	length := len(population)

	if length == 0 || length % 2 != 0{
		panic(fmt.Sprintf("Invalid population length: %v", length))
	}

	gc := GeneticController{
		Iterations: 0,
		Population: population,
	}

	switch sortOrder {
	case "ascending":
		lessFunc = func (slice Candidates, i int, j int) bool {
			return slice[i].Value() < slice[j].Value()
		}
	case "descending":
		lessFunc = func (slice Candidates, i int, j int) bool {
			return slice[i].Value() > slice[j].Value()
		}
	}

	return gc
}

func (gc *GeneticController) SpawnChildren(){
	//mutate highest ranking with second highest ranking. Mutate second with third, etc.
	halfLength := len(gc.Population)/2
	children := make(Candidates, halfLength, halfLength)
	for index, _ := range(children){
		parent1 := gc.Population[index]
		parent2 := gc.Population[index + 1]
		children[index] = parent1.MutateWith(parent2)
	}

	//merge children in
	gc.Population = append(gc.Population[0:halfLength], children...)
	gc.Population.Sort()
}

func (gc *GeneticController) RunTillTrue(condition func() bool) Candidate{
	gc.Population.Sort()
	for {
		if condition(){
			break
		}

		gc.Iterations ++
		gc.SpawnChildren()
	}

	return gc.Population[0]
}