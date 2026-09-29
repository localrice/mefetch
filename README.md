# mefetch

a fetching tool to print your personal profile including sources from github


# Configuration

mefetch stores its configuration files in:

```text
~/.config/mefetch/
├── config.yaml
└── ascii.txt
```

# Initialize:
```
mefetch init
```
Edit the config:
```
mefetch config
```
Example config:
```
name: auto
location: auto

github: your-username
interests: robotics, Linux
discord: your-discord

text-color: "#4ba3f5"
ascii: true
```

github:  GitHub username used for auto fields.
auto: fetches the matching field from GitHub.
Other values: displayed as entered.
text-color: profile label color.
ascii: enables/disables ascii.txt.