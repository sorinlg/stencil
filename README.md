# Stencil
*A templating tool*

## Binary
Release a binary with the following command

```bash
go build -o stencil-`git rev-parse --short HEAD`
```

## IO Spec
- **input**
    - template file
    - environment variables
- **output**
    - artifact (file with replaced token variables)
