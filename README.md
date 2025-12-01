# Confusio

**Confusio** is a Unicode-Normalisation "Confusables" Payload CLI tool written in Go. It generates visually similar Unicode variants (homoglyphs) for a given keyword. This is useful for testing input sanitization, spam filters, and phishing detection mechanisms.

## Features

-   **Zero Dependencies**: Built using only the Go standard library.
-   **Cross-Platform**: Runs on Linux, macOS, and Windows.
-   **Rich Output**: Generates ≥ 30 variants per run.
-   **JSON Support**: Optional JSON output for easy integration with other tools.
-   **Entropy Sorting**: Sort variants by entropy (complexity/disorder).

## Installation

### From Source

Requirements: Go 1.21+

```bash
git clone https://github.com/yourusername/confusio.git
cd confusio
make build
```

The binary will be created in the current directory.

## Usage

```bash
./confusio [options] <keyword>
```

### Options

-   `-j`: Output as a JSON array.
-   `-e`: Sort output by entropy (increasing).

### Examples

**Basic Usage:**

```bash
$ ./confusio stripe
ѕtripe
st​ripe
stṛipe
...
```

**JSON Output:**

```bash
$ ./confusio -j paypal
[
  "рaypal",
  "pａypal",
  ...
]
```

**Entropy Sorted:**

```bash
$ ./confusio -e google
```

## Development

To build for all supported platforms (Linux, macOS, Windows):

```bash
make all
```

## License

MIT
