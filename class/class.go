package class

type Class struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity"`
}

func NewClass(name string, startDate string, endDate string, capacity int) (Class, error) {
	return Class{
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  capacity,
	}, nil
}
