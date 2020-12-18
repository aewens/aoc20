package solutions

func init() {
	Map[17] = Solution17
}

type Point4D struct {
	X int
	Y int
	Z int
	W int
}

type Cubes map[Point4D]bool
type Conway4D struct {
	Cubes Cubes
	Min   Point4D
	Max   Point4D
}

func NewConway4D() *Conway4D {
	return &Conway4D{
		Cubes: make(Cubes),
		Min:   Point4D{-1,-1,0,0},
		Max:   Point4D{-1,-1,0,0},
	}
}

func (conway *Conway4D) Parse(y int, line string) {
	z := 0
	w := 0
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
		conway.Cubes[Point4D{x,y,z,w}] = active
	}
}

func (conway *Conway4D) Search(next Cubes, cube Point4D, active bool) Cubes {
	for dz := -1; dz <= 1; dz++ {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				search := Point4D{
					X: cube.X+dx,
					Y: cube.Y+dy,
					Z: cube.Z+dz,
					W: cube.W,
				}

				active, ok := conway.Cubes[search]
				if !ok {
					active = false
				}

				_, ok = next[search]
				if !ok {
					next = conway.Check(next, search, active)
				}

			}
		}
	}
	return next
}

func (conway *Conway4D) Check(next Cubes, cube Point4D, active bool) Cubes {
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

				search := Point4D{
					X: xSearch,
					Y: ySearch,
					Z: zSearch,
					W: cube.W,
				}

				active, ok := conway.Cubes[search]
				if !ok {
					active = false
				}

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

func (conway *Conway4D) SearchW(next Cubes, cube Point4D, active bool) Cubes {
	for dw := -1; dw <= 1; dw++ {
		for dz := -1; dz <= 1; dz++ {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					search := Point4D{
						X: cube.X+dx,
						Y: cube.Y+dy,
						Z: cube.Z+dz,
						W: cube.W+dw,
					}

					active, ok := conway.Cubes[search]
					if !ok {
						active = false
					}

					_, ok = next[search]
					if !ok {
						next = conway.CheckW(next, search, active)
					}
				}
			}
		}
	}
	return next
}

func (conway *Conway4D) CheckW(next Cubes, cube Point4D, active bool) Cubes {
	actives := 0
	for dw := -1; dw <= 1; dw++ {
		wSearch := cube.W+dw
		if wSearch < conway.Min.W {
			conway.Min.W = wSearch
		}
		if wSearch > conway.Max.W {
			conway.Max.W = wSearch
		}
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
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}

					search := Point4D{
						X: xSearch,
						Y: ySearch,
						Z: zSearch,
						W: wSearch,
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

func (conway *Conway4D) Step() {
	next := make(Cubes)
	for cube, active := range conway.Cubes {
		next = conway.Search(next, cube, active)
	}
	conway.Cubes = next
}

func (conway *Conway4D) StepW() {
	next := make(Cubes)
	for cube, active := range conway.Cubes {
		next = conway.SearchW(next, cube, active)
	}
	conway.Cubes = next
}

func (conway *Conway4D) Count() int {
	actives := 0
	for _, active := range conway.Cubes {
		if active {
			actives = actives + 1
		}
	}
	return actives
}

func (conway *Conway4D) Display() {
	Display(-1, conway.Min)
	Display(-1, conway.Max)
	for w := conway.Min.W; w <= conway.Max.W; w++ {
		Display(-2, w)
		for z := conway.Min.Z; z <= conway.Max.Z; z++ {
			Display(-3, z)
			for y := conway.Min.Y; y <= conway.Max.Y; y++ {
				row := ""
				for x := conway.Min.X; x <= conway.Max.X; x++ {
					search := Point4D{x,y,z,w}
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
		Display(0, "====")
	}
}

func (conway *Conway4D) Run(steps int) int {
	for i := 0; i < steps; i++ {
		conway.Step()
	}
	return conway.Count()
}

func (conway *Conway4D) RunW(steps int) int {
	for i := 0; i < steps; i++ {
		conway.StepW()
	}
	return conway.Count()
}

func Solution17(lines chan string) {
	conway1 := NewConway4D()
	conway2 := NewConway4D()
	y := 0
	for line := range lines {
		conway1.Parse(y, line)
		conway2.Parse(y, line)
		y = y + 1
	}
	Display(1, conway1.Run(6))
	Display(2, conway2.RunW(6))
}
