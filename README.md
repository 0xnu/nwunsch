## nwunsch

[![Release](https://img.shields.io/github/release/0xnu/nwunsch.svg)](https://github.com/0xnu/nwunsch/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/0xnu/nwunsch)](https://goreportcard.com/report/github.com/0xnu/nwunsch)
[![Go Reference](https://pkg.go.dev/badge/github.com/0xnu/nwunsch.svg)](https://pkg.go.dev/github.com/0xnu/nwunsch)
[![License](https://img.shields.io/github/license/0xnu/nwunsch)](/LICENSE)

[Needleman–Wunsch](https://en.wikipedia.org/wiki/Needleman%E2%80%93Wunsch_algorithm) algorithm implementation in Go.

### Features

- **Sequence Comparison**: The package enables the comparison of two sequences (seq1 and seq2) by considering all possible alignments and choosing the best one.

- **Scoring System**: The scoring system is customisable with match, mismatch, and gap parameters, allowing users to define the reward for matched characters and penalties for mismatches and gaps.

- **Matrix Generation and Population**: It creates a dynamic programming matrix of sequence lengths and populates it accordingly based on optimal alignment values.

- **Traceback Functionality**: It includes a traceback procedure which identifies the optimal path through the matrix, producing the final alignment of the input sequences.

- **Efficient Evaluation**: It utilises a max function to determine the maximum score between match, mismatch, and gap; this feature guarantees the efficiency of the sequence alignment process.

### Installation

Install the `nwunsch` package in your `Go` project by installing it with the following command:

```shell
go get github.com/0xnu/nwunsch
```

### Usage

Check out the [examples](./examples/) directory for how to use the `nwunsch` package.

### Testing

To run the tests for the `nwunsch` package, use the following command:

```sh
go test
```

### License

This project is licensed under the [MIT License](./LICENSE).

### Copyright

(c) 2024 [Finbarrs Oketunji](https://finbarrs.eu).