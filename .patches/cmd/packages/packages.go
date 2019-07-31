---
+++
@@ -0,0 +1,210 @@
+package packages
+
+import (
+	"errors"
+	"io/ioutil"
+	"log"
+	"os"
+	"strings"
+
+	"github.com/spf13/cobra"
+	"gopkg.in/yaml.v3"
+)
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
+// FileIsPackage checks if a given file is (or points to) a package directory
+func FileIsPackage(f os.FileInfo) bool {
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
+// GetState will return a current and complete state of the packages both built and known
+// This includes both .eopkg files, and .yaml files. When a .eopkg file is found, it takes precedence.
+// .eopkg results are cached on disk and tied to the hash of the file providing the information.
+// When a current .eopkg file is found in a package directory, an packages of the correct build number shall be checked for state.
+//
+// Eg. If dtkwm has `dtkwm` and `dtkwm-devel` .eopkg files, they will be added separately to the state.
+// But if dtkwm is not yet built, only, `dtkwm` will make it into the state.
+func GetState() map[string]SolusPackage {
+
+	state := make(map[string]SolusPackage)
+
+	for _, n := range ListAllPackageDirectories("./") {
+
+		if packageIsBuilt(n) {
+			for k, v := range getEeopkgState(n) {
+				state[k] = SolusPackage{Attributes: v, Built: true}
+			}
+		} else {
+			state[n] = SolusPackage{Attributes: getYAMLState(n), Built: false}
+		}
+	}
+
+	return state
+}
+
+func getYAMLState(packageName string) PackageFile {
+
+	yamlData := PackageFile{}
+
+	yamlFile, err := ioutil.ReadFile(packageName + "/package.yml")
+
+	if err != nil {
+		println(packageName)
+		log.Fatalf("error: %v", err)
+	}
+
+	err = yaml.Unmarshal(yamlFile, &yamlData)
+	if err != nil {
+		println(packageName)
+		log.Fatalf("error: %v", err)
+	}
+
+	return PackageFile{
+		Name:      packageName,
+		Version:   yamlData.Version,
+		Release:   yamlData.Release,
+		Builddeps: interpretDeps(yamlData.Builddeps),
+		Rundeps:   interpretDeps(yamlData.Rundeps)}
+}
+
+func interpretDeps(deps []string) []string {
+
+	var interpretted []string
+
+	yamlData := BuilddepMap{}
+	yamlFile, err := ioutil.ReadFile("pkgconfig_dictionary.yml")
+
+	err = yaml.Unmarshal(yamlFile, &yamlData)
+	if err != nil {
+		log.Fatalf("error: %v", err)
+	}
+
+	for _, dep := range deps {
+		realdep := dep
+		for _, target := range yamlData.Dictionary {
+			for _, pkgconfig := range target.Pkgconfigs {
+
+				if "pkgconfig("+pkgconfig+")" == dep {
+					realdep = target.Name
+					break
+				}
+			}
+			if realdep != dep {
+				break
+			}
+		}
+
+		interpretted = append(interpretted, realdep)
+
+	}
+
+	packageList := ListAllPackageDirectories("./")
+
+	filtered := []string{}
+
+	for _, d := range interpretted {
+		trimmedName := testTrimmedName(d)
+		for _, provided := range packageList {
+			if trimmedName == provided {
+				filtered = append(filtered, d)
+			}
+		}
+	}
+
+	return filtered
+
+}
+
+func packageNameTrimmer(packageName string) (nameCandidates []string) {
+
+	// TODO: Track down a definitive list of patterns to work with.
+	trimList := [...]string{
+		"-devel",
+		"-doc",
+	}
+	for _, suffix := range trimList {
+		nameCandidates = append(nameCandidates, strings.TrimSuffix(packageName, suffix))
+	}
+
+	return nameCandidates
+
+}
+
+func testTrimmedName(packageName string) string {
+	for _, trimmedName := range packageNameTrimmer(packageName) {
+		if trimmedName != packageName {
+			return trimmedName
+		}
+	}
+	return packageName
+}
+
+// ListAllPackageDirectories lists all package directory names in the current directory
+func ListAllPackageDirectories(directory string) []string {
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
+		if FileIsPackage(f) {
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
+// InputCheckPackage determines whether at least one valid package has been provided
+func InputCheckPackage(cmd *cobra.Command, args []string) error {
+	if len(args) < 1 {
+		return errors.New("At least 1 argument is required")
+	}
+	return cobra.OnlyValidArgs(cmd, args)
+}
