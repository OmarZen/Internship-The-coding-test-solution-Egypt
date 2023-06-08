# Fibonacci Sequence - README

This repository contains three different implementations of the Fibonacci sequence, solving the problem F(n) = F(n-3) + F(n-2), where F(0) = 0, F(1) = 1, and F(2) = 2. The implementations are written in C and demonstrate three different approaches: recursive, iterative, and dynamic programming.

## Development Environment

The source code provided in this repository was developed using the following environment:

- Operating System: [Specify your operating system here]
- Compiler: [Specify your compiler here, e.g., gcc]
- Version: [Specify the compiler version if applicable]

## Compilation

To compile the source code, follow the steps below:

1. Open a terminal or command prompt.
2. Navigate to the directory where the source code is located.
3. Use the appropriate compiler command to compile the code. For example:
   - For the recursive approach: `gcc recursive_fibonacci.c -o recursive_fibonacci`
   - For the iterative approach: `gcc iterative_fibonacci.c -o iterative_fibonacci`
   - For the dynamic programming approach: `gcc dynamic_fibonacci.c -o dynamic_fibonacci`

## Running the Program

To run the compiled program, follow the steps below:

1. In the same terminal or command prompt, execute the compiled program. For example:
   - For the recursive approach: `./recursive_fibonacci`
   - For the iterative approach: `./iterative_fibonacci`
   - For the dynamic programming approach: `./dynamic_fibonacci`

2. The program will display the Fibonacci number for the specified input value of `n` according to the selected approach.

## Explanation of the Differences

The three different approaches used in this implementation have the following advantages and disadvantages:

1. Recursive Approach:
   - Advantage: The recursive approach provides a straightforward and intuitive implementation of the recurrence relation. It closely reflects the mathematical definition of the relation.
   - Disadvantage: The recursive approach can be inefficient for large values of `n`. It often leads to redundant computations, as the same subproblems are solved multiple times.

2. Iterative Approach:
   - Advantage: The iterative approach avoids redundant computations and provides a more efficient solution compared to the recursive approach. It uses a loop to calculate the Fibonacci sequence iteratively.
   - Disadvantage: The iterative approach might be slightly more complex to understand compared to the recursive approach, as it involves managing multiple variables and maintaining the state across iterations.

3. Dynamic Programming Approach:
   - Advantage: The dynamic programming approach also avoids redundant computations and provides an efficient solution. It uses an array (`dp`) to store the values of subproblems, eliminating the need to recalculate them.
   - Disadvantage: The dynamic programming approach requires additional memory to store the array `dp`, which can be a concern for large values of `n`. It might require more space compared to the iterative approach.

In summary, the choice of approach depends on the specific requirements of the problem and the trade-offs between simplicity and efficiency.