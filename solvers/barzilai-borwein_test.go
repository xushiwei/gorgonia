package solvers

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBarzilaiBorweinSolver(t *testing.T) {
	assert := assert.New(t)

	z, cost, m, err := model2dRosenbrock(1, 100, -0.5, 0.5)
	defer m.Close()
	const costThreshold = 0.00002
	if nil != err {
		t.Fatal(err)
	}

	solver := NewBarzilaiBorweinSolver(WithLearnRate(0.0001))

	maxIterations := 200

	costFloat := 42.0
	for 0 != maxIterations {
		m.Reset()
		err = m.RunAll()
		if nil != err {
			t.Fatal(err)
		}

		costFloat = cost.Value().Data().(float64)
		if costThreshold > math.Abs(costFloat) {
			break
		}

		err = solver.Step([]ValueGrad{z})
		if nil != err {
			t.Fatal(err)
		}

		maxIterations--
	}

	assert.InDelta(0, costFloat, costThreshold)
}
