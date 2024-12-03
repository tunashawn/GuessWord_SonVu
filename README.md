### How to run:
For daily mode: `go run main.go daily size`

For random mode: `go run main.go random size seed`

replace `size` is an integer >= 2, `seed` with any integer



### Analyze problem:

API:
* https://wordle.votee.dev:8000/wordseg 
* https://wordle.votee.dev:8000/daily?guess=aiueo&size=5 seems to have a fixed word for each size
* https://wordle.votee.dev:8000/random?guess=aiueo&size=5&seed=1 has a fixed word with length = size for each seed. It is 
completely random without seed param so to use this user must input a seed number
* https://wordle.votee.dev:8000/word/hello?guess=hello just check if two provided words are the same
* uppercase or lowercase does not matter

Result from API:
* "absent" if the correct word does not contain the letter, 
* "present" if it does contain, 
* "correct" if the letter is in correct position

### Stategy 1:
Call API to guess with all letters I have in alphabet until I guess the right word. For example word "ant", guesses are
aaa, nnn, ttt, etc... result in a with correct position in 0, n : 2, t : 3 => combine to get the correct word.

* Pros: can independently guess all words without using third party api
* Cons: high api call to guess.

### Strategy 2:
Because there are no evidence support that all words will be in English, it is not wise to check with the popular letter
in English. Though for now datamuse seems to only support English.
The easier way is to check with letters a, i, u, e, o, y as most word in language use alphabet contain at least one of these
letters.
After confirmed existing letters, I will use brute force to find the correct position of that letters, then fill other 
with "?" and use the api from datamuse to get a list of possible matches. After that I just have to guess with those 
possible match.
In case these letters are absent, then I'll check with other letters.

* Pros: better guess for most cases and lower guess count.
* Cons: corner cases like some words does not contain any vowels such as "hmm" will need addition steps similar to 
strategy 1 to fill some letter before asking for top possible match from datamuse api.

For this repository, I'll use the strategy 2. 



### Credit:

* API to guess a word with some given letter from https://www.datamuse.com/api/
* Few code was used from ChatGPT to fasten development time
