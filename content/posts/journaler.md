+++
title = 'Journaler'
date = 2024-08-26T01:13:13-04:00
draft = false
+++

I like journaling using Neovim at times, so I created a new script called `journaler`. It removes the toil of:
- having to `cd` into your journal directory
- `touch` the file with a datestamp name I briefly had to look at the clock for
- opening the file using my desired editor

It is also the first time I turn a Python script into a binary so that was pretty cool. It was a very simple and easy mini-project.

Here is the [link to the repo](https://github.com/lv/journaler).


## Future plans

- Add git support
  - `-c` flag: Commits the journal directory
  - `-p` flag: Pushes the repo to `origin`
