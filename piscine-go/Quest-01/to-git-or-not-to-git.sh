curl https://zone01normandie.org/assets/superhero/all.json | jq '.[] | select(.id == 170) .name' 
curl https://zone01normandie.org/assets/superhero/all.json | jq '.[] | select(.id == 170) .powerstats .power' 
curl https://zone01normandie.org/assets/superhero/all.json | jq '.[] | select(.id == 170) .appearance .gender' 