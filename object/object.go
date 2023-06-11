package object

type ObjectType string


const (
    Integer_OBJ = "Integer"
    Boolean_OBJ = "Boolean"
    Null_OBj = "Null"
)


type Object interface {
    Type() ObjectType
    Inspect() string
}

