# HosterRed collection of system checks
A collection of various OS checks for system administrators written in GoLang. Please keep in mind that every binary requires `bash` to be installed on the system.

## Reboot after kernel update check for Linux
### Installation
```
wget https://github.com/yaroslav-gwit/system-checks/releases/download/0.02-alpha/hr_kern_reboot
chmod +x hr_kern_reboot
mv hr_kern_reboot /bin/hr_kern_reboot

hr_kern_reboot
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

## Update checker for Linux
```
wget https://github.com/yaroslav-gwit/system-checks/releases/download/0.02-alpha/hr_update_checker
chmod +x hr_update_checker
mv hr_kern_reboot /bin/hr_update_checker

hr_update_checker
```
