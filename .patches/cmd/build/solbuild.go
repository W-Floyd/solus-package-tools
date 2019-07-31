---
+++
@@ -0,0 +1,142 @@
+package build
+
+import (
+	"bufio"
+	"io/ioutil"
+	"log"
+	"os"
+	"os/exec"
+	"strconv"
+	"strings"
+
+	"github.com/W-Floyd/solus-package-tools/solus-package-util/cmd/packages"
+)
+
+func buildPrepare(targetPackage string, state *map[string]packages.SolusPackage) {
+	cleanCache()
+	copyToCache(targetPackage, state)
+
+	*state = packages.GetState()
+}
+
+func cleanCache() {
+
+	files, err := ioutil.ReadDir("/var/lib/solbuild/local/")
+	if err != nil {
+		log.Fatal(err)
+	}
+
+	for _, file := range files {
+		if strings.HasSuffix(file.Name(), ".eopkg") {
+			err := os.Remove("/var/lib/solbuild/local/" + file.Name())
+			if err != nil {
+				log.Fatal(err)
+			}
+		}
+
+	}
+
+}
+
+func copyToCache(targetPackage string, state *map[string](packages.SolusPackage)) {
+
+	fileList := make(map[string](packages.EopkgCopy))
+
+	for _, buildDep := range (*state)[targetPackage].Attributes.Builddeps {
+
+		directory := packages.TestTrimmedName(buildDep)
+		filename := buildDep + "-" + (*state)[buildDep].Attributes.Version + "-" + strconv.Itoa((*state)[buildDep].Attributes.Release) + "-1-x86_64.eopkg"
+
+		fileList[buildDep] = packages.EopkgCopy{
+			Filename:  filename,
+			Directory: directory,
+		}
+
+		for key, value := range listRundepRecurse(buildDep, state) {
+			fileList[key] = value
+		}
+	}
+
+	for _, value := range fileList {
+
+		b, err := ioutil.ReadFile("./" + value.Directory + "/" + value.Filename)
+		if err != nil {
+			log.Fatal(err)
+		}
+
+		// write the whole body at once
+		err = ioutil.WriteFile("/var/lib/solbuild/local/"+value.Filename, b, 0644)
+		if err != nil {
+			log.Fatal(err)
+		}
+	}
+
+}
+
+// listRundepRecurse returns a map of all rundeps and the file location of each package
+func listRundepRecurse(packageName string, state *map[string]packages.SolusPackage) map[string](packages.EopkgCopy) {
+
+	fileList := make(map[string](packages.EopkgCopy))
+
+	for _, rundep := range (*state)[packageName].Attributes.Rundeps {
+
+		directory := packages.TestTrimmedName(rundep)
+		filename := rundep + "-" + (*state)[rundep].Attributes.Version + "-" + strconv.Itoa((*state)[rundep].Attributes.Release) + "-1-x86_64.eopkg"
+
+		fileList[rundep] = packages.EopkgCopy{
+			Filename:  filename,
+			Directory: directory,
+		}
+
+		for key, value := range listRundepRecurse(rundep, state) {
+			fileList[key] = value
+		}
+
+	}
+
+	return fileList
+
+}
+
+func solbuildOffload(targetPackage string, state *map[string]packages.SolusPackage) {
+
+	var cmd *exec.Cmd
+
+	if len((*state)[targetPackage].Attributes.Builddeps) > 0 {
+
+		cmd = exec.Command("make", "local")
+
+	} else {
+
+		cmd = exec.Command("make")
+	}
+
+	cmd.Dir = "./" + packages.TestTrimmedName(targetPackage) + "/"
+
+	stdout, _ := cmd.StdoutPipe()
+
+	err := cmd.Start()
+
+	scanner := bufio.NewScanner(stdout)
+	for scanner.Scan() {
+		m := scanner.Text()
+		log.Printf(m)
+	}
+
+	cmd.Wait()
+
+	if err != nil {
+		(*state)[targetPackage] = packages.SolusPackage{
+			Attributes: (*state)[targetPackage].Attributes,
+			Built:      false,
+			Failed:     true,
+		}
+	} else {
+		(*state)[targetPackage] = packages.SolusPackage{
+			Attributes: (*state)[targetPackage].Attributes,
+			Built:      true,
+			Failed:     false,
+		}
+	}
+
+}
