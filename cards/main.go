package main

func main() {
    //var card string = "Ace of Spades"
    card := deck{"Ace of Diamonds", newCard()}
    card = append(card, "Six of Spades")
    card.print()
}

func newCard() string {
    return "Five of Diamond"
}
