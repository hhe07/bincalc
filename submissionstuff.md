# Section A: 
- Describe scenario
- Provide evidence of initial requirements meeting
- Describe proposed products
- Specific performance criteria

Solution title: bincalc

User Name: Edward Wawrzynek (Student's friend)

## Criterion A:
### Defining the Problem
My friend Edward **(client)** makes frequent use of bitwise math while programming, especially in hardware-design related contexts. For example, when this project was started, he was working on a chess competition framework for our school's robotics club. To save space, the chess boards were represented in integer form, with each bit representing whether a space was occupied or not. However, this made it necessary to calculate various "magic numbers" using bitwise math, and he was not satisfied with his calculator program's programmer mode. So, I decided to try to make a command line application to satisfy this need, with my computer science teacher **(advisor)** providing assistance and supervision. With further consultation from Edward, I decided to try the format, as it seemed the most convenient for his workflow. My first consultations with Edward are summarized in Appendix A.

### Rationale for the Proposed Product
I will program this project in Go. I am fairly familiar with the language, and I believe this project will allow me to deepen my understanding of it. Also, it has similar performance to C++, while maintaining decent ease of use. It also has support for data structures and object-oriented programming, which may prove useful. Finally, it is compiled, so I can distribute my code as an executable, saving the user some hassle.

I will make the project a command-line tool, as the client is familiar and comfortable with using them, and it can easily be used through an IDE's terminal or drop-down terminal.

### Success Criteria
- The program is run through the command line, and takes user input (either as arguments or as a live shell) and provide a result.
- Inputted numbers can either be in hexadecimal, binary, or decimal form, but must have a "tag" (e.g. ``0x`` for hexadecimal, ``0b`` for binary, none for decimal) before them.
- Other than the four mathematical operators and exponentiation, the bitwise AND, OR, NOT, and left/right shift are accepted.
- The correct solution is provided via terminal output.
- Parenthesis and more complex expressions (more than 2 inputs and 1 operator) are supported



# Section B:
- Record of Tasks (Fill out form)

My program consists of two important functions and the runtime loop. The first function, the "tokenizer", takes acceptable infix input and separates it into its correct components. The second function, the "shunting-yard", is based off of Djikstra's Shunting Yard Algorithm and evaluates tokens in order. The runtime loop accepts user input, combines the two functions, and does a few optimisations.

### Tokenizer
The tokenizer iteraters through each character in the input ``string``, and checks what "type" of character the current one is. It adds characters to the current token if they are of the same "type". If the current "type" is different from the previous "type", that indicates a separation between tokens. Then, the current token is appended to an output list, and the token variable is cleared.

The tokenizer also performs checks on number of parenthesis (if left and right quantities are not equal, then there is a mismatch), and whether two-character repeated operators (most of the bitwise ones) are correctly entered.

Below is an example of the tokenization of "2 + (23 << 0x6c)".

Diagram 1: Simplified tokenization

Diagram 2: Codetrace
// TODO: Diagram (Done)

### Shunting-Yard

The "Shunting Yard" algorithm is designed to convert infix notation into Reverse Polish Notation. However, a Reverse Polish Notation evaluator can be integrated with the algorithm for efficiency. 

The evaluator can be integrated into a Shunting Yard algorithm by replacing all of the cases where operators are transferred to the return stack with an evaluation of the operator on the last elements of the stack instead.

Another modification made is the conversion of the token "x", which is used to represent the previous result, to the actual previous result.

Diagram 3: Shunting-Yard Algorithm
// TODO: Diagram (needs layout) Talk about RPN?

### Runtime Loop

The runtime loop accepts user input, tokenizes it, inputs the tokens into the shunting yard, and returns the result. It also does some "shortcuts", such as skipping evaluation if the input is just a newline, or printing out the previous result if "x" is entered.

Diagram 4: Runtime Loop
// TODO: diagram

### Criteria to Test
Criteria | Method | Result/Adjustments

Takes User Input| Check that keyboard input works in final program (user can type)                               | Works
Inputted numbers correct | Check that the final program supports hex/binary numbers and converts them to decimal properly | Works
Operators Accepted       | Check that each of the operators is accepted and works as intended.                            | Works
Correct solution  | See above.                                                                                     | Works

// TODO: correct solutions for example: list



Tokenizer works          | Unit tests
#### Changes to tokenizer:
- Change nesting parenthesis behaviour: prior grouped paren next to each other

Shuntyard works          | Final testing of random expressions

Subfunctions work        | Unit tests



# Section C:
- structure
- algorithmic thinking
- techniques
- existing tools
- extensibility
- must include references

Types of Techniques and Where They Were Employed
- algorithms
  - Tokenizer
  - Shunting Yard
- methods
  - Tokenizer helpers
- data structures
  - Stack
  - Unit Testing Class
  - Magic numbers?
- classes/objects
 - Stack
- user interface
  - Final loop
## Tokenizer
The tokenizer separates a string into meaningful parts for the shunting yard algorithm (described below). It does this by looking at every character within the string, and determining its type. If the type matches the previous type, the character is added to the current token. If it does not, the token is added to the end of the output stack, and is then cleared. Additional checks for parenthesis count and correct input of double-character operators are included as well.

A sample codetrace is provided in the previous section, as "Diagram 2".

## Shunting Yard
The Shunting Yard algorithm, an algorithm invented by Dr. Edsger Djikstra, is a algorithm for converting an expression in infix notation (normal mathematical/expression notation) to reverse Polish notation (a notation that does not have parentheses and is easier to write evaluators for). This algorithm takes the output off of the Tokenizer and produces the solution to the user's expression.

A diagram of the algorithm's function is included in Section B, but, briefly, the algorithm works with two stacks: an output stack and an operator stack. Every iteration looks at a new token off of a list of tokens. Numbers are added to the bottom of the output stack, whereas operators are added to the top of the operator stack. The top operator from the operator stack is moved to the output stack when the top operator has greater predecence than the current token, after parentheses are closed, and at the end of looking at tokens. 

// TODO: better word for looking at, cite

The version of the shunting yard algorithm I used in my program is modified to evaluate operators and functions instead of adding them to the output stack. It does this using the ``operatorEval`` function, which uses the operator/function on the last two elements of the output stack, and replaces them with the solution. 

### Libraries Used

The ``strconv`` library is used to convert a string number into an integer, and the ``strings`` library is used to lowercase all tokens. The ``isNumber`` function, which returns whether a string is a number and the number contains permitted characters, also uses the ``strings`` library for string trimming and to determine its "prefix".

## Tokenizer Helpers
There are many miscellaneous functions that assist in the Tokenizer's operation.

``charType`` returns the ``TokenType`` of the character given, to allow the Tokenizer to determine when to clear the ``Token`` variable. Many other functions (which will be explained below) in turn make up this function.

``inRangeInc`` is used in order to define ``charIsNumber``, which determines whether a character is "numerical" by checking if it fits in the ranges of acceptable numerical chars.

``charIsFunc`` checks over a list of permitted functions and returns whether the character is present in it. The list of permitted functions can be expanded, if the user wishes to add a new function.

``charIsOperator`` works similarly to ``charIsFunc``, but instead checks over a list of operators.

``charIsLParen`` and ``charIsRParen`` check if a character is a left or right parenthesis, respectively. They can be modified to support different parenthesis formats or more than one type by either changing the character used for comparison, or by iterating over a list of supported parenthesis as with the other functions.

## Magic Numbers

I defined two magic number data structures in my program: ``TokenType`` and ``NumberType``. They also demonstrate object-oriented programming by inheriting from the ``int`` type. Both are used for a function's internal record-keeping, with the former used for ``Tokenizer`` and the latter used for ``isNumber``. These structures can be expanded by adding a different term, or creating objects that inherit from them.

## Unit Testing Class
In order to simplify one of my unit tests, I defined the ``OperatorTest`` data structure, which stores result, operator, and input necessary to test operator evaluation under one object. If new operators are added, a new data structure can be defined to create a new unit test.

## Stack

I used stacks primarily in the shunting yard algorithm. The "implementation" was essentially the use of a go slice within specific cases (appending to top, removing from top, etc). It both represents a data structure, storing the tokens/numbers within the shunting yard, and an object, due to the functions surrounding it that enabled my append/removal operations. 

Otherwise, I used slices and maps extensively, especially in my unit tests, in order to store test cases. There, they were solely data structures, as no functions were used in order to complete modifications on their contents.

## User Interface

While my program's user interface is not graphical, it is something that the user can type into. The program displays the characters ``> `` and expects a user input. Upon pressing enter, the user will receive an output. Some conveniences are also included, such as the ability to type ``exit`` for a clean closure, or shortcuts for newlines and ``x`` that skip the shunting yard entirely.

### Libraries Used
The user interface's input works by reading off of ``stdin`` (accessed via ``os``) with ``bufio``, and then takes that input in order to calculate the answer. That answer is then printed through ``stdin`` via the ``fmt`` library. 


# Section D:
- Functionality (video)
- documentation

# Section E:
## Success Criteria
- The program is run through the command line, and takes user input (either as arguments or as a live shell) and provide a result. (yes: takes input as a live shell)

- Inputted numbers can either be in hexadecimal, binary, or decimal form, but must have a "tag" (e.g. ``0x`` for hexadecimal, ``0b`` for binary, none for decimal) before them. (yes: takes all these forms and converts them to decimal)

- Other than the four mathematical operators and exponentiation, the bitwise AND, OR, NOT, and left/right shift are accepted. (yes: can support these operators)

- The correct solution is provided via terminal output. (yes: correct solutions provided, all of the test cases in the Testing Plan successfully run)

- Parenthesis and more complex expressions (more than 2 inputs and 1 operator) are supported (yes: complex expressions supported)
## Evaluation/feedback
As demonstrated in Appendix C, my client was satisifed with the final product. The main functionality (ability to do complex calculations with bitwise operators) satisfied his expectations, although, as noted in "Areas of improvement", he did wish that results could be displayed in binary as well. 
## Areas of improvement
With my client, I came up with the following areas of improvement or further development:
- cleaner code in some areas, e.g. figuring out how pointers could be used to simplify operatorEval, or passing around the constants for data efficiency/clarity
- improved error management/handling: potential for error without closing program
- command-line options: help menu, verbosity toggle or debugging mode
- extra features: user-defined functions, ability to store variables/use environment variables, cleaner ways to add new functions, different parenthesis
- support for user-defined options for single or double commands, input header, etc. 
- ability to display results in binary 

## FIRST CONSULTATION
We talked about potential IAs and narrowed it down to some sort of expression evaluator with support for binary. This was in the context of the chess competition that Edward was writing. The chess boards were stored in "bitboards", with each bit in a integer representing a position on the board. However, certain tasks (such determining masks to prevent shifts from overflowing the bitboard) required use of a binary calculator. His binary calculator required another window outside of the IDE, and did not support complex operations or expressions, which slowed his progress significantly. Since I had developed some command-line apps before, and Edward also enjoys using them (they're easier to access through an IDE), we thought of the idea of making a command-line binary calculator, and agreed it was a good way to solve the problem. We agreed upon some basic features (evaluation of the most common binary operations, and support for hexadecimal/binary numbers as well as decimals), and upon the basic format of the program.

## Second Consultation
We discussed my progress: I finished the shunting yard algorithm and reverse Polish notation evaluator, neither of which were tested. The shunting yard algorithm was recommended by Edward and approved by my advisor in order to add complexity and demonstrate algorithmic thinking. Edward suggested combining the evaluator and the shunting yard algorithm for better efficiency, and explained how to. We tested the shunting yard algorithm a little, and gradually found problems with the wording of the pseudocode I based my program off of. The issues were in precedence determination: I thought the pseudocode asked to determine precedence compared to the entire operator stack, but it turns out it was supposed to be applied to the first element only. This was fixed successfully. We also briefly discussed how I might test the various other functions I wrote, and he discussed some cases I might need to handle in the tokenizer (like mismatched parentheses).

## Final Consultation
My client was satisfied with the product. He felt that it met the success criteria and was useable. He liked the simple terminal interface and the operator support, and appreciated the peace-of-mind provided by the scope of the unit-testing. The improvements he wanted the most were the ability to display the result in binary, support for single-character bitwise operators, and perhaps some command-line options to enable the two or more options.
