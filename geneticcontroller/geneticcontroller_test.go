package geneticcontroller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/BarryWilliams/GeneticAlgorithm/geneticcontroller"
)

type FakeCandidate struct {
	value float64
}

func NewFakeCandidate(value float64) Candidate {
	fc := FakeCandidate{
		value: float64(value),
	}
	return &fc
}

func (c *FakeCandidate) Value() float64 {
	return c.value
}

//this fake mutator should cause each generation to have better children
func (parent1 *FakeCandidate) MutateWith(parent2 Candidate) Candidate {
	childValue := (parent1.Value() + parent2.Value())/2.0 + 1.0
	return NewFakeCandidate(childValue)
}

var _ = Describe("Genetic Controller", func() {

	Context("Init error conditions", func(){
		It("Panics if the population size is not divisible by 2",func(){
			population := []Candidate{}
			population = append(population, NewFakeCandidate(0))
			Ω(func(){NewGeneticController(population, "descending")}).Should(Panic())
		})

		It("Panics if the population size is 0",func(){
			population := []Candidate{}
			Ω(func(){NewGeneticController(population, "descending")}).Should(Panic())
		})
	})

	Context("Simple Population", func() {

		var (
			population []Candidate
			gc GeneticController
		)

		BeforeEach(func() {
			population = []Candidate{}
			population = append(population, NewFakeCandidate(1))
			population = append(population, NewFakeCandidate(0))
			population = append(population, NewFakeCandidate(2))
			population = append(population, NewFakeCandidate(-1))

		})

		It("Orders the population by rank ascending", func() {
			gc = NewGeneticController(population,"ascending")
			gc.Population.Sort()

			Expect(gc.Population[0].Value()).To(Equal(float64(-1)))
			Expect(gc.Population[1].Value()).To(Equal(float64(0)))
			Expect(gc.Population[2].Value()).To(Equal(float64(1)))
			Expect(gc.Population[3].Value()).To(Equal(float64(2)))
		})

		It("Orders the population by rank descending", func() {
			gc = NewGeneticController(population,"descending")
			gc.Population.Sort()

			Expect(gc.Population[0].Value()).To(Equal(float64(2)))
			Expect(gc.Population[1].Value()).To(Equal(float64(1)))
			Expect(gc.Population[2].Value()).To(Equal(float64(0)))
			Expect(gc.Population[3].Value()).To(Equal(float64(-1)))
		})

		It("Mutates the population", func(){
			gc = NewGeneticController(population,"descending")
			gc.Population.Sort()
			gc.SpawnChildren()

			Expect(gc.Population[0].Value()).To(Equal(float64(2.5)))
			Expect(gc.Population[1].Value()).To(Equal(float64(2)))
			Expect(gc.Population[2].Value()).To(Equal(float64(1.5)))
			Expect(gc.Population[3].Value()).To(Equal(float64(1)))
		})

		It("Evaluates for the highest candidate after 50 iterations", func() {
			gc = NewGeneticController(population,"descending")
			bestCandidate := gc.RunTillTrue(
				func() bool {
					return gc.Iterations > 50
				})
			Expect(bestCandidate.Value()).To(BeNumerically(">", 25))
		})

		It("Evaluates for the lowest candidate after 50 iterations", func() {
			gc = NewGeneticController(population,"ascending")
			bestCandidate := gc.RunTillTrue(
				func() bool {
					return gc.Iterations > 50
				})
			Expect(bestCandidate.Value()).To(Equal(float64(-1)))
		})

		It("Stops evaluating when a value greater than 25 is reached", func() {
			gc = NewGeneticController(population,"descending")
			bestCandidate := gc.RunTillTrue(
				func() bool {
					return gc.Population[0].Value() > 25
				})
			Expect(bestCandidate.Value()).To(BeNumerically(">", 25))
		})
	})
})
