# Tedious 📝

> Because managing your to-dos shouldn't be more tedious than actually doing them.

## What is this?

Let's be honest: most task management apps are like that friend who over-explains everything. You just wanted to jot down "buy milk," but now you're choosing priority levels, tags, categories, sub-categories, and somehow picking an emoji that represents your emotional connection to dairy products.

**Tedious** is different. It's a terminal-based (TUI) todo list that strips away all the visual noise and gets straight to the point. No fancy dashboards, no notification badges giving you anxiety, no gamification trying to trick you into productivity. Just you, your tasks, and a clean terminal interface.

### But wait, there's more!

Tedious comes with built-in time tracking that's actually *easy* to use. Track how long tasks take with a single keystroke, and finally answer that age-old question: "Why did this 'quick' task take 3 hours?" Spoiler: it always takes longer than you think. But hey, at least now you'll have data to back up your excuses... er, estimates.

## Installation
```bash
# Clone the repository
git clone https://github.com/arthurlaramachado/tedious.git
cd tedious

# Run directly
go run .

# Or build and run
go build -o tedious
./tedious
```

And boom, you're managing tasks like a terminal wizard. 🧙‍♂️

## Keyboard Shortcuts

Because clicking is so 2010.

### Navigation
| Key | Action |
|-----|--------|
| `↑` | Move cursor up |
| `↓` / `Enter` | Move cursor down (creates new task if at the end of a non-empty task) |

### Task Management
| Key | Action |
|-----|--------|
| `Ctrl+N` | Create a new task at current position |
| `Ctrl+D` | Delete current task |
| `Backspace` | Delete last character (deletes task if empty) |
| `Tab` | Convert task to subtask |
| `Space` / `Any key` | Type characters into your task |

### Task Reordering
| Key | Action |
|-----|--------|
| `Ctrl+K` | Move current task up ↑ |
| `Ctrl+J` | Move current task down ↓ |

*Vim users will feel right at home. Others will wonder why we didn't use arrows. Trust us, this is better.*

### Time Tracking
| Key | Action |
|-----|--------|
| `Ctrl+B` | Toggle task state (Unmarked → In Progress → Completed → Unmarked) |

When you mark a task as "In Progress," time tracking starts automatically. Mark it "Completed" and it stops. It's like a stopwatch, but you don't have to pretend you're timing a race.

**⚠️ Heads up:** If you cycle a completed task back to unmarked, the timer resets. So don't fat-finger that `Ctrl+B` unless you want to start fresh!

### Application Control
| Key | Action |
|-----|--------|
| `Ctrl+C` | Save and exit (your tasks are safe, unlike your browser tabs) |

## Features

- ✅ Clean, distraction-free interface
- ⏱️ Automatic time tracking for tasks
- 🎯 Subtask support (tasks can have tasks)
- ⌨️ Keyboard-driven workflow (your mouse can finally rest)
- 🎨 Minimal visual clutter (no animations, no gradients, just getting things done)

## Task States

Tasks cycle through three states:

1. **Unmarked** (☐) - Not started yet
2. **In Progress** (⧗) - Currently working on it (timer running!)
3. **Completed** (✓) - Done and dusted (with time taken recorded)

Press `Ctrl+B` to cycle through states and watch your productivity data grow!

## Why "Tedious"?

The name is ironic. Managing your tasks *shouldn't* be tedious. That's the whole point. Also, all the good names were taken.

## Contributing

Found a bug? Want to add a feature? PRs are welcome! 

**Our philosophy:** Keep it clean, keep it simple. We love improvements that make Tedious more useful, but we're committed to staying distraction-free. Think "elegantly minimal" rather than "feature-packed dashboard." ✨

Before submitting a PR, please:
- Keep the terminal UI clean and uncluttered
- Maintain keyboard-first navigation
- Test on both macOS and Linux if possible
- Add documentation for new features

When in doubt, open an issue first to discuss your idea!

## License
This project is licensed under the MIT License - see the 
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) for details.

## Credits

Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) - because building TUIs should be fun, not tedious.

---

*Now stop reading this README and go finish your actual tasks.* 🚀