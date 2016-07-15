#!/bin/bash
set -euo pipefail

command="copy CategoryTable (isLeaf, parent, item, description) from stdin with (format csv, header);"

psql blitzlabs blitzlabs --command "$command" <<ENDOFDATA

truncate CategoryTable;
insert into CategoryTable
     (parent, item, isleaf) values
     ('root', 'Consumer', false)
    ,('root', 'Technology, Media, Telecom', false)
    ,('root', 'Healthcare', false)
    ,('root', 'Financial Services', false)
    ,('root', 'Energy and Utilities', false)
    ,('root', 'Legal', false)
    ,('root', 'Accounting', false)
    ,('root', 'Travel', false)

    ,('Technology, Media, Telecom', 'Software', true)

    ,('Software', 'software', true)
    ,('Software', 'ios', true)
    ,('Software', 'technology', true)
    ,('Software', 'android', true)
    ,('Software', 'go', true)
    ,('Software', 'go-lang', true)
    ,('Software', 'mobile', true)
    ,('Software', 'programming', true)
    ,('Software', 'programmer', true)
    ,('Software', 'software engineer', true)
    ,('Software', 'iphone', true)
    ,('Software', 'quality assurance', true)
    ,('Software', 'qa', true)
    ;

ENDOFDATA
