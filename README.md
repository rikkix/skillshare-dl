# SkillShare Downloader
This project aims to help skillshare premium users easier to access their 
lessons when network is unstable or not available.  

! DO **NOT** use this project for piracy !  

## Feature
- 5 threads by default
- ...

## Install
### Linux AMD64
```bash
$ sudo wget https://github.com/iochen/skillshare-dl/releases/latest/download/skillshare-dl_amd64_linux -O /usr/bin/skillshare-dl
$ sudo chmod +x /usr/bin/skillshare-dl 
```
### Other Platforms
1. Download from https://github.com/iochen/skillshare-dl/releases/latest
2. rename to `skillshare-dl`
3. give execute permission to it
4. add or put it to your path

## Quick start
1. Login to your skillshare premium account on browser, press **F12** and type  
```javascript
document.cookie
```
in **Console**    

2. Copy and save the output to a file such as `cookie.txt`

3. run(or build it yourself) **skillshare-dl**, with command like below
```bash
$ skillshare-dl -cookie cookie.txt -id 970659408
# or
$ skillshare-dl -cookie cookie.txt -id 970659408 -id 652554100 -id ...
```

## Thanks
[kallqvist/skillshare-downloader](https://github.com/kallqvist/skillshare-downloader)
