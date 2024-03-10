// Package actions holds a set of actions that can be run on an Gunreal project.
package actions

import (
	"fmt"

	"github.com/cristiandonosoc/gunreal/pkg/unreal"

	goversion "github.com/hashicorp/go-version"
)

type GunrealActions struct {
	BuildUBT func() error
	RunUBT   func(args []string) error

	project *unreal.Project
}

// NewGunrealActions generates the set of actions available for a particular config.
// It will take into account the version of the engine to correctly account for any differences.
func NewGunrealActions(project *unreal.Project) (*GunrealActions, error) {
	if err := validate(project); err != nil {
		return nil, fmt.Errorf("validating project: %w", err)
	}

	actions := &GunrealActions{
		project: project,
	}

	if err := obtainUBTFunctions(actions); err != nil {
		return nil, fmt.Errorf("obtaining UBT functions: %w", err)
	}

	return actions, nil
}

func validate(project *unreal.Project) error {
	// Ensure the editor is valid.
	if project.UnrealEditor == nil {
		return fmt.Errorf("actions requires an associated engine. Is gunreal.yml correctly set?")
	}

	constraints := goversion.MustConstraints(goversion.NewConstraint(">= 5.2, <= 5.3"))
	if !constraints.Check(project.UnrealEditor.Version) {
		return fmt.Errorf("editor version %q does not comply constraints %q", project.UnrealEditor.Version, constraints)
	}

	return nil
}

func obtainUBTFunctions(actions *GunrealActions) error {
	return nil
}
