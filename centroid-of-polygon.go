package main

import (
	"fmt"

	"github.com/georgea93/gull/point"
	"github.com/golang/geo/s2"
)

func compute2DPolygonCentroid(vertices point.Points, vertexCount int) point.Point {
	centroid := point.Point{X: 0, Y: 0}
	signedArea := 0.0
	x0 := 0.0
	y0 := 0.0
	x1 := 0.0
	y1 := 0.0
	a := 0.0

	// For all vertices except last
	i := 0
	for i < vertexCount-1 {
		x0 = vertices[i].X
		y0 = vertices[i].Y
		x1 = vertices[i+1].X
		y1 = vertices[i+1].Y
		a = x0*y1 - x1*y0
		signedArea += a
		centroid.X += (x0 + x1) * a
		centroid.Y += (y0 + y1) * a
		i++
	}

	// Do last vertex separately to avoid performing an expensive
	// modulus operation in each iteration.
	x0 = vertices[i].X
	y0 = vertices[i].Y
	x1 = vertices[0].X
	y1 = vertices[0].Y
	a = x0*y1 - x1*y0
	signedArea += a
	centroid.X += (x0 + x1) * a
	centroid.Y += (y0 + y1) * a

	signedArea *= 0.5
	centroid.X /= (6.0 * signedArea)
	centroid.Y /= (6.0 * signedArea)
	return centroid
}

func main() {
	/*polygon := []point.Point{{X: 1.0, Y: 0.0}, {X: 1.0, Y: 1.0}, {X: 0.0, Y: 1.0}, {X: 0.0, Y: 0.0}}
	vertexCount := len(polygon)
	fmt.Print("\nPolyGon:", polygon, " Centroid:", compute2DPolygonCentroid(polygon, vertexCount))
	HullObj := *hull.FromPoints(polygon)
	fmt.Print("\n", HullObj)*/

	FindConvexHull()
}

//FindConvexHull LS
func FindConvexHull() {
	query := s2.NewConvexHullQuery()
	p := s2.PointFromCoords(0, 0, 1)
	q := s2.PointFromCoords(0, 1, 1)
	r := s2.PointFromCoords(1, 1, 1)
	s := s2.PointFromCoords(1, 0, 1)
	query.AddPoint(p)
	query.AddPoint(q)
	query.AddPoint(r)
	query.AddPoint(s)
	result := query.ConvexHull()
	fmt.Print(result.Centroid())
}
