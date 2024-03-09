// Binary unreal_lister is a little test application to test out the unreal project querying code.
package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/cristiandonosoc/gunreal/pkg/unreal"
)

func internalMain() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("Usage: unreal_lister [PROJECT_DIR]")
	}

	projectDir := os.Args[1]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	start := time.Now()

	project, err := unreal.NewProjectFromPath(projectDir)
	if err != nil {
		return fmt.Errorf("reading project at %q: %w", projectDir, err)
	}

	if err := project.IndexModules(ctx); err != nil {
		return fmt.Errorf("indexing unreal project at %q: %w", projectDir, err)
	}

	duration := time.Since(start)

	// Go over the modules, but in a sorted fashion.
	modules := make([]*unreal.Module, 0, len(project.Modules))

	for _, module := range project.Modules {
		modules = append(modules, module)

		// We prefetch the modules.
		_, err := module.LoadUHTFiles(unreal.Platform_Windows, true)
		if err != nil {
			return fmt.Errorf("loading UHT files for module %q: %w", module.Name, err)
		}
	}
	sort.Slice(modules, func(i, j int) bool {
		return modules[i].Name < modules[j].Name
	})

	for _, module := range modules {
		uhtFiles, err := module.LoadUHTFiles(unreal.Platform_Windows, false)
		if err != nil {
			return fmt.Errorf("loading UHT files for module %q: %w", module.Name, err)
		}

		fmt.Println("---------------------------------------------------------")
		fmt.Println("MODULE:", module.Name)
		fmt.Println("BASE DIR:", module.BaseDir)
		fmt.Println("BUILD FILE:", module.BuildFile)
		fmt.Println("FILES:", len(module.Files))
		// for _, file := range module.Files {
		// 	fmt.Println("-", file)
		// }
		fmt.Println("UHT FILES:", len(uhtFiles))
		// for _, file := range uhtFiles {
		// 	fmt.Println("-", file)
		// }
	}

	fmt.Printf("Indexing took %v to execute.\n", duration)

	return nil
}

func main() {
	if err := internalMain(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
