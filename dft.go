package main

import raylib "github.com/gen2brain/raylib-go/raylib"
import (
	"math"
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
	real float64
	imaginary float64
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
				WIDTH / 2 - float64(samples[n].X),
				HEIGHT / 2 - float64(samples[n].Y),
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
			real: sum.real,
			imaginary: sum.imaginary,
			frequency: float64(k), // TODO: Redundant - this is just the index of X
			amplitude: math.Sqrt(sum.real * sum.real + sum.imaginary * sum.imaginary),
			phase: math.Atan2(sum.imaginary, sum.real),
		}
	}

	return X
}