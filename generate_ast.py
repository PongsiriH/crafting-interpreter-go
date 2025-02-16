exprTypes = [
    "Binary   : Expr left, Token operator, Expr right",
    "Grouping : Expr expression",
    "Literal  : Object value",
    "Unary    : Token operator, Expr right",
]

with open("expr.go", "w") as f:
    f.write("package main\n\n")

    # base expr interface
    f.write("""
type Expr[T any] interface {
    Accept(visitor Visitor[T]) T
}
            """)

    # visitor interface
    visitor_str = ""
    for exprType in exprTypes:
        className, fields = exprType.split(":")
        className = className.strip()
        visitor_str += f"\tVisit{className}Expr(expr *{className}Expr[T]) T\n"
    f.write(f"""
type Visitor[T any] interface {{
{ visitor_str }}}
        """)

    for exprType in exprTypes:
        className, fields = exprType.split(":")
        className = className.strip()
        fields = [field.strip() for field in fields.split(",")]

        # fields
        fields_str = ""
        for field in fields:
            tokenType, name = field.split(" ")
            tokenType = tokenType.capitalize()
            if tokenType == "Expr":
                tokenType = "Expr[T]"
            if tokenType == "Object":
                tokenType = "T"
            name = name.capitalize()
            fields_str += f"\t{name} {tokenType}\n"
        f.write(f"""
type {className}Expr[T any] struct {{
{fields_str}}}
                """)

        f.write(f"""
func (expr * {className}Expr[T]) Accept(visitor Visitor[T]) T {{
    return visitor.Visit{className}Expr(expr)}}
                """)
