# Math Interpreter

To view the Github Repository <a href=https://github.com/colorfulparadox/Custom-Math-Interpreter target=_blank>click here</a>


An interpreter reads input and executes instructions without converting it into ASM or intermediate machine code.

This project was to take a simple context-free grammar (CFG), analyze input, parse through the input, and generate instructions to be executed.

The CFG can be seen below:
```
NUM  -> integer+
ID   -> char | char char | ε
OP   -> + | - | * | / 

expr -> NUM | ID
expr -> expr OP expr    
expr -> ( expr OP expr )

var-init   -> let ID = expr;
var-assign ->  ID = expr;
var-access ->  ID

output -> print(var-access|expr)
```
<sub>If you are unfamiliar with what a CFG is, I recommend you read Introduction To Computer Theory by Michael Sipser.</sub>

## Implementations and Examples

To begin developing an interpreter I first needed to write a Lexical Analyzer. A Lexical Analyzer takes in my input and then tokenizes and creates lexemes. From here we can dive into the parser. Currently, my parser could use some work. It currently does syntax, semantic, and code generation. Normally the parser does syntax and then after you would do semantic checking along with code generation. I wish to refactor these functions eventually.

The other big part of this project was implementing operator precedence. Operator precedence is externally making sure that the order of operations is correctly followed.

Another smaller implementation detail is that I used C++ to develop this. I also use Make to compile the project. I followed OOP principles such as polymorphism and composition to build out larger objects.

Below is an example of the input that could be passed to the interpreter.
```
let a = 2 + 2;
let b = 5;
b = b + a * (48/24);
print(b);
```
This would output 13.

If you wish to look at the code the repository is linked at the top of this document.