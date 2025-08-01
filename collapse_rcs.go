package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"strings"
)

const (
	qCount  = 100   // number of qubits
	depth   = 50    // circuit depth
	samples = 10000 // number of collapse samples
)

// Qubit structure with alpha and beta amplitudes
type Qubit struct {
	Alpha complex128
	Beta  complex128
}

// Normalize ensures the qubit remains normalized
func (q *Qubit) Normalize() {
	norm := math.Sqrt(cmplx.Abs(q.Alpha)*cmplx.Abs(q.Alpha) + cmplx.Abs(q.Beta)*cmplx.Abs(q.Beta))
	if norm == 0 {
		q.Alpha = 1
		q.Beta = 0
		return
	}
	q.Alpha /= complex(norm, 0)
	q.Beta /= complex(norm, 0)
}

// ApplyDeterministicGate simulates a collapse-respecting deterministic unitary
func (q *Qubit) ApplyDeterministicGate(theta float64) {
	phase := cmplx.Exp(complex(0, theta))
	q.Alpha *= phase
	q.Beta *= cmplx.Conj(phase)
	q.Normalize()
}

// CollapseBit deterministically collapses the qubit using the argument phase
func (q *Qubit) CollapseBit() byte {
	realPart := real(q.Alpha * cmplx.Conj(q.Alpha))
	if realPart >= 0.5 {
		return '0'
	}
	return '1'
}

// CollapseOneSample simulates the full evolution and collapse of one sample
func collapseOneSample(qCount, depth int) string {
	state := make([]Qubit, qCount)
	for i := range state {
		angle := 2 * math.Pi * float64(i) / float64(qCount)
		state[i] = Qubit{
			Alpha: cmplx.Exp(complex(0, angle)),
			Beta:  complex(0, 1),
		}
		state[i].Normalize()
	}

	for d := 0; d < depth; d++ {
		for i := range state {
			theta := 2 * math.Pi * float64(d+i) / float64(depth+qCount)
			state[i].ApplyDeterministicGate(theta)
		}
	}

	var b strings.Builder
	for _, q := range state {
		b.WriteByte(q.CollapseBit())
	}
	return b.String()
}

// WriteCSV writes all bitstrings to a CSV file
func writeCSV(filename string, bitstrings []string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	_ = writer.Write([]string{"Bitstring"})
	for _, b := range bitstrings {
		_ = writer.Write([]string{b})
	}
	return nil
}

func main() {
	fmt.Printf("Collapse Supremacy Simulation\nQubits: %d, Depth: %d, Samples: %d\n", qCount, depth, samples)

	allBitstrings := make([]string, samples)
	for i := 0; i < samples; i++ {
		bitstr := collapseOneSample(qCount, depth)
		allBitstrings[i] = bitstr
		if i%500 == 0 {
			fmt.Printf("Progress: %d/%d samples\n", i, samples)
		}
	}

	err := writeCSV("collapse_rcs_output.csv", allBitstrings)
	if err != nil {
		fmt.Println("Error writing CSV:", err)
		return
	}

	fmt.Println("Collapse simulation complete. Output saved to collapse_rcs_output.csv")
}

