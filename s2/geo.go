package s2

type Geo struct {
	level int
}

func NewGeo(level int) *Geo {
	return &Geo{
		level: level,
	}
}
