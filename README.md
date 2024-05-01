# Password Hash Cracker

## Overview
The Password Hash Cracker is a simple yet powerful tool designed to help security professionals test the strength of passwords stored in hash form. It supports `MD5` and `SHA1` hash algorithms, commonly used for storing passwords in databases. By attempting to match provided hashes with a list of potential passwords, this tool can be an essential part of a security audit or penetration test.

## Features
- **Support for MD5 and SHA1 hashes**: Commonly used hash algorithms in password storage.
- **Multithreading**: Leverages Go's concurrency model for faster processing.
- **Safe output**: Sanitizes output to prevent terminal disruption from special characters.

## Installation

### Prerequisites
- Go (1.15 or later recommended)
- Access to a Unix-like terminal (if running on Windows, Cygwin or a similar environment is required)

### Setup
Clone the repository to your local machine:

```bash
git clone https://github.com/SlimJUC/md5-sha1-cracker.git
cd md5-sha1-cracker
```

### Build the binary 

```bash
go build crack.go
```

### Start 

```bash
./crack
```

You will be prompted to enter:

The stored hash of a password.
The file path to a list of potential plaintext passwords (one per line).
The tool will process each password, attempting to generate the hash and compare it with the provided hash. It updates the current testing password on a single line to keep the output clean.

### Example

```bash
Enter the stored password hash: ec77751b745f211b220c1ff228ed7f7e
Enter the file name containing the list of passwords: /path/to/password/list.txt
Testing password: password123
```
If a match is found, the tool will output:

```bash
Password password123 is correct
```

### Limitations

The tool is intended for educational and legal use in security assessments where permission has been granted.
Only supports MD5 and SHA1 hashes. For more complex algorithms or newer hash functions, additional development is needed.
