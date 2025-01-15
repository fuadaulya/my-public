### The Problem

An `m` x `n` map grid is provided to you. Cells are marked either. Cells are marked either `#` or `.` where `#` represents `land` and `.` represents `water`.
You can assume that the following conditions are all true:

- Grid cells are connected horizontally/vertically/diagonally
- \*The map grid is surrounded by land

**Function Description**
Create a function that returns the number of lakes represented by a set of `water`.

Input Format

- The first line contains a single integer, `m`,the map's width.
- The second line contains a single integer, `n`, the map's height.
- The following lines represent the map grid.

**Constraints**

```
1 <= 'm' <= 100
1 <= 'n' <= 100
```

**Sample Input**

```
5
4
#.###
..###
##.##
####.
```

**Sample Input**

```
2
```

**Explanation**
**Lake = 1**

```
#+###
++###
##+##
####.
```

**Lake = 2**

```
#.###
..###
##.##
####+
```
