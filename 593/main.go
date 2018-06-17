package main;
import "math"
import "sort"

type Vector struct { x, y float64 }

type Point_ struct { x, y float64 }

func (first_vector Vector) Equal(other_vector Vector) bool {
	delta_x := first_vector.x - other_vector.x
	delta_y := first_vector.y - other_vector.y
	if(delta_x * delta_x + delta_y * delta_y < 0.000001) {
		return true
	}
	return false
}

func (vector Vector) length() float64 {
	return math.Sqrt(vector.x * vector.x + vector.y * vector.y)
}

func (first_vector Vector) Dot(other_vector Vector) float64 {
	return first_vector.x * other_vector.x + first_vector.y * other_vector.y
}

func (first_point Point_) Equal(other_point Point_) bool {
	delta_x := first_point.x - other_point.x
	delta_y := first_point.y - other_point.y
	if(delta_x * delta_x + delta_y * delta_y < 0.000001) {
		return true
	}
	return false
}

type By func(p1, p2 *Point_) bool

func (by By) Sort(points []Point_) {
	ps := &pointSorter{
		points: points,
		by:      by}
	sort.Sort(ps)
}

type pointSorter struct {
	points []Point_
	by      func(p1, p2 *Point_) bool
}

func (s *pointSorter) Len() int {
	return len(s.points)
}

func (s *pointSorter) Swap(i, j int) {
	s.points[i], s.points[j] = s.points[j], s.points[i]
}

func (s *pointSorter) Less(i, j int) bool {
	return s.by(&s.points[i], &s.points[j])
}

func (from Point_) To(to Point_) Vector {
	return Vector{to.x-from.x,to.y-from.y}
}

func (from Point_) Add(path Vector) Point_ {
	return Point_{from.x+path.x,from.y+path.y}
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	points := make([]Point_,4)
	points[0].x = float64(p1[0])
	points[0].y = float64(p1[1])
	points[1].x = float64(p2[0])
	points[1].y = float64(p2[1])
	points[2].x = float64(p3[0])
	points[2].y = float64(p3[1])
	points[3].x = float64(p4[0])
	points[3].y = float64(p4[1])

	By(func(p1, p2 *Point_) bool {
		if p1.x < p2.x {
			return true
		} else if p1.x == p2.x && p1.y < p2.y {
			return true
		}
		return false
	}).Sort(points)	
	if(points[0].Equal(points[1]) || points[1].Equal(points[2]) || points[0].Equal(points[3])) {
		return false
	}
	edge1 := points[0].To(points[1])
	edge2 := points[0].To(points[2])
	if(math.Abs(edge1.length()-edge2.length())<0.000001 && 
	points[2].Add(edge1).Equal(points[3]) && math.Abs(edge1.Dot(edge2)) < 0.000001) {
		return true
	}
	return false
}

func main() {
	p0 := []int{0,0}
	p1 := []int{0,0}
	p2 := []int{0,0}
	p3 := []int{0,0}
	print(validSquare(p1,p2,p0,p3))
}