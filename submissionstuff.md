# Section A: 
- Describe scenario
- Provide evidence of initial requirements meeting
- Describe proposed products
- Specific performance criteria

Solution title: bincalc

User Name: Edward Wawrzynek (Student's friend)

## Criterion A:
### Defining the Problem
My friend Edward **(client)** makes frequent use of bitwise math while programming, especially in hardware-design related contexts. For example, when this project was started, he was working on a chess competition framework for our school's robotics club. To save space, the chess boards were represented in integer form, with each bit representing whether a space was occupied or not. However, this made it necessary to calculate various "magic numbers" using bitwise math, and he was not satisfied with his calculator program's programmer mode. So, I decided to try to make a command line application to satisfy this need, with my computer science teacher **(advisor)** providing assistance and supervision. With further consultation from Edward, I decided to try to make a command-line application, as it seemed the most convenient for his workflow. 

### Rationale for the Proposed Product
I will program this project in Go. I am fairly familiar with the language, and I believe this project will allow me to deepen my understanding of it. Also, it has similar performance to C++, while maintaining decent ease of use. It also has support for data structures and object-oriented programming, which may prove useful. Finally, it is compiled, so I can distribute my code as an executable, saving the user some hassle.

### Success Criteria
- The program is run through the command line, and takes user input (either as arguments or as a live shell) and provide a result.
- Inputted numbers can either be in hexadecimal, binary, or decimal form, but must have a "tag" (e.g. ``0x`` for hexadecimal, ``0b`` for binary, none for decimal) before them.
- Other than the four mathematical operators and exponentiation, the bitwise AND, OR, NOT, and left/right shift are accepted. 
- The correct solution is provided via terminal output.



# Section B:
- Record of Tasks (Fill out form)

My program consists of two important functions and the runtime loop. The first function, the "tokenizer", takes acceptable infix input and separates it into its correct components. The second function, the "shunting-yard", is based off of Djikstra's Shunting Yard Algorithm and evaluates tokens in order. The runtime loop accepts user input, combines the two functions, and does a few optimisations.

### Tokenizer
The tokenizer iteraters through each character in the input ``string``, and checks what "type" of character the current one is. It adds characters to the current token if they are of the same "type". If the current "type" is different from the previous "type", that indicates a separation between tokens. Then, the current token is appended to an output list, and the token variable is cleared.

The tokenizer also performs checks on number of parenthesis (if left and right quantities are not equal, then there is a mismatch), and whether two-character repeated operators (most of the bitwise ones) are correctly entered.

Below is an example of the tokenization of "2 + (23 << 0x6c)".

// TODO: Diagram

### Shunting-Yard

The "Shunting Yard" algorithm is designed to convert infix notation into Reverse Polish Notation. However, a Reverse Polish Notation evaluator can be integrated with the algorithm for efficiency. 

The evaluator can be integrated into a Shunting Yard algorithm by replacing all of the situations where operators are transferred to the return stack with an evaluation of the operator.

One modification made is the conversion of the token "x", which is used to represent the previous result, to the actual previous result.

// TODO: Diagram

### Runtime Loop

The runtime loop accepts user input, tokenizes it, inputs the tokens into the shunting yard, and returns the result. It also does some "shortcuts", such as skipping evaluation if the input is just a newline, or printing out the previous result if "x" is entered.

### Criteria to Test
Criteria | Method | Result/Adjustments

Takes User Input         | Check that keyboard input works in final program (user can type)                               | Works
Inputted numbers correct | Check that the final program supports hex/binary numbers and converts them to decimal properly | Works
Operators Accepted       | Check that each of the operators is accepted and works as intended.                            | Works
Correct solution         | See above.                                                                                     | Works

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
  
- algorithms
  - Shunting Yard
  - Tokenizer
- methods
  - Tokenizer helpers
- data structures
- classes/objects
  - Magic numbers?
- user interface
  - Final loop

# Section D:
- Functionality (video)
- Extensibility (documentation, describe extensible components)

# Section E:
Evaluation/feedback

## TODO: Cheese appendicies

### Rationale for Solution
Binary math on paper is generally rather time-consuming and requires a distraction of a programmer's workflow. Whereas a graphical user interface may require mouse input and another desktop window, a command-line interface can usually be easily accessed through a code editing program and is keyboard-only. Therefore, the main rationale of digitising this solution is convenience. With a digital (and more specifically command-line) implementation, Edward could do calculations faster than he currently can.



The solution will solve my client's needs, as demonstrated in Appendix 1. He feels that the solution I proposed includes the features he wants, while also providing some measure of convenience
