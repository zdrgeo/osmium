package view

import "path/filepath"

func viewPath(basePath, analysisName, name string) string {
	return filepath.Join(basePath, "osmium", "analysis", analysisName, "view", name)
}
