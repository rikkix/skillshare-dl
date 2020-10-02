# SkillShare Downloader
This project aims to help skillshare premium users easier to access their 
lessons when network is unstable or not available.  

! DO NOT use this project for piracy !  

## Feature
- 5 threads by default
- ...

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
```

## Thanks
[kallqvist/skillshare-downloader](https://github.com/kallqvist/skillshare-downloader)
