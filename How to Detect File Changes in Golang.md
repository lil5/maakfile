[![Steve Domino](https://miro.medium.com/fit/c/56/56/1*zvG5UTm1h9lVsd9rYVh3_g.png)](https://medium.com/@skdomino?source=post_page-----5b5a247cf71f-----------------------------------)

![](https://miro.medium.com/max/1400/1*DA0M0c9PP1w0YihC3Niukg.png)

Image by [deer](http://simpledesktops.com/browse/desktops/2013/jul/11/squares/)

Go has quickly become one of my favorite languages to write in. I often find my self smiling at how easy Go makes almost every task. Detecting file changes is no exception. Using the `filepath` package from the standard library along with `[fsnotify](https://github.com/fsnotify/fsnotify)`, makes this a fairly simple task.

Out of the box, `fsnotify` provides all the functionality needed to watch a single file or directory:

Unfortunately `fsnotify` doesn’t support recursion, so `filepath` will have to fill in the gaps:

Since `fsnotify` is able to watch all the files in a directory, `filepath` is used to walk the entire document tree looking only for directories to attach watchers to.