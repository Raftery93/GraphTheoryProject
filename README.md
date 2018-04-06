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


## Project Plan / Project Development
I began my project plan by sitting down and writing out a development process via an algorithm. This gave me an initial framework to work on top of. My algorithm included building a runner/main, which the user would launch, and this would be contained within a loop so the user does not have to constantly run the program. This main would be backed by a couple of classes which would do the most of the calculations and functions.

This led me to actually creating the assets I needed for the program to process regular expressions. I used the video tutorials to build these classes, which I later edited to add in extra commands and functions. See the bottom of this readme for references. Once I had these classes finished, I had to 'hook' them up to the main page.

After this, I had a somewhat functioning program with multiple bugs. I created a list of these bugs and I assigned them a priority. For example, one of the high priority bugs was that the user input was being taken in wrong, thus I created a seperate function to resolve this issue. One of the low priority bugs was displaying the details and information to the user in a user friendly fashion.

## Research
When I began this project, I started by looking into [GoLangs Regexp documentation](https://golang.org/pkg/regexp/). This gave me a better understanding on what a regular expression actually is, and how it is implemented. A regular expression is a sequence of symbols and characters expressing a string or pattern to be searched for within a longer piece of text. This led me to researching how to transform a regular expression into an [NFA](http://infolab.stanford.edu/~ullman/ialc/spr10/slides/fa3.pdf).

Once I got a solid grasp on what a regular expression is, and how to transform it into an NFA, I looked at how to write a program that can process a string through a regular expression function, and how to determine if the string matches. I found an extremely useful paper on this: [Regular Expression Matching](https://swtch.com/~rsc/regexp/regexp1.html). I also used video references which are held below under references.

My final embarkment was how do I transform a regular expression into an NFA. I used one of the video references below for this, along with [Thompsons Construction](https://www.cs.york.ac.uk/fp/lsa/lectures/REToC.pdf).



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
