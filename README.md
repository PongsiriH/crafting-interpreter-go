What I feel through building each thing
### Scanner
- This section is quite tedious.
- One nice pattern is having StartIndex pointing at starting charcter of the "token" and CurrentIndex to point to the last character of the "token", then just take substring from SourceString. This approach make it avoid a lot of overhead in building string and much easier to backtrack if needed. 
- One thing to be careful of when working with bytes is how it is read and converted. For instance, newline "\n" can be translated to byte(10) or byte('\\') + byte('n') depending on how you convert it. Different OS also handle these special characters differently.

### Expr and Visitor 
- Having **visitor pattern** that should be able to return many types in static type language is quite difficult. Maybe I did it the wrong way? or maybe this pattern just does not work well in Go?

### Parser
- Well, this section is probably the most challenging one so far with recursive calls between Expression, Equality, Comparision, Term, Factor, Unary and Primary.
- This is when my codebase is becoming quite complex and probably need to start adding error handlings.

### Statements
- This section pretty much make our interpreter look like an actual one... Previously it just seem like calculator that only accept single expression.
- Now, the interpreter can have a list of statements and executable.

### Printing, Declaring variable and Assigning variable, Scoping
- Defined printing statements which, as the name implied, print to the stdout.
- Defined `environment` which store variables which we can `Define`, `Assign`, and `Read`.
- *Scoping by recursively pointing back to the parent environment. In Java, author use `finally` keywords, but we use `defer` here... which feel nice somehow :).

### Block, If-else
- I am thinking of enforcing curling braces {} after `if` statement instead of allowing single-line without curly braces.. should I?
### Logical expression, While-loop, For-loop
- FOR LOOP IS JUST SO TEDIUS!
- I think I am getting used to this now? Maybe not so much... we have like 4 chapters left for golox so let not be overconfident.
