What I learn through building each thing
### Scanner
- This section is quite tedious. One nice pattern is having StartIndex pointing at starting charcter of the "token" and CurrentIndex to point to the last character of the "token", then just take substring from SourceString.

### Expr and Visitor 
- Having **visitor pattern** that should be able to return many types in static type language is quite difficult. Maybe I did it the wrong way? or maybe this pattern just does not work well in Go?

### Parser
- Well, this section is probably the most challenging one so far with recursive calls between Expression, Equality, Comparision, Term, Factor, Unary and Primary.
- This is when my codebase is becoming quite complex and probably need to start adding error handlings.
