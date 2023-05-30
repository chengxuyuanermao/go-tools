package graph

import "math"

// 写一个golang函数，判断两个坐标的点所形成的线，是否与坐标系内的某个圆形相交
type Point struct {
	x float64
	y float64
}

type Circle struct {
	center Point
	radius float64
}

func isIntersect(p1, p2 Point, circle Circle) bool {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	a := dx*dx + dy*dy
	b := 2 * (dx*(p1.x-circle.center.x) + dy*(p1.y-circle.center.y))
	c := (p1.x-circle.center.x)*(p1.x-circle.center.x) +
		(p1.y-circle.center.y)*(p1.y-circle.center.y) -
		circle.radius*circle.radius
	delta := b*b - 4*a*c

	if delta < 0 {
		return false
	}

	t1 := (-b - math.Sqrt(delta)) / (2 * a)
	t2 := (-b + math.Sqrt(delta)) / (2 * a)

	if t1 >= 0 && t1 <= 1 {
		return true
	}

	if t2 >= 0 && t2 <= 1 {
		return true
	}

	return false
}
