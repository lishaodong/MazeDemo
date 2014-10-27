package  internal

type Room struct {
	IP string
	port int

}

func NewRoome() (*Room){
	return &Room{}
}

func (r *Room) Open(){}
