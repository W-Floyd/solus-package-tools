---
+++
@@ -0,0 +1,64 @@
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
+package eopkg
+
+import (
+	"encoding/json"
+	"fmt"
+	"log"
+
+	"github.com/getsolus/libeopkg"
+)
+
+func ExtractMetaData(filename string) (metadata *libeopkg.Metadata, err error) {
+	pkg, err := libeopkg.Open(filename)
+
+	if err != nil {
+		return metadata, err
+	}
+
+	err = pkg.ReadMetadata()
+
+	if err != nil {
+		return metadata, err
+	}
+
+	return pkg.Meta, nil
+}
+
+// PrintMetaDataJSON prints a pretty JSON output of a package metadata
+func PrintMetaDataJSON(filename string) {
+	metadata, err := ExtractMetaData(filename)
+
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	prettyJSON, err := json.MarshalIndent(*metadata, "", "\t")
+
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	fmt.Print(string(prettyJSON))
+
+}
