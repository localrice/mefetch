# mefetch

<p align="center">
    <img src="./docs/images/mefetch_banner.png" alt="mefetch" border="0">
</p>

A fetching tool to display your personal profile and information from GitHub.

<p align="center">
    <img src="./docs/images/mefetch.png" alt="mefetch" border="0">
</p>

## Configuration

Initialize the configuration:

```bash
mefetch init
```

Edit the configuration:

```bash
mefetch config
```

Configuration files are stored in:

```text
~/.config/mefetch/
├── config.yaml
└── ascii.txt
```

### Example config file

```yaml
name: auto
location: auto

github: your-username
interests: cats, Linux
discord: your-discord

text-color: "#4ba3f5"
ascii: true
```

## Options

* `auto`: Automatically fetches the field from your GitHub profile.
* `github`:Your GitHub username. Used as the source for `auto` fields.
* Other values: Displayed exactly as configured.
* `text-color`: Color used for profile labels.
* `ascii`: Enables or disables the custom ASCII art from `ascii.txt`.

### GitHub

Fields set to `auto` are populated using the GitHub API.
Any profile information supported by the GitHub API can be used when available, including profile details, links, and other public account fields.

For example:

```yaml
name: auto
location: auto
github: your-username
```

This lets mefetch automatically use the corresponding information from your GitHub profile while keeping other fields manually configurable.
