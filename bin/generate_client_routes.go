package main

import (
	"fmt"
	"os"
	"strings"

	"puzzlr.gg/src/server/routes"
)

const (
	generatedRouterPath = "src/client/routes.generated.ts"
)

func main() {
	type processedRoute struct {
		name   string
		path   string
		params []string
	}

	var processed []processedRoute
	for _, route := range routes.PageRoutes {
		path, params, err := chiPathToReactRouter(route.Path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to convert route %q: %v\n", route.Path, err)
			os.Exit(1)
		}
		processed = append(processed, processedRoute{name: string(route.Name), path: path, params: params})
	}

	var sb strings.Builder
	sb.WriteString("// AUTO-GENERATED — do not edit manually\n")
	sb.WriteString("// Run `go run bin/generate_client_routes.go` to regenerate\n\n")

	sb.WriteString("export const routes = {\n")
	for _, r := range processed {
		if len(r.params) == 0 {
			sb.WriteString(fmt.Sprintf("  %s: {\n", r.name))
			sb.WriteString(fmt.Sprintf("    path: %q,\n", r.path))
			sb.WriteString(fmt.Sprintf("    build: () => %q,\n", r.path))
			sb.WriteString("  },\n")
		} else {
			paramList := make([]string, len(r.params))
			for i, p := range r.params {
				paramList[i] = fmt.Sprintf("%s: string", p)
			}

			tmpl := r.path
			for _, p := range r.params {
				tmpl = strings.Replace(tmpl, ":"+p, "${"+p+"}", 1)
			}

			sb.WriteString(fmt.Sprintf("  %s: {\n", r.name))
			sb.WriteString(fmt.Sprintf("    path: %q,\n", r.path))
			sb.WriteString(fmt.Sprintf("    build: (%s) => `%s`,\n", strings.Join(paramList, ", "), tmpl))
			sb.WriteString("  },\n")
		}
	}
	sb.WriteString("} as const;\n")

	err := os.WriteFile(generatedRouterPath, []byte(sb.String()), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write generated routes: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Generated %s", generatedRouterPath)
}

// chiPathToReactRouter converts Chi-style path params ({id}) to React Router-style (:id).
func chiPathToReactRouter(path string) (string, []string, error) {
	if path == "" || path[0] != '/' {
		return "", nil, fmt.Errorf("path %q must start with '/'", path)
	}
	if strings.Contains(path, "*") {
		return "", nil, fmt.Errorf("wildcard paths are not supported: %q", path)
	}

	var params []string
	result := strings.Builder{}
	for i := 0; i < len(path); i++ {
		if path[i] == '{' {
			end := strings.IndexByte(path[i+1:], '}')
			if end == -1 {
				return "", nil, fmt.Errorf("path %q contains an unterminated placeholder", path)
			}
			end += i + 1

			placeholder := path[i+1 : end]
			parts := strings.SplitN(placeholder, ":", 2)
			name := parts[0]
			if name == "" {
				return "", nil, fmt.Errorf("path %q contains an anonymous placeholder", path)
			}
			if len(parts) == 2 {
				return "", nil, fmt.Errorf("path %q uses a regex-constrained placeholder that React Router cannot express", path)
			}

			params = append(params, name)
			result.WriteByte(':')
			result.WriteString(name)
			i = end
		} else {
			result.WriteByte(path[i])
		}
	}
	return result.String(), params, nil
}
