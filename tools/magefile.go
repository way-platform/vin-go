//go:build mage

package main

import (
	"io/fs"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Build

// Build runs a full CI build.
func Build() {
	mg.SerialDeps(
		Download,
		Lint,
		Test,
		Tidy,
		CLI,
		Diff,
	)
}

// Lint runs the Go linter.
func Lint() error {
	return forEachGoMod(func(dir string) error {
		return tool(dir, "golangci-lint", "run", "--path-prefix", dir, "--build-tags", "mage").Run()
	})
}

// Test runs the Go tests.
func Test() error {
	return cmd(root(), "go", "test", "-v", "-cover", "./...").Run()
}

// Download downloads the Go dependencies.
func Download() error {
	return forEachGoMod(func(dir string) error {
		return cmd(dir, "go", "mod", "download").Run()
	})
}

// Tidy tidies the Go mod files.
func Tidy() error {
	return forEachGoMod(func(dir string) error {
		return cmd(dir, "go", "mod", "tidy", "-v").Run()
	})
}

// Diff checks for git diffs.
func Diff() error {
	return cmd(root(), "git", "diff", "--exit-code").Run()
}

// CLI builds the CLI.
func CLI() error {
	return cmd(root("cmd/vin"), "go", "install", ".").Run()
}

// VPIC downloads the vPIC database.
func VPIC() error {
	const fileName = "vPICList_lite_2025_11.bak.zip"
	const url = "https://vpic.nhtsa.dot.gov/api/" + fileName
	if err := cmd(root(), "curl", "--create-dirs", "-L", "-o", root("data", "vpic", fileName), url).Run(); err != nil {
		return err
	}
	return cmd(root("data", "vpic"), "unzip", fileName).Run()
}

// KBA downloads the KBA (Kraftfahrt-Bundesamt) manufacturer database.
func KBA() error {
	const fileName = "sv32_pdf_en.pdf"
	const url = "https://www.kba.de/SharedDocs/Downloads/EN/SV/" + fileName + "?__blob=publicationFile&v=2"
	slog.Info("downloading KBA PDF", "url", url, "fileName", fileName)
	if err := cmd(root(), "curl", "--create-dirs", "-L", "-o", root("data", "kba", fileName), url).Run(); err != nil {
		return err
	}
	slog.Info("splitting KBA PDF into pages", "path", root("data", "kba", "pages"))
	if err := os.RemoveAll(root("data", "kba", "pages")); err != nil {
		return err
	}
	if err := os.MkdirAll(root("data", "kba", "pages"), 0o700); err != nil {
		return err
	}
	if err := tool(root("data", "kba"), "pdfcpu", "split", fileName, "pages").Run(); err != nil {
		return err
	}
	pagesDir := root("data", "kba", "pages")
	return filepath.WalkDir(pagesDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".pdf" {
			return nil
		}
		outPath := path[0 : len(path)-len(filepath.Ext(path))]
		slog.Info("converting PDF to PNG", "path", path, "outPath", outPath)
		return cmd(pagesDir, "pdftoppm", "-png", "-singlefile", path, outPath).Run()
	})
}

// KBAToCSV converts the KBA pages to CSV.
func KBAToCSV() error {
	return filepath.WalkDir(root("data", "kba", "pages"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".png" {
			return nil
		}
		filename := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		pageNumber, err := strconv.Atoi(strings.Split(filename, "_")[3])
		if err != nil {
			return err
		}
		if pageNumber < 6 || pageNumber > 147 {
			slog.Info("skipping page", "pageNumber", pageNumber)
			return nil
		}
		csvPath := root("data", "kba", "pages", filename+".csv")
		if _, err := os.Stat(csvPath); err == nil {
			slog.Info("already exists", "pageNumber", pageNumber, "csvPath", csvPath)
			return nil
		}
		return cmd(root("tools", "cmd", "kba-to-csv"), "go", "run", ".", path).Run()
	})
}

// KBACollate collates the KBA pages into a single CSV file.
func KBACollate() error {
	return cmd(
		root("tools", "cmd", "kba-collate"),
		"go", "run", ".",
		"-d", root("data", "kba", "pages"),
		"-o", root("data", "kba", "wmi.csv"),
	).Run()
}

// WikibooksToCSV converts the Wikibooks WMI data to CSV.
func WikibooksToCSV() error {
	return cmd(
		root("tools", "cmd", "wikibooks-to-csv"),
		"go", "run", ".",
		"-i", root("data", "wikibooks", "wmi.md"),
		"-o", root("data", "wikibooks", "wmi.csv"),
	).Run()
}

func forEachGoMod(f func(dir string) error) error {
	return filepath.WalkDir(root(), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || d.Name() != "go.mod" {
			return nil
		}
		return f(filepath.Dir(path))
	})
}

func cmd(dir string, command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func tool(dir string, tool string, args ...string) *exec.Cmd {
	cmdArgs := []string{"tool", "-modfile", filepath.Join(root(), "tools", "go.mod"), tool}
	return cmd(dir, "go", append(cmdArgs, args...)...)
}

func root(subdirs ...string) string {
	result, err := sh.Output("git", "rev-parse", "--show-toplevel")
	if err != nil {
		panic(err)
	}
	return filepath.Join(append([]string{result}, subdirs...)...)
}
