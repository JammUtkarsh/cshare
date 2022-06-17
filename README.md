## Cshare(Clip Share) (_Under Development_)

![Cshare Logo](https://github.com/JammUtkarsh/cshare/blob/main/misc/Cshare.png?raw=true)

Cshare is a CLI app that lets you send clips/texts across multiple devices.

# Project overview

# 💡 Inspiration

While I was building my local self-hosted server for streaming and file storage. I decided to install [Manjaro](https://manjaro.org/) as my base OS.
Since I would be operating it remotely, I wanted to make sure I can connect it via ssh. 

Manjaro is a security tight OS. Since I was still learning, I wasn't able to connect with it via ssh.
This was because I had to share the ssh key of laptop to my server. Then I had known no way to share a key securely. 
Maybe I was still a beginner for connecting with Manjaro with SSH or I just wanted a simple solution to share data(int text) across devices.

This is how I was inspired to make this app. 

# ⚙ Mechanism

- Install the client-side app on the devices you want to share clip on. 
- Follow the [user guide](https://github.com/JammUtkarsh/cshare/blob/main/UserGuide.md) and run the necessary commands.
- `cp` command  will send an _POST_  request to the server. 
- `paste` command will make a _GET_ request from the server.
- This server can be self-hosted or could be centrally provided. 

# 🔧 How we built it

- For the CLI structure we have used [Cobra](https://github.com/spf13/cobra) Framework.
- For Web Server Config, we have used [Gorilla/Mux](https://github.com/gorilla/mux) Framework.
- Rest of the architecture is yet to be decided.

## 💪 Challenges we ran into

- To avoid framework lock-in, choosing the right framework.

[//]: # (## 📌 Accomplishments that we're proud of)

[//]: # ()
[//]: # (## 📚 What we learned)

[//]: # ()
[//]: # (## ⏭ What's next for Cshare)

[//]: # (- Todo)
