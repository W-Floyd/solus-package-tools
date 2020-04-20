---
+++
@@ -0,0 +1,103 @@
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
+	"archive/zip"
+	"bytes"
+	jsonreal "encoding/json"
+	"fmt"
+	"log"
+
+	xj "github.com/basgys/goxml2json"
+)
+
+// extractMetaDataXML streams the metadata XML file from the archive
+func extractMetaDataXML(fileName string) (metadata *bytes.Buffer, err error) {
+
+	r, err := zip.OpenReader(fileName)
+
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	defer r.Close()
+
+	for _, f := range r.File {
+
+		if f.Name == "metadata.xml" {
+
+			xml, err := f.Open()
+
+			if err != nil {
+				log.Fatal(err)
+			}
+
+			buf := new(bytes.Buffer)
+
+			buf.ReadFrom(xml)
+
+			xml.Close()
+
+			return buf, nil
+
+		}
+	}
+
+	return metadata, fmt.Errorf("zip: metadata not found in file %s", fileName)
+
+}
+
+func extractMetaDataJSON(filename string) (metadata *bytes.Buffer, err error) {
+	r, err := extractMetaDataXML(filename)
+
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	json, err := xj.Convert(r)
+
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	return json, nil
+}
+
+// PrintMetaDataJSON prints a pretty JSON output of a package metadata
+func PrintMetaDataJSON(filename string) {
+	r, err := extractMetaDataJSON(filename)
+
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	var prettyJSON bytes.Buffer
+	err = jsonreal.Indent(&prettyJSON, r.Bytes(), "", "\t")
+
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	fmt.Print(prettyJSON.String())
+
+}
