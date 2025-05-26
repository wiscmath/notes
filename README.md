## Introduction

Mathematical algorithm implementation and benchmarks. [zml](https://github.com/wiscmath/zml.) library is the basis for
easy-to-use, high-performance mathematical software from WiscMath.

## Benchmarks
* Tests for execution time.
* __Hardware__: MacOS 12.4, Chip: Apple M1 Max, Total Number of Cores: 10 (8 performance and 2 efficiency), Memory: 64
  GB.

## Algorithm Implementations

<table><thead>
  <tr>
    <th>Algorithm</th>
    <th>Implementation</th>
    <th>Result</th>
  </tr></thead>
<tbody>
  <tr>
    <td rowspan="9">Euler</td>
    <td rowspan="3">Python</td>
    <td>n1 = 1</td>
  </tr>
  <tr>
    <td>n2 = 2</td>
  </tr>
  <tr>
    <td>n3 = 3</td>
  </tr>
  <tr>
    <td rowspan="3">Go</td>
    <td>n4 = 4</td>
  </tr>
  <tr>
    <td>n5 = 5</td>
  </tr>
  <tr>
    <td>n6 = 6</td>
  </tr>
  <tr>
    <td rowspan="3">Zig</td>
    <td>n7 = 7</td>
  </tr>
  <tr>
    <td>n8 = 8</td>
  </tr>
  <tr>
    <td>n9 = 9</td>
  </tr>
  <tr>
    <td rowspan="6">Archimedes's Constant</td>
    <td rowspan="2">Python</td>
    <td>n10 = 10</td>
  </tr>
  <tr>
    <td>n11 = 11</td>
  </tr>
  <tr>
    <td rowspan="2">Go</td>
    <td>n12 = 12</td>
  </tr>
  <tr>
    <td>n13 = 13</td>
  </tr>
  <tr>
    <td rowspan="2">Zig</td>
    <td>n14 = 14</td>
  </tr>
  <tr>
    <td>n15 = 15</td>
  </tr>
</tbody>
</table>

<table>
<thead>
  <tr>
    <th>Algorithm</th>
    <th>Implementation</th>
    <th>Result</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td rowspan="9"><a href="euler.md">Euler's Number</a></td>
    <td rowspan="3"><a href="python/euler.py">euler.py</a></td>
    <td><em>euler1</em>, 1542 <em>ns</em></td>
  </tr>
  <tr>
    <td><em>euler2</em>, 1750 <em>ns</em></td>
  </tr>
  <tr>
    <td><em>euler3</em>, 0.013572167 <em>s</em></td>
  </tr>
    <td rowspan="3"><a href="go/euler.go">euler.go</a></td>
    <td><em>euler1</em>, 4.421 <em>ns/op</em></td>
  <tr>
    <td><em>euler2</em>, 42.87 <em>ns/op</em></td>
  </tr>
   <tr><td>n6 = 6</td></tr>
  <tr>
    <td rowspan="3"><a href="https://github.com/wiscmath/zml/blob/main/src/euler.zig">euler.zig(zml)</a></td>
    <td>n7 = 7</td>
  </tr>
  <tr>
    <td>n8 = 8</td>
  </tr>
</tbody>
</table>       

## Reference Materials 
* [Implementation and Synthesis of Math Library Functions](https://dl.acm.org/doi/pdf/10.1145/3632874)
* [From Mathematics to Generic Programming](http://www.fm2gp.com)
* [CS053 The Matrix in Computer Science](https://cs.brown.edu/courses/cs053/current/lectures.htm)

## Other Libraries
* https://github.com/nomemory/neat-matrix-library
* https://github.com/ziglang/zig/blob/master/lib/std/math.zig
* https://github.com/ziglibs/zlm
* https://github.com/griush/zm
* https://mpmath.org/
* https://flintlib.org/
* https://www.sagemath.org/
* https://numpy.org/
* https://scipy.org/
* https://matplotlib.org/
* https://docs.python.org/3/library/math.html
* https://pkg.go.dev/math
* https://rust-lang-nursery.github.io/rust-cookbook/science/mathematics.html

## Implementation of math functions
* https://stackoverflow.com/questions/2284860/how-does-c-compute-sin-and-other-math-functions
* https://sourceware.org/git/?p=glibc.git;a=blob;f=sysdeps/ieee754/dbl-64/s_sin.c;hb=HEAD#l194
* https://git.musl-libc.org/cgit/musl/tree/src/math/sin.c
* https://github.com/golang/go/blob/master/src/math/sin.go#L185
* https://www.netlib.org/fdlibm/s_sin.c

