package executor

import (
	"fmt"
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

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{
			Value: node.Value,
		}
	case *ast.Program:
		return evalProgram(node.Statements, env)
	case *ast.ExprStatement:
		return Eval(node.Expr, env)
	case *ast.Boolean:
		if node.Value {
			return kTrue
		}
		return kFalse

	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.IfExpression:
		return evalIfExpression(node, env)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.FuncLiteral:
		params := node.Parameters
		body := node.Body
		return &object.Function{Parameters: params,
			Body: body, Env: env}
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpression(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)
	}

	return nil
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	function, ok := fn.(*object.Function)
	if !ok {
		return newError("not a function: %s", fn.Type())
	}
	extendedEnv := extendFunctionEnv(function, args)
	evaluated := Eval(function.Body, extendedEnv)
	return unwrapReturnValue(evaluated)
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnClosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func evalExpression(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return newError("identifier not found: " + node.Value)
	}
	return val
}

func evalProgram(stmts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object
	for _, statements := range stmts {
		result = Eval(statements, env)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func evalBlockStatement(b *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range b.Statements {
		result = Eval(statement, env)
		if result != nil && (result.Type() == object.RETURN_VALUE_OBJ ||
			result.Type() == object.ERROR_OBJ) {
			return result
		}
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
		return newError("unknown operator: %s%s", operator, right.Type())
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
		return newError("unknown operator: -%s", right.Type())
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
		return newError("unknown operator: %s %s %s", left.Type(),
			operator, right.Type())
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
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	cond := Eval(ie.Condition, env)
	if isTruthy(cond) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return kNull
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case kNull:
		return false
	case kTrue:
		return true
	case kFalse:
		return false
	default:
		return true
	}
}

func newError(s string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(s, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
