# go-elastictable [![Build Status](https://travis-ci.org/rdsubhas/go-elastictable.svg?branch=master)](https://travis-ci.org/rdsubhas/go-elastictable) [![godoc](https://godoc.org/github.com/rdsubhas/go-elastictable?status.svg)](https://godoc.org/github.com/rdsubhas/go-elastictable)

Go command line tables with elastic column resizing to adapt terminal width.

## Approach

Naive algorithms simply take `width-of-table/width-of-terminal` and apply a uniform factor to grow/shrink the columns. This results in shoddy output, small columns become worse when wrapped.

This library optimizes for **visual balance** by doing the following:

- Closely pack wrapped lines:
    ```
    # don't:
    | don't waste whitespace in the second |
    | line                                 |

    # do:
    | when wrapping a column |
    | try to pack the lines  |
    | as closely as possible |
    ```

- Balance row heights:
    ```
    # don't:
    | COL-1       | COL-2       | COL-3       |
    +-------------+-------------+-------------+
    | small col   | small col   | really big  |
    |             |             | column that |
    |             |             | spans many  |
    |             |             | rows        |

    # do:
    | COL-1 | COL-2 | COL-3                   |
    +-------+-------+-------------------------+
    | small | small | really big column that  |
    | col   | col   | spans many rows         |
    ```

- Number of lines depends on wrapping behavior and is hard to predict in advance. So instead of fully computing expensive re-wraps for every pass, it rather just guesses the number of lines using simple rune widths. This works well in most cases, except there will be some trade-offs in certain scenarios.

## Compatibility Note

This library is currently under active development. Please use a dependency manager (like godep) to pin specific commit versions.
