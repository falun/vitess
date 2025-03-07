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
	"log"

	"github.com/dave/jennifer/jen"
)

// cloneGen creates the deep clone methods for the AST. It works by discovering the types that it needs to support,
// starting from a root interface type. While creating the clone method for this root interface, more types that need
// to be cloned are discovered. This continues type by type until all necessary types have been traversed.
type cloneGen struct {
	exceptType string
}

var _ generator2 = (*cloneGen)(nil)

func newCloneGen(exceptType string) *cloneGen {
	return &cloneGen{
		exceptType: exceptType,
	}
}

const cloneName = "Clone"

// readValueOfType produces code to read the expression of type `t`, and adds the type to the todo-list
func (c *cloneGen) readValueOfType(t types.Type, expr jen.Code, spi generatorSPI) jen.Code {
	switch t.Underlying().(type) {
	case *types.Basic:
		return expr
	case *types.Interface:
		if types.TypeString(t, noQualifier) == "interface{}" {
			// these fields have to be taken care of manually
			return expr
		}
	}
	spi.addType(t)
	return jen.Id(cloneName + printableTypeName(t)).Call(expr)
}

func (c *cloneGen) structMethod(t types.Type, _ *types.Struct, spi generatorSPI) error {
	typeString := types.TypeString(t, noQualifier)
	funcName := cloneName + printableTypeName(t)
	spi.addFunc(funcName, clone,
		jen.Func().Id(funcName).Call(jen.Id("n").Id(typeString)).Id(typeString).Block(
			jen.Return(jen.Op("*").Add(c.readValueOfType(types.NewPointer(t), jen.Op("&").Id("n"), spi))),
		))
	return nil
}

func (c *cloneGen) sliceMethod(t types.Type, slice *types.Slice, spi generatorSPI) error {
	typeString := types.TypeString(t, noQualifier)
	name := printableTypeName(t)
	funcName := cloneName + name

	spi.addFunc(funcName, clone,
		//func (n Bytes) Clone() Bytes {
		jen.Func().Id(funcName).Call(jen.Id("n").Id(typeString)).Id(typeString).Block(
			//	res := make(Bytes, len(n))
			jen.Id("res").Op(":=").Id("make").Call(jen.Id(typeString), jen.Lit(0), jen.Id("len").Call(jen.Id("n"))),
			c.copySliceElement(slice.Elem(), spi),
			//	return res
			jen.Return(jen.Id("res")),
		))
	return nil
}

func (c *cloneGen) basicMethod(t types.Type, basic *types.Basic, spi generatorSPI) error {
	return nil
}

func (c *cloneGen) copySliceElement(elType types.Type, spi generatorSPI) jen.Code {
	if isBasic(elType) {
		//	copy(res, n)
		return jen.Id("copy").Call(jen.Id("res"), jen.Id("n"))
	}

	//for _, x := range n {
	//  res = append(res, CloneAST(x))
	//}
	spi.addType(elType)

	return jen.For(jen.List(jen.Op("_"), jen.Id("x"))).Op(":=").Range().Id("n").Block(
		jen.Id("res").Op("=").Id("append").Call(jen.Id("res"), c.readValueOfType(elType, jen.Id("x"), spi)),
	)
}

