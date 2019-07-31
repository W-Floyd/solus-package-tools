---
+++
@@ -0,0 +1,111 @@
+package packages
+
+import (
+	"encoding/xml"
+	"io/ioutil"
+	"log"
+	"os/exec"
+	"regexp"
+	"strconv"
+)
+
+// From my old bash script, this was used:
+// sed 's#.*-\([0-9]*\)-1-x86_64\.eopkg$#\1#'
+var eopkgRegex = regexp.MustCompile("^.*-([0-9]*)-1-x86_64\\.eopkg$")
+
+func getEeopkgState(packageName string) map[string]PackageFile {
+	packages := make(map[string]PackageFile)
+
+	yamlData := getYAMLState(packageName)
+
+	for _, fileName := range listCurrentEopkgFiles(packageName) {
+
+		hash := hashFile("./" + packageName + "/" + fileName)
+
+		if val, ok := PackageCacheStore.Packages[hash]; ok {
+			packages[PackageCacheStore.Packages[hash].Name] = val
+		} else {
+
+			xmlData := getEopkgFileAttributes(fileName, packageName)
+
+			xmlDataParsed := Eopkg{}
+
+			xml.Unmarshal(xmlData, &xmlDataParsed)
+
+			release, _ := strconv.Atoi(xmlDataParsed.Package.History[0].Release)
+
+			foundPackage := PackageFile{
+				Name:      xmlDataParsed.Package.Name,
+				Version:   xmlDataParsed.Package.History[0].Version,
+				Release:   release,
+				Rundeps:   interpretDeps(xmlDataParsed.Package.RuntimeDependencies),
+				Builddeps: yamlData.Builddeps}
+
+			PackageCacheStore.Packages[hash] = foundPackage
+
+			packages[xmlDataParsed.Package.Name] = foundPackage
+
+		}
+
+	}
+
+	return packages
+}
+
+func getEopkgFileAttributes(fileName string, packageName string) []byte {
+
+	out, err := exec.Command("eopkg", "info", "./"+packageName+"/"+fileName, "--xml").Output()
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	return out
+
+}
+
+// PackageIsBuilt checks if the provided package name (that is, top level package) is built
+func packageIsBuilt(packageName string) bool {
+
+	if len(listCurrentEopkgFiles(packageName)) > 0 {
+		return true
+	}
+
+	return false
+
+}
+
+func listCurrentEopkgFiles(packageName string) (fileList []string) {
+
+	files, err := ioutil.ReadDir("./" + packageName)
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	currentRelease := getYAMLState(packageName).Release
+
+	for _, f := range files {
+
+		fileRelease := getFilenameRelease(f.Name())
+
+		if fileRelease > 0 {
+			if fileRelease == currentRelease {
+				fileList = append(fileList, f.Name())
+			}
+		}
+
+	}
+
+	return fileList
+}
+
+func getFilenameRelease(fileName string) int {
+
+	fileRelease := eopkgRegex.FindStringSubmatch(fileName)
+
+	if len(fileRelease) > 0 {
+		value, _ := strconv.Atoi(fileRelease[1])
+		return value
+	}
+	return -1
+
+}
