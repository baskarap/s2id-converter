package s2id_converter

type Geo interface {
	GetS2ID(lat float64, lng float64) (string, error)
}
