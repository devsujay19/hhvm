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

trait A {
  private function foo(): void {}
}

trait B {
  use A;

  public function bar(): void {
    $this->foo();
  }
}
