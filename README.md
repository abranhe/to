<p align="center">
  <a href="http://go.mdc.blue">
        <img src="https://mdc.blue/badge.svg"
            alt="MDC Blue"></a>
<!-- 
  Author: Carlos Abraham 
  github.com/19cah
  www.19cah.com
-->
</p>

<h2 align="center"> How to make a URL redirect</h2>


-  Go to **[_redirects](https://github.com/MDCblue/redirect/tree/master/_redirects)**
-  Create an **.html** file (preferably with the name of the link)
    - ex: *github.html* (where *github* is the link that you want to make)
- Inside the file you need to copy the code below and paste it on you file, change it with you content.

``` html
---
permalink: /github <!-- Link wil be avilable at go.mdc.blue/github -->
destination: http://github.com/mdcblue  <!-- The link where you want to redirect -->
---
```


### Example:

[go.mdc.blue/github](http://go.mdc.blue/github) will redirect you to [github.com/MDCblue](https://github.com/MDCblue)

**NOTE:** when you are done, please include you path on the [path list](https://github.com/MDCblue/redirect/blob/master/list/README.md)
