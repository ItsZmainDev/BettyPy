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
- [License](#license)
- [Authors](#authors)

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

### Installation
To install BettyPy, execute the following commands:

```bash
$ git clone https://github.com/ItsZmainDev/BettyPy.git
$ cd BettyPy
$ ./install.sh
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
1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m 'Add your feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

---

### License
This project is licensed under the Mozilla Public License 2.0. See the [LICENSE](LICENSE) file for details.

---

### Authors &copy;

- **Martin Clement** - [Github Profile](https://github.com/ItsZmainDev)