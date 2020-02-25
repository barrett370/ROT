package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type numberSet struct {
	values []float64
	bounds []float64
}

func newNumberSet(values []float64, weight []float64) (*numberSet, error) {
	if len(values) != len(weight) {
		return nil, fmt.Errorf("values and weight should have the same length")
	}
	s := &numberSet{
		values: values,
		bounds: weight,
	}
	sort.Sort(s)

	sum := float64(0)
	for i, weight := range s.bounds {
		sum += weight
		s.bounds[i] = sum
	}
	if sum-1 > 1e9 {
		return nil, fmt.Errorf("sum of weight should be 1, but was %f", sum)
	}
	return s, nil
}

func (s *numberSet) Len() int { return len(s.values) }
func (s *numberSet) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
	s.bounds[i], s.bounds[j] = s.bounds[j], s.bounds[i]
}
func (s *numberSet) Less(i, j int) bool { return s.bounds[i] < s.bounds[j] }

// Generator is a struct that can returns a random number chosen from a set
// of numbers where each has a specified probability.
type Generator struct {
	randSource rand.Source
	size       int
	numberSet
}

// NewGenerator return a Generator. It returns an error if len(weight) != len(values),
// or if the sum of weights is != 1.
// Two Generators with same seed, values and weight will always produce the same sequence
// of random number
func NewGenerator(seed int64, values []float64, weight []float64) (*Generator, error) {
	s, err := newNumberSet(values, weight)
	if err != nil {
		return nil, err
	}
	return &Generator{
		randSource: rand.NewSource(seed),
		size:       len(values),
		numberSet:  *s,
	}, nil
}

// Random returns a random number from the generator number set.
func (g *Generator) Random() float64 {
	r := float64(g.randSource.Int63()) / (1 << 63)
	i := sort.Search(g.size, func(i int) bool {
		return g.bounds[i] >= r
	})
	return g.values[i]
}