func (c *cloneGen) interfaceMethod(t types.Type, iface *types.Interface, spi generatorSPI) error {

	//func CloneAST(in AST) AST {
	//	if in == nil {
	//	return nil
	//}
	//	switch in := in.(type) {
	//case *RefContainer:
	//	return in.CloneRefOfRefContainer()
	//}
	//	// this should never happen
	//	return nil
	//}

	typeString := types.TypeString(t, noQualifier)
	typeName := printableTypeName(t)

	stmts := []jen.Code{ifNilReturnNil("in")}

	var cases []jen.Code
	_ = findImplementations(spi.scope(), iface, func(t types.Type) error {
		typeString := types.TypeString(t, noQualifier)

		// case Type: return CloneType(in)
		block := jen.Case(jen.Id(typeString)).Block(jen.Return(c.readValueOfType(t, jen.Id("in"), spi)))
		switch t := t.(type) {
		case *types.Pointer:
			_, isIface := t.Elem().(*types.Interface)
			if !isIface {
				cases = append(cases, block)
			}

		case *types.Named:
			_, isIface := t.Underlying().(*types.Interface)
			if !isIface {
				cases = append(cases, block)
			}

		default:
			log.Fatalf("unexpected type encountered: %s", typeString)
		}

		return nil
	})

	cases = append(cases,
		jen.Default().Block(
			jen.Comment("this should never happen"),
			jen.Return(jen.Nil()),
		))

	//	switch n := node.(type) {
	stmts = append(stmts, jen.Switch(jen.Id("in").Op(":=").Id("in").Assert(jen.Id("type")).Block(
		cases...,
	)))

	funcName := cloneName + typeName
	funcDecl := jen.Func().Id(funcName).Call(jen.Id("in").Id(typeString)).Id(typeString).Block(stmts...)
	spi.addFunc(funcName, clone, funcDecl)
	return nil
}

func (c *cloneGen) ptrToBasicMethod(t types.Type, _ *types.Basic, spi generatorSPI) error {
	ptr := t.Underlying().(*types.Pointer)
	return c.ptrToOtherMethod(t, ptr, spi)
}

func (c *cloneGen) ptrToOtherMethod(t types.Type, ptr *types.Pointer, spi generatorSPI) error {
	receiveType := types.TypeString(t, noQualifier)

	funcName := "Clone" + printableTypeName(t)
	spi.addFunc(funcName, clone,
		jen.Func().Id(funcName).Call(jen.Id("n").Id(receiveType)).Id(receiveType).Block(
			ifNilReturnNil("n"),
			jen.Id("out").Op(":=").Add(c.readValueOfType(ptr.Elem(), jen.Op("*").Id("n"), spi)),
			jen.Return(jen.Op("&").Id("out")),
		))
	return nil
}

func ifNilReturnNil(id string) *jen.Statement {
	return jen.If(jen.Id(id).Op("==").Nil()).Block(jen.Return(jen.Nil()))
}

func isBasic(t types.Type) bool {
	_, x := t.Underlying().(*types.Basic)
	return x
}

func (c *cloneGen) ptrToStructMethod(t types.Type, strct *types.Struct, spi generatorSPI) error {
	receiveType := types.TypeString(t, noQualifier)
	funcName := cloneName + printableTypeName(t)

	//func CloneRefOfType(n *Type) *Type
	funcDeclaration := jen.Func().Id(funcName).Call(jen.Id("n").Id(receiveType)).Id(receiveType)

	if receiveType == c.exceptType {
		spi.addFunc(funcName, clone, funcDeclaration.Block(
			jen.Return(jen.Id("n")),
		))
		return nil
	}

	var fields []jen.Code
	for i := 0; i < strct.NumFields(); i++ {
		field := strct.Field(i)
		if isBasic(field.Type()) || field.Name() == "_" {
			continue
		}
		// out.Field = CloneType(n.Field)
		fields = append(fields,
			jen.Id("out").Dot(field.Name()).Op("=").Add(c.readValueOfType(field.Type(), jen.Id("n").Dot(field.Name()), spi)))
	}

	stmts := []jen.Code{
		// if n == nil { return nil }
		ifNilReturnNil("n"),
		// 	out := *n
		jen.Id("out").Op(":=").Op("*").Id("n"),
	}

	// handle all fields with CloneAble types
	stmts = append(stmts, fields...)

	stmts = append(stmts,
		// return &out
		jen.Return(jen.Op("&").Id("out")),
	)

	spi.addFunc(funcName, clone,
		funcDeclaration.Block(stmts...),
	)
	return nil
}
