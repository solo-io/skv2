package funcs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib/codegen/funcs"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var _ = Describe("Funcs", func() {

	Context("MakeHomogenousSnapshotFuncs", func() {
		It("generates the expected functions", func() {
			selectFromGroups := map[string][]model.Group{
				"github.com/solo-io/my-module": {
					{
						GroupVersion: schema.GroupVersion{
							Group:   "a.solo.io",
							Version: "v1",
						},
						Module:  "github.com/solo-io/my-module",
						ApiRoot: "pkg/api",
						Resources: []model.Resource{
							{
								Kind: "Basket",
							},
							{
								Kind: "Box",
							},
							{
								Kind: "Egg",
							},
						},
					},
					{
						GroupVersion: schema.GroupVersion{
							Group:   "foo.solo.io",
							Version: "v1",
						},
						Module:  "github.com/solo-io/my-module",
						ApiRoot: "pkg/api",
						Resources: []model.Resource{
							{
								Kind: "Paint",
							},
							{
								Kind: "House",
							},
							{
								Kind: "Color",
							},
						},
					},
					{
						GroupVersion: schema.GroupVersion{
							Group:   "b.solo.io",
							Version: "v3",
						},
						Module:  "github.com/solo-io/my-module",
						ApiRoot: "pkg/api",
						Resources: []model.Resource{
							{
								Kind: "Chicken",
							},
							{
								Kind: "Duck",
							},
						},
					},
				},
			}
			resourcesToSelect := map[schema.GroupVersion][]string{
				schema.GroupVersion{
					Group:   "a.solo.io",
					Version: "v1",
				}: {"Egg", "Basket"},
				schema.GroupVersion{
					Group:   "b.solo.io",
					Version: "v3",
				}: {"Chicken"},
			}

			funcMap := funcs.MakeHomogenousSnapshotFuncs("my.snapshot.name", "pkg/api/output/filename.go",
				selectFromGroups, resourcesToSelect)

			// check that all the functions return expected values
			Expect(funcMap["snapshot_name"].(func() string)()).To(Equal("my.snapshot.name"))
			Expect(funcMap["package"].(func() string)()).To(Equal("output"))

			importedGroups, err := funcMap["imported_groups"].(func() ([]model.Group, error))()
			Expect(err).NotTo(HaveOccurred())
			Expect(importedGroups).To(HaveLen(2))
			// we can't do an equality check on the entire result, since the resources and groups
			// have circular references to each other. so just compare the fields we care about
			Expect(importedGroups[0].GroupVersion).To(Equal(schema.GroupVersion{
				Group:   "a.solo.io",
				Version: "v1",
			}))
			Expect(importedGroups[0].Module).To(Equal("github.com/solo-io/my-module"))
			Expect(importedGroups[0].ApiRoot).To(Equal("pkg/api"))
			Expect(importedGroups[0].Resources).To(HaveLen(2))
			Expect(importedGroups[0].Resources[0].Kind).To(Equal("Basket"))
			Expect(importedGroups[0].Resources[0].Group.GroupVersion).To(Equal(schema.GroupVersion{
				Group:   "a.solo.io",
				Version: "v1",
			}))
			Expect(importedGroups[0].Resources[1].Kind).To(Equal("Egg"))
			Expect(importedGroups[0].Resources[1].Group.GroupVersion).To(Equal(schema.GroupVersion{
				Group:   "a.solo.io",
				Version: "v1",
			}))

			Expect(importedGroups[1].GroupVersion).To(Equal(schema.GroupVersion{
				Group:   "b.solo.io",
				Version: "v3",
			}))
			Expect(importedGroups[1].Module).To(Equal("github.com/solo-io/my-module"))
			Expect(importedGroups[1].ApiRoot).To(Equal("pkg/api"))
			Expect(importedGroups[1].Resources).To(HaveLen(1))
			Expect(importedGroups[1].Resources[0].Kind).To(Equal("Chicken"))
			Expect(importedGroups[1].Resources[0].Group.GroupVersion).To(Equal(schema.GroupVersion{
				Group:   "b.solo.io",
				Version: "v3",
			}))

			Expect(funcMap["client_import_path"].(func(group model.Group) string)(model.Group{
				GroupVersion: schema.GroupVersion{
					Group:   "b.solo.io",
					Version: "v3",
				},
			})).To(Equal("github.com/solo-io/my-module/pkg/api/b.solo.io/v3"))
			Expect(funcMap["set_import_path"].(func(group model.Group) string)(model.Group{
				GroupVersion: schema.GroupVersion{
					Group:   "b.solo.io",
					Version: "v3",
				},
			})).To(Equal("github.com/solo-io/my-module/pkg/api/b.solo.io/v3/sets"))
			Expect(funcMap["controller_import_path"].(func(group model.Group) string)(model.Group{
				GroupVersion: schema.GroupVersion{
					Group:   "b.solo.io",
					Version: "v3",
				},
			})).To(Equal("github.com/solo-io/my-module/pkg/api/b.solo.io/v3/controller"))
		})

	})
})
