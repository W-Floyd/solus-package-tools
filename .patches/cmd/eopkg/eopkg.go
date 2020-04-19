---
+++
@@ -0,0 +1,14 @@
+package eopkg
+
+import "github.com/W-Floyd/solus-package-tools/solus-package-tools/cmd/packages"
+
+// ListEopkgFiles lists all .eopkg files that are in the directory for a given package
+func ListEopkgFiles(packageName string) (packageFiles []string) {
+
+	if !packages.FileIsPackage() {
+		return packageFiles
+	}
+
+	return packageFiles
+
+}
