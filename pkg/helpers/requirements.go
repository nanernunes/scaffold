package helpers

import (
	"os/exec"
)

type Requirements struct {
	Packages []string
}

func NewRequirements(packages ...string) *Requirements {
	return &Requirements{Packages: packages}
}

func (r *Requirements) IsInstalled(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func (r *Requirements) MissingPackages() (missing []string) {
	for _, requirement := range r.Packages {
		if !r.IsInstalled(requirement) {
			missing = append(missing, requirement)
		}
	}

	return
}
