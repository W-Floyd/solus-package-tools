---
+++
@@ -0,0 +1,75 @@
+package build
+
+import (
+	"fmt"
+
+	"github.com/W-Floyd/solus-package-tools/solus-package-util/cmd/packages"
+)
+
+var globalState = packages.GetState()
+
+// ProcessQueue takes a reference to a process queue and iterates through it until it's empty
+func ProcessQueue(buildQueue *[]string) {
+
+	for len(*buildQueue) > 0 {
+
+		newQueue := []string{}
+
+		for _, targetPackage := range *buildQueue {
+
+			if globalState[targetPackage].Built || packages.IsPackageFailed(targetPackage, &globalState) {
+				continue
+			}
+
+			for _, builddep := range globalState[targetPackage].Attributes.Builddeps {
+				newQueue = append(newQueue, builddep)
+				newQueue = append(newQueue, packages.RundepRecurse(builddep, &globalState)...)
+			}
+
+			if !packages.IsPackageBuildable(targetPackage, &globalState) {
+				newQueue = append(newQueue, targetPackage)
+				//fmt.Println(targetPackage + " is unbuildable.")
+				continue
+			}
+
+			buildPrepare(targetPackage, &globalState)
+			solbuildOffload(targetPackage, &globalState)
+
+			globalState = packages.GetState()
+
+		}
+
+		*buildQueue = newQueue
+
+		if !checkBuildQueueForBuildability(*buildQueue, &globalState) {
+			fmt.Println("Queue is unbuildable")
+			/* 			for _, dep := range globalState {
+				json, _ := json.Marshal(dep)
+				fmt.Println(string(json))
+			} */
+			break
+		}
+
+	}
+
+	packages.WritePackageCache(&packages.PackageCacheStore)
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
+
+func checkBuildQueueForBuildability(queue []string, state *map[string]packages.SolusPackage) bool {
+	for _, targetPackage := range queue {
+		if packages.IsPackageBuildable(targetPackage, state) {
+			return true
+		}
+	}
+	return false
+}
