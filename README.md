# BettyPy - Python Style Guide for Betty

![Build Status](https://img.shields.io/badge/build-passing-brightgreen)
![License](https://img.shields.io/badge/license-MPL%202.0-blue)
![Python](https://img.shields.io/badge/python-3.x-blue)
![Version](https://img.shields.io/badge/version-1.0.0-orange)
![Linux](https://img.shields.io/badge/os-linux-blue)
![macOS](https://img.shields.io/badge/os-macOS-blue)
![Windows](https://img.shields.io/badge/os-windows-yellow)

BettyPy is a Python style guide for the Betty programming language. It enforces best practices and ensures clean, readable, and maintainable code.

---

### Table of Contents
- [Description](#description)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
- [Requirements](#requirements)
- [Contributing](#contributing)
- [Acknowledgments](#acknowledgments)
- [Authors](#authors)
- [License](#license)

---

### Description

BettyPy provides a set of rules and guidelines for writing Python code in the Betty programming language. It helps developers maintain consistent coding standards and improve code quality.

---

### Features
| Feature        | Description                                                                 |
|----------------|-----------------------------------------------------------------------------|
| Empty files    | Detects and warns about empty files in the project.                        |
| Line Length    | Ensures that lines do not exceed 80 characters for better readability.     |
| Indentation    | Enforces 4 spaces per indentation level to maintain consistent formatting. |
| Docstrings     | Checks for missing or improperly formatted docstrings in code.            |
| Comments       | Ensures comments are clear, properly formatted, and start with a space.   |
| Annotations    | Detects TODO, FIXME, and other annotations in the code.                   |
| Imports        | Identifies unused imports and suggests their removal.                     |
| Last Line      | Ensures the file ends with a newline character.                           |
| Function Names | Verifies that function names follow the snake_case naming convention.     |

---

---

### Installation

##### Linux
To install BettyPy on Linux, follow these steps:

```bash
$ git clone https://github.com/ItsZmainDev/BettyPy.git
$ cd BettyPy
$ sudo ./install.sh
```

##### Windows
To install BettyPy on Windows, follow these steps:
```bash
$ git clone https://github.com/ItsZmainDev/BettyPy.git
$ cd BettyPy
$ install_windows.bat
```

---

### Usage
Run BettyPy on your Python file with the following command:

```bash
$ bettypy <path_to_your_python_file>
```

---

### Example
#### Example Python file
```python
# This is an example Python file

def function_example():
	# print("cette ligne est trop longue et dépasse 80 caractères, ce qui est inacceptable dans le style de codage de Betty.")
    print("cette ligne n'est pas indentée correctement et doit être corrigée.")

####test
```

#### Command to run BettyPy
```bash
$ bettypy example.py
```

#### Output of BettyPy
```bash
root@eddc83f1bb46:/home/ubuntu/BettyPy# bettypy tests/testfile.py 
[2025-04-14 22:30:46] [INFO] tests/testfile.py - Starting analysis...
[2025-04-14 22:30:46] [WARNING] tests/testfile.py:4: line too long (125 characters)
[2025-04-14 22:30:46] [WARNING] tests/testfile.py:5: line too long (82 characters)
[2025-04-14 22:30:46] [WARNING] tests/testfile.py:7: comment should start with a space after #
[2025-04-14 22:30:46] [WARNING] tests/testfile.py:7: comment too short (4 characters)
[2025-04-14 22:30:46] [WARNING] tests/testfile.py: file should end with a new line
[2025-04-14 22:30:46] [WARNING] tests/testfile.py:3: Function 'def function_example():' is missing a docstring.
[2025-04-14 22:30:46] [WARNING] tests/testfile.py:: line should be indented with tabs, not spaces. Please use tabs for indentation.
[2025-04-14 22:30:46] [ERROR] Analysis completed with 7 errors.
root@eddc83f1bb46:/home/ubuntu/BettyPy# 
```

---

### Requirements
- Linux or macOS (Windows support is experimental)

---

### Contributing

We welcome contributions to improve BettyPy! Follow these steps to get started:

#### Option 1: Fork and Submit a Pull Request
1. Fork the repository on GitHub.
2. Create a new branch for your feature or bug fix:
   ```bash
   $ git checkout -b my-feature
   ```
3. Make your changes and commit them with a clear message:
   ```bash
   $ git commit -m "Add my feature"
   ```
4. Push your changes to your fork:
   ```bash
   $ git push origin my-feature
   ```
5. Open a pull request to the main repository.
6. Your changes will be reviewed, and if approved, merged into the main branch.

#### Option 2: Report an Issue
1. Found a bug or have a feature request? Create an issue in the GitHub repository.
2. Provide a clear and detailed description of the problem or suggestion.
3. Include relevant details, such as steps to reproduce the issue or examples of the desired feature.
4. If you'd like to work on the issue, assign it to yourself or request assignment.
5. Once your work is complete, submit a pull request referencing the issue.

#### Review Process
- All pull requests are reviewed by project maintainers.
- Feedback will be provided if changes are needed.
- Once approved, your contribution will be merged, and you’ll be credited for your work.

Thank you for helping improve BettyPy!

---

### Acknowledgments

<p align="center">
  <strong>Holberton School</strong>
</p>

<p align="center">
  <img src="https://user-images.strikinglycdn.com/res/hrscywv4p/image/upload/c_limit,fl_lossy,h_630,w_1200,f_auto,q_auto/79001/331125_630361.png" alt="Holberton School Logo" width="150">
</p>

<p align="center">
  This project was inspired by the Betty style guide for C, developed at <strong>Holberton School</strong>. BettyPy was created as a personal project to bring similar coding standards to Python. If it can help students or developers, all the better!
</p>

---

### Authors &copy;

- **Martin Clement** - [Github Profile](https://github.com/ItsZmainDev)

---

### License
This project is licensed under the Mozilla Public License 2.0. See the [LICENSE](LICENSE) file for details.