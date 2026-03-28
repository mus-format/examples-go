# DTM Constants

DTM (Data Type Metadata) is a unique identifier for each type that implements 
an interface.

## User-defined DTMs

The user can define DTMs manually:

```go
const (
  FooDTM com.DTM = 1
  BarDTM com.DTM = 2
)
```

## Generated DTMs

If the user doesn't provide all the necessary DTMs, generate them automatically 
starting from 1:

```go
const (
  FooDTM com.DTM = iota + 1
  BarDTM
)
```

## DTM Conflicts

Resolving DTM conflicts is the user's responsibility.
