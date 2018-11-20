#! /bin/bash
# created by tanck, it's part of a feature tracking system

features_go="../vendor/tanck/eagleeye/features.go"
features_go_templ="../vendor/tanck/eagleeye/features.go_templ"

# generate a new feature file
cp -f $features_go_templ $features_go

find .. -name '*.go' -type f -exec  grep '//#-CK_FEATURE_MAGIC-#' {} \; | while read line 
do 
  eval $(echo $line |awk 'split($0, a, "\"") {$5 = a[2]}  {printf("author=%s; name=%s; date=%s; desc=\"%s\";",$2,$3,$4,$5)}')
  #echo $author,$name,$date,$desc

  if [ ! -n "$author" -o ! -n "$name" -o ! -n "$date" -o ! -n "$desc" ]; then
    continue
  fi

  # whether the name & date & desc existes
  t_desc=$(echo $desc | sed 's/-/\\-/g')
  t_name=$(echo $name | sed 's/-/\\-/g')
  t_date=$(echo $date | sed 's/-/\\-/g')
  t_author=$(echo $author | sed 's/-/\\-/g')
  feature=$(cat $features_go | grep "$t_desc" | grep $t_name | grep $t_date | grep $t_author)

  if [ ! -n "$feature" ]; then
    newfeature="{Author: \"$author\", Name: \"$name\", Date: \"$date\", Desc: \"$desc\"},"
    #echo $newfeature

    # subtitue / to \/
    snewfeature=$(echo $newfeature | sed 's/\//\\\//g')

    # modify features.go
    sed -i "s/\/\*\#FEATURE_MAGIC\#\*\//$snewfeature \/\*\#FEATURE_MAGIC\#\*\//g" $features_go
    echo "feature $name added!"
  else
    echo "feature $name already added!"
  fi
  
done


