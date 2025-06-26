# Bookworm Book Digest CLI Tool

## Overview

This project provides a command-line interface (CLI) tool designed for bookworms to manage and analyze their book collections. Inspired by Fadi and Peggy, who want to find common books and discover new reads, this tool helps users digest their book inventories. It reinforces concepts of CLI development, file I/O, JSON parsing, and Go's dynamic data structures like maps and slices.

## Features

- **Input Handling**: Reads a list of bookworms and their respective book collections from a JSON file.
- **Common Book Discovery**: Identifies and lists books that appear on more than one bookworm's bookshelf.
- **Standard Output**: Prints the common books to the console for easy review.
- **Book Recommendation (Bonus)**: Recommends new books for each bookworm based on their shared reading interests with other bookworms. This feature leverages Go's `map` and `slice` data structures.

## Limitations

- **Single Author Assumption**: For simplicity, each book is assumed to have only one author.
- **Input File Size**: The input JSON file is assumed not to exceed one megabyte in size.

## How to Use

1. **Prepare your JSON input file**: The tool expects a JSON file containing bookworm data, structured to include each bookworm's name and their list of books. Each book should have a title and a single author.
2. **Run the CLI tool**: Execute the compiled Go program, passing your JSON file as an argument.
3. **View Results**: The common books will be printed to your standard output. If the bonus feature is enabled, book recommendations will also be displayed.

## Usage

```bash
./bookworms-cli <path-to-json-file>
```

## Input Format

The tool expects a JSON file with the following structure:

```json
{
  "bookworms": [
    {
      "name": "Fadi",
      "books": [
        {
          "title": "Book Title",
          "author": "Author Name"
        }
      ]
    },
    {
      "name": "Peggy",
      "books": [
        {
          "title": "Another Book",
          "author": "Another Author"
        }
      ]
    }
  ]
}
```

## Output

The tool will display:

1. **Common Books**: Books that appear on multiple bookshelves
2. **Recommendations**: Suggested books for each bookworm based on their reading overlap with others

## Project Requirements (Technical)

- Develop a CLI tool in Go.
- Read and parse a JSON input file using standard Go libraries.
- Implement logic to find common books among different collections.
- Implement logic for book recommendations (bonus).

## Technical Implementation

- Built with Go's standard libraries
- Utilizes Go's map and slice data structures for efficient data processing
- JSON parsing using Go's built-in JSON package
- File reading capabilities using Go's standard I/O libraries

## Installation

1. Clone this repository
2. Build the executable:
   ```bash
   go build -o bookworms-cli
   ```
3. Run with your JSON book collection file

## Example

```bash
./bookworms-cli my-book-collections.json
```

Expected output:
```
Common Books Found:
- "The Great Gatsby" by F. Scott Fitzgerald
- "1984" by George Orwell

Recommendations:
For Fadi: "To Kill a Mockingbird" by Harper Lee
For Peggy: "Pride and Prejudice" by Jane Austen
```

## Contributing

This project is part of a learning exercise focusing on:
- Command-line interface development in Go
- JSON file processing
- Data structure manipulation with maps and slices
- File I/O operations

Feel free to contribute improvements or additional features!

