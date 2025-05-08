# Chooising a Web Conformance Guidelines
* law suits happen
* there are 2 section 508 and WCAG
* 508 is meant for federal project
* wcag - globally accepted
* 78 cateogries, 13 sections (POUR)
pERCEIVABLE - users must see
Operable- user be able to operate 
Understandble - users should understand
Robust - site should work everywhere

A- lowest level min reqs
AA - mid leve
AAA - highest


choose WCAG 2.1 AA it includes section 508

# HTML
* it lies in markup, so much accessibility is free
* using !doctype is good for accessibility
![1660828419548](image/README/1660828419548.png)
* no ids are duplicated

* set the lang here
* ![1660828503051](image/README/1660828503051.png)

* set these
![1660828558885](image/README/1660828558885.png)

* always include title for each page
* make sure viewport for mobile scale is like this
![1660828610090](image/README/1660828610090.png)

* use relative em
1.4.1 - text should sacle to 200%, nor horizontal scrollbars
use landmark block 
A 2.4.1 
* header,nav,main,footer,aside
* main can only be used once per page

* for nav give aria-label
__site__ - site level nave
__page__ - page level nav

* every page should only have 1 h1 tag should be first element in main section
* for headings you have to go in order h1 - h6, 

## Lists
* ordered, unordered and description list
level A 1.3.1
* you can use css but lose acsiibility

## Nav and Skip links
* A 1.4.5 - use text instead of images to convey meaning when possible

* AA 3.2.3-  Constient Nav
* AA 2.4.5- more than one to fidnd stuff in your website
* AA 3.2.4 - Consistent
* A 2.4.4 - purpose of link should be by link text alone
  * avoid Read more
  * visually hidden css when you dont want to change the text
![1660829468973](image/README/1660829468973.png)

* A 2.4.1 - skip link
  * frustruating for readers to have to tab through section to get to main content
  goes in header
```html
<header>
<a class="skip-link" href="#main">Skip to main </a>
</header>
...
<main id="main" class="main" tabindex="-1">
```

* some skip link code
![1660829707796](image/README/1660829707796.png)

![1660829732452](image/README/1660829732452.png)

## Tables
* caption tag title for table
  must be first child of table
 * thead,tfoot,tbody
 th - table header
 scope headers
  row, to indicate its the title for the row
  col, to indicate its an entry

* table tag is the container
* thead - a row contianer for thead
* tr - table row
* td , table entry

![1660830058418](image/README/1660830058418.png)


*complex table
  * colspan - sets the width of a column
  * use visuallyHidden for table summary
  * table summary used for complex pages
![1660830201323](image/README/1660830201323.png)

* avoid complex tables

* tfoot tag goes before body
## Forms
* no placeholders
AA 1.4.11- contrast ratio for 3.1
A 1.3.1. - proper labels provide formats for dates and currencies
A 4.1.2- your same field must be made with the same level of accessibility
A 1.3.3 - 
* no keyboard traps, 
* turn off, remap or activate keyboard shortcuts
* color cant be 

* use field set to group fields, and legenend is the label for the fieldset
use aria-invalid attribute
* never specify tabindex greater than one it will lead to unpredicatble results
* tabindex -1 takes out of the natural order, tabindex = 0 places in natural order
* use your own custom css focus,
* error prevention, web pages that cause legal commitments or financial transcactions must be reversible,checked and confirmed
* users should be able to turn off,adjust or extend the time limit 

# Media

## Intro
1.1.1. - all non text has a text alt

## Images
* all img use must use alt tag
* images for decoration use empty alt
* describe the meaning of the image

## Background Images via CSS
* on spans we just have icons, wanna avoid image requests, so use the visually hidden class

## SVG
![1661032016120](image/README/1661032016120.png)

![1661032126702](image/README/1661032126702.png)

## Audio
* time-based media
podcasts,
* create a transcript, 
* speech reconginition, google,apple windows 
* descrbe whole speech

## Video
* caption
  * __open caption__ embedded in video
  * __closed caption__ toggle
* srt - feature to synchronize audio with video
* never allow anything to flash 3 times

# Responsive Web Design and Accessibility

## Switching Context
* when main nav switches from main nav to dropdown the dropdown options for navigation must not cause navigation select, so if someone is using the key board let them choose a dropdown option and then navigate as necessary

## Order of Content/Focus
* visual order must match the dom order
* screen readers are used with the managifying glass
* autofocus is only good for google search

## Mobile Devices
* don't lock orientation
* all functions with gestures can be opeprated with a single pointer
* must not have pointer lock
* if there is motion control there should be an alternative do do the same control as well

## Additional Responsive Guideless
* watch for offscreen content, people may have to tab through your site and if it has too many tabs may leave,
use html5 hidden attribute
  then in order of avaiialbleity
  display:none,visibility:hidden,aria-hidden="true"
* use  relative units
AA 1.4.12- Text Spacing
![1661033286012](image/README/1661033286012.png)
```css
html,body{
  font-size:1rem;
  line-height:1.5;
  letter-spacing:.12rem;
  word-spacing:.16rem;
}

p + p{
  margin-top:2rem
}
```
AA 1.4.10
content should resize to 320 css and 256 css w/o losing content
