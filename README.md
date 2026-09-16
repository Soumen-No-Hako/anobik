# Anobik
a small chatgpt, gemini type agentic search engine and assistant

# Purpose
Using any open weight model, custom tools to get data and info from scratch. No API. (For now, since I am poor)

# Requirement
| Name | Version | Required? |
|---|---|---|
| Ollama || Yes
| Go || Yes
| git || No(optional)

# Set-up

Will create an installer later.

1. [Install Ollama](https://docs.ollama.com/quickstart)
2. clone this repo - git clone https://github.com/Soumen-No-Hako/anobik.git OR git clone git@github.com:Soumen-No-Hako/anobik.git
3. Run/get some open-weights model's gguf from ollama site or huggingface e.g llama3.2, gemma4 etc.
```bash
ollama run {model-name}
```
4. Run these type of commands to create the Agents/tools
```bash
ollama create {Agent_Name}:{version_number} -f ./Orchestro/orchestro-modelfile
```
5. Run you agents
```bash
ollama run {Agent_name}:{Version_number}
```
