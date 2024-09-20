package model

type MealType int

const (
	Breakfast MealType = iota
	Lunch
	Dinner
)

func (m MealType) String() string {
	switch m {
	case Breakfast:
		return "Breakfast"
	case Lunch:
		return "Lunch"
	case Dinner:
		return "Dinner"
	}
	return "Unknown"
}
