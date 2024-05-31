package main

import "fmt"

func main() {

	var principal_amount, loan_term, monthly_loan_payment float64

	

	print("Enter the Amount: ")
	fmt.Scan(&principal_amount)

	println("Enter the loan term in months")
	print("Enter the loan term: ")
	fmt.Scan(&loan_term)

	monthly_loan_payment = (principal_amount + (principal_amount*10/100*loan_term))/loan_term

	fmt.Print("Monthly Loan Payment is"," ",monthly_loan_payment)

}