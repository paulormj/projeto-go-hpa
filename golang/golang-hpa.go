package main

import(
    "math"
	"fmt"
)
func main(){
	executa_calculo()
	fmt.Printf("%.1f",executa_calculo())
	fmt.Printf("Code.education Rocks!!")
}

func executa_calculo() float64{
	var x float64 =0.0001
	for i:=1;i<100000000;i++ {
		x+= math.Sqrt(x)
	}

	return x
}