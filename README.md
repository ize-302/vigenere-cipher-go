# Golang implementation of the Vigenère Cipher

The Vigenère cipher is a method of encrypting alphabetic text by using a simple form of polyalphabetic substitution. It uses a keyword, where each letter of the keyword provides a shift for the corresponding letter of the plaintext.

### How the Vigenère Cipher Works

#### Choose a keyword

This keyword is repeated to match the length of the plaintext. For example, if the keyword is "key" and the plaintext is "hello world", the keyword is repeated as "keyke ykeyk".

```sh
# Example
hello world
-----------
keyke ykeyk
```

#### Encrypting the message

- Each letter of the plaintext is shifted according to the corresponding letter of the repeated keyword.
- The shift is determined by the position of the keyword letter in the alphabet

```sh
# Representation of alphabelts with numbers starting from 0 to 25
a  b  c  d  e  f  g  h  i  j  k  l  m  n  o  p  q  r  s  t  u  v  w  x  y  z
0  1  2  3  4  5  6  7  8  9  10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25
```

##### Example

```sh
Text - hello world
Keyword - key

# The encryption process:
h (7) + k (10) = 17 => r
e (4) + e (4) = 8 => i
l (11) + y (24) = 35 => 35 % 26 = 9 => j
l (11) + k (10) = 21 => v
o (14) + e (4) = 18 => s
# we dont transform spaces, numbers and symbols
w (22) + y (24) = 46 => 46 % 26 = 20 => u
o (14) + k (10) = 24 => y
r (17) + e (4) = 21 => v
l (11) + y (24) = 35 => 35 % 26 = 9 => j
d (3) + k (10) = 13 => n

"hello world" becomes "rijvs uyvjn"
```

#### Decrypting the message:

To decrypt, the keyword is used in reverse to subtract the shift values.

##### Example

```sh
Text - rijvs uyvjn
Keyword - key

# The encryption process:
r (17) - k (10) = 7 => h
i (8) - e (4) = 4 => e
j (9) - y (24) = -15 => -15 % 26 = 11 => l
v (21) - k (10) = 11 => l
s (18) - e (4) = 14 => o
# we dont transform spaces, numbers and symbols
u (20) - y (24) = -4 => -4 % 26 = 22 => w
y (24) - k (10) = 14 => o
v (21) - e (4) = 17 => r
j (9) - y (24) = -15 => -15 % 26 = 11 => l
n (13) - k (10) = 3 => d

"rijvs uyvjn" becomes "hello world"
```

#### Properties and Security

- Polygraphic nature: Unlike the Caesar cipher, which is monoalphabetic, the Vigenère cipher uses multiple Caesar ciphers based on the keyword, making it more resistant to simple frequency analysis.
- Key length: The security of the Vigenère cipher is significantly influenced by the length and randomness of the keyword. The longer and more random the keyword, the more secure the cipher.
- Repetition of keyword: A repeating keyword introduces patterns that can be exploited by cryptanalysts, especially with methods like the Kasiski examination and frequency analysis.

#### Practical Use

The Vigenère cipher is a classical cipher and was widely used until the 19th century. However, with the advent of modern cryptography and the development of more secure encryption methods, it has become mostly obsolete for serious security applications. It is still useful for educational purposes to understand the basics of polyalphabetic ciphers and cryptographic principles.

#### Research materials

- [Vigenère cipher on Wikipedia](https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher)
- [Vigenere Cipher 1 by MathAfterMath](https://www.youtube.com/watch?v=E352JJ8xv48)
- [Vigenere Cipher 2 by MathAfterMath](https://www.youtube.com/watch?v=VK3qc5Lv8QQ)
