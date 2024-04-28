package storage

type (
	Storage[T ~string, V any]          map[T]V
	StorageInterface[T ~string, V any] interface {
		Set(t T, v V)
		Has(t T) bool
		Get(t T) V
	}
)

func (s *Storage[T, V]) Set(t T, v V) {
	(*s)[t] = v
}

func (s *Storage[T, V]) Has(t T) bool {
	_, ok := (*s)[t]

	return ok
}

func (s *Storage[T, V]) Get(t T) V {
	return (*s)[t]
}
