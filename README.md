# NFA-Builder

**Brian Doyle - G00330969**
---
A program which can build a non-deterministic finite automaton (NFA) from a regular expression, and also can use the NFA to check if the regular expression matches any given string of text.

Program written for a project in Graph Theory.

### Problem Statement
You  must  write  a  program  in  the  Go  programming  language  [2]  that  can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text.  You must write the program from scratch and cannot use the regexp
package from the Go standard library nor any other external library.


A  regular  expression  is  a  string  containing  a  series  of  characters,  some
of  which  may  have  a  special  meaning.   For  example,  the  three  characters “.”, “|”, and “∗” 
have the special meanings “concatenate”, “or”, and “Kleene star” respectively.  So, 0.1 means
a 0 followed by a 1, 0|1 means a 0 or a 1,and 1∗means any number of 1’s.  
These special characters must be used in your submission.


### Breakdown of problem
The problem can be broken down into multiple parts:

First the [Shunting Yard Algorithm] must be used to convert infix regular expression to postfix regular expression.

Next [Thompson's Construction] is used to construct a non-deterministic finite automaton (NFA) from the regular expression.

Finally after taking in the string to compare to the regular expression, the NFA must be navigated to determine if it is a match.


Here are some great videos that explain the process in depth.

Shunting Yard Algorithm: (https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e)

Thompson's Construction: (https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b)

Compare String to NFA: (https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d)


[Shunting Yard Algorithm]: https://brilliant.org/wiki/shunting-yard-algorithm/
[Thompson's Construction]: https://www.cs.york.ac.uk/fp/lsa/lectures/REToC.pdf
