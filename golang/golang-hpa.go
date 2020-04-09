package main

import(
    "math"
	"fmt"
    "net/http"
)
func main(){
	//executa_calculo()
	http.HandleFunc("/", index)
    http.ListenAndServe(":8000", nil) 
	//fmt.Printf("%.1f",executa_calculo())
	//fmt.Printf("Code.education Rocks!!")
}
func index(w http.ResponseWriter, r *http.Request) {
	 executa_calculo()
	 fmt.Printf("Code.education Rocks!!")
	 w.WriteHeader(http.StatusOK)
	 //tpl.Execute(w, data)
}

func executa_calculo() float64{
	var x float64 =0.0001
	for i:=1;i<100000000;i++ {
		x+= math.Sqrt(x)
	}

	return x
}