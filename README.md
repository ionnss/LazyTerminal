# Lazy Terminal

![LazyTerminal Logo](static/logo.png)

**LazyTerminal** is a simple and efficient command-line interface (CLI) tool designed to interact with OpenAI's ChatGPT directly from your terminal. With **LazyTerminal**, you can send queries and receive intelligent responses without relying on additional applications or unnecessary logins.

---

## Features

- **Send Queries:** Use the `lzt` command to ask ChatGPT questions directly.
- **Token Management:**
  - Automatically saves your OpenAI API token in a `.env` file.
  - Update the token anytime using the `lzt -ut` command.
- **Global Command Installation:** Make `lzt` a global command for convenience.
- **Customizable:** Ready for enhancements like advanced configurations and logging.

---

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

## Requirements

- **Go**: Version 1.17 or newer.
- **OpenAI API Key:** Obtain your key from [OpenAI's API platform](https://platform.openai.com/).

---

## Installation

### **Clone the Repository**
```bash
git clone https://github.com/your-username/lazyterminal.git
cd lazyterminal
```

### **Install Dependencies**
```bash
go mod tidy
```

### **Build the Executable**
Create the executable file for LazyTerminal:
```bash
go build -o lzt
```

### **Install as a Global Command**
1. Move the executable to a directory in your system's `PATH`:
   - On Linux/macOS:
     ```bash
     sudo mv lzt /usr/local/bin/
     ```
   - On Windows:
     Add the directory containing the `lzt` executable to your system's `Path` environment variable.

2. Test the command:
   ```bash
   lzt "This is a test"
   ```

Now, `lzt` is accessible globally as a terminal command.

---

## Usage

### **Initial Setup: Insert Your OpenAI API Token**
When you first run the program, it will ask you to input your OpenAI API token:
```bash
Enter your OpenAI API token: sk-xxxxxx
```

Once provided, the token will be securely saved in the `.env` file for future use.

### **Send Queries**
Ask a question using the `lzt` command:
```bash
lzt "What is the capital of Brazil?"
```

The program will respond with:
```bash
LazyTerminal > The capital of Brazil is Brasília.
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

## Troubleshooting

### Common Errors

1. **Exceeded API Quota**:
   - If you see the error `You exceeded your current quota`, check your OpenAI account usage and ensure your API key has sufficient credits.

2. **Permission Issues Running Executable**:
   - If you cannot execute the `lzt` file, ensure it has execute permissions:
     ```bash
     chmod +x lzt
     ```

3. **Command Not Found (Global Command)**:
   - Ensure the `lzt` executable is in a directory listed in your system's `PATH`.
   - Verify your `PATH` configuration:
     ```bash
     echo $PATH
     ```

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

- **Developed by:** [ionnss - Leonardo Zara](https://github.com/ionnss)
- **Mascot Design:** A curious and playful sloth in `static/logo.png`.

---

Let me know if there’s anything else you’d like to modify or add! 🚀
