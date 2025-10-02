# Go Utils

Este repositório contém uma coleção de pacotes Go com funções utilitárias para manipulação de arrays, cache, arquivos, pastas e JSON.

## Instalação

Para usar os pacotes em seu projeto, você pode usar o `go get`:

```bash
go get github.com/pedro-makoski/go-utils
```

## Pacotes

### array

O pacote `array` fornece funções para manipulação de slices.

**Funções:**

- **`Contains[T comparable](arr []T, element T) bool`**: Verifica se um slice contém um determinado elemento.
- **`SameListExcept[T any](list []T, indexToNotInclude int) []T`**: Retorna um novo slice com todos os elementos do slice original, exceto o elemento no índice especificado.

### cache

O pacote `cache` oferece funcionalidades para cache de resultados de funções.

**Funções:**

- **`AssignWithCacheAndReturnIfExists(comandUnique string, params []any, assignFn func(retornos []any)) bool`**: Verifica se o resultado de uma função com determinados parâmetros está em cache e, em caso afirmativo, atribui os valores de retorno em cache à função `assignFn`.
- **`Register(configs RegisterConfigs)`**: Registra o resultado de uma função no cache.

### file

O pacote `file` contém utilitários para manipulação de arquivos.

**Funções:**

- **`CopyFile(src, dst string) goerror.Error`**: Copia um arquivo de um local de origem para um local de destino.
- **`GetAbs(pathRelative string) (string, goerror.Error)`**: Retorna o caminho absoluto de um caminho relativo.
- **`IsFile(path string) bool`**: Verifica se um caminho corresponde a um arquivo.
- **`ReadFile(path string) (string, goerror.Error)`**: Lê o conteúdo de um arquivo e o retorna como uma string.
- **`CacheReadFile(path string) (string, goerror.Error)`**: Lê o conteúdo de um arquivo, utilizando o cache para evitar leituras repetidas.
- **`ReadFileInfo(path string) (*FileInfo, goerror.Error)`**: Lê as informações de um arquivo, incluindo nome, tamanho, modo, data de modificação e conteúdo.
- **`RemoveFile(path string) goerror.Error`**: Remove um arquivo.
- **`RenameFile(oldPath, newPath string) goerror.Error`**: Renomeia ou move um arquivo.
- **`DoesThisFileExists(path string) bool`**: Verifica se um arquivo existe.
- **`WriteFile(path string, newContent string, permission os.FileMode) goerror.Error`**: Escreve conteúdo em um arquivo.
- **`AppendInFile(path string, newLine string) goerror.Error`**: Adiciona uma nova linha ao final de um arquivo.

### folder

O pacote `folder` fornece utilitários para manipulação de pastas.

**Funções Não Recursivas**

Estas funções operam apenas no primeiro nível do diretório especificado.

- **`GetAllPaths(root string) ([]string, goerror.Error)`**: Retorna todos os caminhos (arquivos e pastas) dentro do primeiro nível de uma pasta.
- **`GetAllPathsWithExt(ext string, root string) ([]string, goerror.Error)`**: Retorna todos os caminhos de arquivos com uma determinada extensão e todas as pastas dentro do primeiro nível.
- **`GetAllFolders(root string) ([]string, goerror.Error)`**: Retorna todos os caminhos de pastas dentro do primeiro nível de uma pasta.
- **`GetAllFiles(root string) ([]string, goerror.Error)`**: Retorna todos os caminhos de arquivos dentro do primeiro nível de uma pasta.
- **`CachedGetAllPaths(path string) ([]string, goerror.Error)`**: Versão com cache de `GetAllPaths`.
- **`CachedGetAllPathsWithExt(ext string, path string) ([]string, goerror.Error)`**: Versão com cache de `GetAllPathsWithExt`.
- **`CachedGetAllFolders(root string) ([]string, goerror.Error)`**: Versão com cache de `GetAllFolders`.
- **`CachedGetAllFiles(root string) ([]string, goerror.Error)`**: Versão com cache de `GetAllFiles`.

**Funções Recursivas**

Estas funções operam em todo o aninhamento de diretórios a partir do caminho raiz.

- **`CopyFolder(srcRoot, dstRoot string) goerror.Error`**: Copia uma pasta e todo o seu conteúdo (recursivamente) para um novo local.
- **`GetAllPathsRecursive(root string) ([]string, goerror.Error)`**: Retorna todos os caminhos (arquivos e pastas) recursivamente.
- **`GetAllPathsWithExtRecursive(ext string, root string) ([]string, goerror.Error)`**: Retorna todos os caminhos de arquivos com uma determinada extensão recursivamente.
- **`GetAllFoldersRecursive(root string) ([]string, goerror.Error)`**: Retorna todos os caminhos de pastas recursivamente.
- **`GetAllFilesRecursive(root string) ([]string, goerror.Error)`**: Retorna todos os caminhos de arquivos recursivamente.
- **`CachedGetAllPathsRecursive(path string) ([]string, goerror.Error)`**: Versão com cache de `GetAllPathsRecursive`.
- **`CachedGetAllPathsWithExtRecursive(ext string, path string) ([]string, goerror.Error)`**: Versão com cache de `GetAllPathsWithExtRecursive`.
- **`CachedGetAllFoldersRecursive(root string) ([]string, goerror.Error)`**: Versão com cache de `GetAllFoldersRecursive`.
- **`CachedGetAllFilesRecursive(root string) ([]string, goerror.Error)`**: Versão com cache de `GetAllFilesRecursive`.

