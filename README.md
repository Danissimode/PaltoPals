<br/>
<div align="center">
  <a href="https://github.com/Danissimode/Palto">
    <img src="https://Paltopals.dev/gh.svg?v=2" alt="Paltopals Logo" width="100%">
  </a>
  <h3 align="center">Paltopals</h3>
  <p align="center">
    Your intelligent pair programmer directly within your Palto sessions.
    <br/>
    <br/>
    <a href="https://github.com/Danissimode/Palto/blob/main/LICENSE"><img alt="License" src="https://img.shields.io/github/license/Danissimode/Paltopals?style=flat-square"></a>
    <a href="https://github.com/Danissimode/Palto/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/Danissimode/Paltopals?style=flat-square"></a>
    <a href="https://github.com/Danissimode/Palto/issues"><img alt="Issues" src="https://img.shields.io/github/issues/Danissimode/Paltopals?style=flat-square"></a>
    <br/>
    <br/>
    <br/>
    <a href="https://Paltopals.dev/screenshots" target="_blank">Screenshots</a> |
    <a href="https://github.com/Danissimode/Palto/issues/new?labels=bug&template=bug_report.md" target="_blank">Report Bug</a> |
    <a href="https://github.com/Danissimode/Palto/issues/new?labels=enhancement&template=feature_request.md" target="_blank">Request Feature</a>
    <br/>
    <br/>
    <a href="https://Paltopals.dev/Palto-getting-started/" target="_blank">Palto Getting Started</a> |
    <a href="https://Paltopals.dev/Palto-config/" target="_blank">Palto Config Generator</a> |
    <a href="https://Paltopals.dev/Palto-shortcuts/" target="_blank">Palto Shortcuts</a>
  </p>
</div>

## Table of Contents

