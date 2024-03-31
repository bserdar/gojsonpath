// Code generated from Jsonpath.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Jsonpath

import "github.com/antlr4-go/antlr/v4"

// JsonpathListener is a complete listener for a parse tree produced by JsonpathParser.
type JsonpathListener interface {
	antlr.ParseTreeListener

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterSlice is called when entering the slice production.
	EnterSlice(c *SliceContext)

	// EnterIndexExpression is called when entering the indexExpression production.
	EnterIndexExpression(c *IndexExpressionContext)

	// EnterTernaryExpression is called when entering the TernaryExpression production.
	EnterTernaryExpression(c *TernaryExpressionContext)

	// EnterChainExpression is called when entering the ChainExpression production.
	EnterChainExpression(c *ChainExpressionContext)

	// EnterLogicalAndExpression is called when entering the LogicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterPowerExpression is called when entering the PowerExpression production.
	EnterPowerExpression(c *PowerExpressionContext)

	// EnterObjectLiteralExpression is called when entering the ObjectLiteralExpression production.
	EnterObjectLiteralExpression(c *ObjectLiteralExpressionContext)

	// EnterInExpression is called when entering the InExpression production.
	EnterInExpression(c *InExpressionContext)

	// EnterLogicalOrExpression is called when entering the LogicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterNotExpression is called when entering the NotExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterSelectorExpression is called when entering the SelectorExpression production.
	EnterSelectorExpression(c *SelectorExpressionContext)

	// EnterRecursiveDescentTermExpression is called when entering the RecursiveDescentTermExpression production.
	EnterRecursiveDescentTermExpression(c *RecursiveDescentTermExpressionContext)

	// EnterArgumentsExpression is called when entering the ArgumentsExpression production.
	EnterArgumentsExpression(c *ArgumentsExpressionContext)

	// EnterUnaryMinusExpression is called when entering the UnaryMinusExpression production.
	EnterUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// EnterUnaryPlusExpression is called when entering the UnaryPlusExpression production.
	EnterUnaryPlusExpression(c *UnaryPlusExpressionContext)

	// EnterFilterExpression is called when entering the FilterExpression production.
	EnterFilterExpression(c *FilterExpressionContext)

	// EnterEqualityExpression is called when entering the EqualityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterBitXOrExpression is called when entering the BitXOrExpression production.
	EnterBitXOrExpression(c *BitXOrExpressionContext)

	// EnterMultiplicativeExpression is called when entering the MultiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterParenthesizedExpression is called when entering the ParenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterAdditiveExpression is called when entering the AdditiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterRelationalExpression is called when entering the RelationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterRecursiveDescentExpression is called when entering the RecursiveDescentExpression production.
	EnterRecursiveDescentExpression(c *RecursiveDescentExpressionContext)

	// EnterBitNotExpression is called when entering the BitNotExpression production.
	EnterBitNotExpression(c *BitNotExpressionContext)

	// EnterLiteralExpression is called when entering the LiteralExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterArrayLiteralExpression is called when entering the ArrayLiteralExpression production.
	EnterArrayLiteralExpression(c *ArrayLiteralExpressionContext)

	// EnterMemberIndexExpression is called when entering the MemberIndexExpression production.
	EnterMemberIndexExpression(c *MemberIndexExpressionContext)

	// EnterIdentifierExpression is called when entering the IdentifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterBitAndExpression is called when entering the BitAndExpression production.
	EnterBitAndExpression(c *BitAndExpressionContext)

	// EnterBitOrExpression is called when entering the BitOrExpression production.
	EnterBitOrExpression(c *BitOrExpressionContext)

	// EnterRecursiveDescentMemberIndexExpression is called when entering the RecursiveDescentMemberIndexExpression production.
	EnterRecursiveDescentMemberIndexExpression(c *RecursiveDescentMemberIndexExpressionContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterElementList is called when entering the elementList production.
	EnterElementList(c *ElementListContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterPropertyExpressionAssignment is called when entering the PropertyExpressionAssignment production.
	EnterPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext)

	// EnterComputedPropertyExpressionAssignment is called when entering the ComputedPropertyExpressionAssignment production.
	EnterComputedPropertyExpressionAssignment(c *ComputedPropertyExpressionAssignmentContext)

	// EnterPropertyName is called when entering the propertyName production.
	EnterPropertyName(c *PropertyNameContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitSlice is called when exiting the slice production.
	ExitSlice(c *SliceContext)

	// ExitIndexExpression is called when exiting the indexExpression production.
	ExitIndexExpression(c *IndexExpressionContext)

	// ExitTernaryExpression is called when exiting the TernaryExpression production.
	ExitTernaryExpression(c *TernaryExpressionContext)

	// ExitChainExpression is called when exiting the ChainExpression production.
	ExitChainExpression(c *ChainExpressionContext)

	// ExitLogicalAndExpression is called when exiting the LogicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitPowerExpression is called when exiting the PowerExpression production.
	ExitPowerExpression(c *PowerExpressionContext)

	// ExitObjectLiteralExpression is called when exiting the ObjectLiteralExpression production.
	ExitObjectLiteralExpression(c *ObjectLiteralExpressionContext)

	// ExitInExpression is called when exiting the InExpression production.
	ExitInExpression(c *InExpressionContext)

	// ExitLogicalOrExpression is called when exiting the LogicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitNotExpression is called when exiting the NotExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitSelectorExpression is called when exiting the SelectorExpression production.
	ExitSelectorExpression(c *SelectorExpressionContext)

	// ExitRecursiveDescentTermExpression is called when exiting the RecursiveDescentTermExpression production.
	ExitRecursiveDescentTermExpression(c *RecursiveDescentTermExpressionContext)

	// ExitArgumentsExpression is called when exiting the ArgumentsExpression production.
	ExitArgumentsExpression(c *ArgumentsExpressionContext)

	// ExitUnaryMinusExpression is called when exiting the UnaryMinusExpression production.
	ExitUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// ExitUnaryPlusExpression is called when exiting the UnaryPlusExpression production.
	ExitUnaryPlusExpression(c *UnaryPlusExpressionContext)

	// ExitFilterExpression is called when exiting the FilterExpression production.
	ExitFilterExpression(c *FilterExpressionContext)

	// ExitEqualityExpression is called when exiting the EqualityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitBitXOrExpression is called when exiting the BitXOrExpression production.
	ExitBitXOrExpression(c *BitXOrExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the MultiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitParenthesizedExpression is called when exiting the ParenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitAdditiveExpression is called when exiting the AdditiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitRelationalExpression is called when exiting the RelationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitRecursiveDescentExpression is called when exiting the RecursiveDescentExpression production.
	ExitRecursiveDescentExpression(c *RecursiveDescentExpressionContext)

	// ExitBitNotExpression is called when exiting the BitNotExpression production.
	ExitBitNotExpression(c *BitNotExpressionContext)

	// ExitLiteralExpression is called when exiting the LiteralExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitArrayLiteralExpression is called when exiting the ArrayLiteralExpression production.
	ExitArrayLiteralExpression(c *ArrayLiteralExpressionContext)

	// ExitMemberIndexExpression is called when exiting the MemberIndexExpression production.
	ExitMemberIndexExpression(c *MemberIndexExpressionContext)

	// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitBitAndExpression is called when exiting the BitAndExpression production.
	ExitBitAndExpression(c *BitAndExpressionContext)

	// ExitBitOrExpression is called when exiting the BitOrExpression production.
	ExitBitOrExpression(c *BitOrExpressionContext)

	// ExitRecursiveDescentMemberIndexExpression is called when exiting the RecursiveDescentMemberIndexExpression production.
	ExitRecursiveDescentMemberIndexExpression(c *RecursiveDescentMemberIndexExpressionContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitElementList is called when exiting the elementList production.
	ExitElementList(c *ElementListContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitPropertyExpressionAssignment is called when exiting the PropertyExpressionAssignment production.
	ExitPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext)

	// ExitComputedPropertyExpressionAssignment is called when exiting the ComputedPropertyExpressionAssignment production.
	ExitComputedPropertyExpressionAssignment(c *ComputedPropertyExpressionAssignmentContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)
}
