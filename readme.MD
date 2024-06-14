# OneDriveFileAvailabilityChecker

Utility to check for the availability status of a file in a OneDrive folder  
Only works on Windows

## Compile

```
go build
```

## Usage

```
OneDriveFileAvailabilityChecker.exe PATH
```

## Output

only as a print to the stdout for now

```
status "Available offline", "Always available on this device", "Available on this device", "Available when online":
  "available: STATUS"

status "Sync pending":
	"pending: STATUS"

status "Error":
	"not_available: STATUS"

any other status:
	"unknown_status: STATUS"
```

| STATUS is the actual text of status from the file