- [About The Project](#about-the-project)
  - [Human-Inspired Interface](#human-inspired-interface)
- [Installation](#installation)
  - [Quick Install](#quick-install)
  - [Homebrew](#homebrew)
  - [Manual Download](#manual-download)
- [Post-Installation Setup](#post-installation-setup)
- [Paltopals Layout](#Paltopals-layout)
- [Observe Mode](#observe-mode)
- [Prepare Mode](#prepare-mode)
- [Watch Mode](#watch-mode)
  - [Activating Watch Mode](#activating-watch-mode)
  - [Example Use Cases](#example-use-cases)
- [Squashing](#squashing)
  - [What is Squashing?](#what-is-squashing)
  - [Manual Squashing](#manual-squashing)
- [Core Commands](#core-commands)
- [Command-Line Usage](#command-line-usage)
- [Configuration](#configuration)
  - [Environment Variables](#environment-variables)
  - [Session-Specific Configuration](#session-specific-configuration)
  - [Using Other AI Providers](#using-other-ai-providers)
- [Contributing](#contributing)
- [License](#license)

## About The Project

![Product Demo](https://Paltopals.dev/demo.png)

Paltopals is an intelligent terminal assistant that lives inside your Palto sessions. Unlike other CLI AI tools, Paltopals observes and understands the content of your Palto panes, providing assistance without requiring you to change your workflow or interrupt your terminal sessions.

Think of Paltopals as a _pair programmer_ that sits beside you, watching your terminal environment exactly as you see it. It can understand what you're working on across multiple panes, help solve problems and execute commands on your behalf in a dedicated execution pane.

### Human-Inspired Interface

Paltopals's design philosophy mirrors the way humans collaborate at the terminal. Just as a colleague sitting next to you would observe your screen, understand context from what's visible, and help accordingly, Paltopals:

1. **Observes**: Reads the visible content in all your panes
2. **Communicates**: Uses a dedicated chat pane for interaction
3. **Acts**: Can execute commands in a separate execution pane (with your permission)

This approach provides powerful AI assistance while respecting your existing workflow and maintaining the familiar terminal environment you're already comfortable with.

## Installation

Paltopals requires only Palto to be installed on your system. It's designed to work on Unix-based operating systems including Linux and macOS.

### Quick Install

The fastest way to install Paltopals is using the installation script:

```bash
# install Palto if not already installed
curl -fsSL https://get.Paltopals.dev | bash
```

This installs Paltopals to `/usr/local/bin/Paltopals` by default. If you need to install to a different location or want to see what the script does before running it, you can view the source at [get.Paltopals.dev](https://get.Paltopals.dev).

### Homebrew

If you use Homebrew, you can install Paltopals with:

```bash
brew install Paltopals
```

### Manual Download

You can also download pre-built binaries from the [GitHub releases page](https://github.com/Danissimode/Palto/releases).

After downloading, make the binary executable and move it to a directory in your PATH:

```bash
chmod +x ./Paltopals
sudo mv ./Paltopals /usr/local/bin/
```

## Post-Installation Setup

After installing Paltopals, you need to configure your API key to start using it:

1. **Set the API Key**  
   Paltopals uses the OpenRouter endpoint by default. Set your API key by adding the following to your shell configuration (e.g., `~/.bashrc`, `~/.zshrc`):

   ```bash
   export Paltopals_OPENROUTER_API_KEY="your-api-key-here"
   ```

2. **Start Paltopals**

   ```bash
   Paltopals
   ```

## Paltopals Layout

![Panes](https://Paltopals.dev/shots/panes.png?lastmode=1)

Paltopals is designed to operate within a single Palto window, with one instance of
Paltopals running per window and organizes your workspace using the following pane structure:

1. **Chat Pane**: This is where you interact with the AI. It features a REPL-like interface with syntax highlighting, auto-completion, and readline shortcuts.

2. **Exec Pane**: Paltopals selects (or creates) a pane where commands can be executed.

3. **Read-Only Panes**: All other panes in the current window serve as additional context. Paltopals can read their content but does not interact with them.

## Observe Mode

![Observe Mode](https://Paltopals.dev/shots/demo-observe.png)
_Paltopals sent the first ping command and is waiting for the countdown to check for the next step_

Paltopals operates by default in "observe mode". Here's how the interaction flow works:

1. **User types a message** in the Chat Pane.

2. **Paltopals captures context** from all visible panes in your current Palto window (excluding the Chat Pane itself). This includes:

   - Current command with arguments
   - Detected shell type
   - User's operating system
   - Current content of each pane

3. **Paltopals processes your request** by sending user's message, the current pane context, and chat history to the AI.

4. **The AI responds** with information, which may include a suggested command to run.

5. **If a command is suggested**, Paltopals will:

   - Check if the command matches whitelist or blacklist patterns
   - Ask for your confirmation (unless the command is whitelisted)
   - Execute the command in the designated Exec Pane if approved
   - Wait for the `wait_interval` (default: 5 seconds) (You can pause/resume the countdown with `space` or `enter` to stop the countdown)
   - Capture the new output from all panes
   - Send the updated context back to the AI to continue helping you

6. **The conversation continues** until your task is complete.

![Observe Mode Flowchart](https://Paltopals.dev/shots/observe-mode.png)

## Prepare Mode

![Prepare Mode](https://Paltopals.dev/shots/demo-prepare.png?lastmode=1)
_Paltopals customized the pane prompt and sent the first ping command. Instead of the countdown, it's waiting for command completion_

Prepare mode is an optional feature that enhances Paltopals's ability to work with your terminal by customizing
your shell prompt and tracking command execution with better precision. This
enhancement eliminates the need for arbitrary wait intervals and provides the AI
with more detailed information about your commands and their results.

When you enable Prepare Mode, Paltopals will:

1. **Detects your current shell** in the execution pane (supports bash, zsh, and fish)
2. **Customizes your shell prompt** to include special markers that Paltopals can recognize
3. **Will track command execution history** including exit codes, and per-command outputs
4. **Will detect command completion** instead of using fixed wait time intervals

To activate Prepare Mode, simply use:

```
Paltopals » /prepare
```

**Prepared Fish Example:**

```shell
$ function fish_prompt; set -l s $status; printf '%s@%s:%s[%s][%d]» ' $USER (hostname -s) (prompt_pwd) (date +"%H:%M") $s; end
username@hostname:~/r/Paltopals[21:05][0]»
```

## Watch Mode

![Watch Mode](https://Paltopals.dev/shots/demo-watch.png)
_Paltopals watching user shell commands and better alternatives_

Watch Mode transforms Paltopals into a proactive assistant that continuously
monitors your terminal activity and provides suggestions based on what you're
doing.

### Activating Watch Mode

To enable Watch Mode, use the `/watch` command followed by a description of what you want Paltopals to look for:

```
Paltopals » /watch spot and suggest more efficient alternatives to my shell commands
```

When activated, Paltopals will:

1. Start capturing the content of all panes in your current Palto window at regular intervals (`wait_interval` configuration)
2. Analyze content based on your specified watch goal and provide suggestions when appropriate

### Example Use Cases

Watch Mode could be valuable for scenarios such as:

- **Learning shell efficiency**: Get suggestions for more concise commands as you work

  ```
  Paltopals » /watch spot and suggest more efficient alternatives to my shell commands
  ```

- **Detecting common errors**: Receive warnings about potential issues or mistakes

  ```
  Paltopals » /watch flag commands that could expose sensitive data or weaken system security
  ```

- **Log Monitoring and Error Detection**: Have Paltopals monitor log files or terminal output for errors

  ```
  Paltopals » /watch monitor log output for errors, warnings, or critical issues and suggest fixes
  ```

## Squashing

As you work with Paltopals, your conversation history grows, adding to the context
provided to the AI model with each interaction. Different AI models have
different context size limits and pricing structures based on token usage. To
manage this, Paltopals implements a simple context management feature called
"squashing."

### What is Squashing?

Squashing is Paltopals's built-in mechanism for summarizing chat history to manage
token usage.

When your context grows too large, Paltopals condenses previous
messages into a more compact summary.

You can check your current context utilization at any time using the `/info` command:

```bash
Paltopals » /info

Context
────────

Messages            15
Context Size~       16500 tokens
                    ████████░░ 82.5%
Max Size            20000 tokens
```

This example shows that the context is at 82.5% capacity (16,500 tokens out of 20,000). When the context size reaches 80% of the configured maximum (`max_context_size` in your config), Paltopals automatically triggers squashing.

### Manual Squashing

If you'd like to manage your context before reaching the automatic threshold, you can trigger squashing manually with the `/squash` command:

```bash
Paltopals » /squash
```

## Core Commands

| Command                     | Description                                                      |
| --------------------------- | ---------------------------------------------------------------- |
| `/info`                     | Display system information, pane details, and context statistics |
| `/clear`                    | Clear chat history.                                              |
| `/reset`                    | Clear chat history and reset all panes.                          |
| `/config`                   | View current configuration settings                              |
| `/config set <key> <value>` | Override configuration for current session                       |
| `/squash`                   | Manually trigger context summarization                           |
| `/prepare`                  | Initialize Prepared Mode for the Exec Pane                       |
| `/watch <description>`      | Enable Watch Mode with specified goal                            |
| `/exit`                     | Exit Paltopals                                                      |

## Command-Line Usage

You can start `Paltopals` with an initial message or task file from the command line:

- **Direct Message:**

  ```sh
  Paltopals your initial message
  ```

- **Task File:**
  ```sh
  Paltopals -f path/to/your_task.txt
  ```

## Configuration

The configuration can be managed through a YAML file, environment variables, or via runtime commands.

Paltopals looks for its configuration file at `~/.config/Paltopals/config.yaml`.
For a sample configuration file, see [config.example.yaml](https://github.com/Danissimode/Palto/blob/main/config.example.yaml).

### Environment Variables

All configuration options can also be set via environment variables, which take precedence over the config file. Use the prefix `Paltopals_` followed by the uppercase configuration key:

```bash
# Examples
export Paltopals_DEBUG=true
export Paltopals_MAX_CAPTURE_LINES=300
export Paltopals_OPENROUTER_API_KEY="your-api-key-here"
export Paltopals_OPENROUTER_MODEL="..."
```

You can also use environment variables directly within your configuration file values. The application will automatically expand these variables when loading the configuration:

```yaml
# Example config.yaml with environment variable expansion
openrouter:
  api_key: "${OPENAI_API_KEY}"
  base_url: https://api.openai.com/v1
```

### Session-Specific Configuration

You can override some configuration values for your current Paltopals session using the `/config` command:

```bash
# View current configuration
Paltopals » /config

# Override a configuration value for this session
Paltopals » /config set max_capture_lines 300
Paltopals » /config set openrouter.model gpt-4o-mini
```

These changes will persist only for the current session and won't modify your config file.

### Using Other AI Providers

OpenRouter is OpenAI API-compatible, so you can direct Paltopals at OpenAI or any other OpenAI API-compatible endpoint by customizing the `base_url`.

For OpenAI:

```yaml
openrouter:
  api_key: sk-proj-XXX
  model: o4-mini-2025-04-16
  base_url: https://api.openai.com/v1
```

For Anthropic's Claude:

```yaml
openrouter:
  api_key: sk-proj-XXX
  model: claude-3-7-sonnet-20250219
  base_url: https://api.anthropic.com/v1
```

For local Ollama:

```yaml
openrouter:
  api_key: api-key
  model: gemma3:1b
  base_url: http://localhost:11434/v1
```

_Prompts are currently tuned for Gemini 2.5 by default; behavior with other models may vary._

## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request.
You can also simply open an issue.
<br>
Don't forget to give the project a star!

## License

Distributed under the Apache License. See [Apache License](https://github.com/Danissimode/Palto/blob/main/LICENSE) for more information.
