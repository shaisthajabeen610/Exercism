package grains
import (
	"math"
	"errors"
)
//Square is a function used to find the number of grains on the given square
func Square(n int)(uint64, error){
	if n<=0 || n>64{
		return 0, errors.New("true")
	}
	if n>0 || n<=64{
		return uint64(math.Pow(2,float64(n-1))),nil
		}
  return 0, nil
}
//Total function returns the no.of grains on the chess board
func Total()uint64{
	return 1<<64-1
}
