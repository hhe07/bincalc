
# misc. scrapyard

What do I want from Edward?
- How to fix unit tests (done)
- Whether test coverage is suitable or not (done)
- Ask about success criteria (done)
- Meeting evidence cheese?
- Show him tests + testing plan, demonstrate that it works (done)
- Feedback from client: internal improvements, external improvements, etc. (done)
- Evaluate success criteria (done)
- Discuss somewhat illustrative test cases for presentation? (done)

- simple -> complicated expressions (decent amount of operators, parenthesis, etc) 

### Rationale for Solution
Binary math on paper is generally rather time-consuming and requires a distraction of a programmer's workflow. Whereas a graphical user interface may require mouse input and another desktop window, a command-line interface can usually be easily accessed through a code editing program and is keyboard-only. Therefore, the main rationale of digitising this solution is convenience. With a digital (and more specifically command-line) implementation, Edward could do calculations faster than he currently can.



The solution will solve my client's needs, as demonstrated in Appendix 1. He feels that the solution I proposed includes the features he wants, while also providing some measure of convenience

- for token in tokens:
  - if is number:
    - add to end of return stack
  - if is function:
    - add to top of operator stack.
  - if is operator:
    - while there is a top operator in operator stack,
      and (top operator is greater precedence than token or
      operator has equal precedence but token is left associative),
      and the top operator isn't a left parenthesis:
      - move operators from operator stack to end of return stack
    - add operator to top of operator stack
  - if is left parenthesis:
    - add to top of operator stack
  - if is right parenthesis:
    - while the top operator in operator stack isn't 



// 5 or so tests that complete program works
