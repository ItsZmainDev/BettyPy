# BettyPy - Python Style Guide for Betty

### Table of Contents
- [Description](#description)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
- [License](#license)
- [Contributing](#contributing)
- [Authors](#authors)

### Description

BettyPy is a Python style guide for the Betty programming language. It provides a set of rules and guidelines for writing clean, readable, and maintainable code in Betty. The goal of BettyPy is to promote best practices and improve the overall quality of code written in Betty.

### Features
| Feature | Description |
| ------- | ----------- |
| Empty files | Check for empty files. |
| Line Length | Limit lines to 80 characters. |
| Indentation | Use 4 spaces per indentation level. |
| Docstrings | Use docstrings to describe modules, classes, and functions. |
| Comments | Check comment |
| Anotations | Check for anotations: TODO, FIXME, etc. |
| Imports | Check for unused imports. |
| Last Line | Ensure the last line of a file is a newline. |

### Installation
To install BettyPy, execute the following command:

```bash
git clone https://github.com/ItsZmainDev/BettyPy.git
cd BettyPy
```

Then, run the installation script:

```bash
./install.sh
```

### Usage
To use BettyPy, you can run the following command in your terminal:

```bash
bettypy <path_to_your_python_file>
```

### Example
###### Example Python file to test BettyPy
```python
# This is an example Python file

def function_example():
	# print("cette ligne est trop longue et dépasse 80 caractères, ce qui est inacceptable dans le style de codage de Betty.")
    print("cette ligne n'est pas indentée correctement et doit être corrigée.")

####test
```
###### Command to run BettyPy
```bash
bettypy example.py
```
###### Output of BettyPy
```bash
root@eddc83f1bb46:/home/ubuntu/BettyPy# bettypy tests/testfile.py 
[2025-04-12 15:07:25] [INFO] tests/testfile.py - Starting analysis...
[2025-04-12 15:07:25] [WARNING] tests/testfile.py:4: line too long (125 characters)
[2025-04-12 15:07:25] [WARNING] tests/testfile.py:5: line too long (82 characters)
[2025-04-12 15:07:25] [WARNING] tests/testfile.py:7: comment should start with a space after #
[2025-04-12 15:07:25] [WARNING] tests/testfile.py:7: comment too short (4 characters)
[2025-04-12 15:07:25] [WARNING] tests/testfile.py: file should end with a new line
[2025-04-12 15:07:25] [WARNING] tests/testfile.py:3: Function 'def function_example():' is missing a docstring.
[2025-04-12 15:07:25] [ERROR] Analysis completed with 6 errors.
root@eddc83f1bb46:/home/ubuntu/BettyPy# 
```

### License
This project is licensed under the Mozilla Public License 2.0. See the [LICENSE](LICENSE) file for details.

### Contributing
Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request.

### Authors &copy;

- **Martin Clement** - [Github Profile](https://github.com/ItsZmainDev)