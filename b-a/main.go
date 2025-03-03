package main

import (
	
	"fmt"
	"os"
	"strconv"
)

func writeBalanceToFile(balance float64){
	os.WriteFile("balance.txt", []byte(fmt.Sprintf("%f", balance)), 0644)
}
func getBalanceFromFile() (float64,error){
     data,_:=os.ReadFile("balance.txt")
	 balanceText:=string(data)
	 balance,err:=strconv.ParseFloat(balanceText,64)

	

	 ebt:=1000
	 profit:=2000
	 result:=fmt.Sprintf("ebt:%d\nprofit:%d\n",ebt,profit)
	 os.WriteFile("result.txt", []byte(result), 0644)
	 return balance, err

}

func main(){
	acc:=1000.0
	fmt.Println(acc)

    fmt.Print(getBalanceFromFile())
}