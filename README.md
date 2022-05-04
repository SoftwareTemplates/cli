<div align="center">
    <h1>SoftwareTemplates CLI</h1>
<hr>
<strong>A simple and powerful cli for setting up your projects more quickly</strong><br><br>
<img src="https://img.shields.io/github/workflow/status/SoftwareTemplates/cli/Test?style=for-the-badge">
<img src="https://img.shields.io/github/license/SoftwareTemplates/cli?style=for-the-badge"> 
<img src="https://img.shields.io/github/go-mod/go-version/SoftwareTemplates/cli?style=for-the-badge">
</div>
<hr>

<div align="center">
<img src=".media/icon.png" width="100">
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" height="100">
</div>

# Project Information

This project provides a simple cli for setting up new projects more quickly. 
The organisation contains a lot of templates for many different languages and software frameworks.
You can use this cli to easily use the templates and skip the anoying process of 
creating your projects base structure.

# Usage

Init your first project
```bash
    softwareTemplates init --initGit --template test-template
```

### Commands

<strong>1. Init:</strong><br>
```bash
softwareTemplates init # general init command for new projects
```

### Flags

1. `--initGit or --git or --ig`
If this flag is added to the init command, a new git repository will be 
initialized on project creation. If it is not provided, this will be validated by a cli prompt.

2. `--template or -t`
If this flag is provided it will directly choose a template. If the 
template does not exists the cli terminates. If this flag is not provided the 
the cli will open a prompt to select a template

3. `--projectName or -p`
If this flag is added to the init command, the project name will automatically
selected. If this flag is not provided the cli will open a input prompt


# Installation

1. Choose the latest version from the releases tab
2. Select the version that is used for your operating system
3. Move to your program files

<strong>Linux/MacOS: </strong>
```bash
   sudo mv ./softwareTemplates /usr/local/bin/softwareTemplates
```

<strong>Windows:</strong><br>
Move copy your file to `C:\Program Files`

