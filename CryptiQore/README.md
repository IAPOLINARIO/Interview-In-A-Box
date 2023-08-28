# Coding Challenge Interview

## Instructions:

Read the Coding Challenge
The challenge can be solved in any programming language
Once you have read the challenge share the screen with your favorite IDE and this page.
Start solving the problem
If you need to make a diagram please use https://excalidraw.com/

## Problem:

The annual snake festival is upon us, and all the snakes of the kingdom have gathered to participate in the procession. Bitso has been tasked with reporting on the procession, and for this he decides to first keep track of all the snakes. When he sees a snake first, it'll be its Head, and hence he will mark a 'H'. The snakes are long, and when he sees the snake finally slither away, he'll mark a 'T' to denote its tail. In the time in between, when the snake is moving past him, or the time between one snake and the next snake, he marks with '.'s.
Because the snakes come in a procession, and one by one, a valid report would be something like "..H..T...HTH....T.", or "...", or "HT", whereas "T...H..H.T", "H..T..H", "H..H..T..T" would be invalid reports (See explanations at the bottom).

Formally, a snake is represented by a 'H' followed by some (possibly zero) '.'s, and then a 'T'. A valid report is one such that it begins with a (possibly zero length) string of '.'s, and then some (possibly zero) snakes between which there can be some '.'s, and then finally ends with some (possibly zero) '.'s.

## Inputs

Each report contains a string of length L. The string contains only the characters '.', 'H', and 'T'.

## Output

For each report, output the string Valid or Invalid, depending on whether it was a valid report or not.
Examples

```
Input:
..H..T...HTH....T.
...
H..H..T..T
HT
.T...H..H.T
H..T..H

Output

Valid
Valid
Invalid
Valid
Invalid
Invalid
```
