---
+++
@@ -0,0 +1,74 @@
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
+)
+
+func filterPackages(names []string) []string {
+
+	excludeList := []string{"common"}
+
+	var filtered []string
+
+	for _, name := range names {
+		for _, exclude := range excludeList {
+			if exclude != name {
+				filtered = append(filtered, name)
+			}
+		}
+	}
+
+	return filtered
+
+}
+
+func fileIsPackage(f os.FileInfo) bool {
+	target, err := os.Readlink(f.Name())
+
+	if err == nil {
+		f, _ = os.Lstat(target)
+	}
+
+	if !strings.HasPrefix(f.Name(), ".") && f.IsDir() {
+		return true
+	}
+	return false
+}
+
+// List lists all packages in the current directory
+// TODO: Make to list only packages, currently lists all files.
+func List() []string {
+
+	var packages []string
+	var filenames []string
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
+	packages = filterPackages(filenames)
+
+	return packages
+
+}
+
+// InputCheckPackage determines whether at least one valid package has been provided
+func InputCheckPackage(cmd *cobra.Command, args []string) error {
+	if len(args) < 1 {
+		return errors.New("requires at least 1 argument")
+	}
+	return cobra.OnlyValidArgs(cmd, args)
+}
