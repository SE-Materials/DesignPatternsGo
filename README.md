# Characteristics of Go lang that will affect Design Patterns

- No support for traditional OOP classes
    - No static members or constrcutors : affects singleton

- No support for class based inheritance or method overloading
    - Affects pattern like visitor

- No support for exceptions - error handling is via returns values
    - Affects patterns like builder

- No support for abstract classes
    - Affects patterns like Abstract factory and Bridge 

# Design Pattern Categories.

|Creational | Structural | Behavioral |
|-----------|------------|------------|
|Abstract Factory| **Adapter** | Chain of Responsibility |
|**Builder**| Bridge | Command |
|**Factory**| Composite | Interpreter |
|Lazy Initialization | Decorator | **Iterator** |
|Multiton | **Facade** | Mediator |
|Object Pool | Flyweight | Memento |
|Prototype | Proxy | **Observer** |
|**Singeton** | | Strategy |
| | | Visitor |