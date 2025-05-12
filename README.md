## Introduction

Mathematical algorithm implementation and benchmarks. [zml](https://github.com/wiscmath/zml.) library is the basis for
easy-to-use, high-performance mathematical software from WiscMath.

## Benchmarks
* Tests for execution time.
* __Hardware__: MacOS 12.4, Chip: Apple M1 Max, Total Number of Cores: 10 (8 performance and 2 efficiency), Memory: 64
  GB.

## Algorithm Implementations

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
    <td rowspan="6"><a href="euler.md">Euler's Number</a></td>
    <td rowspan="2"><a href="python/euler.py">euler.py</a></td>
    <td>n = 1</td>
  </tr>
  <tr>
    <td>n = 2</td>
  </tr>
  <tr>
    <td rowspan="2"><a href="go/euler.go">euler.go</a></td>
    <td>n = 1</td>
  </tr>
  <tr>
    <td>n = 2</td>
  </tr>
  <tr>
    <td rowspan="2"><a href="https://github.com/wiscmath/zml/blob/main/src/euler.zig">euler.zig</a></td>
    <td>n=1</td>
  </tr>
  <tr>
    <td>n=2</td>
  </tr>
  <tr>
    <td rowspan="6"><a href="archimedes.md">Archimedes's Constant</a></td>
    <td rowspan="2">Python</td>
    <td>n = 1</td>
  </tr>
  <tr>
    <td>n = 2</td>
  </tr>
  <tr>
    <td rowspan="2">Go</td>
    <td>n = 1</td>
  </tr>
  <tr>
    <td>n = 2</td>
  </tr>
  <tr>
    <td rowspan="2">Zig</td>
    <td>n = 1</td>
  </tr>
  <tr>
    <td>n = 2</td>
  </tr>
</tbody>
</table>       
