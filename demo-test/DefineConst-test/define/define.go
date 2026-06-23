package define

const (
	Class_1_Name1 = "ClassName1"
	Class_1_Name2 = "ClassName2"
	Class_1_Name3 = "ClassName3"
	Class_1_Name4 = "ClassName4"
)

const (
	Class_2_Name1 = "ClassName1"
	Class_2_Name2 = "ClassName2"
	Class_2_Name3 = "ClassName3"
	Class_2_Name4 = "ClassName4"
)

type Class struct {
	Name1 string `json:"name1" define:"name1"`
	Name2 string `json:"name2" define:"name2"`
	Name3 string `json:"name3" define:"name3"`
	Name4 string `json:"name4" define:"name4"`
}

var ClassName = Class{
	"name1",
	"name2",
	"name3",
	"name4",
}
