package evaluator

import (
	"../object"
)

func evalSetsInfixExpression(operator string,
	left, right object.Object,
) object.Object {

	switch operator {
	case "+":
	case "|":
		return  &object.Sets{Elements: *unionSets(
			&left.(*object.Sets).Elements,
			&right.(*object.Sets).Elements,
		)}
	case "&":
		return &object.Sets{Elements: *crossSets(
		&left.(*object.Sets).Elements,
		&right.(*object.Sets).Elements,
		)}
	case "\\":
		return &object.Sets{Elements: *diffSets(
			&left.(*object.Sets).Elements,
			&right.(*object.Sets).Elements,
		)}

	}

	return newError("unknown operator: %s %s %s",
		left.Type(), operator, right.Type())
}

func unionSets(a *[]object.Object, b *[]object.Object) *[]object.Object {
	has := make(map[string]bool)
	elements := new([]object.Object)
	for _, element := range append(*a, *b...) {
		if _, ok := has[element.Inspect()]; !ok {
			*elements = append(*elements, element)
			has[element.Inspect()] = true
		}
	}
	return  elements
}

func crossSets(a *[]object.Object, b *[]object.Object) *[]object.Object {
	has := make(map[string]int)
	elements := new([]object.Object)
	for _, element := range append(*a, *b...) {
		if count, ok := has[element.Inspect()]; ok && count == 1 {
			*elements = append(*elements, element)
		}
		has[element.Inspect()] ++
	}

	return  elements
}

func diffSets(a *[]object.Object, b *[]object.Object) *[]object.Object {

	elements := new([]object.Object)
	for _, element := range *a {
		has := false
		for _, elementB := range *b {
			if element.Inspect() == elementB.Inspect() {
				has = true
				break
			}
		}

		if !has {
			*elements = append(*elements, element)
		}

	}

	return  elements
}