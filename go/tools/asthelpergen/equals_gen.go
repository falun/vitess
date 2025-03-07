/*
Copyright 2021 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package asthelpergen

import (
	"go/types"

	"github.com/dave/jennifer/jen"
)

const equalsName = "Equals"

type equalsGen struct{}

var _ generator2 = (*equalsGen)(nil)

func (e equalsGen) interfaceMethod(t types.Type, iface *types.Interface, spi generatorSPI) error {
	/*
		func EqualsAST(inA, inB AST) bool {
			if inA == inB {
				return true
			}
			if inA == nil || inB8 == nil {
				return false
			}
			switch a := inA.(type) {
			case *SubImpl:
				b, ok := inB.(*SubImpl)
				if !ok {
					return false
				}
				return EqualsSubImpl(a, b)
			}
			return false
		}
	*/
	stmts := []jen.Code{
		jen.If(jen.Id("inA == nil").Op("&&").Id("inB == nil")).Block(jen.Return(jen.True())),
		jen.If(jen.Id("inA == nil").Op("||").Id("inB == nil")).Block(jen.Return(jen.False())),
	}

	var cases []jen.Code
	_ = spi.findImplementations(iface, func(t types.Type) error {
		if _, ok := t.Underlying().(*types.Interface); ok {
			return nil
		}
		typeString := types.TypeString(t, noQualifier)
		caseBlock := jen.Case(jen.Id(typeString)).Block(
			jen.Id("b, ok := inB.").Call(jen.Id(typeString)),
			jen.If(jen.Id("!ok")).Block(jen.Return(jen.False())),
			jen.Return(compareValueType(t, jen.Id("a"), jen.Id("b"), true, spi)),
		)
		cases = append(cases, caseBlock)
		return nil
	})

	cases = append(cases,
		jen.Default().Block(
			jen.Comment("this should never happen"),
			jen.Return(jen.False()),
		))

	stmts = append(stmts, jen.Switch(jen.Id("a := inA.(type)").Block(
		cases...,
	)))

	typeString := types.TypeString(t, noQualifier)
	funcName := equalsName + printableTypeName(t)
	funcDecl := jen.Func().Id(funcName).Call(jen.List(jen.Id("inA"), jen.Id("inB")).Id(typeString)).Bool().Block(stmts...)
	spi.addFunc(funcName, equals, funcDecl)

	return nil
}

func compareValueType(t types.Type, a, b *jen.Statement, eq bool, spi generatorSPI) *jen.Statement {
	switch t.Underlying().(type) {
	case *types.Basic:
		if eq {
			return a.Op("==").Add(b)
		}
		return a.Op("!=").Add(b)
	}
	spi.addType(t)
	var neg = "!"
	if eq {
		neg = ""
	}
	return jen.Id(neg+equalsName+printableTypeName(t)).Call(a, b)
}

func (e equalsGen) structMethod(t types.Type, strct *types.Struct, spi generatorSPI) error {
	/*
		func EqualsRefOfRefContainer(inA RefContainer, inB RefContainer) bool {
			return EqualsRefOfLeaf(inA.ASTImplementationType, inB.ASTImplementationType) &&
				EqualsAST(inA.ASTType, inB.ASTType) && inA.NotASTType == inB.NotASTType
		}

	*/

	typeString := types.TypeString(t, noQualifier)
	funcName := equalsName + printableTypeName(t)
	funcDecl := jen.Func().Id(funcName).Call(jen.List(jen.Id("a"), jen.Id("b")).Id(typeString)).Bool().
		Block(jen.Return(compareAllStructFields(strct, spi)))
	spi.addFunc(funcName, equals, funcDecl)

	return nil
}

