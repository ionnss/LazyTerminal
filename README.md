# LazyTerminal

![LazyTerminal Logo](static/logo.png)

**LazyTerminal** is a simple and fun CLI tool that allows users to send queries to OpenAI's ChatGPT directly from their terminal. Designed with simplicity in mind, it includes features like token management and a sloth mascot for added charm.

## Project Structure

```
├── api.go            # Handles API interactions with OpenAI
├── go.mod            # Go module file
├── go.sum            # Go dependencies lock file
├── main.go           # CLI logic and command handling
└── static
    └── logo.png      # Project logo (a cute sloth!)
```

---

## Features

- **Send Queries:** Use the `lzt` command to ask ChatGPT questions directly.
- **Token Management:**
  - Automatically saves your OpenAI API token in a `.env` file.
  - Update the token anytime using the `lzt -ut` command.
- **Customizable:** Ready for enhancements like advanced configurations and logging.

---

## Requirements

- **Go**: Version 1.23 or newer.
- **OpenAI API Key:** Obtain your key from [OpenAI's API platform](https://platform.openai.com/).

---

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/lazyterminal.git
   cd lazyterminal
   ```

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Set Up Your API Token**:
   - When prompted, provide your OpenAI API token. It will be saved in a `.env` file for future use.

4. **Run the Application**:
   ```bash
   go run main.go
   ```

---

## Usage

### **Basic Command**
Ask a question using the `lzt` command:
```bash
lzt "What is the capital of France?"
```

The program will respond with:
```bash
LazyTerminal > The capital of France is Paris.
```

### **Update Token**
If you need to update your API token, use the `lzt -ut` command:
```bash
lzt -ut
```
You will be prompted to enter a new token, which will be saved in the `.env` file.

### **Exit the Program**
Type `exit` to close the program:
```bash
exit
```

---

## Customization

- **Logo:** A playful sloth logo (`static/logo.png`) is included for branding.
- **API Parameters:** The model used is `gpt-3.5-turbo` by default. You can modify the `api.go` file to support other models.

---

## Future Improvements

- **Interactive Mode:** Continuous conversations without repeating the `lzt` prefix.
- **Logging:** Save query and response history to a local file.
- **Advanced Configurations:** Add support for parameters like temperature and response length.
- **Cross-Platform Binaries:** Distribute executables for Linux, macOS, and Windows.

---

## Contributing

Contributions are welcome! Feel free to fork the repository and submit a pull request with your improvements or ideas.

---

## License

This project is licensed under the MIT License. See the [LICENSE](https://opensource.org/licenses/MIT) file for details.

---

## Credits

- **Developed by:** [ionnss](https://www.github.com/ionnss)
- **Mascot Design:** A curious and playful sloth in `static/logo.png`.

