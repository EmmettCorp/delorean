WIP
---

# DeLorean 

## What is DeLorean.

DeLorean is a btrfs manager. It allows you to create, list, delete snapshots for any btrfs subvolumes on your system (even hot pluggable devices).

Delorean doesn't force you to name your subvolume specifically. You can choose any name you want.

Delorean helps managing snapshots of any mounted btrfs subvolume.

Delorean allows to rollback subvolumes, that are children of the top level subvolume, using ui. Check [Flat layout](https://btrfs.wiki.kernel.org/index.php/SysadminGuide#Flat) from btrfs wiki.

## Features

- cli-ui
- mouse support
- support snapshots managing for hot pluggable devices 
- easy rollback of top level children


## UI

<img src="assets/scrnsht.png" width="700">

## Alternatives

There are mature and awesome tools that help you to manage btrfs as well but with different from delorean approach.

Each of them (delorean as well) has it's own restrictions and advantages. Choose more appropriate to your cases.

[Timeshift](https://github.com/teejee2008/timeshift)

[snapper](https://github.com/openSUSE/snapper) 