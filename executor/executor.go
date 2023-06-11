package executor

import (
	"github.com/fansehep/monkey-lang/ast"
	"github.com/fansehep/monkey-lang/object"
)

var (
	kTrue = &object.Boolean{
		Value: true,
	}
	kFalse = &object.Boolean{
		Value: false,
	}
	kNull = &object.Null{}
)

func GlobalBoolObject(ue bool) *object.Boolean {
    if ue {
        return kTrue
    } 
    return kFalse
}


func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{
			Value: node.Value,
		}
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExprStatement:
		return Eval(node.Expr)
	case *ast.Boolean:
		if node.Value {
			return kTrue
		}
		return kFalse
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	}
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statements := range stmts {
		result = Eval(statements)
	}
	return result
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return kNull
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case kTrue:
		return kFalse
	case kFalse:
		return kTrue
	case kNull:
		return kTrue
	default:
		return kFalse
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.Integer_OBJ {
		return kNull
	}
	value := right.(*object.Integer).Value
	return &object.Integer{
		Value: -value,
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{
			Value: leftVal + rightVal,
		}
	case "-":
		return &object.Integer{
			Value: leftVal - rightVal,
		}
	case "*":
		return &object.Integer{
			Value: leftVal * rightVal,
		}
	case "/":
		return &object.Integer{
			Value: leftVal / rightVal,
		}
    case "<":
        return GlobalBoolObject(leftVal < rightVal)
    case ">":
        return GlobalBoolObject(leftVal > rightVal)
    case "==":
        return GlobalBoolObject(leftVal == rightVal)
    case "!=":
        return GlobalBoolObject(leftVal != rightVal)
    case "<=":
        return GlobalBoolObject(leftVal <= rightVal)
    case ">=":
        return GlobalBoolObject(leftVal >= rightVal)
	default:
		return kNull
	}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.Integer_OBJ && right.Type() == object.Integer_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
    case operator == "==":
        return GlobalBoolObject(left == right)
    case operator == "!=":
        return GlobalBoolObject(left != right)
    default:
		return kNull
	}
}