**Outras Funções**

- **`CreateFolder(path string) goerror.Error`**: Cria uma nova pasta (e pais, se necessário).
- **`DoesThisFolderExists(path string) bool`**: Verifica se uma pasta existe.
- **`IsFolder(path string) bool`**: Verifica se um caminho corresponde a uma pasta.
- **`RemoveFolder(path string) goerror.Error`**: Remove uma pasta e todo o seu conteúdo.
- **`RenameFolder(oldPath, newPath string) goerror.Error`**: Renomeia ou move uma pasta.

### folderfile

O pacote `folderfile` oferece funções para interagir com arquivos e pastas de forma intercambiável.

**Funções:**

- **`Copy(srcPath, dstPath string) goerror.Error`**: Copia um arquivo ou pasta. A cópia de pastas é recursiva.
- **`Remove(path string) goerror.Error`**: Remove um arquivo ou pasta.
- **`RenameOrMove(oldPath, newPath string) goerror.Error`**: Renomeia ou move um arquivo ou pasta.

### json

O pacote `json` contém utilitários para trabalhar com JSON.

**Funções:**

- **`GetJsonFile[format any](path string, obj *format) goerror.Error`**: Lê um arquivo JSON e o converte em um objeto Go.
- **`Convert[format any](content string, obj *format) goerror.Error`**: Converte uma string JSON em um objeto Go.
- **`WriteJSONFile[Format any](path string, object Format, prefix string, indent string, permissionFile os.FileMode) goerror.Error`**: Escreve um objeto Go em um arquivo JSON.
- **`GetStringOnJSONFile[Format any](object Format, prefix string, indent string) (string, goerror.Error)`**: Converte um objeto Go em uma string JSON.

### path

O pacote `path` fornece uma estrutura `Path` para facilitar a manipulação de caminhos de arquivo.

**Estrutura:**

- **`Path`**: Representa um caminho de arquivo ou pasta.

**Funções:**

- **`GenPath(path string) Path`**: Cria uma nova instância de `Path`.
- **`Path.IsFile() bool`**: Verifica se o caminho é um arquivo.
- **`Path.IsFolder() bool`**: Verifica se o caminho é uma pasta.
- **`Path.ContainsFile(fileName string) (bool, goerror.Error)`**: Verifica se a pasta contém um arquivo com o nome especificado (não recursivo).
- **`Path.ContainsFileName(fileName string) (bool, goerror.Error)`**: Verifica se a pasta contém um arquivo com o nome de arquivo especificado, ignorando o caminho (não recursivo).
- **`Path.ContainsFolder(folderName string) (bool, goerror.Error)`**: Verifica se a pasta contém uma subpasta com o nome especificado (não recursivo).
- **`Path.Copy(destinoPath string) goerror.Error`**: Copia o arquivo ou pasta para um novo destino (recursivo para pastas).
- **`Path.ThisPathExists() bool`**: Verifica se o caminho existe.
- **`Path.GetNameOfThis() string`**: Retorna o nome do arquivo ou pasta.
- **`Path.GetParentPath() Path`**: Retorna o caminho do diretório pai.
- **`Path.HasParentPath() bool`**: Verifica se o caminho tem um diretório pai.
- **`Path.Read() (string, goerror.Error)`**: Lê o conteúdo do arquivo.
- **`Path.GetInsidePaths() ([]string, goerror.Error)`**: Retorna os caminhos dentro do primeiro nível da pasta.
- **`Path.CachedGetInsidePaths() ([]string, goerror.Error)`**: Versão com cache de `GetInsidePaths`.
- **`Path.CachedGetInsidePathsWithExt(ext string) ([]string, goerror.Error)`**: Retorna os caminhos com uma determinada extensão dentro do primeiro nível da pasta, utilizando cache.
- **`Path.CachedGetAllFolders() ([]string, goerror.Error)`**: Retorna todas as pastas dentro do primeiro nível da pasta, utilizando cache.
- **`Path.CachedGetAllFiles() ([]string, goerror.Error)`**: Retorna todos os arquivos dentro do primeiro nível da pasta, utilizando cache.
- **`Path.Remove() goerror.Error`**: Remove o arquivo ou pasta.
- **`Path.RenameOrMove(newPath string) goerror.Error`**: Renomeia ou move o arquivo ou pasta.
- **`Path.WriteInFile(content string, permission file.FilePermissions) goerror.Error`**: Escreve conteúdo no arquivo.