# will_compare_v2

An updated version of [will-compare](https://github.com/rosemadr/will-compare). A project to compare Shakespeare play texts.

## Why compare Shakespeare?
Shakespeare's plays are actually very variant. Many of the plays have mulitple source texts, sometimes including not only bound manuscripts but more "live" fragments of player's lines. Add to this that the plays are now over four hundred years old and have been editied, copied, translated, published and performed many times in those years and you find that "Shakespeare's play" is actually a 
rather textually unstable entitiy.

## Project Aims

### Must Have (MVP)
* Be able to give the program two Hamlet playtexts (raw text files) and have it:
  1. Standardise the character names, whitespace and any other data we <i>don't</i> want it to compare.
  2. Output a html file which displays the differences to the texts.
* Be fully tested using automated testing. Built using TDD, as much as possible.

### Should Have
* Ability to manage the full Shakespeare dramatic corpus. A user should be able to pick which play they wish to compare and the program will standardise it.
* Ability for the user to select whether they can see word differences or punctation differences. These differences should be be indicated visually using textual display.

### Could Have (Longer Term Additional Features)
* GUI or <i>minimal</i> cli tool to enable a lay tech user to run on their computer and display the html in their browser.
* Ability to automatically recognise which plays it is presented with.
* Ability to, in addition to the above word vs punctuation differences, also display the difference between grammatical (same/similar words in different orders) and structural differences (scene and character additions).
* Webapp.
