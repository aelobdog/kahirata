rm kahirata_temp
./kahirata
[[ $1 == "xclip" ]] && cat kahirata_temp | xclip -i -selection clipboard
[[ $1 == "clip" ]] && cat kahirata_temp | clip.exe
rm kahirata_temp
