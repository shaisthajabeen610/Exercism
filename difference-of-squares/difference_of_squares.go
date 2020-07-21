package diffsquares
import "math"
//SumOfSquares is a function used to find the sum of squares
func SumOfSquares(n int)int{
	sum:=0
	for i:=1;i<=n;i++{
		sum+=i*i
	}
	return sum
}
//SquareOfSum is a function used to find the square of the sum
func SquareOfSum(n int)int{
	sum2:=0
	for i:=1;i<=n;i++{
		sum2+=i
	}
	return sum2*sum2
}

//Difference is a function used to find the difference 
func Difference(n int)int{
	res:=math.Abs(float64(SumOfSquares(n))-float64(SquareOfSum(n)))
	return int(res)
}
	
