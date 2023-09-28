package main

import raylib "github.com/gen2brain/raylib-go/raylib"
import (
	"math"
	"sort"
)

type Complex struct {
	real float64
	imaginary float64
}

func Add(a Complex, b Complex) Complex {
	return Complex {
		a.real + b.real,
		a.imaginary + b.imaginary,
	}
}

func Multiply(a Complex, b Complex) Complex {
	return Complex {
		a.real * b.real - a.imaginary * b.imaginary,
		a.real * b.imaginary + a.imaginary * b.real,
	}
}

type DiscreteFourierTransform struct {
	frequency float64
	amplitude float64
	phase float64
}

// Xₖ => (1 / N) * ( Σ Xₙ * [ cos( (2π/N)•k•n ) - i•sin( (2π/N)•k•n ) ] )
func dft(samples []raylib.Vector2) []DiscreteFourierTransform {
	var X []DiscreteFourierTransform = make([]DiscreteFourierTransform, len(samples))
	var N = len(samples)

	for k := 0; k < N; k++ {
		var sum = Complex { 0, 0 }

		for n := 0; n < N; n++ {
			var Xn = Complex {
				float64(samples[n].X) - WIDTH / 2,
				float64(samples[n].Y) - HEIGHT / 2,
			}

			// ø = (2π/N)•k•n
			var theta = (2 * math.Pi * float64(k * n)) / float64(N)

			// Xₙ * (cos(ø) - i•sin(ø))
			var inc = Multiply(Xn, Complex { math.Cos(theta), -math.Sin(theta)})
			sum = Add(sum, inc)
		}

		// 1/N * the whole sum
		sum.real = sum.real / float64(N)
		sum.imaginary = sum.imaginary / float64(N)

		X[k] = DiscreteFourierTransform {
			frequency: float64(k),
			amplitude: math.Sqrt(sum.real * sum.real + sum.imaginary * sum.imaginary),
			phase: math.Atan2(sum.imaginary, sum.real),
		}
	}

	// Sort by amplitude so the circles will decrease in size
	// Frequency (k) is stored so sorting shouldn't matter
	sort.Slice(X, func(i, j int) bool {
		return X[i].amplitude > X[j].amplitude
	})

	return X
}