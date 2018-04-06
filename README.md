# GoLang Automaton
This is a go application that can:

1) build a postfix expression from an infix expression

2) build a non-deterministic finite automaton (NFA) from a regular postfix expression.

3) check does a string match a regular expression.


## Description
The aim of this project is to write a program in the Go programming language that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text. This program does not use the regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some of which may have a special meaning.
For example, this program can determine between characters ".", "|", "*", "+" and "?" which have the special meanings
"concatenate", "or", "Kleene star", "match-one-or-more quantifier" and "match-zero-or-one quantifier" respectively.

## Testing the Program
Infix to Postfix:

• a.b.c* -> ab.c*.

• (a.(b|d))* -> abd|.*

• a.(b|d).c* -> abd|.c*.

• a.(b.b)+.c -> abb.+.c.

Postfix to NFA:

•ab.c*| will match cccc

## Downloading and Running the program
This program uses the Go programming language.

If you do not currently have Go installed click on the following link to download:
[INSTALL GO](https://golang.org/dl/)

To clone the repository, open command prompt, and enter:
```
git clone https://github.com/Raftery93/GraphTheoryProject
```
To run this program, you need to:
**Build and Run**
 Navigate to where you the project was cloned to and enter:
```
go build runner.go
```
This will create a .exe file in your current directory.This will create an executable file. To run this file, Enter:
```
runner.exe
```

## Design & Information
This program was developed using the Go programming language, Visual Studio Code and Git.

### Technologies Used
- GoLang
- Visual Studio Code
- Git

## Resources
https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e

https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b

https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d

https://www.cs.york.ac.uk/fp/lsa/lectures/REToC.pdf

### Contributor
Conor Raftery
