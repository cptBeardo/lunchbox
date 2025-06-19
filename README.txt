This is my first un-guided project through Boot.dev.  The initial purpose of this
project will be to keep inventory of food in the house and to create a menu
of items that I would like to make for different meals during the week.

In addition, it will keep track of staples that need to be replenished, as well as
taking recipes that we would like to make on a given night and create a shopping list
based on both.


(while this is the initial intention, it may morph into being an inventory tool overall
used to keep track of musical instruments checked out to students, or for a small retail
business keeping track of sales. Or, better yet, an inventory tracker for my piano tech
business)

==========================================================================================
After consideration, the simplest course of action will be to make this a database
of school owned musical instruments and whether they are checked out to a student, out for 
repair, or in-house ready for use.  Will need the following:
- Instrument type
- Serial number
- Repair status
    - date out
    - date in
- Check-Out status
- Student who checked it out
    - date out
    - condition
        - new
        - good
        - fair
    - date in
        - good
        - fair
        - bad (needs repair)

===========================================================================================
Utilizing a static HTML page for use as browser support (and for personal business later on)

===========================================================================================
decided to use Go instead of C because it is MUCH easier to use. Need to implement a CSV file
and will be using python to grab and insert the information to that file later on.
