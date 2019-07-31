---
+++
@@ -0,0 +1,67 @@
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
+	Failed     bool
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
+////////////////////////////////////////////////////////////////////////////////
+
+// Eopkg holds the important parts of the structure of a .eopkg file from XML
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
+
+////////////////////////////////////////////////////////////////////////////////
+
+// PackageCache holds a map of package information, as indexed by the hash of the referenced file.
+type PackageCache struct {
+	Packages map[string]PackageFile
+}
+
+////////////////////////////////////////////////////////////////////////////////
+
+// EopkgCopy holds information for when we transfer files to eopkg cache
+type EopkgCopy struct {
+	Filename  string
+	Directory string
+}
