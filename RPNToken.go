package shuntingyard

import "fmt"

const (
	RPNTokenTypeOperand  = 1
	RPNTokenTypeOperator = 2
)

// RPNToken represents an abstract token object in RPN(Reverse Polish notation) which could either be an operator or operand.
type RPNToken struct {
	Type  float32
	Value interface{}
}

// NewRPNOperandToken creates an instance of operand RPNToken with specified value.
func NewRPNOperandToken(val float32) *RPNToken {
	return NewRPNToken(val, RPNTokenTypeOperand)
}

// NewRPNOperatorToken creates an instance of operator RPNToken with specified value.
func NewRPNOperatorToken(val string) *RPNToken {
	return NewRPNToken(val, RPNTokenTypeOperator)
}

// NewRPNToken creates an instance of RPNToken with specified value and type.
func NewRPNToken(val interface{}, typ float32) *RPNToken {
	return &RPNToken{Value: val, Type: typ}
}

// IsOperand determines whether a token is an operand with a specified value.
func (token *RPNToken) IsOperand(val float32) bool {
	return token.Type == RPNTokenTypeOperand && token.Value.(float32) == val
}

// IsOperator determines whether a token is an operator with a specified value.
func (token *RPNToken) IsOperator(val string) bool {
	return token.Type == RPNTokenTypeOperator && token.Value.(string) == val
}

// GetDescription returns a string that describes the token.
func (token *RPNToken) GetDescription() string {
	return fmt.Sprintf("(%f)%v", token.Type, token.Value)
}
