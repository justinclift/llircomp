A simple cli utility to parse two .ll files, and display the differences.

Based on the initial Go code from Step 3 of:

  https://github.com/llir/llvm/issues/86#issuecomment-498357924

At the moment it's extremely basic, just to assist with checking things
semi-automatically instead of having to do everything manually.

It doesn't take command line arguments yet, instead the two file names
it compares are hard coded as `target.ll` and `foo.ll`.

PRs to improve and expand things are welcome. :smile:

Example output:

```
$ llircomp
Differences to investigate:
  * Source filename
      Target: 'target.c'
        Work: ''

  * Data layout
      Target: 'e-m:e-i64:64-f80:128-n8:16:32:64-S128'
        Work: ''

  * Target triple
      Target: 'x86_64-pc-linux-gnu'
        Work: ''

  * Function 'foo'
    * # of Attributes
        Target: '1'
          Work: '0'

  * Function 'llvm.dbg.declare'
    * # of Attributes
        Target: '1'
          Work: '0'

  * Function 'main'
    * # of Attributes
        Target: '1'
          Work: '0'

  * Metadata element !27 is different
      Target: '!DILocation(line: 8, column: 5, scope: !23)'
        Work: '!DILocation(line: 8, column: 5, scope: !22)'

  * Named metadata definition 'llvm.dbg.cu'... no differences detected
  * Named metadata definition 'llvm.module.flags'...
    * Node '1' differs
        Target: '!4'
          Work: '!6'

    * Node '2' differs
        Target: '!5'
          Work: '!7'

  * Named metadata definition 'llvm.ident'...
    * Node '0' differs
        Target: '!6'
          Work: '!7'
```
