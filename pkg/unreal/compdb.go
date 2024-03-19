package unreal

import "fmt"

var (
	gCompdb_ubtArgs = []string{
		"-ProjectFiles",
		"-VSCode",
		"-Game",
		"-Engine",
		//"-Progress",
		"-NoIntelliSense",
	}
)

func (p *Project) GenerateCompDB() error {
	// Use UBT to generate the VSCode compilation database.
	if err := p.UBT(gCompdb_ubtArgs); err != nil {
		return fmt.Errorf("generating project files: %w", err)
	}

	return nil
}