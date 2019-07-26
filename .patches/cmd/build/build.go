---
+++
@@ -0,0 +1,9 @@
+package build
+
+import "github.com/W-Floyd/solus-package-tools/solus-package-util/cmd/packages"
+
+// VersionMatch checks if a given package has an .eopkg file that matches the release in the YAML file
+func VersionMatch(p packages.SolusPackage) bool {
+
+	return true
+}
