---
+++
@@ -0,0 +1,170 @@
+/*
+Copyright © 2020 William Floyd <william.png2000@gmail.com>
+
+Permission is hereby granted, free of charge, to any person obtaining a copy
+of this software and associated documentation files (the "Software"), to deal
+in the Software without restriction, including without limitation the rights
+to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
+copies of the Software, and to permit persons to whom the Software is
+furnished to do so, subject to the following conditions:
+
+The above copyright notice and this permission notice shall be included in
+all copies or substantial portions of the Software.
+
+THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
+IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
+FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
+AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
+LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
+OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
+THE SOFTWARE.
+*/
+package packages
+
+import (
+	"bytes"
+	"encoding/json"
+	"fmt"
+	"io/ioutil"
+	"log"
+	"os"
+	"regexp"
+	"strconv"
+	"strings"
+
+	"github.com/Jeffail/gabs"
+
+	yamltojson "github.com/ghodss/yaml"
+)
+
+// fileIsPackage checks if a given file is (or points to) a package directory
+func fileIsPackage(f os.FileInfo) bool {
+
+	// Try to read the package as a symlink
+	target, err := os.Readlink(f.Name())
+
+	// If it is a symlink, change the file to point to the symlink target
+	if err == nil {
+		f, _ = os.Lstat(target)
+	}
+
+	// If it is a directory, and does not start with `.`
+	if !strings.HasPrefix(f.Name(), ".") && f.IsDir() {
+
+		// If a `package.yml` exists
+		if _, err := os.Stat(f.Name() + "/package.yml"); !os.IsNotExist(err) {
+			return true
+		}
+
+	}
+	return false
+}
+
+// ListNames lists all package names in the current directory
+func ListNames(directory string) []string {
+
+	var filenames []string
+	var packageNames []string
+
+	files, err := ioutil.ReadDir(directory)
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	for _, f := range files {
+		if fileIsPackage(f) {
+			filenames = append(filenames, f.Name())
+		}
+	}
+
+	packageNames = filterPackages(filenames)
+
+	return packageNames
+
+}
+
+func filterPackages(names []string) []string {
+
+	// Directories that should be ignored, and cannot not be a package.
+	excludeList := []string{"common"}
+
+	var filtered []string
+
+	// For every name fed in
+	for _, name := range names {
+		shouldExclude := false
+		// For every name to exclude
+		for _, exclude := range excludeList {
+			// Mark for exclusion if part of the list
+			if exclude == name {
+				shouldExclude = true
+				break
+			}
+		}
+
+		// If not to be excluded, add to list
+		if !shouldExclude {
+			filtered = append(filtered, name)
+		}
+
+	}
+
+	return filtered
+
+}
+
+func ListCurrentEopkgFiles(packagePath string) []string {
+	files, err := ioutil.ReadDir(packagePath)
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	currentRelease, err := GetCurrentRelease(packagePath)
+
+	if err != nil {
+		return nil
+	}
+
+	r := regexp.MustCompile(strconv.FormatInt(currentRelease, 10) + `-1-x86_64\.eopkg$`)
+
+	matchedFiles := []string{}
+
+	for _, f := range files {
+
+		if r.MatchString(f.Name()) {
+			matchedFiles = append(matchedFiles, f.Name())
+		}
+	}
+
+	return matchedFiles
+}
+
+func GetCurrentRelease(packagePath string) (release int64, err error) {
+	yamlData, err := ioutil.ReadFile(packagePath + "/package.yml")
+
+	if err != nil {
+		return -1, err
+	}
+
+	jsonData, err := yamltojson.YAMLToJSON(yamlData)
+
+	if err != nil {
+		return -1, err
+	}
+
+	dec := json.NewDecoder(bytes.NewReader(jsonData))
+	dec.UseNumber()
+
+	jsonParsed, err := gabs.ParseJSONDecoder(dec)
+
+	if err != nil {
+		return -1, err
+	}
+
+	value, err := jsonParsed.Path("release").Data().(json.Number).Int64()
+
+	if err == nil {
+		return value, nil
+	}
+	return -1, fmt.Errorf("Release value not found")
+}
