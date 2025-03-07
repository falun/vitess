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

const visitName = "Visit"

type visitGen struct{}

var _ generator2 = (*visitGen)(nil)

func shouldAdd(t types.Type, i *types.Interface) bool {
	return types.Implements(t, i)
}

func (e visitGen) interfaceMethod(t types.Type, iface *types.Interface, spi generatorSPI) error {
	if !shouldAdd(t, spi.iface()) {
		return nil
	}
	/*
		func VisitAST(in AST) (bool, error) {
			if in == nil {
				return false, nil
			}
			switch a := inA.(type) {
			case *SubImpl:
				return VisitSubImpl(a, b)
			default:
				return false, nil
			}
		}
	*/
	stmts := []jen.Code{
		jen.If(jen.Id("in == nil").Block(returnNil())),
	}

	var cases []jen.Code
	_ = spi.findImplementations(iface, func(t types.Type) error {
		if _, ok := t.Underlying().(*types.Interface); ok {
			return nil
		}
		typeString := types.TypeString(t, noQualifier)
		funcName := visitName + printableTypeName(t)
		spi.addType(t)
		caseBlock := jen.Case(jen.Id(typeString)).Block(
			jen.Return(jen.Id(funcName).Call(jen.Id("in"), jen.Id("f"))),
		)
		cases = append(cases, caseBlock)
		return nil
	})

	cases = append(cases,
		jen.Default().Block(
			jen.Comment("this should never happen"),
			returnNil(),
		))

	stmts = append(stmts, jen.Switch(jen.Id("in := in.(type)").Block(
		cases...,
	)))

	visitFunc(t, stmts, spi)
	return nil
}

func returnNil() jen.Code {
	return jen.Return(jen.Nil())
}

func (e visitGen) structMethod(t types.Type, strct *types.Struct, spi generatorSPI) error {
	if !shouldAdd(t, spi.iface()) {
		return nil
	}

	/*
		func VisitRefOfRefContainer(in *RefContainer, f func(node AST) (kontinue bool, err error)) (bool, error) {
			if cont, err := f(in); err != nil || !cont {
				return false, err
			}
			if k, err := VisitRefOfLeaf(in.ASTImplementationType, f); err != nil || !k {
				return false, err
			}
			if k, err := VisitAST(in.ASTType, f); err != nil || !k {
				return false, err
			}
			return true, nil
		}
	*/

	stmts := visitAllStructFields(strct, spi)
	visitFunc(t, stmts, spi)

	return nil
}

func (e visitGen) ptrToStructMethod(t types.Type, strct *types.Struct, spi generatorSPI) error {
	if !shouldAdd(t, spi.iface()) {
		return nil
	}

	/*
		func VisitRefOfRefContainer(in *RefContainer, f func(node AST) (kontinue bool, err error)) (bool, error) {
			if in == nil {
				return true, nil
			}
			if cont, err := f(in); err != nil || !cont {
				return false, err
			}
			if k, err := VisitRefOfLeaf(in.ASTImplementationType, f); err != nil || !k {
				return false, err
			}
			if k, err := VisitAST(in.ASTType, f); err != nil || !k {
				return false, err
			}
			return true, nil
		}
	*/

	stmts := []jen.Code{
		jen.If(jen.Id("in == nil").Block(returnNil())),
	}
	stmts = append(stmts, visitAllStructFields(strct, spi)...)
	visitFunc(t, stmts, spi)

	return nil
}

func (e visitGen) ptrToBasicMethod(t types.Type, _ *types.Basic, spi generatorSPI) error {
	if !shouldAdd(t, spi.iface()) {
		return nil
	}

	stmts := []jen.Code{
		jen.Comment("ptrToBasicMethod"),
	}

	visitFunc(t, stmts, spi)

	return nil
}

func (e visitGen) sliceMethod(t types.Type, slice *types.Slice, spi generatorSPI) error {
	if !shouldAdd(t, spi.iface()) {
		return nil
	}

	if !shouldAdd(slice.Elem(), spi.iface()) {
		return e.visitNoChildren(t, spi)
	}

	stmts := []jen.Code{
		jen.If(jen.Id("in == nil").Block(returnNil())),
		visitIn(),
		jen.For(jen.Id("_, el := range in")).Block(
			visitChild(slice.Elem(), jen.Id("el")),
		),
		returnNil(),
	}

	visitFunc(t, stmts, spi)

	return nil
}

func (e visitGen) ptrToOtherMethod(t types.Type, _ *types.Pointer, spi generatorSPI) error {
	if !shouldAdd(t, spi.iface()) {
		return nil
	}

	stmts := []jen.Code{
		jen.Comment("ptrToOtherMethod "),
	}

	visitFunc(t, stmts, spi)

	return nil
}

func (e visitGen) basicMethod(t types.Type, basic *types.Basic, spi generatorSPI) error {
	if !shouldAdd(t, spi.iface()) {
		return nil
	}

	return e.visitNoChildren(t, spi)
}

func (e visitGen) visitNoChildren(t types.Type, spi generatorSPI) error {
	stmts := []jen.Code{
		jen.Id("_, err := f(in)"),
		jen.Return(jen.Err()),
	}

	visitFunc(t, stmts, spi)

	return nil
}

func visitAllStructFields(strct *types.Struct, spi generatorSPI) []jen.Code {
	output := []jen.Code{
		visitIn(),
	}
	for i := 0; i < strct.NumFields(); i++ {
		field := strct.Field(i)
		if types.Implements(field.Type(), spi.iface()) {
			spi.addType(field.Type())
			visitField := visitChild(field.Type(), jen.Id("in").Dot(field.Name()))
			output = append(output, visitField)
			continue
		}
		slice, isSlice := field.Type().(*types.Slice)
		if isSlice && types.Implements(slice.Elem(), spi.iface()) {
			spi.addType(slice.Elem())
			output = append(output, jen.For(jen.Id("_, el := range in."+field.Name())).Block(
				visitChild(slice.Elem(), jen.Id("el")),
			))
		}
	}
	output = append(output, returnNil())
	return output
}

func visitChild(t types.Type, id jen.Code) *jen.Statement {
	funcName := visitName + printableTypeName(t)
	visitField := jen.If(
		jen.Id("err := ").Id(funcName).Call(id, jen.Id("f")),
		jen.Id("err != nil "),
	).Block(jen.Return(jen.Err()))
	return visitField
}

func visitIn() *jen.Statement {
	return jen.If(
		jen.Id("cont, err := ").Id("f").Call(jen.Id("in")),
		jen.Id("err != nil || !cont"),
	).Block(jen.Return(jen.Err()))
}

func visitFunc(t types.Type, stmts []jen.Code, spi generatorSPI) {
	typeString := types.TypeString(t, noQualifier)
	funcName := visitName + printableTypeName(t)
	code := jen.Func().Id(funcName).Call(jen.Id("in").Id(typeString), jen.Id("f Visit")).Error().Block(stmts...)
	spi.addFunc(funcName, visit, code)
}
