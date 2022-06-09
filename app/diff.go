// package willcompare
package main

import (
	"fmt"
	"os"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	text1 = "Lorem ipsum dolor."
	text2 = "Lorem dolor sit amet."
	// pl1 = "Of Mans First Disobedience, and the Fruit
	// Of that Forbidden Tree, whose mortal tast
	// Brought Death into the World, and all our woe,
	// With loss of Eden, till one greater Man
	// Restore us, and regain the blissful Seat,
	// Sing Heav’nly Muse, that on the secret top
	// Of Oreb, or of Sinai, didst inspire
	// That Shepherd, who first taught the chosen Seed,
	// In the Beginning how the Heav’ns and Earth
	// Rose out of Chaos: Or if Sion Hill
	// Delight thee more, and Siloa’s Brook that flow’d
	// Fast by the Oracle of God; I thence
	// Invoke thy aid to my adventrous Song,
	// That with no middle flight intends to soar
	// Above th’ Aonian Mount, while it pursues
	// Things unattempted yet in Prose or Rhime.
	// And chiefly Thou O Spirit, that dost prefer
	// Before all Temples th’ upright heart and pure,
	// Instruct me, for Thou know’st; Thou from the first
	// Wast present, and with mighty wings outspread
	// Dove-like satst brooding on the vast Abyss
	// And mad’st it pregnant: What in me is dark
	// Illumine, what is low raise and support;
	// That to the highth of this great Argument
	// I may assert th’ Eternal Providence,
	// And justifie the wayes of God to men."
	// pl2 = "OF Mans First Disobedience, and the Fruit
	// Of that Forbidden Tree, whose mortal tast
	// Brought Death into the World, and all our woe,
	// With loss of Eden, till one greater Man
	// Restore us, and regain the blissful Seat,
	// Sing Heav'nly Muse, that on the secret top
	// Of Oreb, or of Sinai, didst inspire
	// That Shepherd, who first taught the chosen Seed,
	// In the Beginning how the Heav'ns and Earth
	// Rose out of Chaos: or if Sion Hill
	// Delight thee more, and Siloa's brook that flow'd
	// Fast by the Oracle of God; I thence
	// Invoke thy aid to my adventrous Song,
	// That with no middle flight intends to soar
	// Above th' Aonian Mount, while it pursues
	// Things unattempted yet in Prose or Rhime.
	// And chiefly Thou, O Spirit, that dost prefer
	// Before all Temples th' upright heart and pure,
	// Instruct me, for Thou know'st; Thou from the first
	// Wast present, and with mighty wings outspread
	// Dove-like satst brooding on the vast Abyss
	// And mad'st it pregnant: What in me is dark
	// Illumin, what is low raise and support;
	// That to the highth of this great Argument
	// I may assert Eternal Providence,
	// And justifie the wayes of God to men."
)

func main() {
	toHTML(compare(), "test")
}

func compare() []diffmatchpatch.Diff {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(text1, text2, false)

	fmt.Println(diffs)

	fmt.Println(dmp.DiffPrettyText(diffs))

	fmt.Println(dmp.DiffPrettyHtml(diffs))

	return diffs
}

func toHTML(diffData []diffmatchpatch.Diff, fileName string) {
	fileName = fileName + ".html"

	dmp := diffmatchpatch.New()
	asHTML := dmp.DiffPrettyHtml(diffData)

	err := (os.WriteFile(fileName, []byte(asHTML), 0555))
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}
