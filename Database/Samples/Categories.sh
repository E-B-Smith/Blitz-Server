#!/bin/bash
set -euo pipefail

command="copy CategoryTable (isLeaf, parent, item, description) from stdin with (format csv, header);"

psql blitzlabs blitzlabs --command "$command" <<ENDOFDATA
isLeaf,parent,item,description
f,root, Personal            ,"Legal, Financial, Lifestyle, Travel, and more"
f,root, Professional        ,Industry specific intelligence and learning
f,Personal, Legal               ,Seek help from lawyers and clarify your legal questions
f,Personal, Financial Planning,"Tax planning, Accounting,  Trading, and Investing Personal finances"
f,Personal, Health & Fitness  ,"Nutrition, Diet, Training, Athletics, and Well-being"
f,Personal, Travel            ,"Visa requirements, Recommendations, Tips"
f,Personal, Genius Bar,"Home equipment, Computers, TV/Audio, Fashion, etc."
f,Professional, Industry          ,
f,Professinal, Function          ,
f,Industry, Telecom           ,
f,Industry, Retail             ,
f,Industry, Technology        ,
f,Industry, Startups          ,
f,Function, Marketing         ,
f,Function, Sales             ,
f,Function, Software          ,
f,Function, Operations        ,
t,Software,software,
t,Software,ios,
f, Genius Bar,Fashion,
t,Fashion,shoes,
ENDOFDATA
