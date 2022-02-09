# HosterRed collection of system checks
A collection of various OS checks for system administrators written in GoLang. Please keep in mind that every binary requires `bash` to be installed on the system.

## Reboot after kernel update check for Linux
### Installation
```
wget https://github.com/yaroslav-gwit/system-checks/releases/download/0.01-alpha/reboot_after_kern_update

chmod +x reboot_after_kern_update

./reboot_after_kern_update
```
### Example output
If the reboot is not required:
```
All good. You are running the latest kernel version.
```
If the reboot is required:
```
Please reboot to apply the kernel update!

You are running:      3.10.0-1160.53.0.el7.x86_64
The latest installed: 3.10.0-1160.53.1.el7.x86_64
```