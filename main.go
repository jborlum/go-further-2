package main

import "fmt"

// ---------------------------------
// Types used for the example.
// ---------------------------------
// valuePrinter uses non-pointer (T) receiver.
type valuePrinter struct{}

func (p valuePrinter) print(str string) {
	fmt.Printf("%v\n", str)
}

// pointerPrinter uses pointer (*T) receiver.
type pointerPrinter struct{}

func (p *pointerPrinter) print(str string) {
	fmt.Printf("%v\n", str)
}

func main() {
	// Create non-pointer printer instances with value and pointer reciever.
	printer1 := valuePrinter{}
	printer2 := pointerPrinter{}

	text := "Epic poem"

	// To really understand how interfaces work in Go we first need to look
	// at how methods are called.

	// ---------------------------------
	// Calling using method expression.
	// Has the form 'ReceiverType.MethodName' and yields a function type.
	// It only holds the function reference NOT the receiver which you will
	// need to supply manually. Which is why they are so great for
	// demonstrating how receivers work.
	//
	// <method-expr>   ::= <receiver-type> "." <method-name>
	// <receiver-type> ::= <type-name> | "(" "*" <type-name> ")" | "(" <receiver-type> ")"
	// ---------------------------------
	{
		valuePrinter.print(printer1, text)
		f1 := valuePrinter.print // func(valuePrinter, string)
		f1(printer1, text)

		(valuePrinter).print(printer1, text)
		f2 := (valuePrinter).print // func(valuePrinter, string)
		f2(printer1, text)

		// Notice here that you explicitly define the expected receiver to be of pointer-type.
		// This is required to be able to call methods which expect pointer-receivers.
		//
		// pointerPrinter.print(&printer2, text) // Wouldn't work!
		(*pointerPrinter).print(&printer2, text)
		f3 := (*pointerPrinter).print // func(*pointerPrinter, string)
		f3(&printer2, text)
	}

	// ---------------------------------
	// Calling using method values.
	// Has the form 'x.MethodName' where the expression has static type T.
	// T can be an interface or non-interface type.
	// A method value also stores the receiver too, so when you call it,
	// you don't have to pass a receiver to it.
	//
	// This is the way we call most of our methods.
	// ---------------------------------
	{
		// Calling a method expecting a value receiver works as expected.
		printer1.print(text)
		f1 := printer1.print // Value reciever is bound.
		f1(text)

		// A reference to a non-interface method with a value receiver using a
		// pointer will automatically dereference that pointer.
		// x.Mv is equivalent to (*x).Mv where Mv is a method with a value reciever.
		printer1Ptr := &printer1
		printer1Ptr.print(text)
		(*printer1Ptr).print(text)

		// As with method calls, a reference to a non-interface method with a pointer
		// receiver using an addressable value will automatically take the address of that value.
		// x.Mp is equivalent to (&x).Mp where Mp is a method with a pointer reciever.
		printer2.print(text)
		f2 := printer2.print // Pointer to value is bound as reciever.
		f2(text)

		// Calling a method expecting a pointer receiver works as expected.
		printer2Ptr := &printer2
		printer2Ptr.print(text)
		(printer2Ptr).print(text)
	}

	// ---------------------------------
	// Interfaces 101:
	// So what is an interface? An interface is two things: it is a set of methods,
	// but it is also a type. Interfaces are implicitly satisfied by types.
	//
	// An interface value in Go is a value like any other. The type contains two fields:
	// - A pointer to a table of methods implemented by the underlaying type (vtable).
	// - A pointer to the actual data.
	//
	// Interface definition does not prescribe whether an implementor should implement
	// the interface using a pointer receiver or a value receiver. When you are given
	// an interface value, there’s no guarantee whether the underlying type is or
	// isn’t a pointer.
	// This can be quite confusing when learning Go.
	// ---------------------------------
	type printer interface {
		print(string)
	}

	// This works because a pointer type can access the methods of its associated value type, but not vice versa.
	var iPrinter1Value printer = printer1
	var iPrinter1Ptr printer = &printer1

	// pointerPrinter does not implement printer because it requires a pointer reciever.
	// var iPrinter2Value printer = printer2 // Wouldn't work!
	var iPrinter2Ptr printer = &printer2

	// ---------------------------------
	// Examples using method expressions.
	// It is possible to call interface type just like non-interface types using
	// method expressions.
	// ---------------------------------
	{
		printer.print(printer1, text)
		f1 := printer.print // func(valuePrinter, string)
		f1(printer1, text)

		(printer).print(printer1, text)
		f2 := (printer).print // func(valuePrinter, string)
		f2(printer1, text)

		// However unlike non-interface types it is not possible to
		// specify the reciever type directly. Passed recievers will have
		// to match they implementing methods reciever types.
		//
		// printer.print(printer2, text) // Wouldn't work!
		printer.print(&printer2, text)
	}

	// ---------------------------------
	// Examples using method values.
	// ---------------------------------
	{
		// Calling a method on a interface works as expected.
		iPrinter1Value.print(text)
		iPrinter1Ptr.print(text)

		f1 := iPrinter1Value.print // Value reciever is bound.
		f2 := iPrinter1Ptr.print   // Pointer reciever is bound.
		f1(text)
		f2(text)

		iPrinter2Ptr.print(text)
		f3 := iPrinter2Ptr.print // Pointer reciever is bound.
		f3(text)
	}

	// ---------------------------------
	// Calling via interfaces.
	// ---------------------------------
	{
		// Function expecting an interface value of type printer.
		printFunc := func(str string, p printer) {
			p.print(str)
		}

		// Passing interfaces bound to non-interface type with value reciever
		// works as expected.
		printFunc(text, printer1)
		printFunc(text, iPrinter1Value)
		printFunc(text, iPrinter1Ptr)

		// Because the interface doesn't know about reciever types of the implementing
		// type the correct type have to be passed.
		//
		// printFunc(text2, printer2) // Wouldn't work!
		printFunc(text, &printer2)
		printFunc(text, iPrinter2Ptr)
	}
}
