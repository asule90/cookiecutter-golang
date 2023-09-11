#!/bin/bash
#

# Code coverage generation
COVERAGE_DIR="${COVERAGE_DIR:-coverage}"
PKG_LIST=$(go list ./...)

# Create the coverage files directory
mkdir -p "$COVERAGE_DIR";

# Create a coverage file for each package
for package in ${PKG_LIST}; do
    go test -covermode=count -coverprofile "${COVERAGE_DIR}/${package##*/}.cov" "$package" ;
done ;

# Merge the coverage profile files
echo 'mode: count' > "${COVERAGE_DIR}"/coverage.cov ;
tail -q -n +2 "${COVERAGE_DIR}"/*.cov >> "${COVERAGE_DIR}"/coverage.cov ;

# Display the global code coverage
# go tool cover -func="${COVERAGE_DIR}"/coverage.cov ;

# If needed, generate HTML report
if [ "$1" == "html" ]; then
    go tool cover -html="${COVERAGE_DIR}"/coverage.cov -o coverage.html ;
fi

# Remove the coverage files directory
rm -rf "$COVERAGE_DIR";

# Break test if fail
for s in $(go list ./...); do if ! go test -v -p 1 $s; then exit 1; fi; done