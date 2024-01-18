package refiner

import "golang.org/x/exp/constraints"

type AllowedTypes interface {
	constraints.Ordered
}

func Refine[T AllowedTypes](slice1, slice2 []T) []T {
	combinedSlice := append(slice1, slice2...)
	temp := make(map[T]bool)
	var refined []T
	for _, v := range combinedSlice {
		_, ok := temp[v]
		if !ok {
			temp[v] = true
			refined = append(refined, v)
		}
	}
	return refined
}
