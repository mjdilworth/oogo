# Go and Object-Oriented Programming

Whilst go has no classes, objects, exceptions, templates, type hierarchy it is very much an object oriented language.

Go does have built in concurrency and garbage collection.

## Go has Types
Structs - user defined types with methods

Type Creature struct {
Name string
Real bool
}

func (c Creature) Dump() {
  fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}
## Embedding
type FlyingCreature struct {
Creature
WingSpan int
}
dragon := &FlyingCreature{
    Creature{"Dragon", false, },
    15,
}
fmt.Println(dragon.Name)
fmt.Println(dragon.Real)
fmt.Println(dragon.WingSpan)

## Interfaces
Interfaces are types that declare sets of methods. Similarly to interfaces in other languages, they have no implementation.
Typess that implement all the interface methods automatically implement the interface. There is no inheritance or subclassing or "implements" keyword. In the following code snippet, type Foo implements the Fooer interface (by convention, Go interface names end with "er").

type Fooer interface {
  Foo1()
  Foo2()
  Foo3()
}
 
type Foo struct {
}
 
func (f Foo) Foo1() {
    fmt.Println("Foo1() here")
}
 
func (f Foo) Foo2() {
    fmt.Println("Foo2() here")
}
 
func (f Foo) Foo3() {
    fmt.Println("Foo3() here")
}

## encapsulation, inheritance, and polymorphism

## Encapsulation
Done at package level. Names starting with lower case are only visible within the package. You can hide anything in a private package and just expose specific types, interfaces, and factory functions
## Inheritance
Modern languages and object-oriented thinking now favor composition over inheritance.
For all intents and purposes, composition by embedding an anonymous type is equivalent to implementation inheritance. An embedded struct is just as fragile as a base class. You can also embed an interface, which is equivalent to inheriting from an interface in languages like Java or C++. It can even lead to a runtime error that is not discovered at compile time if the embedding type doesn't implement all the interface methods.
## Polymorphism
Treat objects of different types uniformly as long as they adhere to the same interface.
