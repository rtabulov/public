## revWstr

### Instructions

Write a program that takes a string as a parameter, and prints its words in reverse.

- A word is a sequence of **alphanumerical** characters.

- If the number of parameters is different from 1, the program will display `\n`.

- In the parameters that are going to be tested, there will not be any additional spaces. (meaning that there will not be additionnal spaces at the beginning or at the end of the string, and words will always be separated by exactly one space).

Examples of outputs :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test "the time of contempt precedes that of indifference"
indifference of that precedes contempt of time the
student@ubuntu:~/piscine/test$ ./test "abcdefghijklm"
abcdefghijklm
student@ubuntu:~/piscine/test$ ./test "he stared at the mountain"
mountain the at stared he
student@ubuntu:~/piscine/test$ ./test

student@ubuntu:~/piscine/test$
```