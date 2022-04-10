package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	// TODO
	return nil
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	// TODO
	return nil
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	// TODO
	return nil
}

func evalIfExpresssion(ie *ast.IfExpression, env *object.Environment) object.Object {
	// TODO
	return nil
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	// TODO
	return nil
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	// TODO
	return nil
}

func evalIndexExpression(left, index object.Object) object.Object {
	// TODO
	return nil
}
