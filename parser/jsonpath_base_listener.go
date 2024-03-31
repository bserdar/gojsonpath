// Code generated from Jsonpath.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Jsonpath

import "github.com/antlr4-go/antlr/v4"

// BaseJsonpathListener is a complete listener for a parse tree produced by JsonpathParser.
type BaseJsonpathListener struct{}

var _ JsonpathListener = &BaseJsonpathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJsonpathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJsonpathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJsonpathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJsonpathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPath is called when production path is entered.
func (s *BaseJsonpathListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseJsonpathListener) ExitPath(ctx *PathContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseJsonpathListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseJsonpathListener) ExitSelector(ctx *SelectorContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseJsonpathListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseJsonpathListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterSlice is called when production slice is entered.
func (s *BaseJsonpathListener) EnterSlice(ctx *SliceContext) {}

// ExitSlice is called when production slice is exited.
func (s *BaseJsonpathListener) ExitSlice(ctx *SliceContext) {}

// EnterIndexExpression is called when production indexExpression is entered.
func (s *BaseJsonpathListener) EnterIndexExpression(ctx *IndexExpressionContext) {}

// ExitIndexExpression is called when production indexExpression is exited.
func (s *BaseJsonpathListener) ExitIndexExpression(ctx *IndexExpressionContext) {}

// EnterTernaryExpression is called when production TernaryExpression is entered.
func (s *BaseJsonpathListener) EnterTernaryExpression(ctx *TernaryExpressionContext) {}

// ExitTernaryExpression is called when production TernaryExpression is exited.
func (s *BaseJsonpathListener) ExitTernaryExpression(ctx *TernaryExpressionContext) {}

// EnterChainExpression is called when production ChainExpression is entered.
func (s *BaseJsonpathListener) EnterChainExpression(ctx *ChainExpressionContext) {}

// ExitChainExpression is called when production ChainExpression is exited.
func (s *BaseJsonpathListener) ExitChainExpression(ctx *ChainExpressionContext) {}

// EnterLogicalAndExpression is called when production LogicalAndExpression is entered.
func (s *BaseJsonpathListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production LogicalAndExpression is exited.
func (s *BaseJsonpathListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterPowerExpression is called when production PowerExpression is entered.
func (s *BaseJsonpathListener) EnterPowerExpression(ctx *PowerExpressionContext) {}

// ExitPowerExpression is called when production PowerExpression is exited.
func (s *BaseJsonpathListener) ExitPowerExpression(ctx *PowerExpressionContext) {}

// EnterObjectLiteralExpression is called when production ObjectLiteralExpression is entered.
func (s *BaseJsonpathListener) EnterObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) {}

// ExitObjectLiteralExpression is called when production ObjectLiteralExpression is exited.
func (s *BaseJsonpathListener) ExitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) {}

// EnterInExpression is called when production InExpression is entered.
func (s *BaseJsonpathListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production InExpression is exited.
func (s *BaseJsonpathListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterLogicalOrExpression is called when production LogicalOrExpression is entered.
func (s *BaseJsonpathListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production LogicalOrExpression is exited.
func (s *BaseJsonpathListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterNotExpression is called when production NotExpression is entered.
func (s *BaseJsonpathListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production NotExpression is exited.
func (s *BaseJsonpathListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterSelectorExpression is called when production SelectorExpression is entered.
func (s *BaseJsonpathListener) EnterSelectorExpression(ctx *SelectorExpressionContext) {}

// ExitSelectorExpression is called when production SelectorExpression is exited.
func (s *BaseJsonpathListener) ExitSelectorExpression(ctx *SelectorExpressionContext) {}

// EnterRecursiveDescentTermExpression is called when production RecursiveDescentTermExpression is entered.
func (s *BaseJsonpathListener) EnterRecursiveDescentTermExpression(ctx *RecursiveDescentTermExpressionContext) {
}

// ExitRecursiveDescentTermExpression is called when production RecursiveDescentTermExpression is exited.
func (s *BaseJsonpathListener) ExitRecursiveDescentTermExpression(ctx *RecursiveDescentTermExpressionContext) {
}

// EnterArgumentsExpression is called when production ArgumentsExpression is entered.
func (s *BaseJsonpathListener) EnterArgumentsExpression(ctx *ArgumentsExpressionContext) {}

// ExitArgumentsExpression is called when production ArgumentsExpression is exited.
func (s *BaseJsonpathListener) ExitArgumentsExpression(ctx *ArgumentsExpressionContext) {}

// EnterUnaryMinusExpression is called when production UnaryMinusExpression is entered.
func (s *BaseJsonpathListener) EnterUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// ExitUnaryMinusExpression is called when production UnaryMinusExpression is exited.
func (s *BaseJsonpathListener) ExitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// EnterUnaryPlusExpression is called when production UnaryPlusExpression is entered.
func (s *BaseJsonpathListener) EnterUnaryPlusExpression(ctx *UnaryPlusExpressionContext) {}

// ExitUnaryPlusExpression is called when production UnaryPlusExpression is exited.
func (s *BaseJsonpathListener) ExitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) {}

// EnterFilterExpression is called when production FilterExpression is entered.
func (s *BaseJsonpathListener) EnterFilterExpression(ctx *FilterExpressionContext) {}

// ExitFilterExpression is called when production FilterExpression is exited.
func (s *BaseJsonpathListener) ExitFilterExpression(ctx *FilterExpressionContext) {}

// EnterEqualityExpression is called when production EqualityExpression is entered.
func (s *BaseJsonpathListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production EqualityExpression is exited.
func (s *BaseJsonpathListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterBitXOrExpression is called when production BitXOrExpression is entered.
func (s *BaseJsonpathListener) EnterBitXOrExpression(ctx *BitXOrExpressionContext) {}

// ExitBitXOrExpression is called when production BitXOrExpression is exited.
func (s *BaseJsonpathListener) ExitBitXOrExpression(ctx *BitXOrExpressionContext) {}

// EnterMultiplicativeExpression is called when production MultiplicativeExpression is entered.
func (s *BaseJsonpathListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production MultiplicativeExpression is exited.
func (s *BaseJsonpathListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterParenthesizedExpression is called when production ParenthesizedExpression is entered.
func (s *BaseJsonpathListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// ExitParenthesizedExpression is called when production ParenthesizedExpression is exited.
func (s *BaseJsonpathListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterAdditiveExpression is called when production AdditiveExpression is entered.
func (s *BaseJsonpathListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production AdditiveExpression is exited.
func (s *BaseJsonpathListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterRelationalExpression is called when production RelationalExpression is entered.
func (s *BaseJsonpathListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production RelationalExpression is exited.
func (s *BaseJsonpathListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterRecursiveDescentExpression is called when production RecursiveDescentExpression is entered.
func (s *BaseJsonpathListener) EnterRecursiveDescentExpression(ctx *RecursiveDescentExpressionContext) {
}

// ExitRecursiveDescentExpression is called when production RecursiveDescentExpression is exited.
func (s *BaseJsonpathListener) ExitRecursiveDescentExpression(ctx *RecursiveDescentExpressionContext) {
}

// EnterBitNotExpression is called when production BitNotExpression is entered.
func (s *BaseJsonpathListener) EnterBitNotExpression(ctx *BitNotExpressionContext) {}

// ExitBitNotExpression is called when production BitNotExpression is exited.
func (s *BaseJsonpathListener) ExitBitNotExpression(ctx *BitNotExpressionContext) {}

// EnterLiteralExpression is called when production LiteralExpression is entered.
func (s *BaseJsonpathListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production LiteralExpression is exited.
func (s *BaseJsonpathListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterArrayLiteralExpression is called when production ArrayLiteralExpression is entered.
func (s *BaseJsonpathListener) EnterArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {}

// ExitArrayLiteralExpression is called when production ArrayLiteralExpression is exited.
func (s *BaseJsonpathListener) ExitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {}

// EnterMemberIndexExpression is called when production MemberIndexExpression is entered.
func (s *BaseJsonpathListener) EnterMemberIndexExpression(ctx *MemberIndexExpressionContext) {}

// ExitMemberIndexExpression is called when production MemberIndexExpression is exited.
func (s *BaseJsonpathListener) ExitMemberIndexExpression(ctx *MemberIndexExpressionContext) {}

// EnterIdentifierExpression is called when production IdentifierExpression is entered.
func (s *BaseJsonpathListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production IdentifierExpression is exited.
func (s *BaseJsonpathListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterBitAndExpression is called when production BitAndExpression is entered.
func (s *BaseJsonpathListener) EnterBitAndExpression(ctx *BitAndExpressionContext) {}

// ExitBitAndExpression is called when production BitAndExpression is exited.
func (s *BaseJsonpathListener) ExitBitAndExpression(ctx *BitAndExpressionContext) {}

// EnterBitOrExpression is called when production BitOrExpression is entered.
func (s *BaseJsonpathListener) EnterBitOrExpression(ctx *BitOrExpressionContext) {}

// ExitBitOrExpression is called when production BitOrExpression is exited.
func (s *BaseJsonpathListener) ExitBitOrExpression(ctx *BitOrExpressionContext) {}

// EnterRecursiveDescentMemberIndexExpression is called when production RecursiveDescentMemberIndexExpression is entered.
func (s *BaseJsonpathListener) EnterRecursiveDescentMemberIndexExpression(ctx *RecursiveDescentMemberIndexExpressionContext) {
}

// ExitRecursiveDescentMemberIndexExpression is called when production RecursiveDescentMemberIndexExpression is exited.
func (s *BaseJsonpathListener) ExitRecursiveDescentMemberIndexExpression(ctx *RecursiveDescentMemberIndexExpressionContext) {
}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseJsonpathListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseJsonpathListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterElementList is called when production elementList is entered.
func (s *BaseJsonpathListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BaseJsonpathListener) ExitElementList(ctx *ElementListContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BaseJsonpathListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BaseJsonpathListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterPropertyExpressionAssignment is called when production PropertyExpressionAssignment is entered.
func (s *BaseJsonpathListener) EnterPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) {
}

// ExitPropertyExpressionAssignment is called when production PropertyExpressionAssignment is exited.
func (s *BaseJsonpathListener) ExitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) {
}

// EnterComputedPropertyExpressionAssignment is called when production ComputedPropertyExpressionAssignment is entered.
func (s *BaseJsonpathListener) EnterComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) {
}

// ExitComputedPropertyExpressionAssignment is called when production ComputedPropertyExpressionAssignment is exited.
func (s *BaseJsonpathListener) ExitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) {
}

// EnterPropertyName is called when production propertyName is entered.
func (s *BaseJsonpathListener) EnterPropertyName(ctx *PropertyNameContext) {}

// ExitPropertyName is called when production propertyName is exited.
func (s *BaseJsonpathListener) ExitPropertyName(ctx *PropertyNameContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseJsonpathListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseJsonpathListener) ExitLiteral(ctx *LiteralContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseJsonpathListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseJsonpathListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseJsonpathListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseJsonpathListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}
