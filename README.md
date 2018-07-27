<h2 align="center"> How to make a URL redirect</h2>


-  Go to **[_redirects](https://github.com/abranhe/to/tree/master/_redirects)**
-  Create an **.html** file (preferably with the name of the link)
    - ex: *github.html* (where *github* is the link that you want to make)
- Inside the file you need to copy the code below and paste it on you file, change it with you content.

``` html
---
permalink: /github <!-- Link wil be avilable at go.abranhe.com/github -->
destination: http://github.com/abranhe  <!-- The link where you want to redirect -->
---
```


### Example:

[go.abranhe.com/github](http://go.abranhe.com/github) will redirect you to [github.com/abranhe](https://github.com/abranhe)
