---
+++
@@ -0,0 +1,60 @@
+package packages
+
+import "encoding/xml"
+
+// PackageFile holds all information for a package, as parsed from YAML or .eopkg
+type PackageFile struct {
+	Name      string
+	Version   string
+	Release   int
+	Builddeps []string
+	Rundeps   []string
+}
+
+// SolusPackage holds all information related to a package
+type SolusPackage struct {
+	Attributes PackageFile
+	Built      bool
+}
+
+// builddepEntry holds one set of mappings for
+type builddepEntry struct {
+	Name       string
+	Pkgconfigs []string
+}
+
+// BuilddepMap holds the
+type BuilddepMap struct {
+	Dictionary []builddepEntry
+}
+
+// GlobalState holds the overall state of the whole of the repository, including things such as whether a package is currently built or not.
+type GlobalState struct {
+	Packages []packageState
+}
+
+type packageState struct {
+	PackageValues PackageFile
+}
+
+////////////////////////////////////////////////////////////////////////////////
+
+// Eopkg holds the structure of a .eopkg file from XML
+type Eopkg struct {
+	XMLName xml.Name     `xml:"PISI"`
+	Package eopkgPackage `xml:"Package"`
+}
+
+type eopkgPackage struct {
+	XMLName xml.Name `xml:"Package"`
+	Name    string   `xml:"Name"`
+
+	RuntimeDependencies []string      `xml:"RuntimeDependencies>Dependency"`
+	History             []eopkgUpdate `xml:"History>Update"`
+}
+
+type eopkgUpdate struct {
+	XMLName xml.Name `xml:"Update"`
+	Release string   `xml:"release,attr"`
+	Version string   `xml:"Version"`
+}
