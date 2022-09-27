package main

import (
	"fmt"
	"math/rand"
	"strings"
	// "github.com/google/uuid"
)

// 1. There are two roles Accountant and Customer
// 2. Accountants can create a bank account for the customer by taking their email and phone
// number and returning them bank account details(Including bank account ID and
// randomly generated password)
// 3. Customers can deposit and withdraw money to their respective bank accounts by
// providing their bank account number and generated password.
// 4. Customers can also see their current bank account balance by providing their bank
// account number and generated password

type Accountant interface {
	create() BankAccountDetails
}

type Customer interface {
	deposit(amt int) float32
	withdraw(amt int) float32
	balanceCheck() float32
	history(d1, d2 string)
}

type CustomerDetails struct {
	bankDetails BankAccountDetails
	cred        AcoountImplCredentials
}

type BankAccountDetails struct {
	accountId int
	password  string
	balance   float32
}
type AdditionalAccountDetails struct {
}

type AcoountImplCredentials struct {
	email, phone string
}

func (cred AcoountImplCredentials) create() BankAccountDetails {
	newAccount := new(BankAccountDetails)
	newAccount.accountId = rand.Intn(1000)

	newAccount.balance = 0
	newAccount.password = generatePassword(8, 1, 1, 1)

	accountMap[cred] = *newAccount

	return *newAccount
}

func (details BankAccountDetails) deposit(amt float32) float32 {
	for key, value := range accountMap {
		if value.accountId == details.accountId &&
			value.password == details.password {
			value.balance = value.balance + amt
			accountMap[key] = value
			return value.balance
		}
	}
	return -1
}

func (details BankAccountDetails) withdraw(amt float32) float32 {
	for key, value := range accountMap {
		if value.accountId == details.accountId &&
			value.password == details.password {
			value.balance = value.balance - amt
			accountMap[key] = value
			return value.balance
		}
	}
	return -1
}

func (details BankAccountDetails) balanceCheck() float32 {
	for _, value := range accountMap {
		if value.accountId == details.accountId &&
			value.password == details.password {
			return value.balance
		}
	}
	return -1
}

var accountMap = make(map[AcoountImplCredentials]BankAccountDetails)

func main() {
	accountMap[AcoountImplCredentials{"prashant@josh", "8888715525"}] =
		BankAccountDetails{1, "qwerty", 1000}
	accountMap[AcoountImplCredentials{"suraj@josh", "123456"}] =
		BankAccountDetails{2, "asdfgh", 1500}
	for {
		fmt.Println("Select Role")
		fmt.Printf("1 Accountant\n2 Customer\n")
		var role int
		fmt.Scan(&role)

		switch role {
		case 1:
			{
				fmt.Printf("Enter  email and phone number:")
				var accEmail string
				var accPwd string
				fmt.Scan(&accEmail, &accPwd)
				if accEmail != "account@bank.com" && accPwd != "josh@123" {
					fmt.Println("Invalid credentials")
					continue
				}
				fmt.Printf(" 1 Create Account\n 2 Display all Accounts\n")
				var accountChoice int
				fmt.Scan(&accountChoice)
				if accountChoice == 1 {
					fmt.Printf("Enter customer email and phone number:")
					var email string
					var phone string
					fmt.Scan(&email, &phone)
					if len(email) > 0 {
						cd := AcoountImplCredentials{email, phone}
						var accountDetails BankAccountDetails = cd.create()
						accountMap[cd] = accountDetails

						fmt.Println("Account created successfully:", accountDetails)
						DisplayAllAccount()
					}

				} else if accountChoice == 2 {
					DisplayAllAccount()
				}
			}
		case 2:
			{
				fmt.Printf(" 1 Deposit \n 2 Withdraw\n 3 View Balance\n")
				var customerChoice int
				fmt.Scan(&customerChoice)

				fmt.Printf("Enter accountId and password:")
				var accountId int
				var password string
				fmt.Scan(&accountId, &password)

				// _, exists := accountMap[accountId] // Just checks for key existence

				for key, value := range accountMap {
					if value.accountId == accountId &&
						value.password == password {

						if customerChoice == 1 {
							var accountDetails BankAccountDetails = accountMap[key]
							fmt.Println("Enter amount to be deposit:")
							var amt float32
							fmt.Scan(&amt)
							accountDetails.deposit(amt)
							DisplayAllAccount()

						} else if customerChoice == 2 {
							var accountDetails BankAccountDetails = accountMap[key]
							fmt.Println("Enter amount to be withdraw:")
							var amt float32
							fmt.Scan(&amt)
							accountDetails.withdraw(amt)
							DisplayAllAccount()

						} else if customerChoice == 3 {
							var accountDetails BankAccountDetails = accountMap[key]
							var bal float32 = accountDetails.balanceCheck()
							fmt.Println("BALANCE:", bal)
						}
					}
				}
			}
		case 0:
			break
		default:
			{
				fmt.Println("Invalid,press 0 to exit ")
				break
			}
		}
		if role == 0 {
			break
		}
	}

}

func DisplayAllAccount() {
	fmt.Println("-----------------")
	for key, value := range accountMap {
		fmt.Printf("\nAccountId: %v Account Deatils: %v", key, value)
	}
	fmt.Println("\n-----------------")

}

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
