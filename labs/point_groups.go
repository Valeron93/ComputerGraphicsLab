package lab

import "sync"

func FindGroups(dataset []Dot[int]) [][]Dot[int] {

	set := make(map[Dot[int]]bool)

	for _, element := range dataset {
		set[element] = true
	}

	result := make([][]Dot[int], 0)

	for len(set) > 0 {

		for k := range set {
			result = append(result, Delete(set, k))
			delete(set, k)
			break
		}
	}
	return result
}

func Delete(set map[Dot[int]]bool, p Dot[int]) []Dot[int] {

	result := make([]Dot[int], 0)

	for _, point := range Neighbors(p) {
		if _, has := set[point]; has {
			delete(set, point)

			result = append(result, point)
			result = append(result, Delete(set, point)...)
		}
	}

	return result
}

func Neighbors(p Dot[int]) []Dot[int]{

	return []Dot[int] {
		{p.X - 1, p.Y},
		{p.X + 1, p.Y},
		{p.X, p.Y - 1},
		{p.X, p.Y + 1},
	}
} 

func AveragePoint(group []Dot[int]) Dot[int] {

	size := float64(len(group))

	sum_x, sum_y := 0.0, 0.0

	for _, point := range group {
		sum_x += float64(point.X) / size
		sum_y += float64(point.Y) / size
	}

	return Dot[int] {
		X: int(sum_x),
		Y: int(sum_y),
	}
}

func AveragePointsGroups(groups [][]Dot[int]) []Dot[int] {

	result := make([]Dot[int], len(groups))

	wait := sync.WaitGroup{}

	for idx, group := range groups {
		wait.Add(1)

		go func(idx int, group []Dot[int]) {
			defer wait.Done()

			result[idx] = AveragePoint(group)
		
		}(idx, group)
	}

	wait.Wait()
	return result
}

