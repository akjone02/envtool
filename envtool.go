/*
This program is designed to compare the contents of two linux directories and list out differences.
ENVTOOL was written by ANDREW JONES as part of a coding competancy test for ANTIMATTER.
This program was built in an Ubuntu 20.04 VM using Visual Studio Code and GOlang.
This is my first attempt at coding with GO.
*/

package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//Main function now allows specifying a directory to search
func main() {
	ptrDir := flag.String("d", ".", "Specify a directory to be parsed")
	ptrOut := flag.String("o", "output.log", "The location of the output file")

	ptrComp := flag.Bool("c", false, "Compare two files")
	ptrF1 := flag.String("f1", "", "Use in conjunction with -c and -f2 flags, the first file to be compared")
	ptrF2 := flag.String("f2", "", "Use in conjunction with -c and -f1 flags, the second file to be compared")
	flag.Parse()

	if *ptrComp {
		compareFiles(*ptrF1, *ptrF2)
	} else {
		getDirContents(*ptrDir, *ptrOut)
	}
}

//This function takes two input files and compares them to see if they're the same
func compareFiles(strFile1 string, strFile2 string) {
	switch {
	case !fileExist(strFile1):
		fmt.Println("Could not open first file at: " + strFile1)
	case !fileExist(strFile2):
		fmt.Println("Could not open second file at: " + strFile2)
	default:
		if getHash(strFile1) == getHash(strFile2) { //Check the hash first. If identical, no need to parse files
			fmt.Println("The two capture files are the same")
			fmt.Println("No files are different between these two directories")
		} else { //File hash differs, need to parse the files and see what changed
			fmt.Println("Files are not the same!  Directory differences found.")
			strFile1Text := scanTextFile(strFile1)
			strFile2Text := scanTextFile(strFile2)

			fmt.Println("Files found in directory 1 but not directory 2:")
			s1 := findSliceNameDiff(strFile1Text, strFile2Text)
			for _, i := range s1 {
				fmt.Println(i)
			}
			fmt.Println("Files found in directory 2 but not directory 1:")
			s2 := findSliceNameDiff(strFile2Text, strFile1Text)
			for _, i := range s2 {
				fmt.Println(i)
			}
			fmt.Println("Files that are different between directories 1 and 2:")
			s3 := findSliceHashDiff(strFile1Text, strFile2Text)
			for _, i := range s3 {
				fmt.Println(i)
			}

			//Need to add in a function to scan for same filename but different md5hash
		}
	}
}

//Scans through a text file and slices it
func scanTextFile(strFile string) []string {
	file, err := os.Open(strFile)
	defer file.Close()
	errCheck(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var strFileText []string
	for scanner.Scan() {
		strFileText = append(strFileText, scanner.Text())
	}
	return strFileText
}

//Function to find and return strings where the hash differs between the two directories
func findSliceHashDiff(strSlice1 []string, strSlice2 []string) []string {
	var strDiff []string

	for _, s1 := range strSlice1 { //for each element of the first slice...
		s1Hash := s1[len(s1)-34:]
		boolDiff := false
		for _, s2 := range strSlice2 { //iterate through each elemeny of the second slice
			s2Hash := s2[len(s2)-34:]
			if s1Hash == s2Hash {
				boolDiff = true
				break //Found an identical string, break out
			}
		}
		//We checked each element of the 2nd slice and found no match
		if !boolDiff {
			strDiff = append(strDiff, s1)
		}
	}
	return strDiff

}

//Updated diff function that will check only for filenames that differ
func findSliceNameDiff(strSlice1 []string, strSlice2 []string) []string {
	var strDiff []string

	for _, s1 := range strSlice1 { //for each element of the first slice...
		s1 = s1[:len(s1)-34]
		boolDiff := false
		for _, s2 := range strSlice2 { //iterate through each element of the second slice
			s2 = s2[:len(s2)-34]
			if s1 == s2 {
				boolDiff = true
				break //Found an identical string, break out
			}
		}
		//We checked each element of the 2nd slice and found no match
		if !boolDiff {
			strDiff = append(strDiff, s1)
		}
	}
	return strDiff

}

//Basic error check, since I'm now checking in multiple places
func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func fileExist(strFile string) bool {
	_, err := os.Stat(strFile)

	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

//Gets the hash of a file
func getHash(strFile string) string {
	data, err := ioutil.ReadFile(strFile)
	errCheck(err)
	bHash := md5.Sum(data)
	return hex.EncodeToString(bHash[:])
}

//Basics of function for parsing directory
//Currently does not walk down subdirs
func getDirContents(strDir string, strOut string) {
	files, err := ioutil.ReadDir(strDir)
	if err != nil {

	}

	outputFile, err := os.Create(strOut)
	defer outputFile.Close()

	strFileName := ""
	strHash := ""
	intCounter := 0

	for _, file := range files {
		if !file.IsDir() {
			strFileName = file.Name()
			strHash = getHash(strDir + "/" + strFileName)
			fmt.Println(strFileName + " - " + strHash)
			outputFile.WriteString(file.Name() + " - " + strHash + "\n")
			intCounter++
		}
	}
	fmt.Println("Finished parsing directory: " + strDir)
	fmt.Println("Total files hashed: " + strconv.Itoa(intCounter))
	fmt.Println("Full output saved to: " + strOut)
}
