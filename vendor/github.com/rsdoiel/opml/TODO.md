
## Next

+ [ ] create a command line tool that reads an OPML and appends a element to the list (e.g. adds a feed URL to an OPML list of feeds)
    + basic verbs would be insert (insert a new list element), append (a new list element), replace (replace a list element) delete (a list element), and find (return the path to an element by name or attribute value)
        + append, insert, replace takes a path and the value to update with
        + delete takes a path to the item to be deleted, if you try to delete the root list you get back an empty list
        + find takes an attribute name and value and returns a path
    + the tree the location for the verb to operate would be assumed to be the root or the index number of the items as a path
        + append "URL" would append the url to the root of the list
        + append "URL" /3 would append the UR" to the with item in the root list
        + append "URL" /3/2 would apend the URL the third items' second entry 
        + find ITEN_NAME append NEW_ITEM would find the item in the OPML tree, then append the content
+ [ ] create a tool that can read an OPML file, harvest the current feeds, index for browsing and support an option to send interesting articles to Pocket
+ [ ] Add a opml.Walk() function to package
+ [ ] Add support to process Frontier's fttb into OPML
    + See http://scripting.com/fatpages/about.html, http://scripting.com/fatpages/faq.html and http://scripting.com/fatpages/outline.html
    + fttb is a "fatpages" document, it is a Base 64 encoded document like is done with email.

## Someday, maybe

+ opmlviewer - a cli/terminal based opml viewer
+ omplfilter
    + might function something like Unix find command
+ ompledit - CRUD operations to individual Outline elements and element lists
+ Review River5 by Dave Winer for insights into what additional functions are needed in package 

## Completed

+ [x] Add opml2json 
+ [x] Add support for custom attributes
+ [x] Add Bash script to fetch Dave Winer's userland samples at http://scripting.com/misc/userlandSamples.zip
