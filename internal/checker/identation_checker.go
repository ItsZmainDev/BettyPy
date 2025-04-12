package checker

type IdentationChecker struct {
    Content  []string
    Filename string
}

func NewIdentationChecker(content []string, filename string) *IdentationChecker {
    return &IdentationChecker{Content: content, Filename: filename}
}

func (c *IdentationChecker) Run() []string {
    var errors []string

	for i, line := range c.Content {
		if len(line) > 0 && line[0] == ' ' {
			errors = append(errors, c.Filename+":"+string(i+1)+": line should be indented with tabs, not spaces. Please use tabs for indentation.")
		}
		if len(line) > 0 && line[len(line)-1] == ' ' {
			errors = append(errors, c.Filename+":"+string(i+1)+": line should not have trailing spaces. Please remove trailing spaces at the end of the line.")
		}
	}
    
    return errors
}