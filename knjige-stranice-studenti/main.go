package main

// Služi za nalaženje krajnje desne vrednosti prve iteracije binarne pretrage
func sumaStranica(knjige []int) int {
	suma := 0

	for _, stranice := range knjige {
		suma += stranice
	}

	return suma
}

func knjige(knjige []int, brStudenata int) int {
	l := 0
	r := sumaStranica(knjige) // smisliti normalniju vrednost

	for l < r {
		mid := l + (r-l+1)/2
		if ok(knjige, brStudenata, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func ok(knjige []int, brStudenata int, stranice int) bool {
	procitano := 0
	potrebnoStudenata := 1

	for i := 0; i < len(knjige); i += 1 {
		if knjige[i] > stranice {
			return false
		}

		if procitano+knjige[i] <= stranice {
			procitano += knjige[i]
		} else {
			potrebnoStudenata += 1
			procitano = knjige[i]
		}
	}
	return potrebnoStudenata <= brStudenata
}

func main() {
	brStudenata := 2
	nizKnjiga := []int{12, 34, 67, 90}

	println(knjige(nizKnjiga, brStudenata))
}
