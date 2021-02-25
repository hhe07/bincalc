- Basic flow
  - Parse input into (term) (operation) (term)
  - Convert terms into binary
  - Use operation function on terms
  - return result
- Parsing
  - Separate by spaces using built-in Go args system
  - Check that there are 3 arguments
  - Check that the arguments are of the right type
    - String starting with 0x/0b/0d or is LA (last answer)
      - String doesn't contain more digits than a restriction?
      - String contains valid digits
    - Operation is less than 2 chars long, is +*-/ >> << && || (insert not)
  - return string, string, string
- Binary conversion
  - Fairly self-explanatory, use traditional method based on the first term, or sub in previous answer
- Use function
  - Self-explanatory. >> and << are easy
  - The others might require research, but maybe consider using dec as base instead of binary to make it easier?
- Return result
  - Store answer to file in binary, convert result to right type



Classes

- Number: has attributes binary, hex, decimal
  - Constructor autoconverts from string
  - One function to convert throughout







Additional features:

- Save calculations?
- Have "variable creation" thing? (save to file in pair string, number)
- Take an outer !() not function?
- Take NOT by an extra parameter before each of the numbers
- More functions?
- User defined sequences?

KEEP LAST ANSWER


Shunting Yard:
/* This implementation does not implement composite functions,functions with variable number of arguments, and unary operators. */

while there are tokens to be read:
    read a token.
    if the token is a number, then:
        push it to the output queue.
    else if the token is a function then:
        push it onto the operator stack 
    else if the token is an operator then:
        while ((there is an operator at the top of the operator stack)
              and ((the operator at the top of the operator stack has greater precedence)
                  or (the operator at the top of the operator stack has equal precedence and the token is left associative))
              and (the operator at the top of the operator stack is not a left parenthesis)):
            pop operators from the operator stack onto the output queue.
        push it onto the operator stack.
    else if the token is a left parenthesis (i.e. "("), then:
        push it onto the operator stack.
    else if the token is a right parenthesis (i.e. ")"), then:
        while the operator at the top of the operator stack is not a left parenthesis:
            pop the operator from the operator stack onto the output queue.
        /* If the stack runs out without finding a left parenthesis, then there are mismatched parentheses. */
        if there is a left parenthesis at the top of the operator stack, then:
            pop the operator from the operator stack and discard it
        if there is a function token at the top of the operator stack, then:
            pop the function from the operator stack onto the output queue.
/* After while loop, if operator stack not null, pop everything to output queue */
if there are no more tokens to read then:
    while there are still operator tokens on the stack:
        /* If the operator token on the top of the stack is a parenthesis, then there are mismatched parentheses. */
        pop the operator from the operator stack onto the output queue.
exit.