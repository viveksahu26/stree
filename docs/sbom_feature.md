# Representing SBOM Dependencies as a Directory Tree or File Tree

## Why(Motivation) ??

An SBOM (Software Bill of Materials) provides detailed information about the components and dependencies of a software system. Often, this document contains numerous dependencies and sub-dependencies, making it time-consuming and tedious to manually navigate through them. For users primarily interested in visualizing these relationships, a directory or file tree structure—similar to the output of the tree command—offers a more intuitive and visual approach, allowing quick access to dependencies without the need to read through the entire document.

## How to Implement ??

- Provide SBOM via arguments.
  - Using Command: `$ stree sbom <file>`
- Get a SBOM.
  - validate the file path
  - Read the SBOM file.
- Parse the SBOM.
- Convert the parse SBOM information into JSON format.
- Retrieve the particular field from JSON structure.
- Now create a directory tree or file tree structure.
