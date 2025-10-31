package main

import (
	"estoque/internal/services"
	"fmt"
)

func main() {
	fmt.Println("Sistema de estoque")

	// estoque := services.NewEstoque()
	// itens := []models.Item{
	// 	{ID: 1, Name: "Fone", Quantity: 10, Price: 100},
	// 	{ID: 2, Name: "Camiseta", Quantity: 2, Price: 55.99},
	// 	{ID: 3, Name: "Mouse", Quantity: 6, Price: 12.99},
	// }

	// for _, item := range itens {
	// 	err := estoque.AddItem(item, "Bruno")

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// }

	// //fmt.Println(estoque)

	// for _, item := range estoque.ListItems() {
	// 	fmt.Printf("\nID: %d | Item: %s | Quantidade: %d | Preço: %.2f", item.ID, item.Name, item.Quantity, item.Price)
	// }

	// fmt.Println("\nValor total do estoque R$:", estoque.CalculateTotalCost())

	// itemParaBuscar, err := services.FindBy(itens, func(item models.Item) bool {
	// 	return item.Name == "Camiseta"
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Print("Item encontrado: ", itemParaBuscar)

	alura := services.Fornecedor{
		CNPJ:    "123456",
		Contato: "123",
		Cidade:  "SP",
	}

	fmt.Println(alura.GetInfo())

	if alura.VerificarDispinibilidade(10, 15) {
		fmt.Println("Possui disponibilidade")
	} else {
		fmt.Println("Não possui disponibilidade")
	}

	// itemByName, err := services.FindByName(itens, "Camiseta")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Print(itemByName)

	//fmt.Println(estoque.ViewAuditLog())

	// logs := estoque.ViewAuditLog()

	// for _, log := range logs {
	// 	fmt.Printf("\n[%s] Ação: %s - Usuário: %s - Item ID: %d - Quantidade: %d - Motivo: %s", log.Timestamp.Format("01/02 15:04:05"),
	// 		log.Action, log.User, log.ItemID, log.Quantity, log.Reason)
	// }

	//fmt.Println(estoque.ListItems())

	// item := models.Item{
	// 	ID:       1,
	// 	Name:     "Fone",
	// 	Quantity: -5,
	// 	Price:    10,
	// }

	// err := estoque.AddItem(item)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// item1 := models.Item{
	// 	ID:       1,
	// 	Name:     "Primeiro Produto",
	// 	Quantity: 10,
	// 	Price:    19.99,
	// }

	// item2 := models.Item{
	// 	ID:       2,
	// 	Name:     "Segundo Produto",
	// 	Quantity: 2,
	// 	Price:    9.99,
	// }

	// person1 := models.Person{
	// 	FirstName: "Bruno",
	// 	LastName:  "Luz",
	// 	Age:       38,
	// }

	// fmt.Println(item1.Info())
	// fmt.Println(item2.Info())
	// fmt.Println(person1.String())
}
