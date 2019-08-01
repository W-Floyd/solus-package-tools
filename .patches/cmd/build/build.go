---
+++
@@ -0,0 +1,86 @@
+package build
+
+import (
+	"fmt"
+	"log"
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
+		for _, realPackage := range *buildQueue {
+
+			targetPackage := packages.TestTrimmedName(realPackage)
+
+			globalState = packages.GetState()
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
+			// Keep in queue for now, to hopefully build later.
+			if !packages.IsPackageBuildable(targetPackage, &globalState) {
+				newQueue = append(newQueue, targetPackage)
+				continue
+			}
+
+			log.Println("Attributes of " + targetPackage)
+			log.Println(globalState[targetPackage].Attributes.Builddeps)
+
+			buildPrepare(targetPackage, &globalState)
+			solbuildOffload(targetPackage, &globalState)
+
+		}
+
+		*buildQueue = newQueue
+
+		if len(*buildQueue) == 0 {
+			continue
+		}
+
+		if !checkBuildQueueForBuildability(*buildQueue, &globalState) {
+			fmt.Println("Queue is unbuildable")
+			for _, dep := range *buildQueue {
+				log.Println(dep)
+			}
+			break
+		}
+
+	}
+
+	log.Println("Done Building")
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
