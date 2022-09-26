package goVersion

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"sort"
)

func Use() {
	v1, err := version.NewVersion("1.2")
	if err != nil {
		fmt.Println(err)
	}
	v2, err := version.NewVersion("1.4+metadata")

	// 比较版本
	if v1.LessThan(v2) {
		fmt.Println("v1 is less than v2")
	}

	// 检查并约束版本
	constraints, err := version.NewConstraint(">= 1.0, < 1.4")
	if constraints.Check(v1) {
		fmt.Printf("%s satisify constraints %s", v1, constraints)
		fmt.Println()
	}

	// 对一组版本号进行排序
	versionRaw := []string{"1.1", "0.7.1", "1.4-beta", "1.4", "2"}
	versions := make([]*version.Version, len(versionRaw))
	for i, raw := range versionRaw {
		v, _ := version.NewVersion(raw)
		versions[i] = v
	}
	sort.Sort(version.Collection(versions))
	fmt.Println(versions, "aaa")
}
