package main;

type Point struct {
    X int
    Y int
}

type Slope struct {
	Dx int
	Dy int
}
func gcd(m int,n int) int {
	if m == 0 {
		return n
	}
    return gcd(n % m,m);
}
func NewSlope(pointA Point,pointB Point) Slope {
	dx := pointB.X - pointA.X
	dy := pointB.Y - pointA.Y;
	if dx == 0 && dy == 0 {
		return Slope{0,0}
	}
	if dx < 0 {
		dx = -dx
		dy = -dy
	}
	gcd := gcd(dx,dy)
	dx /= gcd
	dy /= gcd
	return Slope{dx,dy}
}

func maxPoints(points []Point) int {
	result := 0
    for startPointIndex,startPoint := range points {
		slopes := make(map[Slope]int)
		for _,endPoint := range points[startPointIndex:] {
			slope := NewSlope(startPoint,endPoint)
			slopes[slope]++
		}
		if len(slopes) == 1 && slopes[Slope{0,0}] != 0 {
			if slopes[Slope{0,0}] > result {
				result = slopes[Slope{0,0}]
			}
		} else {
			for slope := range slopes {
				if slope != (Slope{0,0}) {
					if slopes[slope] + slopes[Slope{0,0}] > result {
						result = slopes[slope] + slopes[Slope{0,0}]
					}
				}
			}
		}
	}
	return result
}

func main() {
	println(maxPoints([]Point {
		Point{0,0},
		Point{1,1},
		Point{1,1},
	}));
}