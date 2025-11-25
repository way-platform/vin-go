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

// VPICManual downloads the vPIC user manual.
func VPICManual() error {
	const filename = "vpic-user-manual-2023.pdf"
	const url = "https://crashstats.nhtsa.dot.gov/Api/Public/ViewPublication/813697"
	slog.Info("downloading vPIC user manual", "url", url, "filename", filename)
	if err := cmd(root(), "curl", "--create-dirs", "-L", "-o", root("data", "vpic", "manual", filename), url).Run(); err != nil {
		return err
	}
	pagesPath := root("data", "vpic", "manual", "pages")
	if err := os.RemoveAll(pagesPath); err != nil {
		return err
	}
	if err := os.MkdirAll(pagesPath, 0o700); err != nil {
		return err
	}
	// Split the PDF into pages using pdftoppm.
	return cmd(root("data", "vpic", "manual"), "pdftoppm", "-png", filename, filepath.Join(pagesPath, "manual")).Run()
}

// KBAWMIPDF downloads the KBA (Kraftfahrt-Bundesamt) WMI PDF.
func KBAWMIPDF() error {
	const fileName = "sv32_pdf_en.pdf"
	const url = "https://www.kba.de/SharedDocs/Downloads/EN/SV/" + fileName + "?__blob=publicationFile&v=2"
	slog.Info("downloading KBA PDF", "url", url, "fileName", fileName)
	if err := cmd(root(), "curl", "--create-dirs", "-L", "-o", root("data", "kba", fileName), url).Run(); err != nil {
		return err
	}
	pagesDir := root("data", "kba", "wmi")
	slog.Info("splitting KBA PDF into WMI pages", "path", pagesDir)
	if err := os.RemoveAll(pagesDir); err != nil {
		return err
	}
	if err := os.MkdirAll(pagesDir, 0o700); err != nil {
		return err
	}
	if err := tool(root("data", "kba"), "pdfcpu", "split", fileName, pagesDir).Run(); err != nil {
		return err
	}
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

// KBAWMIToCSV converts the KBA WMI pages to CSV.
func KBAWMIToCSV() error {
	return filepath.WalkDir(root("data", "kba", "wmi"), func(path string, d fs.DirEntry, err error) error {
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
		csvPath := root("data", "kba", "wmi", filename+".csv")
		if _, err := os.Stat(csvPath); err == nil {
			slog.Info("already exists", "pageNumber", pageNumber, "csvPath", csvPath)
			return nil
		}
		return cmd(root("tools", "cmd", "kba-to-csv"), "go", "run", ".", path).Run()
	})
}

// VPICManualToMarkdown converts the VPIC manual pages to markdown.
func VPICManualToMarkdown() error {
	return filepath.WalkDir(root("data", "vpic", "manual", "pages"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".png" {
			return nil
		}
		filename := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		markdownPath := root("data", "vpic", "manual", "pages", filename+".md")
		if _, err := os.Stat(markdownPath); err == nil {
			slog.Info("already exists", "markdownPath", markdownPath)
			return nil
		}
		return cmd(root("tools", "cmd", "vpic-manual-to-markdown"), "go", "run", ".", path).Run()
	})
}

// KBAWMICollate collates the KBA pages into a single CSV file.
func KBAWMICollate() error {
	return cmd(
		root("tools", "cmd", "kba-collate"),
		"go", "run", ".",
		"-d", root("data", "kba", "wmi"),
		"-o", root("data", "kba", "wmi.csv"),
	).Run()
}

// KBAExcel downloads the KBA Excel files.
func KBAExcel() error {
	const trucksUrl = "https://www.kba.de/SharedDocs/Downloads/DE/Statistik/Fahrzeuge/FZ2/fz2_2024.xlsx?__blob=publicationFile&v=2"
	const trucksFileName = "fz2_2024.xlsx"
	slog.Info("downloading KBA trucks Excel file", "trucksUrl", trucksUrl)
	if err := cmd(root(), "curl", "--create-dirs", "-L", "-o", root("data", "kba", trucksFileName), trucksUrl).Run(); err != nil {
		return err
	}
	const carsURL = "https://www.kba.de/SharedDocs/Downloads/DE/Statistik/Fahrzeuge/FZ17/fz17_2024.xlsx?__blob=publicationFile&v=2"
	const carsFileName = "fz17_2024.xlsx"
	slog.Info("downloading KBA cars Excel file", "carsUrl", carsURL)
	if err := cmd(root(), "curl", "--create-dirs", "-L", "-o", root("data", "kba", carsFileName), carsURL).Run(); err != nil {
		return err
	}
	const manufacturersUrl = "https://www.kba.de/SharedDocs/Downloads/DE/Statistik/Fahrzeuge/FZ6/fz6_2024.xlsx?__blob=publicationFile&v=2"
	const manufacturersFileName = "fz6_2024.xlsx"
	slog.Info("downloading KBA manufacturers Excel file", "manufacturersUrl", manufacturersUrl)
	if err := cmd(root(), "curl", "--create-dirs", "-L", "-o", root("data", "kba", manufacturersFileName), manufacturersUrl).Run(); err != nil {
		return err
	}
	return nil
}

// ACEA downloads the ACEA data.
func ACEA() error {
	const url = "https://www.acea.auto/files/2021_by-manuf-and-type_EUEFTAUK.xlsx"
	const fileName = "2021_by-manuf-and-type_EUEFTAUK.xlsx"
	slog.Info("downloading ACEA Excel file", "url", url, "fileName", fileName)
	if err := cmd(root(), "curl", "--create-dirs", "-L", "-o", root("data", "acea", fileName), url).Run(); err != nil {
		return err
	}
	if err := os.RemoveAll(root("data", "acea", "xsls")); err != nil {
		return err
	}
	if err := os.MkdirAll(root("data", "acea", "xlsx"), 0o700); err != nil {
		return err
	}
	if err := cmd(root("data", "acea"), "ssconvert", "-S", fileName, root("data", "acea", "xlsx", fileName+".csv")).Run(); err != nil {
		return err
	}
	const registrationsUrl = "https://www.acea.auto/files/Press_release_car_registrations_September_2025.pdf"
	const registrationsFileName = "Press_release_car_registrations_September_2025.pdf"
	slog.Info("downloading ACEA registrations PDF file", "registrationsUrl", registrationsUrl, "fileName", registrationsFileName)
	if err := os.RemoveAll(root("data", "acea", "pdf")); err != nil {
		return err
	}
	if err := os.MkdirAll(root("data", "acea", "pdf"), 0o700); err != nil {
		return err
	}
	if err := cmd(root(), "curl", "--create-dirs", "-L", "-o", root("data", "acea", registrationsFileName), registrationsUrl).Run(); err != nil {
		return err
	}
	if err := cmd(root("data", "acea"), "pdftoppm", "-png", registrationsFileName, root("data", "acea", "pdf", registrationsFileName)).Run(); err != nil {
		return err
	}
	return nil
}

// KBAExcelToCSV converts the KBA Excel files to CSV.
func KBAExcelToCSV() error {
	if err := os.RemoveAll(root("data", "kba", "xlsx")); err != nil {
		return err
	}
	if err := os.MkdirAll(root("data", "kba", "xlsx"), 0o700); err != nil {
		return err
	}
	return filepath.WalkDir(root("data", "kba"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".xlsx" {
			return nil
		}
		return cmd(root("data", "kba"), "ssconvert", "-S", path, filepath.Join("xlsx", strings.TrimSuffix(d.Name(), filepath.Ext(path))+".csv")).Run()
	})
}

// KBACountVehicles counts the vehicles in the KBA data.
func KBACountVehicles() error {
	return cmd(
		root("tools", "cmd", "kba-count-vehicles"),
		"go", "run", ".",
		"-i", root("data", "kba", "xlsx", "fz6_2024.csv.3"),
		"-o", root("data", "kba", "vehicle-count.csv"),
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

// WikipediaToCSV converts the Wikipedia WMI data to CSV.
func WikipediaToCSV() error {
	return cmd(
		root("tools", "cmd", "wikipedia-to-csv"),
		"go", "run", ".",
		"-i", root("data", "wikipedia", "wmi.md"),
		"-o", root("data", "wikipedia", "wmi.csv"),
	).Run()
}

// WMI creates the WMI index.
func WMI() error {
	return cmd(
		root("tools", "cmd", "wmi-index"),
		"go", "run", ".",
		"-vpic", root("data", "vpic", "wmi.csv"),
		"-kba", root("data", "kba", "wmi.csv"),
		"-wikipedia", root("data", "wikipedia", "wmi.csv"),
		"-wikibooks", root("data", "wikibooks", "wmi.csv"),
		"-vpic-wmi-make", root("data", "vpic", "wmi-make.csv"),
		"-o", root("data", "wmi.csv"),
	).Run()
}

// WMIGenGo generates the wmi_data.go file from the WMI index.
func WMIGenGo() error {
	return cmd(
		root("tools", "cmd", "wmi-gen-go"),
		"go", "run", ".",
		"-csv", root("data", "wmi.csv"),
		"-output", root("wmi_data.go"),
		"-package", "vin",
	).Run()
}

// MercedeseVINCodingSummary downloads the Mercedes-Benz VIN coding summary.
func MercedesVINCodingSummary() error {
	const url = "https://vpic.nhtsa.dot.gov/mid/home/displayfile/4371bcbc-1a7f-4bc3-90bc-b1e560fff309"
	const fileName = "mercedes-vin-coding-summary.pdf"
	slog.Info("downloading Mercedes-Benz VIN coding summary", "url", url, "fileName", fileName)
	if err := cmd(root(), "curl", "--create-dirs", "-L", "-o", root("docs", "mercedes", "vin-coding-summary", fileName), url).Run(); err != nil {
		return err
	}
	pagesPath := root("docs", "mercedes", "vin-coding-summary", "pages")
	if err := os.RemoveAll(pagesPath); err != nil {
		return err
	}
	if err := os.MkdirAll(pagesPath, 0o700); err != nil {
		return err
	}
	return cmd(root("docs", "mercedes", "vin-coding-summary"), "pdftoppm", "-png", fileName, filepath.Join(pagesPath, "vin-coding-summary")).Run()
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
