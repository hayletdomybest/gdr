package rename

import (
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/mod/modfile"
)

func RenameModule(projectPath, newName string) error {
	oldName, err := getOldModuleName(projectPath)
	if err != nil {
		return fmt.Errorf("failed to get old module name: %v", err)
	}

	err = updateGoMod(projectPath, newName)
	if err != nil {
		return fmt.Errorf("failed to update go.mod: %v", err)
	}

	err = updateImports(projectPath, oldName, newName)
	if err != nil {
		return fmt.Errorf("failed to update imports: %v", err)
	}

	fmt.Printf("Successfully renamed module from %s to %s\n", oldName, newName)
	return nil
}

func getOldModuleName(projectPath string) (string, error) {
	goModPath := filepath.Join(projectPath, "go.mod")
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`module\s+(.+)`)
	matches := re.FindSubmatch(content)
	if len(matches) < 2 {
		return "", fmt.Errorf("could not find module name in go.mod")
	}

	return string(matches[1]), nil
}

func updateGoMod(projectPath, newName string) error {
	goModPath := filepath.Join(projectPath, "go.mod")

	// Read the go.mod file
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return fmt.Errorf("error reading go.mod: %v", err)
	}

	// Parse the go.mod file
	f, err := modfile.Parse(goModPath, content, nil)
	if err != nil {
		return fmt.Errorf("error parsing go.mod: %v", err)
	}

	// Update the module name
	if err := f.AddModuleStmt(newName); err != nil {
		return fmt.Errorf("error updating module name: %v", err)
	}

	// Format the go.mod file
	newContent, err := f.Format()
	if err != nil {
		return fmt.Errorf("error formatting go.mod: %v", err)
	}

	// Write the updated content
	if err := os.WriteFile(goModPath, newContent, 0644); err != nil {
		return fmt.Errorf("error writing go.mod: %v", err)
	}

	return nil
}

func updateImports(projectPath, oldName, newName string) error {
	return filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Parse the file
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("error parsing file %s: %v", path, err)
		}

		modified := false
		for _, imp := range file.Imports {
			if strings.Contains(imp.Path.Value, oldName) {
				newPath := strings.Replace(imp.Path.Value, oldName, newName, 1)
				imp.Path.Value = newPath
				modified = true
			}
		}

		if !modified {
			return nil
		}

		// Format the modified AST
		var buf strings.Builder
		err = format.Node(&buf, fset, file)
		if err != nil {
			return fmt.Errorf("error formatting file %s: %v", path, err)
		}

		// Write the formatted code back to file
		err = os.WriteFile(path, []byte(buf.String()), 0644)
		if err != nil {
			return fmt.Errorf("error writing to file %s: %v", path, err)
		}

		fmt.Printf("Updated imports in %s\n", path)
		return nil
	})
}
