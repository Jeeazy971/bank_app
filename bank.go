package main

import (
	"fmt"
	"os"
)

var balance float64 = 10000 // Solde initial

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	for {
		fmt.Println("\nBienvenue dans votre Banque !")
		fmt.Println("1. Consulter mon solde")
		fmt.Println("2. Déposer de l'argent")
		fmt.Println("3. Retirer de l'argent")
		fmt.Println("4. Quitter")

		var choice int
		fmt.Print("Entrez votre choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			consultBalance()
		case 2:
			depositMoney()
		case 3:
			withdrawMoney()
		case 4:
			fmt.Println("Merci d'avoir utilisé notre banque. À bientôt !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func consultBalance() {
	fmt.Printf("Votre solde est de : %.2f€\n", balance)
}

func depositMoney() {
	var amount float64
	fmt.Print("Montant à déposer : ")
	fmt.Scanln(&amount)
	if amount <= 0 {
		fmt.Println("Le montant doit être supérieur à 0.")
		return
	}
	balance += amount
	fmt.Printf("Vous avez déposé %.2f€. Votre nouveau solde est de : %.2f€\n", amount, balance)
	writeBalanceToFile(balance)
}

func withdrawMoney() {
	var amount float64
	fmt.Print("Montant à retirer : ")
	fmt.Scanln(&amount)
	if amount <= 0 {
		fmt.Println("Le montant doit être supérieur à 0.")
		return
	}
	if amount > balance {
		fmt.Println("Fonds insuffisants.")
		return
	}
	balance -= amount
	fmt.Printf("Vous avez retiré %.2f€. Votre nouveau solde est de : %.2f€\n", amount, balance)
	writeBalanceToFile(balance)
}
