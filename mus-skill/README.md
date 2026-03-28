# mus-skill Example

This example demonstrates how to use the [mus-skill](https://github.com/mus-format/mus-skill-go) 
agent skill to generate a MUS serializer and tests.

## Setup

The `mus-skill` skill was installed into `.agents/skills` directory using 
[skills](https://github.com/vercel-labs/skills) with the following command:

```bash
npx skills add github.com/mus-format/mus-skill-go
```

## The Prompt

The following prompt was used to generate the `mus.ai.gen.go` and 
`mus.ai.gen_test.go` files:

```text
Generate MUS serializers for the types found in the mus-skill/main.go file.
```

## How it Works

1. **Define Types**: Define Go types with MUS-specific hints in a file (in this 
   case, `main.go`).
2. **AI Action**: Prompt the AI agent.
3. **Verification**: Check the generated tests and run them.

## Running the Example

```bash
go test -v .
go run .
```
