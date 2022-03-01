# Home Network Monitor

Home Network Monitor is a lightweight cross-platform tool to monitor and notify which devices are connected to your home network.

If you want a simple tool to check from time to time whether there are new devices connecting to your network this is your tool.

It sends notifications to your Telegram account from time to time with a full report.

## Installation

Installation is easy. Please follow the steps below:

1. Go to the dist folder and choose the binary file for your system. At this point we have versions compiled for: Linux, MacOS, Windows and Raspbian
2. Copy the file to your system
3. Run the application

## Configuration

To work properly the application has to have the following configuration files:

1. subnets.cfg: this file is a list of subnets you want to monitor. It is a simple list containing the first three (3) octets of the desired IP address. Example: 192.168.1.
2. .env or environment variables. The tool requires three (3) environment variables:
   1. `TG_TOKEN`: telegram API token
   2. `CHAT_ID`: telegram chat ID
   3. `TG_URL`: telegram API URL
