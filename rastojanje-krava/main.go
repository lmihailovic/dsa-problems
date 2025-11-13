package main

func krave(stale []int, brKrava int) int {
	l := 0
	r := stale[len(stale)-1] - stale[0]
	for l < r {
		mid := l + (r-l+1)/2
		if ok(stale, brKrava, mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func ok(stale []int, brKrava int, rastojanje int) bool {
	sledecaKrava := stale[0]
	smesteneKrave := 1
	for i := 1; i < len(stale); i += 1 {
		if sledecaKrava+rastojanje <= stale[i] {
			smesteneKrave += 1
			sledecaKrava = stale[i]
		}
	}
	return smesteneKrave >= brKrava
}

func main() {
	brKrava := 3
	stale := []int{1, 2, 4, 8, 9}

	println(krave(stale, brKrava))
}
