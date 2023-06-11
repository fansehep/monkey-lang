package object


type Null struct {}

func (n *Null) Type() ObjectType {
    return Null_OBj
}

func (n *Null) Inspect() string {
    return "null"
}