func compareAllStructFields(strct *types.Struct, spi generatorSPI) jen.Code {
	var basicsPred []*jen.Statement
	var others []*jen.Statement
	for i := 0; i < strct.NumFields(); i++ {
		field := strct.Field(i)
		if field.Type().Underlying().String() == "interface{}" || field.Name() == "_" {
			// we can safely ignore this, we do not want ast to contain interface{} types.
			continue
		}
		fieldA := jen.Id("a").Dot(field.Name())
		fieldB := jen.Id("b").Dot(field.Name())
		pred := compareValueType(field.Type(), fieldA, fieldB, true, spi)
		if _, ok := field.Type().(*types.Basic); ok {
			basicsPred = append(basicsPred, pred)
			continue
		}
		others = append(others, pred)
	}

	var ret *jen.Statement
	for _, pred := range basicsPred {
		if ret == nil {
			ret = pred
		} else {
			ret = ret.Op("&&").Line().Add(pred)
		}
	}

	for _, pred := range others {
		if ret == nil {
			ret = pred
		} else {
			ret = ret.Op("&&").Line().Add(pred)
		}
	}

	if ret == nil {
		return jen.True()
	}
	return ret
}

func (e equalsGen) ptrToStructMethod(t types.Type, strct *types.Struct, spi generatorSPI) error {
	typeString := types.TypeString(t, noQualifier)
	funcName := equalsName + printableTypeName(t)

	//func EqualsRefOfType(a,b  *Type) *Type
	funcDeclaration := jen.Func().Id(funcName).Call(jen.Id("a"), jen.Id("b").Id(typeString)).Bool()
	stmts := []jen.Code{
		jen.If(jen.Id("a == b")).Block(jen.Return(jen.True())),
		jen.If(jen.Id("a == nil").Op("||").Id("b == nil")).Block(jen.Return(jen.False())),
		jen.Return(compareAllStructFields(strct, spi)),
	}

	spi.addFunc(funcName, equals, funcDeclaration.Block(stmts...))
	return nil
}

func (e equalsGen) ptrToBasicMethod(t types.Type, _ *types.Basic, spi generatorSPI) error {
	/*
		func EqualsRefOfBool(a, b *bool) bool {
			if a == b {
				return true
			}
			if a == nil || b == nil {
				return false
			}
			return *a == *b
		}
	*/
	typeString := types.TypeString(t, noQualifier)
	funcName := equalsName + printableTypeName(t)

	//func EqualsRefOfType(a,b  *Type) *Type
	funcDeclaration := jen.Func().Id(funcName).Call(jen.Id("a"), jen.Id("b").Id(typeString)).Bool()
	stmts := []jen.Code{
		jen.If(jen.Id("a == b")).Block(jen.Return(jen.True())),
		jen.If(jen.Id("a == nil").Op("||").Id("b == nil")).Block(jen.Return(jen.False())),
		jen.Return(jen.Id("*a == *b")),
	}
	spi.addFunc(funcName, equals, funcDeclaration.Block(stmts...))
	return nil
}

func (e equalsGen) sliceMethod(t types.Type, slice *types.Slice, spi generatorSPI) error {
	/*
		func EqualsSliceOfRefOfLeaf(a, b []*Leaf) bool {
			if len(a) != len(b) {
				return false
			}
			for i := 0; i < len(a); i++ {
				if !EqualsRefOfLeaf(a[i], b[i]) {
					return false
				}
			}
			return false
		}
	*/

	stmts := []jen.Code{jen.If(jen.Id("len(a) != len(b)")).Block(jen.Return(jen.False())),
		jen.For(jen.Id("i := 0; i < len(a); i++")).Block(
			jen.If(compareValueType(slice.Elem(), jen.Id("a[i]"), jen.Id("b[i]"), false, spi)).Block(jen.Return(jen.False()))),
		jen.Return(jen.True()),
	}

	typeString := types.TypeString(t, noQualifier)
	funcName := equalsName + printableTypeName(t)
	funcDecl := jen.Func().Id(funcName).Call(jen.List(jen.Id("a"), jen.Id("b")).Id(typeString)).Bool().Block(stmts...)
	spi.addFunc(funcName, equals, funcDecl)
	return nil
}

func (e equalsGen) basicMethod(t types.Type, basic *types.Basic, spi generatorSPI) error {
	return nil
}

func (e equalsGen) ptrToOtherMethod(types.Type, *types.Pointer, generatorSPI) error {
	return nil
}
