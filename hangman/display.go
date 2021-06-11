package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(`
	/$$   /$$                                                                
	| $$  | $$                                                                
	| $$  | $$  /$$$$$$  /$$$$$$$   /$$$$$$  /$$$$$$/$$$$   /$$$$$$  /$$$$$$$ 
	| $$$$$$$$ |____  $$| $$__  $$ /$$__  $$| $$_  $$_  $$ |____  $$| $$__  $$
	| $$__  $$  /$$$$$$$| $$  \ $$| $$  \ $$| $$ \ $$ \ $$  /$$$$$$$| $$  \ $$
	| $$  | $$ /$$__  $$| $$  | $$| $$  | $$| $$ | $$ | $$ /$$__  $$| $$  | $$
	| $$  | $$|  $$$$$$$| $$  | $$|  $$$$$$$| $$ | $$ | $$|  $$$$$$$| $$  | $$
	|__/  |__/ \_______/|__/  |__/ \____  $$|__/ |__/ |__/ \_______/|__/  |__/
																 /$$  \ $$                                  
																|  $$$$$$/                                  
																 \______/   
	`)
}

func Draw(g *Game, guess string) {
	DrawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func DrawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		+---+
		|   |
		O   |
	 /|\  |
	 / \  |
				|
	=========
		`
	case 1:
		draw = `
		+---+
		|   |
		O   |
	 /|\  |
	 /    |
				|
	=========
		`
	case 2:
		draw = `
		+---+
		|   |
		O   |
	 /|\  |
				|
				|
	=========
		`
	case 3:
		draw = `
		+---+
		|   |
		O   |
	 /|   |
				|
				|
	=========
		`
	case 4:
		draw = `
	+---+
	|   |
	O   |
	|   |
			|
			|
=========
		`
	case 5:
		draw = `
	+---+
	|   |
	O   |
			|
			|
			|
=========
		`
	case 6:
		draw = `
	+---+
  |   |
      |
      |
      |
      |
=========
		`
	case 7:
		draw = `
	+---+
			|
			|
			|
			|
			|
=========
		`

	case 8:
		draw = `
		
		`
	}
	fmt.Println(draw)

}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess!\n")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used!\n", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word\n", guess)
	case "lost":
		fmt.Print("You lost! :( the word was: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("YOU WON\n")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
