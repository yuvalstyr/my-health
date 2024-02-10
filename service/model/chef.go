package model

type Chef int

const (
	Me Chef = iota
	Oredred
	Outsourced
	Work
)

func (c Chef) String() string {
	switch c {
	case Me:
		return "Me"
	case Oredred:
		return "Ordered"
	case Outsourced:
		return "Outsourced"
	case Work:
		return "Work"
	}
	return "Unknown"
}
