package solutions

type Solution struct {
	PartOne func()
	PartTwo func()
}

var SolutionMap = map[string]Solution{
	"1": solution1,
	"2": solution2,
	"3": solution3,
}
