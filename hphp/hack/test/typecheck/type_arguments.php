<?hh
/**
 * Copyright (c) 2014, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree.
 *
 *
 */

class X {
}

function foo(int $x): void {
}

function bar(): void {
  $x = new X();
  invariant_violation('', $x->haha());
}
