#!/bin/bash
set -euo pipefail

# Configuration from data/vpic/run-mssql.sh
DB_PASSWORD="Admin@123"
DB_CONTAINER_NAME="mssql_local"
DB_NAME="vpic"

# --- Script Usage ---
usage() {
    echo "Usage: $0 <output_csv_file>"
    echo "  (SQL query to be provided via standard input)"
    echo ""
    echo "Executes an SQL query (from stdin) against the VPIC database running in the '$DB_CONTAINER_NAME' Docker container."
    echo "The result is formatted as CSV."
    echo ""
    echo "Prerequisites:"
    echo "  - Docker must be installed and running."
    echo "  - The database container must be running (use 'data/vpic/run-mssql.sh' to start it)."
    echo ""
    echo "Arguments:"
    echo "  output_csv_file  : Optional. Path to write the CSV output to. If not provided, output is written to stdout."
    exit 1
}

# --- Argument validation ---
if [ "$#" -gt 1 ]; then
    usage
fi

OUTPUT_CSV_FILE="${1:-}" # Use default empty string if not set

# Check if standard input is empty, assuming a query is expected
if [ -t 0 ] && [ -z "$(head -n 1 /dev/stdin)" ]; then # Check if stdin is a terminal and first line is empty
    echo "Error: No SQL query provided via standard input." >&2
    usage
fi

# Prepend SET NOCOUNT ON to prevent the "(X rows affected)" message at the end of the output.
# Read query from stdin
QUERY="SET NOCOUNT ON; $(cat -)"

# --- Check if container is running ---
if ! docker ps -q -f name=^/${DB_CONTAINER_NAME}$ > /dev/null; then
    echo "Error: The Docker container '$DB_CONTAINER_NAME' is not running." >&2
    echo "Please start the database first using the 'data/vpic/run-mssql.sh' script." >&2
    exit 1
fi

# --- Execute Query ---
# The output from sqlcmd with headers has the following format:
# 1. Header row
# 2. A separator line of dashes (e.g., "-----------,---------")
# 3. Data rows
# We pipe the output to `sed '2d'` to remove the separator line.
execute_query() {
    # Using an array for command arguments is safer than using eval or a simple string with spaces.
    local -a cmd_args=(
        "exec"
        "$DB_CONTAINER_NAME"
        "/opt/mssql-tools18/bin/sqlcmd"
        "-S" "localhost"
        "-U" "sa"
        "-P" "$DB_PASSWORD"
        "-d" "$DB_NAME"
        "-Q" "$QUERY"
        "-s" ","  # Use comma as a separator for CSV
        "-W"      # Remove trailing whitespace from columns
        "-C"      # Trust the server certificate
    )

    docker "${cmd_args[@]}" | sed '2d'
}

if [ -n "$OUTPUT_CSV_FILE" ]; then
    echo "Executing query and writing CSV output to '$OUTPUT_CSV_FILE'..." >&2
    execute_query > "$OUTPUT_CSV_FILE"
    echo "Done." >&2
else
    # No output file specified, write to stdout
    execute_query
fi
