package model

type student struct {
	Name  string
	score float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}
func (u *student) Getscore() float64 {
	return u.score

}
