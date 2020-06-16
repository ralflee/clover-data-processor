package model

//Spec spec struct
type Spec struct {
	Name    string
	Columns []*Column
}

//Column Column struct
type Column struct {
	Name  string
	Width int
	Type  string
}
