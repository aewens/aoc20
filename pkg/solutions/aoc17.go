package solutions

func init() {
	Map[17] = Solution17
}

type Point3D struct {
	X int
	Y int
	Z int
}

type Cubes map[Point3D]bool
type Conway3D struct {
	Cubes Cubes
	Min   Point3D
	Max   Point3D
}

func NewConway3D() *Conway3D {
	return &Conway3D{
		Cubes: make(Cubes),
		Min:   Point3D{-1,-1,0},
		Max:   Point3D{-1,-1,0},
	}
}

func (conway *Conway3D) Parse(y int, line string) {
	z := 0
	for x, letter := range line {
		active := letter == '#'
		if y < conway.Min.Y {
			conway.Min.Y = y
		}
		if y > conway.Max.Y {
			conway.Max.Y = y
		}
		if x < conway.Min.X {
			conway.Min.X = x
		}
		if x > conway.Max.X {
			conway.Max.X = x
		}
		conway.Cubes[Point3D{x,y,z}] = active
	}
}

func (conway *Conway3D) Search(next Cubes, cube Point3D, active bool) Cubes {
	for dz := -1; dz <= 1; dz++ {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				search := Point3D{
					X: cube.X+dx,
					Y: cube.Y+dy,
					Z: cube.Z+dz,
				}

				active, ok := conway.Cubes[search]
				if !ok {
					active = false
				}

				_, ok = next[search]
				if !ok {
					next[search] = active
				}

				next = conway.Check(next, search, active)
			}
		}
	}
	return next
}

func (conway *Conway3D) Check(next Cubes, cube Point3D, active bool) Cubes {
	actives := 0
	for dz := -1; dz <= 1; dz++ {
		zSearch := cube.Z+dz
		if zSearch < conway.Min.Z {
			conway.Min.Z = zSearch
		}
		if zSearch > conway.Max.Z {
			conway.Max.Z = zSearch
		}
		for dy := -1; dy <= 1; dy++ {
			ySearch := cube.Y+dy
			if ySearch < conway.Min.Y {
				conway.Min.Y = ySearch
			}
			if ySearch > conway.Max.Y {
				conway.Max.Y = ySearch
			}
			for dx := -1; dx <= 1; dx++ {
				xSearch := cube.X+dx
				if xSearch < conway.Min.X {
					conway.Min.X = xSearch
				}
				if xSearch > conway.Max.X {
					conway.Max.X = xSearch
				}
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}

				search := Point3D{
					X: xSearch,
					Y: ySearch,
					Z: zSearch,
				}

				active, ok := conway.Cubes[search]
				if !ok {
					active = false
				}

				//_, ok = next[search]
				//if !ok {
				//	next[search] = active
				//}

				if active {
					actives = actives + 1
				}
			}
		}
	}

	if active {
		next[cube] = (actives == 2 || actives == 3)
		//Display(-2, []int{cube.X,cube.Y,cube.Z,actives,1})
	} else {
		next[cube] = (actives == 3)
		//Display(-2, []int{cube.X,cube.Y,cube.Z,actives,0})
	}
	return next
}

func (conway *Conway3D) Step() {
	next := make(Cubes)
	for cube, active := range conway.Cubes {
		next = conway.Search(next, cube, active)
	}
	conway.Cubes = next
}

func (conway *Conway3D) Count() int {
	actives := 0
	for _, active := range conway.Cubes {
		if active {
			actives = actives + 1
		}
	}
	return actives
}

func (conway *Conway3D) Display() {
	Display(-1, conway.Min)
	Display(-1, conway.Max)
	for z := conway.Min.Z; z <= conway.Max.Z; z++ {
		Display(0, z)
		for y := conway.Min.Y; y <= conway.Max.Y; y++ {
			row := ""
			for x := conway.Min.X; x <= conway.Max.X; x++ {
				search := Point3D{x,y,z}
				active, ok := conway.Cubes[search]
				if !ok {
					active = false
				}
				if active {
					row = row + "#"
				} else {
					row = row + "."
				}
			}
			Display(0, row)
		}
		Display(0, "----")
	}
}

func (conway *Conway3D) Run(steps int) int {
	for i := 0; i < steps; i++ {
		conway.Step()
	}
	return conway.Count()
}

func Solution17(lines chan string) {
	conway := NewConway3D()
	y := 0
	for line := range lines {
		conway.Parse(y, line)
		y = y + 1
	}
	Display(1, conway.Run(6))
}
