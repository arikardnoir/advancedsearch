# Advanced Search for Go
`advancedsearch` is a pure Golang package to make search while typing "char" by "char". Actually he has two main functions, CompareSingleWord and CompareMultipleWords.

## How it works

Actually it is used in case where the user has a intension to make string search sequencially and not to knows if this word contains XYZ.

### CompareSingleWord
This function receives two params, the first one is the word that is typed, and second one is the word to be compared.

``` bash

CompareSingleWord(query, word)

```
#### Examples 
Non-Applicable Case:
``` bash
query := "Boo"
theWord := "Notebook"

res := CompareSingleWord(query, theWord)
```
At this case we can see the example of what we can't expect from this library. At this example the word "Notebook" contains the word "Boo" but will return `false`, because this methods expect to compare sequencially starting from position `0` from word "Notebook".

Applicable Case:
``` bash
query := "Boo"
theWord := "Booeing"

res := CompareSingleWord(query, theWord)
```
At least in this scenario we can see that the user starts type "Boo" that match with the starts of the word "Booeing" and in this case will return `true`.

### CompareMultipleWords
This function receives two params, the first one is the word that is typed, and second one is an array of word to be compared.

``` bash

CompareMultipleWords(query, words)

```
#### Examples 

Case:
``` bash
func main () {
  query := "Mac"
  theWords := []string{"Macbook", "Mouse", "Macaco", "Balde", "Copo", "Macarofe"}

  result, elements := CompareMultipleWords(query, theWords)
}
```
In this case we can see that the user starts type "Mac" that match with the starts of some words in this array and in this case will return `true` and array with elmenets.

## Usage
``` bash
func main () {
  result := CompareSingleWord("Mac", "Macbook")
  theWords := []string{"Macbook", "Mouse", "Macaco", "Balde", "Copo", "Macarofe"}
  _, elements := CompareMultipleWords("Mac", arrayToPass)

  fmt.printf("Elements: %v", elements)
}
```
`Output`
``` bash
  ["Macaco", "Macarofe", "Macbook"]
```

## Tests
``` bash
  go test
```
