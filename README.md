## :arrow_forward: Basic Discord Bot 



1. Set the enviroment variables in `.ps1` file (use .bashrc or .zshrc if you're using Linux/MacOS)

Powershell (`.ps1`)
```powershell
$env:TOKEN = "token"
$env:BOT_PREFIX = "prefix"
```

Unix/Linux (`.bashrc` o `.zshrc`)
```bash
export TOKEN="token"
export BOT_PREFIX="prefix"
```

2. Load environment variables running `.ps1` in a terminal

Powershell
```powershell
./file.ps1
```

Unix/Linux
```bash
source .bashrc  # for .bashrc files
source .zshrc   # for .zshrc files
```

--- 

#### Acknowledgements :blue_heart:

This project utilizes the packages :link: [DiscordGo](https://github.com/bwmarrin/discordgo) 