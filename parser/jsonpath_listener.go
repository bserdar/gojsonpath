// Code generated from Jsonpath.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Jsonpath

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JsonpathListener is a complete listener for a parse tree produced by JsonpathParser.
type JsonpathListener interface {
	antlr.ParseTreeListener

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterRoot_selector is called when entering the root_selector production.
	EnterRoot_selector(c *Root_selectorContext)

	// EnterCurrent_element_selector is called when entering the current_element_selector production.
	EnterCurrent_element_selector(c *Current_element_selectorContext)

	// EnterChild_selector is called when entering the child_selector production.
	EnterChild_selector(c *Child_selectorContext)

	// EnterRecursive_descent is called when entering the recursive_descent production.
	EnterRecursive_descent(c *Recursive_descentContext)

	// EnterPath_element is called when entering the path_element production.
	EnterPath_element(c *Path_elementContext)

	// EnterBracketed_selector is called when entering the bracketed_selector production.
	EnterBracketed_selector(c *Bracketed_selectorContext)

	// EnterUnion is called when entering the union production.
	EnterUnion(c *UnionContext)

	// EnterUnionPart is called when entering the unionPart production.
	EnterUnionPart(c *UnionPartContext)

	// EnterSlice is called when entering the slice production.
	EnterSlice(c *SliceContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterExpressionSequence is called when entering the expressionSequence production.
	EnterExpressionSequence(c *ExpressionSequenceContext)

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

	// EnterArgumentsExpression is called when entering the ArgumentsExpression production.
	EnterArgumentsExpression(c *ArgumentsExpressionContext)

	// EnterUnaryMinusExpression is called when entering the UnaryMinusExpression production.
	EnterUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// EnterPathExpression is called when entering the PathExpression production.
	EnterPathExpression(c *PathExpressionContext)

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

	// EnterCoalesceExpression is called when entering the CoalesceExpression production.
	EnterCoalesceExpression(c *CoalesceExpressionContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterElementList is called when entering the elementList production.
	EnterElementList(c *ElementListContext)

	// EnterArrayElement is called when entering the arrayElement production.
	EnterArrayElement(c *ArrayElementContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterPropertyExpressionAssignment is called when entering the PropertyExpressionAssignment production.
	EnterPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext)

	// EnterComputedPropertyExpressionAssignment is called when entering the ComputedPropertyExpressionAssignment production.
	EnterComputedPropertyExpressionAssignment(c *ComputedPropertyExpressionAssignmentContext)

	// EnterPropertyShorthand is called when entering the PropertyShorthand production.
	EnterPropertyShorthand(c *PropertyShorthandContext)

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

	// ExitRoot_selector is called when exiting the root_selector production.
	ExitRoot_selector(c *Root_selectorContext)

	// ExitCurrent_element_selector is called when exiting the current_element_selector production.
	ExitCurrent_element_selector(c *Current_element_selectorContext)

	// ExitChild_selector is called when exiting the child_selector production.
	ExitChild_selector(c *Child_selectorContext)

	// ExitRecursive_descent is called when exiting the recursive_descent production.
	ExitRecursive_descent(c *Recursive_descentContext)

	// ExitPath_element is called when exiting the path_element production.
	ExitPath_element(c *Path_elementContext)

	// ExitBracketed_selector is called when exiting the bracketed_selector production.
	ExitBracketed_selector(c *Bracketed_selectorContext)

	// ExitUnion is called when exiting the union production.
	ExitUnion(c *UnionContext)

	// ExitUnionPart is called when exiting the unionPart production.
	ExitUnionPart(c *UnionPartContext)

	// ExitSlice is called when exiting the slice production.
	ExitSlice(c *SliceContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitExpressionSequence is called when exiting the expressionSequence production.
	ExitExpressionSequence(c *ExpressionSequenceContext)

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

	// ExitArgumentsExpression is called when exiting the ArgumentsExpression production.
	ExitArgumentsExpression(c *ArgumentsExpressionContext)

	// ExitUnaryMinusExpression is called when exiting the UnaryMinusExpression production.
	ExitUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// ExitPathExpression is called when exiting the PathExpression production.
	ExitPathExpression(c *PathExpressionContext)

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

	// ExitCoalesceExpression is called when exiting the CoalesceExpression production.
	ExitCoalesceExpression(c *CoalesceExpressionContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitElementList is called when exiting the elementList production.
	ExitElementList(c *ElementListContext)

	// ExitArrayElement is called when exiting the arrayElement production.
	ExitArrayElement(c *ArrayElementContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitPropertyExpressionAssignment is called when exiting the PropertyExpressionAssignment production.
	ExitPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext)

	// ExitComputedPropertyExpressionAssignment is called when exiting the ComputedPropertyExpressionAssignment production.
	ExitComputedPropertyExpressionAssignment(c *ComputedPropertyExpressionAssignmentContext)

	// ExitPropertyShorthand is called when exiting the PropertyShorthand production.
	ExitPropertyShorthand(c *PropertyShorthandContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)
}
