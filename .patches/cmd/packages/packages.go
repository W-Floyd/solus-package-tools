---
+++
@@ -0,0 +1,134 @@
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
+	"gopkg.in/yaml.v2"
+)
+
+type PackageFile struct {
+	Name    string
+	Version string
+	Release int
+}
+
+// SolusPackage holds all information related to a package
+type SolusPackage struct {
+	Name    string
+	Version string
+	Release int
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
+// List lists all packages in the current directory, with their information
+func List() []SolusPackage {
+
+	var packageList []SolusPackage
+
+	for _, n := range ListNames() {
+
+		yamlData := PackageFile{}
+		yamlFile, err := ioutil.ReadFile(n + "/package.yml")
+
+		if err != nil {
+			log.Fatalf("error: %v", err)
+		}
+
+		err = yaml.Unmarshal(yamlFile, &yamlData)
+		if err != nil {
+			log.Fatalf("error: %v", err)
+		}
+
+		packageList = append(packageList, SolusPackage{Name: n, Version: yamlData.Version, Release: yamlData.Release})
+
+	}
+
+	return packageList
+
+}
+
+// ListNames lists all package names in the current directory
+func ListNames() []string {
+
+	var filenames []string
+	var packageNames []string
+
+	files, err := ioutil.ReadDir("./")
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
+// InputCheckPackage determines whether at least one valid package has been provided
+func InputCheckPackage(cmd *cobra.Command, args []string) error {
+	if len(args) < 1 {
+		return errors.New("At least 1 argument is required")
+	}
+	return cobra.OnlyValidArgs(cmd, args)
+}
