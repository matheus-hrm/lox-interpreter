package main

import (
	"fmt"
	"log"
	"strings"
)

type Evaluator struct {
	AST *AST
}

type RuntimeError struct {
	Message string
	Token   Token
}

type env struct {
	values map[string]interface{}
}

func (e *RuntimeError) Error() string {
	return fmt.Sprintf("%s [line %d]", e.Message, e.Token.Line)
}

func NewEvaluator(ast *AST) *Evaluator {
	return &Evaluator{
		AST: ast,
	}
}

func (e *Evaluator) Evaluate() (interface{}, error) {
	var results []interface{}
	for _, node := range e.AST.Nodes {
		res, err := e.evaluateExpr(node)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}

func (e *Evaluator) evaluateExpr(expr Expr) (interface{}, error) {
	switch expr := expr.(type) {
	case *BinaryExpr:
		return e.evaluateBinary(expr)
	case *Grouping:
		return e.evaluateGrouping(expr)
	case Literal:
		return e.evaluateLiteral(&expr)
	case *UnaryExpr:
		return e.evaluateUnary(expr)
	case *LogicalExpr:
		return e.evaluateLogical(expr)
	case *AssignExpr:
		return e.evaluateAssign(expr)
	default:
		log.Printf("Unknown expression type: %T", expr)
		return nil, &RuntimeError{Message: "Unknown expression type", Token: Token{}}
	}
}

func checkNumOp(operator Token, op interface{}) (float64, error) {
	if v, ok := op.(float64); ok {
		return v, nil
	}
	return 0, &RuntimeError{Message: "Operand must be a number", Token: operator}
}

func checkNumOps(operator Token, left, right interface{}) (float64, float64, error) {
	leftNum, err := checkNumOp(operator, left)
	if err != nil {
		return 0, 0, err
	}
	rightNum, err := checkNumOp(operator, right)
	if err != nil {
		return 0, 0, err
	}
	return leftNum, rightNum, nil
}

func (e *Evaluator) evaluateBinary(expr *BinaryExpr) (interface{}, error) {
	if expr.Left == nil || expr.Right == nil {
		return nil, &RuntimeError{Message: "Invalid binary expression", Token: expr.Operator}
	}
	left, err := e.evaluateExpr(expr.Left)
	if err != nil {
		return nil, err
	}
	right, err := e.evaluateExpr(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case "+":
		if isString(left) && isString(right) {
			return left.(string) + right.(string), nil
		}
		leftNum, rightNum, err := checkNumOps(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return leftNum + rightNum, nil
	case "-":
		if isString(left) || isString(right) {
			return nil, &RuntimeError{Message: "Operands must be numbers", Token: expr.Operator}
		}
		leftNum, rightNum, err := checkNumOps(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return leftNum - rightNum, nil
	case "/":
		leftNum, rightNum, err := checkNumOps(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		if rightNum == 0 {
			return nil, &RuntimeError{Message: "Division by zero", Token: expr.Operator}
		}
		return leftNum / rightNum, nil
	case "*":
		leftNum, rightNum, err := checkNumOps(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return leftNum * rightNum, nil
	case ">":
		leftNum, rightNum, err := checkNumOps(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return leftNum > rightNum, nil
	case ">=":
		leftNum, rightNum, err := checkNumOps(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return leftNum >= rightNum, nil
	case "<":
		leftNum, rightNum, err := checkNumOps(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}

		return leftNum < rightNum, nil
	case "<=":
		leftNum, rightNum, err := checkNumOps(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return leftNum <= rightNum, nil
	case "!=":
		return left != right, nil
	case "==":
		return left == right, nil
	}
	return nil, nil
}

func (e *Evaluator) evaluateGrouping(expr *Grouping) (interface{}, error) {
	return e.evaluateExpr(expr.Expression)
}

func (e *Evaluator) evaluateLiteral(expr *Literal) (interface{}, error) {
	switch v := expr.Value.(type) {
	case float64:
		return v, nil
	case string:
		rawString := fmt.Sprintf("%v", strings.Trim(fmt.Sprintf("%v", v), `"`))
		return rawString, nil
	case bool:
		return v, nil
	case nil:
		return nil, nil
	default:
		return nil, &RuntimeError{Message: "Unknown literal type", Token: expr.Value.(Token)}
	}
}

func (e *Evaluator) evaluateUnary(expr *UnaryExpr) (interface{}, error) {
	right, err := e.evaluateExpr(expr.Right)
	if err != nil {
		return nil, err
	}
	switch expr.Operator.Type {
	case "-":
		rightNum, err := checkNumOp(expr.Operator, right)
		if err != nil {
			return nil, err
		}
		return -rightNum, nil
	case "!":
		return !isTruthy(right), nil
	}
	return nil, &RuntimeError{Message: "Unknown unary operator", Token: expr.Operator}
}

func (e *Evaluator) evaluateLogical(expr *LogicalExpr) (interface{}, error) {
	left, err := e.evaluateExpr(expr.Left)
	if err != nil {
		return nil, err
	}
	right, err := e.evaluateExpr(expr.Right)
	if err != nil {
		return nil, err
	}

	if expr.Operator.Type == "or" {
		return isTruthy(left) || isTruthy(right), nil
	}
	return isTruthy(left) && isTruthy(right), nil
}

func (e *Evaluator) evaluateAssign(expr *AssignExpr) (interface{}, error) {
	v, err := e.evaluateExpr(expr.Value)
	if err != nil {
		return nil, err
	}

	return v, nil
}
