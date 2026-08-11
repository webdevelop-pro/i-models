# `filers`

This package defines `FilerFiler` for the hierarchical
`filer_filers_ltree` table. `FolderName` and `Type` constants classify standard
folders and file/folder entries.

`esign-api` is the verified direct consumer. Its API-integration layer uses the
shared filer row and constants when resolving documents needed for signature
workflows. The actual file hierarchy, access control, and storage operations
remain owned by `filer-api`/the calling service.

```go
filer, err := models.RetrieveOne[filers.FilerFiler](
    ctx, repo, map[string]any{"id": filerID},
)
```

See the [root guide](../README.md) for the generic model contract.
