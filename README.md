# Cinema Distributor Permissions Tool

This Go application manages distributor permissions for the cinema distribution business, allowing you to add distributors with included and excluded regions, and check if they have the right to distribute in a specified region.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Usage](#usage)
  - [Add Distributor](#add-distributor)
  - [Check Distributor Permissions](#check-distributor-permissions)
- [Running the Application](#running-the-application)
- [Testing](#testing)
- [License](#license)

## Prerequisites

- Go installed on your machine (version 1.18 or later recommended).
- CSV file containing the regions data (referenced in the code).
- Basic knowledge of how to run Go programs.

## Getting Started

1. **Clone the Repository** (if applicable):
   ```bash
   git clone <repository_url>
   cd go-challenge-2016
   ```

2. **Ensure the CSV File is Available**:
   Download the CSV file from the provided URL and save it to your project directory.
   ```bash
   curl -o convertcsv.csv https://raw.githubusercontent.com/RealImage/challenge2016/refs/heads/master/cities.csv
   ```

## Usage

### Add Distributor

To add a distributor, use the following command:

```bash
go run main.go -action=add -name=<DISTRIBUTOR_NAME> -include=<REGION_CODES> -exclude=<REGION_CODES>
```

- `<DISTRIBUTOR_NAME>`: Name of the distributor (e.g., `DISTRIBUTOR1`).
- `<REGION_CODES>`: Comma-separated codes for included/excluded regions (e.g., `IN,US` for including India and the United States).

**Example**:
```bash
go run main.go -action=add -name=DISTRIBUTOR1 -include=IN,US -exclude=KA-IN,CHN-TN-IN
```

### Check Distributor Permissions

To check if a distributor can distribute in a specific region, use:

```bash
go run main.go -action=check -name=<DISTRIBUTOR_NAME> -region=<REGION_CODE>
```

- `<REGION_CODE>`: Code for the region to check (e.g., `CHICAGO-IL-US`).

**Example**:
```bash
go run main.go -action=check -name=DISTRIBUTOR1 -region=CHICAGO-IL-US
```

## Running the Application

1. Navigate to the project directory where `main.go` is located.
2. Use the commands mentioned in the **Usage** section to add distributors or check permissions.

## Testing

### Example Workflow

1. **Add a Distributor**:
   ```bash
   go run main.go -action=add -name=DISTRIBUTOR1 -include=IN,US -exclude=KA-IN,CHN-TN-IN
   ```

2. **Add Distributor 2 (DISTRIBUTOR2) (under DISTRIBUTOR1)**:
  ```bash
   go run main.go -action=add -name=DISTRIBUTOR2 -include=IN -exclude=TN-IN -parent=DISTRIBUTOR1
  ```
3. **Add Distributor 3 (DISTRIBUTOR3) (under DISTRIBUTOR2)**:
   ```bash
   go run main.go -action=add -name=DISTRIBUTOR3 -include=HBL-KA-IN -parent=DISTRIBUTOR2
   ```
4. **Check Permissions**:
   ```bash
   go run main.go -action=check -name=DISTRIBUTOR1 -region=CHICAGO-IL-US
   ```

5. **Verify Output**: The expected output should confirm if the distributor has permission or not.

### Debugging
If you encounter issues:
- Check if the distributor was added successfully by printing the current distributors.
- Ensure that the region codes used in the check command are correct and exist in the CSV data.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
