package shared

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListSubdirectories(baseDir string) ([]string, error) {
	var subdirs []string

	// Lê as entradas do diretório especificado
	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o diretório %q: %w", baseDir, err)
	}

	// Filtra apenas os diretórios e adiciona seus caminhos completos à lista
	for _, entry := range entries {
		if entry.IsDir() {
			subdirs = append(subdirs, filepath.Join(baseDir, entry.Name()))
		}
	}

	return subdirs, nil
}

// CombineSQLFilesContent percorre os subdiretórios do diretório base e 
// concatena o conteúdo de todos os arquivos SQL encontrados.
func CombineSQLFilesContent() (string, error) {
	baseDir := "internal/infrastructure/database/query/module"

	// Obtém a lista de subdiretórios dentro de baseDir
	directories, err := ListSubdirectories(baseDir)
	if err != nil {
		return "", fmt.Errorf("erro ao listar subdiretórios: %w", err)
	}

	var content string

	// Percorre cada subdiretório encontrado
	for _, dir := range directories {
		// Lê as entradas dentro do diretório atual
		entries, err := os.ReadDir(dir)
		if err != nil {
			return "", fmt.Errorf("erro ao ler o diretório %q: %w", dir, err)
		}

		// Itera sobre os arquivos do diretório
		for _, entry := range entries {
			if entry.IsDir() || filepath.Ext(entry.Name()) != ".sql" {
				continue // Ignora subdiretórios e arquivos que não sejam SQL
			}

			// Obtém o caminho completo do arquivo
			filePath := filepath.Join(dir, entry.Name())

			// Lê o conteúdo do arquivo SQL
			data, err := os.ReadFile(filePath)
			if err != nil {
				return "", fmt.Errorf("erro ao ler o arquivo %q: %w", filePath, err)
			}

			// Concatena o conteúdo do arquivo na variável content
			content += string(data) + "\n\n"
		}
	}

	fmt.Println(content)

	return content, nil
}



