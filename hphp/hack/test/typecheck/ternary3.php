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

function f1(): ?string {
  if (true) {
    return 'x';
  } else {
    return null;
  }
}

function f2(): string {
  return 'y';
}

function t(): string {
  return f1() ?: f2();
}
