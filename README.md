# confusio

Generate visually similar Unicode variants (homoglyphs) for keywords to test input sanitization, spam filters, and phishing detection.

**Key features:**
- Generates 30+ Unicode confusables per keyword
- Zero dependencies — uses only Go standard library
- Cross-platform (Linux, macOS, Windows)
- JSON output for tool integration
- Entropy sorting to rank variants by complexity

## Quick Start

### Build

Requirements: Go 1.21+

```bash
make build
```

Produces `confusio` binary in the current directory.

For cross-platform release binaries:

```bash
make all
```

## Usage

```
confusio [options] <keyword>
```

**Options:**
- `-j` — Output as JSON array
- `-e` — Sort variants by entropy (increasing complexity)

## Examples

**Basic usage:**
```bash
./confusio stripe
```

**JSON output:**
```bash
./confusio -j paypal
```

**Entropy sorted:**
```bash
./confusio -e google
```

## How It Works

Confusio generates homoglyphs by applying Unicode normalization and finding visually similar characters. This is useful for:
- Testing input validation and sanitization
- Evaluating spam filter robustness
- Assessing phishing detection mechanisms
