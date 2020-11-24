rm enhira_temp
./enhira
[[ $1 == "xclip" ]] && cat enhira_temp | xclip -i -selection clipboard
[[ $1 == "clip" ]] && cat enhira_temp | clip.exe
rm enhira_temp
