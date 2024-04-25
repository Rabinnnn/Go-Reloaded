This is a program that can format text in a file (sample.txt) as per the instructions given and write the output in another file (result.txt). 
Here are some of the instructions that it can handle:
* (Up) - this converts the previous word to uppercase.
* (low) - this converts the previous word to lowercase.
* (cap) - this capitalizes the first letter of the previous word.
* (hex) - this converts the previous hexadecimal value to decimal value
* (bin) - this converts the previous binary word to decimal value
The program also ensures that punctuation is applied correctly. For instance, if there is a space between a word and comma it removes the space. It does the same for double quotes and single quotes. Besides that, if it encounters artile "A" or "a" before words that begin with a vowel or letter "h", it replaces the article with "An" or "an" accordingly.
