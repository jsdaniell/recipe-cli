package react_native_expo

import (
	"github.com/jsdaniell/recipe-cli/utils/shell_commands"
)

func InitRoot(projectName string) {
	shell_commands.ExecuteShellCommand("npx expo-cli expo init", projectName)
}
