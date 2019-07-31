---
+++
@@ -0,0 +1,45 @@
+package build
+
+import (
+	"encoding/json"
+	"fmt"
+
+	"github.com/W-Floyd/solus-package-tools/solus-package-util/cmd/packages"
+)
+
+var packageState = packages.GetState()
+
+// ProcessQueue takes a reference to a process queue and iterates through it until it's empty
+func ProcessQueue(buildQueue *[]string) {
+
+	/* 	for len(*buildQueue) > 0 {
+
+		newQueue := *buildQueue
+
+		for _, targetPackage := range *buildQueue {
+
+			fmt.Println(targetPackage)
+			packageState = packages.GetState()
+
+		}
+
+		*buildQueue = newQueue
+
+	} */
+
+	packageState = packages.GetState()
+
+	json, _ := json.Marshal(packageState)
+
+	fmt.Println(string(json))
+
+}
+
+func stringInSlice(a string, list []string) bool {
+	for _, b := range list {
+		if b == a {
+			return true
+		}
+	}
+	return false
+}
