---
+++
@@ -0,0 +1,76 @@
+package packages
+
+import (
+	"crypto/sha256"
+	"io"
+	"io/ioutil"
+	"log"
+	"os"
+
+	"gopkg.in/yaml.v3"
+)
+
+// PackageCacheStore holds the current cache of packages
+var PackageCacheStore = readPackageCache()
+
+var cacheLocation = "./.packagecache"
+
+// WritePackageCache writes the current cache of packages to file
+func WritePackageCache(cache *PackageCache) {
+
+	if _, err := os.Stat(cacheLocation); err == nil {
+		os.Remove(cacheLocation)
+	}
+
+	yamldata, _ := yaml.Marshal(*cache)
+
+	ioutil.WriteFile(cacheLocation, yamldata, 0644)
+
+}
+
+func readPackageCache() PackageCache {
+
+	if _, err := os.Stat(cacheLocation); err == nil {
+
+		var yamlData PackageCache
+
+		yamlFile, err := ioutil.ReadFile(cacheLocation)
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
+		return yamlData
+
+	}
+
+	var cache PackageCache
+
+	cache.Packages = make(map[string]PackageFile)
+
+	return cache
+
+}
+
+func hashFile(fileName string) string {
+	f, err := os.Open(fileName)
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	h := sha256.New()
+	if _, err := io.Copy(h, f); err != nil {
+		log.Fatal(err)
+	}
+
+	hash := string(h.Sum(nil))
+
+	f.Close()
+
+	return hash
+}
