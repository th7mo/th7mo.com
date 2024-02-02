package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	gardenPath := "src/garden"
	templateBeforePath := "src/template/before.html"
	templateAfterPath := "src/template/after.html"
	distPath := "dist"

	// Clean the dist folder
	err := os.RemoveAll(distPath)
	if err != nil {
		fmt.Printf("Error cleaning dist folder: %v\n", err)
		return
	}

	// Recreate the dist folder
	err = os.Mkdir(distPath, 0755)
	if err != nil {
		fmt.Printf("Error creating dist folder: %v\n", err)
		return
	}

	// Read the content of the template.before.html file
	templateBeforeContent, err := ioutil.ReadFile(templateBeforePath)
	if err != nil {
		fmt.Printf("Error reading template/before.html file: %v\n", err)
		return
	}

	// Read the content of the template.after.html file
	templateAfterContent, err := ioutil.ReadFile(templateAfterPath)
	if err != nil {
		fmt.Printf("Error reading template/after.html file: %v\n", err)
		return
	}

	// Process HTML files in the garden folder
	err = filepath.Walk(gardenPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			// Read the content of the current HTML file
			currentContent, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", path, err)
				return nil
			}

			// Extract maturity value from the comment at the beginning of the HTML file
			maturity := extractMaturity(currentContent)

			// Replace {plant} with the name of the HTML file (remove .html, replace dashes with spaces)
			originalFileName := strings.TrimSuffix(info.Name(), ".html")
			plantName := strings.Replace(strings.Replace(originalFileName, "-", " ", -1), "_", " ", -1)

			// Replace {maturity} in templateBeforeContent with the extracted maturity value (fallback to "seedling" if not found)
			currentTemplateBefore := []byte(strings.Replace(string(templateBeforeContent), "{maturity}", fallbackToSeedling(maturity), -1))
			currentTemplateBefore = []byte(strings.Replace(string(currentTemplateBefore), "{plant}", plantName, -1))

			// Remove the entire line containing the maturity comment and the newline after it from the current content
			currentContent = removeMaturityComment(currentContent)

			// Combine template.before, current content, and template.after
			resultContent := fmt.Sprintf(
				"%s%s%s",
				currentTemplateBefore,
				currentContent,
				templateAfterContent,
			)

			// Create the destination path in the dist folder
			destPath := filepath.Join(distPath, info.Name())

			// Write the combined content to the destination file in the dist folder
			err = ioutil.WriteFile(destPath, []byte(resultContent), 0644)
			if err != nil {
				fmt.Printf("Error writing to file %s: %v\n", destPath, err)
			} else {
				fmt.Printf("Processed file: %s\n", destPath)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through the directory: %v\n", err)
	}

	copyDir("src/css", filepath.Join(distPath, "css"))
	copyDir("src/svg", filepath.Join(distPath, "svg"))
	copyFile("src/favicon.ico", filepath.Join(distPath, "favicon.ico"))
	copyFile("src/index.html", filepath.Join(distPath, "index.html"))
	copyFile("src/overview.html", filepath.Join(distPath, "overview.html"))
}

func extractMaturity(content []byte) string {
	// Define a regular expression to match the maturity comment
	re := regexp.MustCompile(`<!--\s*maturity:\s*([^\s]+)\s*-->`)

	// Find the first match in the content
	match := re.FindStringSubmatch(string(content))

	// If a match is found, return the captured maturity value, otherwise return an empty string
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func fallbackToSeedling(maturity string) string {
	if maturity == "" {
		return "seedling"
	}
	return maturity
}

func removeMaturityComment(content []byte) []byte {
	// Define a regular expression to match the entire line containing the maturity comment and the newline after it
	re := regexp.MustCompile(`(?m)^.*<!--\s*maturity:\s*([^\s]+)\s*-->.*\n*`)

	// Remove the entire line containing the maturity comment and the newline after it from the content
	return re.ReplaceAll(content, []byte(""))
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func copyDir(src, dest string) error {
	sourceInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceInfo.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			err = copyDir(sourcePath, destPath)
			if err != nil {
				fmt.Printf("Error copying directory: %v\n", err)
			}
		} else {
			err = copyFile(sourcePath, destPath)
			if err != nil {
				fmt.Printf("Error copying file: %v\n", err)
			}
		}
	}

	return nil
}
