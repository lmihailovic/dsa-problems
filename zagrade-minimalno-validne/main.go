package main

func raspareneZagrade(S []rune) int {
	resenje := 0

	otvoreneObicne := 0
	otvoreneViticaste := 0
	otvoreneUglaste := 0

	for i := 0; i < len(S); i++ {

		if S[i] == '{' {
			otvoreneViticaste += 1
		}

		if S[i] == '[' {
			otvoreneUglaste += 1
		}

		if S[i] == '(' {
			otvoreneObicne += 1
		}

		if S[i] == '}' && otvoreneViticaste == 0 {
			resenje += 1
		} else if S[i] == '}' {
			otvoreneViticaste -= 1
		}

		if S[i] == ']' && otvoreneUglaste == 0 {
			resenje += 1
		} else if S[i] == ']' {
			otvoreneUglaste -= 1
		}

		if S[i] == ')' && otvoreneObicne == 0 {
			resenje += 1
		} else if S[i] == ')' {
			otvoreneObicne -= 1
		}
	}

	return resenje + otvoreneViticaste + otvoreneUglaste + otvoreneObicne

}

func main() {
	zagrade := []rune{'{', '(', '[', ']', '(', '('}

	println(raspareneZagrade(zagrade))
}
