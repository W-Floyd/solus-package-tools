---
+++
@@ -0,0 +1,91 @@
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
+)
+
+// PkgConfigDictionary will be accessible by later functions so that pkgconfigs can be cached automatically whenever a package is built
+var PkgConfigDictionaryVar PkgConfigDictionary
+
+var dictionaryLocation = "pkgconfigDictionary.yaml"
+
+type PkgConfigDictionary struct {
+	Packages *[]PkgEntry
+}
+
+type PkgEntry struct {
+	PkgName   string
+	EopkgFile *map[string]string
+}
+
+func LoadPkgConfigDictionary() (err error) {
+
+	if _, err := os.Stat(dictionaryLocation); err == nil {
+		data, err := ioutil.ReadFile(dictionaryLocation)
+
+		if err != nil {
+			return err
+		}
+
+		err = json.Unmarshal(data, &PkgConfigDictionaryVar)
+
+		if err != nil {
+			return err
+		}
+
+		return nil
+
+	}
+
+	PkgConfigDictionaryVar = PkgConfigDictionary{}
+
+	err = WritePkgConfigDictionary()
+
+	if err != nil {
+		return err
+	}
+
+	return nil
+
+}
+
+func WritePkgConfigDictionary() (err error) {
+
+	data, err := json.Marshal(PkgConfigDictionaryVar)
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
