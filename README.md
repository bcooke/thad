# Thad

> Turn plain-English requests into safe shell commands **and** help you memorize them with built-in spaced-repetition—all from one tiny CLI.

---

## What is Thad?

**Thad** (**T**erminal **H**elper **A**nd **D**riller) is a CLI tool that:
- Converts natural language into safe, ready-to-run shell commands using local or cloud LLMs
- Lets you save commands as flashcards and drill them with spaced repetition
- Ships as a single static Go binary—no external runtime dependencies

---

## Features (in progress)
- [x] Go CLI scaffold with Cobra
- [x] YAML config loader (`~/.config/thad/config.yaml`)
- [x] Sample config in repo
- [ ] "Ask" command: get shell command suggestions from OpenAI or Ollama
- [ ] LLM backend selection (OpenAI, Ollama)
- [ ] Spaced repetition flashcard system
- [ ] Command safety filter & audit log

---

## Quickstart

1. **Clone the repo:**
   ```sh
   git clone https://github.com/brettcooke/thad.git
   cd thad
   ```
2. **Copy and edit the sample config:**
   ```sh
   mkdir -p ~/.config/thad
   cp config.sample.yaml ~/.config/thad/config.yaml
   # Edit ~/.config/thad/config.yaml to set your model provider and API key
   ```
3. **Build and run:**
   ```sh
   go run cmd/thad/main.go version
   go run cmd/thad/main.go ask "find files named *.log"
   ```

---

## Configuration

Edit `~/.config/thad/config.yaml` (see `config.sample.yaml` for all options):

```yaml
shell: zsh
model:
  provider: openai        # openai | anthropic | ollama | …
  api_key: ENV:OPENAI_API_KEY
prompt_preamble: |
  You are an expert shell assistant. Return the shortest working
  command for the user's OS unless they ask for alternatives.
srs:
  algorithm: sm2
  db_path: ~/.local/share/thad/flashcards.db
```

---

## Roadmap (MVP)

| Pri | Feature                             | Status  |
| --- | ----------------------------------- | ------- |
| P0  | Core `ask` + OpenAI / Ollama client | In progress |
| P0  | `run` w/ confirmation & logging     | Planned |
| P0  | YAML config subsystem               | Done    |
| P1  | Flashcard DB, `remember`, `study`   | Planned |
| P1  | Unit tests & golden files           | Planned |
| P2  | `explain`, `stats`, `history`       | Planned |
| P2  | Self-update via GitHub releases     | Planned |
| P3  | Fish / PowerShell adapters          | Planned |
| P3  | Plugin SDK draft                    | Planned |

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

---

## License

MIT 