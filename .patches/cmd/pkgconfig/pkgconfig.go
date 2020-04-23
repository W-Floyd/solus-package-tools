---
+++
@@ -0,0 +1,181 @@
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
+package pkgconfig
+
+import (
+	"encoding/json"
+	"io/ioutil"
+	"os"
+
+	"github.com/W-Floyd/solus-package-tools/solus-package-tools/cmd/eopkg"
+	"github.com/W-Floyd/solus-package-tools/solus-package-tools/cmd/packages"
+	"github.com/getsolus/libeopkg"
+)
+
+// Dictionary will be accessible by later functions so that pkgconfigs can be cached automatically whenever a package is built
+// It also allows us to modify the pkgconfigdictionary globally, as needed (especially in build)
+var Dictionary Dict
+
+var dictionaryLocation = "pkgconfigDictionary.yaml"
+
+// Dict is used to know where to find any given pkgconfig
+type Dict struct {
+	Packages map[string]Entry
+}
+
+// Entry is used to know what specific eopkg file provides which pkgconfig
+type Entry struct {
+	EopkgFile map[string][]string
+}
+
+func LoadDictionary() error {
+
+	if _, err := os.Stat(dictionaryLocation); err == nil {
+		data, err := ioutil.ReadFile(dictionaryLocation)
+
+		if err != nil {
+			return err
+		}
+
+		initDictionary()
+
+		err = json.Unmarshal(data, &Dictionary)
+
+		if err != nil {
+			return err
+		}
+
+		return nil
+
+	}
+
+	initDictionary()
+
+	err := WriteDictionary()
+
+	if err != nil {
+		return err
+	}
+
+	return nil
+
+}
+
+func WriteDictionary() error {
+
+	initDictionary()
+
+	data, err := json.MarshalIndent(Dictionary, "", "\t")
+
+	if err != nil {
+		return err
+	}
+
+	err = ioutil.WriteFile(dictionaryLocation, data, 0644)
+
+	if err != nil {
+		return err
+	}
+
+	return nil
+
+}
+
+func UpdateDictionary() error {
+
+	initDictionary()
+
+	packageList := packages.ListNames("./")
+
+	for _, pName := range packageList {
+		UpdateEntry(pName)
+	}
+
+	return nil
+
+}
+
+func initDictionary() {
+	if Dictionary.Packages == nil {
+		Dictionary = Dict{}
+		Dictionary.Packages = map[string]Entry{}
+	}
+}
+
+func UpdateEntry(pName string) {
+
+	eopkgFiles := packages.ListCurrentEopkgFiles(pName + "/")
+	newEntry := Entry{}
+	for _, fileName := range eopkgFiles {
+		var meta *libeopkg.Metadata
+		meta, err := eopkg.ExtractMetaData(pName + "/" + fileName)
+		if err == nil {
+			provides := (*meta).Package.Provides
+			if provides != nil {
+				if newEntry.EopkgFile == nil {
+					newEntry = Entry{
+						EopkgFile: map[string][]string{},
+					}
+				}
+				newEntry.EopkgFile[fileName] = (*provides).PkgConfig
+			}
+		}
+	}
+	if newEntry.EopkgFile != nil {
+		Dictionary.Packages[pName] = newEntry
+	}
+}
+
+func PkgConfigsToBuildDeps(pkgConfigs []string) (buildDeps []string) {
+
+	for p, entry := range Dictionary.Packages {
+		for eopkgFile, provides := range entry.EopkgFile {
+			needForBuild := false
+			for _, provision := range provides {
+				if needForBuild {
+					break
+				}
+				for _, need := range pkgConfigs {
+					if needForBuild {
+						break
+					}
+					if need == provision {
+						needForBuild = true
+						break
+					}
+				}
+			}
+			if needForBuild {
+				packageMeta, err := eopkg.ExtractMetaData(p + "/" + eopkgFile)
+
+				if err == nil {
+					packageName := (*packageMeta).Package.Name
+					buildDeps = append(buildDeps, packageName)
+				}
+			}
+		}
+
+	}
+
+	return
+
+}
